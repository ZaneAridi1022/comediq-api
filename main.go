package main

import (
	"fmt"
	"net/http"
	"os"

  "github.com/gorilla/mux"
	"github.com/rs/cors"
	db "github.com/comediq-api/database"
)

func main() {
  r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, 
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

  // Wrap router with CORS handler
	handler := c.Handler(r)
	// r.HandleFunc("/getRawData", database.GetRawData)
	// r.HandleFunc("/updateTransaction", database.UpdateTransaction)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	fmt.Printf("Ready and listening on port %v", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), handler)
	if err != nil {
		fmt.Println(err)
	}
	db.Init()
}