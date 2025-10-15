package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/comediq-api/database"
	"github.com/comediq-api/shows/historical"
	"github.com/comediq-api/shows/upcoming"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	err := database.Init(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))
	if err != nil {
		fmt.Println("database failed to initialize", err)
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.HandleFunc("/shows/historical", historical.HandleCreate).Methods("POST")
	router.HandleFunc("/shows/upcoming", upcoming.HandleCreate).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})
	handler := c.Handler(router)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		fmt.Println("PORT environment variable not set")
		os.Exit(1)
	}

	fmt.Println("Listening on port " + port)
	fmt.Println(http.ListenAndServe(":"+port, handler))

	// database.Init()
	// name := "The Tiny Cupboard"
	// database.FixMicUIDS()
	// helpers.ReadCSV()
	//
	// webscrape.Init()
	// handler := c.Handler(r)
}
