package main

import (
	// "fmt"
	// "net/http"
	// "os"

  // "github.com/gorilla/mux"
	// "github.com/rs/cors"
	"github.com/comediq-api/database"
	// "github.com/comediq-api/helpers"
	"github.com/comediq-api/webscrape"
)

func main() {
	database.Init()
	// name := "The Tiny Cupboard"
	// database.FixMicUIDS()
	// helpers.ReadCSV()
	// 
	
	webscrape.Init()

  // r := mux.NewRouter()
	// c := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"}, 
	// 	AllowedMethods: []string{"GET", "POST", "OPTIONS"},
	// 	AllowedHeaders: []string{"Content-Type"},
	// })

	// handler := c.Handler(r)

	// port, ok := os.LookupEnv("PORT")
	// if !ok {
	// 	port = "3000"
	// }

	// fmt.Printf("Ready and listening on port %v", port)
	// err := http.ListenAndServe(fmt.Sprintf(":%v", port), handler)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}