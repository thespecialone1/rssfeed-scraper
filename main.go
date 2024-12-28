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
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("Port is not found in the envPath: ")
	}
	 
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		 // AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		 AllowedOrigins:   []string{"https://*", "http://*"},
		 // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		 AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		 AllowedHeaders:   []string{"*"},
		 ExposedHeaders:   []string{"Link"},
		 AllowCredentials: false,
		 MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	router.Mount("/v1", v1Router)
	v1Router.Get("/err", errHandler)

	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}

	log.Printf("Server is listening on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
}
