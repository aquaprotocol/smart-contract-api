package main

import (
	"github.com/rs/cors"
	"log"
	"net/http"

	sw "./server"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":7000", handler))
}