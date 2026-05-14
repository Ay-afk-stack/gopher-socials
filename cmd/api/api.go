package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ay-afk-stack/gopher-socials/internal/auth"
	"github.com/Ay-afk-stack/gopher-socials/internal/mailer"
	"github.com/Ay-afk-stack/gopher-socials/internal/store"
	"github.com/Ay-afk-stack/gopher-socials/internal/store/cache"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type application struct {
	config config
	store  store.Storage
	cacheStorage cache.Storage
	logger *zap.SugaredLogger
	mailer mailer.Client
	authenticator auth.Authenticator
}

type config struct {
	addr string
	db   dbConfig
	env  string
	mail mailConfig
	frontendURL string
	auth authConfig
	redisCfg redisConfig
}

type redisConfig struct {
	addr string
	pw string
	db int
	enabled bool
}

type authConfig struct {
	basic basicConfig
	token tokenConfig
}

type tokenConfig struct {
	secret string
	exp time.Duration
	issuer string
}

type basicConfig struct {
	user string
	pass string
}

type mailConfig struct {
	resend resendConfig
	exp string
}

type resendConfig struct {
	apiKey string
	fromEmail string
}

type dbConfig struct {
	databaseURL     string
	maxConns        int
	minConns        int
	maxConnIdleTime string
	dbTimeout       string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(30 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.With(app.BasicAuthMiddleware()).Get("/health", app.healthCheckHandler)
		r.Route("/posts", func(r chi.Router) {
			r.Use(app.AuthTokenMiddleware)
			r.Post("/", app.createPostHandlers)

			r.Route("/{id}", func(r chi.Router) {
				r.Use(app.postContextMiddleware)

				r.Get("/", app.getPostHandler)

				r.Delete("/", app.CheckPostOwnerShip("admin", app.deletePostHandler))
				r.Patch("/", app.CheckPostOwnerShip("moderator", app.updatePostHandler))
			})
		})

		r.Route("/users", func(r chi.Router) {
			r.Put("/activate/{token}", app.activateUserHandler)

			r.Route("/{userID}", func(r chi.Router) {
				r.Use(app.AuthTokenMiddleware)

				r.Get("/", app.getUserHandler)
				r.Put("/follow", app.followUserHandler)
				r.Put("/unfollow", app.unFollowUserHandler)
			})

			r.Group(func(r chi.Router) {
				r.Use(app.AuthTokenMiddleware)
				r.Get("/feeds", app.getUserFeedHandler)
			})

		})

		r.Route("/auth", func (r chi.Router)  {
			r.Post("/users", app.registerUserHandler)
			r.Post("/token", app.createTokenHandler)
		})
	})

	return r
}

func (app *application) run(mux http.Handler) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <- quit

		ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
		defer cancel()

		app.logger.Infow("signal caught", "signal", s.String())

		shutdown <- srv.Shutdown(ctx)
	} ()

	app.logger.Infow("server has started", "addr", app.config.addr, "env", app.config.env)

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	err := <- shutdown
	if err != nil {
		return err
	}

	app.logger.Infow("server has stopped", "addr", app.config.addr, "env", app.config.env)

	return nil
}
