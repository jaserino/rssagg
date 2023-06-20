package main

import (	
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	
	// pulls .env variables into current environment to use
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("PORT is not found in the environment")
	}

	// creating new chi router object
	r := chi.NewRouter()

	// connecting router to http server
	srv := &http.Server {
		Handler: r,
		Addr: ":" + portString,
	}
	log.Printf("Server starting on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}