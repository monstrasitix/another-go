package main

import (
	"fmt"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"github.com/monstrasitix/finance/internal/config"
	"github.com/monstrasitix/finance/internal/web"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	mux := chi.NewRouter()
	server := config.NewServer(mux)

	web.Router(mux)

	fmt.Printf("Running on: http://localhost%s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
