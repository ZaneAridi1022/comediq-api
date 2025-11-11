package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/comediq-api/mics/definitions"
	"github.com/comediq-api/mics/events/past"
	"github.com/comediq-api/validation"
	"github.com/comediq-api/venues"
	_ "github.com/joho/godotenv/autoload"

	"github.com/comediq-api/database"
	"github.com/comediq-api/mics/events/scheduled"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	validation.Init()

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
	fmt.Println("database client initialized")

	//test()
	//return

	router := mux.NewRouter()
	router.HandleFunc("/mics/definition", definitions.HandleCreate).Methods("POST")
	router.HandleFunc("/mics/scheduled", scheduled.HandleCreate).Methods("POST")
	router.HandleFunc("/mics/past", past.HandleCreate).Methods("POST")

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
	micDefinitions := generateRandomUnregisteredMicDefinitions()
	for _, mic := range micDefinitions {
		_, err := definitions.Create(&mic)
		if err != nil {
			fmt.Println("database failed to create mic definitions: ", err)
		}
	}

	scheduledMicEvents := generateRandomUnregisteredScheduledMics()
	for _, mic := range scheduledMicEvents {
		_, err := scheduled.Create(&mic)
		if err != nil {
			fmt.Println("database failed to create scheduled mic event: ", err)
		}
	}

	pastMicEvents := generateRandomUnregisteredPastMics()
	for _, mic := range pastMicEvents {
		_, err := past.Create(&mic)
		if err != nil {
			fmt.Println("database failed to create past mic event: ", err)
		}
	}

	micDef := definitions.Mic{
		Name: "Scheduled Mic",
		OccurrenceRules: []definitions.WeekDayInAMonthRule{
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
	id, _ := definitions.Create(&micDef)
	micDef.ID = &id
	err := definitions.AutoSchedule(&micDef)
	if err != nil {
		fmt.Println("failed to auto schedule: ", err)
	}
}

func generateRandomUnregisteredMicDefinitions() []definitions.Mic {
	var micDefinition []definitions.Mic
	for i := 1; i <= 100; i++ {
		micDefinition = append(micDefinition, definitions.Mic{
			VenueRoom:          getRandomVenueRoom(),
			Name:               fmt.Sprintf("Name %d", i),
			SignupInstructions: chanceOfNil(0.3, fmt.Sprintf("Signup instructions #%d", i)),
			Notes:              chanceOfNil(0.3, fmt.Sprintf("Notes #%d", i)),
			AudienceCost:       fmt.Sprintf("$%d", i),
			ComedianCost:       fmt.Sprintf("$%d", i+10),
			StageTime:          chanceOfNil(0.5, fmt.Sprintf("%d minutes", i)),
			Host:               chanceOfNil(0.3, fmt.Sprintf("Host #%d", i)),
			Instagram:          chanceOfNil(0.3, fmt.Sprintf("@%d", i)),
			SMS:                chanceOfNil(0.3, fmt.Sprintf("%d", i*12312312)),
			Verified:           chanceOfNil(0.3, fmt.Sprintf("Verified #%d", i)),
			Active:             randBool(),
			OccurrenceRules:    randomOccurrenceRules(),
		})
	}
	return micDefinition
}

func generateRandomUnregisteredScheduledMics() []scheduled.Mic {
	var mics []scheduled.Mic
	for i := 1; i <= 100; i++ {
		mics = append(mics, scheduled.Mic{
			VenueRoom:          getRandomVenueRoom(),
			Name:               fmt.Sprintf("Name %d", i),
			SignupInstructions: chanceOfNil(0.3, fmt.Sprintf("Signup instructions #%d", i)),
			Notes:              chanceOfNil(0.3, fmt.Sprintf("Notes #%d", i)),
			AudienceCost:       fmt.Sprintf("$%d", i),
			ComedianCost:       fmt.Sprintf("$%d", i),
			StageTime:          chanceOfNil(0.5, fmt.Sprintf("%d minutes", i)),
			Host:               chanceOfNil(0.3, fmt.Sprintf("Host #%d", i)),
			SMS:                chanceOfNil(0.3, fmt.Sprintf("%d", i*12312312)),
			ComedianLineup:     randomComedianLineUp(),
			StartDateAndTime:   time.Now().Add(time.Duration(rand.Intn(500)) * time.Second),
			EndDateAndTime:     time.Now().Add(time.Duration(500+rand.Intn(500)) * time.Second),
		})
	}
	return mics
}

func generateRandomUnregisteredPastMics() []past.Mic {
	var mics []past.Mic
	for i := 1; i <= 100; i++ {
		mics = append(mics, past.Mic{
			VenueRoom:        getRandomVenueRoom(),
			Name:             fmt.Sprintf("Name %d", i),
			Host:             chanceOfNil(0.3, fmt.Sprintf("Host #%d", i)),
			AudienceCost:     fmt.Sprintf("$%d", i),
			ComedianCost:     fmt.Sprintf("$%d", i),
			StageTime:        chanceOfNil(0.5, fmt.Sprintf("%d minutes", i)),
			ComedianLineup:   randomComedianLineUp(),
			StartDateAndTime: time.Now().Add(time.Duration(rand.Intn(500)) * time.Second),
			EndDateAndTime:   time.Now().Add(time.Duration(500+rand.Intn(500)) * time.Second),
		})
	}
	return mics
}

func getRandomVenueRoom() venues.VenueRoom {
	return venues.VenueRoom{
		Venue: venues.Venue{
			Name:          fmt.Sprintf("Venue #%d", rand.Intn(2)),
			Address:       "1234 Example Street",
			City:          "Yorknew City",
			Neighbourhood: chanceOfNil(0.5, "Examplehood"),
			Borough:       chanceOfNil(0.5, "some borough"),
			Type:          chanceOfNil(0.5, "Normal"),
			Contact:       chanceOfNil(0.5, "+1 123 123 1234"),
		},
		Name: chanceOfNil(0.5, fmt.Sprintf("Room #%d", rand.Intn(5)+1)),
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

func randomOccurrenceRules() []definitions.WeekDayInAMonthRule {
	occurrenceRules := []definitions.WeekDayInAMonthRule{}
	for i := 0; i < rand.Intn(10); i++ {
		occurrenceRules = append(occurrenceRules, definitions.WeekDayInAMonthRule{
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
	return parse.AddDate(time.Now().Year(), 0, 0)
}
