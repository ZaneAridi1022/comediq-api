package database

import (
	"fmt"
	"os"
	"encoding/json"
	"github.com/supabase-community/supabase-go"
	"github.com/joho/godotenv"

	types "github.com/comediq-api/types"
	)

var client *supabase.Client

func Init() {
	fmt.Println("Creating DB client")
	godotenv.Load()
	URL := os.Getenv("SUPABASE_URL")
	KEY := os.Getenv("SUPABASE_KEY")
	var err error
	client, err = supabase.NewClient(URL, KEY, &supabase.ClientOptions{})
	
  if err != nil {
    fmt.Println("cannot initalize client", err)
  } else {
		fmt.Println("Client made")
	}
	
}

func GetMics(venueName *string, day *string, StartTime *string, LatestEndTime *string) ([]types.HistoricalMic, int64, error) {
	query := client.From("open_mics_historical").Select("*", "exact", false)
	var mics []types.HistoricalMic
	if venueName != nil {
		query = query.Filter("Venue Name", "eq", *venueName)
	}
	if day != nil {
		query = query.Filter("Day", "eq", *day)
	}
	if StartTime != nil {
		query = query.Filter("Start Time", "eq", *StartTime)
	}
	if LatestEndTime != nil {
		query = query.Filter("Latest End Time", "eq", *LatestEndTime)
	}
	data, count, err := query.Execute()
	if err != nil {
		return nil, 0, err
	} 
	err = json.Unmarshal(data, &mics)
	if err != nil {
		return nil, 0, err
	} 
	return mics, count, nil
}

func UpsertData(value any, tableName string) {
	data, count, err := client.From(tableName).Upsert(value, "", "", "exact").Execute()
	if err != nil {
		fmt.Println("Error upserting mic:", err)
	} else {
		fmt.Println("Rows affected:", count, string(data))
	}
}

func NukeMics() {
	_, count, err := client.From("open_mics_historical").Delete("", "exact").Neq("unique_identifier", "").Execute()
	if err != nil {
		fmt.Println("Error nuking open_mics_historical:", err)
	} else {
		fmt.Println("Nuked open_mics_historical. Rows affected:", count)
	}

	_, count, err = client.From("profile_open_mics").Delete("", "exact").Neq("unique_identifier", "").Execute()
	if err != nil {
		fmt.Println("Error nuking profile_open_mics:", err)
	} else {
		fmt.Println("Nuked profile_open_mics. Rows affected:", count)
	}

	_, count, err = client.From("open_mics_hosts_sms").Delete("", "exact").Neq("unique_identifier", "").Execute()
	if err != nil {
		fmt.Println("Error nuking open_mics_hosts_sms:", err)
	} else {
		fmt.Println("Nuked open_mics_hosts_sms. Rows affected:", count)
	}
}
