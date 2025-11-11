package definitions

import (
	"encoding/json"
	"fmt"

	"github.com/comediq-api/database"
	"github.com/comediq-api/venues"
)

const tableName = "mic_definitions"

type databaseRow struct {
	ID                 *int32                `json:"id,omitempty"`
	VenueRoomID        int32                 `json:"venue_room_id"`
	Name               string                `json:"name"`
	SignupInstructions *string               `json:"signup_instructions"`
	Notes              *string               `json:"notes"`
	AudienceCost       string                `json:"audience_cost"`
	ComedianCost       string                `json:"comedian_cost"`
	StageTime          *string               `json:"stage_time"`
	Host               *string               `json:"host"`
	Instagram          *string               `json:"instagram"`
	SMS                *string               `json:"sms"`
	Verified           *string               `json:"verified"`
	Active             bool                  `json:"active"`
	OccurrenceRules    []WeekDayInAMonthRule `json:"occurrence_rules"`
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
		Verified:           mic.Verified,
		Active:             mic.Active,
		OccurrenceRules:    mic.OccurrenceRules,
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
