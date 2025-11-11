package scheduled

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/comediq-api/database"
	"github.com/comediq-api/venues"
)

const tableName = "scheduled_mic_events"

type databaseRow struct {
	ID                 *int32    `json:"id,omitempty"`
	DefinitionID       *int32    `json:"definition_id,omitempty"`
	VenueRoomID        int32     `json:"venue_room_id"`
	Name               string    `json:"name"`
	SignupInstructions *string   `json:"signup_instructions"`
	Notes              *string   `json:"notes"`
	AudienceCost       string    `json:"audience_cost"`
	ComedianCost       string    `json:"comedian_cost"`
	StageTime          *string   `json:"stage_time"`
	Host               *string   `json:"host"`
	Instagram          *string   `json:"instagram"`
	SMS                *string   `json:"sms"`
	ComedianLineup     []string  `json:"comedian_lineup"`
	StartDateAndTime   time.Time `json:"start_date_and_time"`
	EndDateAndTime     time.Time `json:"end_date_and_time"`
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
		DefinitionID:       mic.DefinitionID,
		VenueRoomID:        venueRoomID,
		Name:               mic.Name,
		SignupInstructions: mic.SignupInstructions,
		Notes:              mic.Notes,
		AudienceCost:       mic.AudienceCost,
		ComedianCost:       mic.ComedianCost,
		StageTime:          mic.StageTime,
		Host:               mic.Host,
		Instagram:          mic.Instagram,
		SMS:                mic.SMS,
		ComedianLineup:     mic.ComedianLineup,
		StartDateAndTime:   mic.StartDateAndTime,
		EndDateAndTime:     mic.EndDateAndTime,
	}
	data, _, err := database.Client.From(tableName).
		Insert(dbRow, false, "", "representation", "").
		Execute()
	if err != nil {
		return 0, err
	}

	return unmarshalID(data)
}

func deleteByDefinitionID(definitionID int32) error {
	_, _, err := database.Client.From(tableName).Delete("", "").
		Match(map[string]string{
			"definition_id": strconv.Itoa(int(definitionID)),
		}).
		Execute()
	if err != nil {
		fmt.Println(err)
	}
	return err
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
