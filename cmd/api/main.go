package main

import (
	"time"

	"github.com/Ay-afk-stack/gopher-socials/internal/auth"
	"github.com/Ay-afk-stack/gopher-socials/internal/db"
	"github.com/Ay-afk-stack/gopher-socials/internal/env"
	"github.com/Ay-afk-stack/gopher-socials/internal/mailer"
	"github.com/Ay-afk-stack/gopher-socials/internal/store"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const version = "0.0.1"

func main() {
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	if err := godotenv.Load(); err != nil {
		logger.Fatal("error loading .env file: ", err)
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:4000"),
		db: dbConfig{
			databaseURL:     env.GetString("DB_URL", "postgres://postgres:postgres@localhost:5432/social?sslmode=disable"),
			maxConns:        env.GetInt("DB_MAX_CONNS", 20),
			minConns:        env.GetInt("DB_MIN__CONNS", 5),
			maxConnIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15min"),
			dbTimeout:       env.GetString("DB_TIMEOUT", "10s"),
		},
		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp: env.GetString("MAIL_TOKEN_EXP", "5m"),
			resend: resendConfig{
				apiKey: env.GetString("RESEND_API_KEY", ""),
				fromEmail: env.GetString("FROM_MAIL", "onboarding@resend.dev"),
			},
		},
		auth: authConfig{
			basic: basicConfig{
				user: env.GetString("AUTH_BASIC_USER", "admin"),
				pass: env.GetString("AUTH_BASIC_PASS", "admin"),
			},
			token: tokenConfig{
				secret: env.GetString("AUTH_JWT_SECRET", "example"),
				exp: time.Hour * 24 * 3,
				issuer: env.GetString("AUTH_TOKEN_ISSUER", "gopher_socials"),
			},
		},
	}

	pool, err := db.New(
		cfg.db.databaseURL,
		cfg.db.maxConns,
		cfg.db.minConns,
		cfg.db.maxConnIdleTime,
		cfg.db.dbTimeout,
	)
	if err != nil {
		logger.Fatal(err)
	}
	defer pool.Close()

	logger.Info("database connection pool established!")

	store := store.NewStorage(pool)

	mailer := mailer.NewResendMailer(cfg.mail.resend.apiKey, cfg.mail.resend.fromEmail)

	jwtAuthenticator := auth.NewJWTAuthenticator(cfg.auth.token.secret, cfg.auth.token.issuer, cfg.auth.token.issuer)

	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
		mailer: mailer,
		authenticator: jwtAuthenticator,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
