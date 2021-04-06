package main

import (
	sw "./server"
	"github.com/rs/cors"
	"log"
	"net/http"
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