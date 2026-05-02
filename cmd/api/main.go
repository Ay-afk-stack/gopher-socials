package main

import (
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

	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
		mailer: mailer,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
