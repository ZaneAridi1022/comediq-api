package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/comediq-api/shows/definition"
	"github.com/comediq-api/utils"
	"github.com/comediq-api/venues"
	_ "github.com/joho/godotenv/autoload"

	"github.com/comediq-api/database"
	"github.com/comediq-api/shows/historical"
	"github.com/comediq-api/shows/upcoming"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	url, key := os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY")
	if url == "" || key == "" {
		fmt.Println("SUPABASE_URL or SUPABASE_KEY environment variables not set")
		os.Exit(1)
	}

	err := database.Init(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))
	if err != nil {
		fmt.Println("database failed to initialize", err)
		os.Exit(1)
	}

	test()
	return

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

func test() {
	showDefinitions := generateRandomUnregisteredShowDefinitions()
	for _, show := range showDefinitions {
		_, err := definition.Create(&show)
		if err != nil {
			fmt.Println("database failed to create show definition", err)
		}
	}

	upcomingShows := generateRandomUnregisteredUpcomingShows()
	for _, show := range upcomingShows {
		_, err := upcoming.Create(&show)
		if err != nil {
			fmt.Println("database failed to create upcoming show", err)
		}
	}

	historicalShows := generateRandomUnregisteredHistoricalShows()
	for _, show := range historicalShows {
		_, err := historical.Create(&show)
		if err != nil {
			fmt.Println("database failed to create historical show", err)
		}
	}

	showDef := definition.Show{
		ID: utils.Int32(1),
		VenueRoom: venues.VenueRoom{
			Venue: venues.Venue{
				Name:          "a",
				Address:       "a",
				City:          "a",
				Neighbourhood: "a",
				Borough:       "a",
				Type:          "a",
				Contact:       "a",
			},
			Name: "b",
		},
		Name: "Scheduled Show",
		OccurrenceRules: []definition.WeekDayInAMonthRule{
			{
				WeeksInAdvance:     5,
				WeeksOfMonthPolicy: [5]bool{true, true, false, true, true},
				WeekDay:            int32(time.Monday),
				StartTime:          timeParse("1:00 AM"),
				EndTime:            timeParse("1:00 PM"),
			},
			{
				WeeksInAdvance:     2,
				WeeksOfMonthPolicy: [5]bool{false, false, false, false, false},
				WeekDay:            int32(time.Tuesday),
				StartTime:          timeParse("2:00 AM"),
				EndTime:            timeParse("2:00 PM"),
			},
			{
				WeeksInAdvance:     4,
				WeeksOfMonthPolicy: [5]bool{true, false, true, false, true},
				WeekDay:            int32(time.Wednesday),
				StartTime:          timeParse("3:00 AM"),
				EndTime:            timeParse("3:00 PM"),
			},
		},
	}
	err := definition.AutoSchedule(&showDef)
	if err != nil {
		fmt.Println("failed to auto schedule", err)
	}
}

func generateRandomUnregisteredHistoricalShows() []historical.Show {
	var shows []historical.Show
	for i := 1; i <= 100; i++ {
		shows = append(shows, historical.Show{
			VenueRoom:        getRandomVenueRoom(),
			Name:             fmt.Sprintf("Name %d", i),
			Host:             chanceOfNil(0.3, fmt.Sprintf("Host #%d", i)),
			AudienceCost:     chanceOfNil(0.3, fmt.Sprintf("$%d", i)),
			ComedianCost:     chanceOfNil(0.5, fmt.Sprintf("$%d", i)),
			StageTime:        chanceOfNil(0.5, fmt.Sprintf("%d minutes", i)),
			ComedianLineup:   randomComedianLineUp(),
			StartDateAndTime: time.Now().Add(time.Duration(rand.Intn(500)) * time.Second),
			EndDateAndTime:   time.Now().Add(time.Duration(500+rand.Intn(500)) * time.Second),
			Open:             randBool(),
		})
	}
	return shows
}

func generateRandomUnregisteredShowDefinitions() []definition.Show {
	var showDefinitions []definition.Show
	for i := 1; i <= 100; i++ {
		showDefinitions = append(showDefinitions, definition.Show{
			VenueRoom:          getRandomVenueRoom(),
			Name:               fmt.Sprintf("Name %d", i),
			SignupInstructions: fmt.Sprintf("Signup instructions #%d", i),
			Notes:              chanceOfNil(0.3, fmt.Sprintf("Notes #%d", i)),
			AudienceCost:       chanceOfNil(0.3, fmt.Sprintf("$%d", i)),
			ComedianCost:       chanceOfNil(0.5, fmt.Sprintf("$%d", i+10)),
			StageTime:          chanceOfNil(0.5, fmt.Sprintf("%d minutes", i)),
			Host:               chanceOfNil(0.3, fmt.Sprintf("Host #%d", i)),
			SMS:                chanceOfNil(0.3, fmt.Sprintf("%d", i*12312312)),
			Verified:           fmt.Sprintf("Verified #%d", i),
			Open:               randBool(),
			OccurrenceRules:    randomOccurrenceRules(),
		})
	}
	return showDefinitions
}

func generateRandomUnregisteredUpcomingShows() []upcoming.Show {
	var shows []upcoming.Show
	for i := 1; i <= 100; i++ {
		shows = append(shows, upcoming.Show{
			VenueRoom:          getRandomVenueRoom(),
			Name:               fmt.Sprintf("Name %d", i),
			SignupInstructions: fmt.Sprintf("Signup instructions #%d", i),
			Notes:              chanceOfNil(0.3, fmt.Sprintf("Notes #%d", i)),
			AudienceCost:       chanceOfNil(0.3, fmt.Sprintf("$%d", i)),
			ComedianCost:       chanceOfNil(0.5, fmt.Sprintf("$%d", i)),
			StageTime:          chanceOfNil(0.5, fmt.Sprintf("%d minutes", i)),
			Host:               chanceOfNil(0.3, fmt.Sprintf("Host #%d", i)),
			SMS:                chanceOfNil(0.3, fmt.Sprintf("%d", i*12312312)),
			Verified:           fmt.Sprintf("Verified #%d", i),
			ComedianLineup:     randomComedianLineUp(),
			StartDateAndTime:   time.Now().Add(time.Duration(rand.Intn(500)) * time.Second),
			EndDateAndTime:     time.Now().Add(time.Duration(500+rand.Intn(500)) * time.Second),
			Active:             randBool(),
			Open:               randBool(),
		})
	}
	return shows
}

func getRandomVenueRoom() venues.VenueRoom {
	return venues.VenueRoom{
		Venue: venues.Venue{
			Name:          fmt.Sprintf("Venue #%d", rand.Intn(2)),
			Address:       "1234 Example Street",
			City:          "Yorknew City",
			Neighbourhood: "Examplehood",
			Borough:       "some borough",
			Type:          "Normal",
			Contact:       "+1 123 123 1234",
		},
		Name: fmt.Sprintf("Room #%d", rand.Intn(5)+1),
	}
}

func chanceOfNil[T any](chance float64, val T) *T {
	if rand.Float64() < chance {
		return nil
	}
	return &val
}

func randomComedianLineUp() []string {
	comedianLineUp := []string{}
	for i := 0; i < rand.Intn(5); i++ {
		comedianLineUp = append(comedianLineUp, "Name")
	}
	return comedianLineUp
}

func randomOccurrenceRules() []definition.WeekDayInAMonthRule {
	occurrenceRules := []definition.WeekDayInAMonthRule{}
	for i := 0; i < rand.Intn(10); i++ {
		occurrenceRules = append(occurrenceRules, definition.WeekDayInAMonthRule{
			WeeksInAdvance:     int32(rand.Intn(5) + 1),
			WeeksOfMonthPolicy: [5]bool{randBool(), randBool(), randBool(), randBool(), randBool()},
			WeekDay:            int32(rand.Intn(7)),
			StartTime:          time.Now().Add(time.Duration(rand.Intn(500)) * time.Second),
			EndTime:            time.Now().Add(time.Duration(500+rand.Intn(500)) * time.Second),
		})
	}
	return occurrenceRules
}

func randBool() bool {
	return rand.Intn(2) == 1
}

func timeParse(s string) time.Time {
	parse, _ := time.Parse("3:04 PM", s)
	return parse
}
