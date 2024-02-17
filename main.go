package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ekediala/aggregator/internal/database"
	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/cors"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	db, err := sql.Open("postgres", dbUrl)

	log.Print("Database up and running...")

	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	dbConfig := apiConfig{
		DB: dbQueries,
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()

	router.Mount("/api/v1", v1Router)

	v1Router.Get("/healthz", handlerReadiness)

	v1Router.Get("/error", errorHandler)

	v1Router.Post("/users", dbConfig.handlerUserCreate)

	v1Router.Post("/feeds", dbConfig.middlewareAuth(dbConfig.handlerFeedCreate))

	v1Router.Post("/feed_follows", dbConfig.middlewareAuth(dbConfig.handlerFeedFollowCreate))

	v1Router.Get("/feed_follows", dbConfig.middlewareAuth(dbConfig.handleGetUserFeedFollows))

	v1Router.Delete("/feed_follows/{feedFollowId}", dbConfig.middlewareAuth(dbConfig.handlerDeleteFeedFollow))

	v1Router.Get("/feeds", dbConfig.handlerGetAllFeeds)

	v1Router.Get("/user", dbConfig.middlewareAuth(dbConfig.handleGetOneUser))

	v1Router.Get("/posts", dbConfig.middlewareAuth(dbConfig.handlerGetUserPosts))

	server := &http.Server{Handler: router, Addr: ":" + port}

	log.Printf("Listening on port %v...", port)

	const SCRAPING_CONCURRENCY = 5

	go startScraping(dbConfig.DB, SCRAPING_CONCURRENCY, time.Hour)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
