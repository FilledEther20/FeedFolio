package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println("Hi")
	port_string := os.Getenv("PORT")
	if port_string == "" {
		log.Fatal("PORT is not found in the environment")
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	//Now we can create a subrouter that can be mounted on the primary(main) router to handle requests
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerError)

	router.Mount("/v1", v1Router)

	log.Println("Port starting on:", port_string)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port_string,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PORT:", port_string)
}
