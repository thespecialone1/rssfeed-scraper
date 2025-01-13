package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"githhub.com/thespecialone1/rssfeed-scraper/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

    _ "github.com/lib/pq"	//"_(underscore) means "include this code in my program even tho i'm not calling it directly""
)
type apiConfig struct {
	DB	*database.Queries
}

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("Port is not found in the envPath: ")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == ""{
		log.Fatal("DB_URL  is not found in the envPath: ")
	}
	conn, err:= sql.Open("postgres", dbURL)
	if err!=nil{
		log.Fatal("Can't connect to database:", err)
	}

	apiCfg:= apiConfig {
		DB:	database.New(conn),
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
	v1Router.Post("/users", apiCfg.handlerCreateUser)

	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}

	log.Printf("Server is listening on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
}
