package main

import (
	"log"

	"github.com/Ay-afk-stack/gopher-socials/internal/db"
	"github.com/Ay-afk-stack/gopher-socials/internal/env"
	"github.com/Ay-afk-stack/gopher-socials/internal/store"
)

func main() {
	databaseURL := env.GetString("DB_URL", "postgres://postgres:postgres@localhost:5432/social?sslmode=disable")

	pool, err := db.New(databaseURL, 20, 5, "30m", "10s")
	if err != nil {
		log.Println("error during seeding:", err)
		return
	}

	store := store.NewStorage(pool)

	db.Seed(store)

	log.Println("seeding completed")
}
