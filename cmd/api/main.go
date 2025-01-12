package main

import (
	"log"

	"github.com/HarshBardolia01/gophers-connect/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
