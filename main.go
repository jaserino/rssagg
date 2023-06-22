package main

import (	
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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

	// cors to make a request from server to a browser
	r.Use(cors.Handler(cors.Options{	
		AllowedOrigins:   []string{"https://*", "http://*"},	
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, 
	  }))

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/err",handlerErr)


	r.Mount("/v1", v1Router)

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