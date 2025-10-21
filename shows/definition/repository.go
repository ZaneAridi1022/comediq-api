package definition

import (
	"encoding/json"
	"fmt"

	"github.com/comediq-api/database"
	"github.com/comediq-api/venues"
)

const tableName = "show_definitions"

type databaseRow struct {
	ID                 *int32  `json:"id,omitempty"`
	VenueRoomID        int32   `json:"venue_room_id"`
	Name               string  `json:"name"`
	SignupInstructions string  `json:"signup_instructions"`
	Notes              *string `json:"notes"`
	AudienceCost       *string `json:"audience_cost"`
	ComedianCost       *string `json:"comedian_cost"`
	StageTime          *string `json:"stage_time"`
	Host               *string `json:"host"`
	SMS                *string `json:"sms"`
	Verified           string  `json:"verified"`
	// OccurrenceRules     [7]OneMonthDayRule `json:"occurrence_rules"`
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
		VenueRoomID:  venueRoomID,
		Name:         show.Name,
		AudienceCost: show.AudienceCost,
		ComedianCost: show.ComedianCost,
		StageTime:    show.StageTime,
		Host:         show.Host,
		SMS:          show.SMS,
		Verified:     show.Verified,
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
