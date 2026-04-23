package main

import (
	"log"

	"github.com/Ay-afk-stack/gopher-socials/internal/env"
	"github.com/Ay-afk-stack/gopher-socials/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file: ", err)
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store: store,
		}

	mux := app.mount()

	log.Fatal(app.run(mux))
}