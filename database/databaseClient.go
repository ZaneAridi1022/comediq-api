package db

import (
	"fmt"
	"os"
	"github.com/supabase-community/supabase-go"
	"github.com/joho/godotenv"
	)

var client *supabase.Client

func Init() {
	fmt.Println("Creating DB client")
	godotenv.Load()
	URL := os.Getenv("SUPABASE_URL")
	KEY := os.Getenv("SUPABASE_KEY")
	client, err := supabase.NewClient(URL, KEY, &supabase.ClientOptions{})
	
  if err != nil {
    fmt.Println("cannot initalize client", err)
  } else {
		fmt.Println("Client made")
	}

  data, count, err := client.From("open_mics_historical").Select("*", "exact", false).Execute()
	if err != nil {	
		fmt.Println("error with query", err)
	} else {
		fmt.Println(count)
		fmt.Println(string(data))
	}
	
	
}