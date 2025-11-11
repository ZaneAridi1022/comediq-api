package past

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/comediq-api/database"
	"github.com/comediq-api/venues"
)

const tableName = "past_mic_events"

type databaseRow struct {
	ID               *int32    `json:"id,omitempty"`
	VenueRoomID      int32     `json:"venue_room_id"`
	Name             string    `json:"name"`
	Host             *string   `json:"host"`
	AudienceCost     string    `json:"audience_cost"`
	ComedianCost     string    `json:"comedian_cost"`
	StageTime        *string   `json:"stage_time"`
	ComedianLineup   []string  `json:"comedian_lineup"`
	StartDateAndTime time.Time `json:"start_date_and_time"`
	EndDateAndTime   time.Time `json:"end_date_and_time"`
}

func create(mic *Mic) (int32, error) {
	if mic.ID != nil {
		return 0, fmt.Errorf("mic already has an ID (does it exist already?)")
	}

	venueRoomID, err := venues.UpsertVenueRoom(&mic.VenueRoom)
	if err != nil {
		return 0, err
	}

	dbRow := databaseRow{
		VenueRoomID:      venueRoomID,
		Name:             mic.Name,
		Host:             mic.Host,
		AudienceCost:     mic.AudienceCost,
		ComedianCost:     mic.ComedianCost,
		StageTime:        mic.StageTime,
		ComedianLineup:   mic.ComedianLineup,
		StartDateAndTime: mic.StartDateAndTime,
		EndDateAndTime:   mic.EndDateAndTime,
	}
	data, _, err := database.Client.From(tableName).
		Insert(dbRow, false, "", "representation", "").
		Execute()
	if err != nil {
		return 0, err
	}

	return unmarshalID(data)
}

func unmarshalID(data []byte) (int32, error) {
	var result []struct {
		ID int32 `json:"id"`
	}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, fmt.Errorf("database operation returned nothing")
	}
	return result[0].ID, nil
}
