package historical

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/comediq-api/database"
	"github.com/comediq-api/venues"
)

const tableName = "historical_shows"

type databaseRow struct {
	ID               *int32    `json:"id,omitempty"`
	VenueRoomID      int32     `json:"venue_room_id"`
	Name             string    `json:"name"`
	Host             *string   `json:"host"`
	AudienceCost     *string   `json:"audience_cost"`
	ComedianCost     *string   `json:"comedian_cost"`
	StageTime        *string   `json:"stage_time"`
	ComedianLineup   []string  `json:"comedian_lineup"`
	StartDateAndTime time.Time `json:"start_date_and_time"`
	EndDateAndTime   time.Time `json:"end_date_and_time"`
	Open             bool      `json:"open"`
}

func create(show *Show) (int32, error) {
	if show.ID != nil {
		return 0, fmt.Errorf("show already has an ID (does it exist already?)")
	}

	venueRoomID, err := venues.UpsertVenueRoom(&show.VenueRoom)
	if err != nil {
		return 0, err
	}

	dbRow := databaseRow{
		VenueRoomID:      venueRoomID,
		Name:             show.Name,
		Host:             show.Host,
		AudienceCost:     show.AudienceCost,
		ComedianCost:     show.ComedianCost,
		StageTime:        show.StageTime,
		ComedianLineup:   show.ComedianLineup,
		StartDateAndTime: show.StartDateAndTime,
		EndDateAndTime:   show.EndDateAndTime,
		Open:             show.Open,
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
