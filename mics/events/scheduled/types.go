package scheduled

import (
	"time"

	"github.com/comediq-api/venues"
)

type Mic struct {
	ID                 *int32           `json:"id"`
	DefinitionID       *int32           `json:"definition_id"`
	VenueRoom          venues.VenueRoom `json:"venue_room" validation:"required"`
	Name               string           `json:"name" validation:"required,min=1"`
	SignupInstructions *string          `json:"signup_instructions"`
	Notes              *string          `json:"notes"`
	AudienceCost       string           `json:"audience_cost" validation:"required,min=1"`
	ComedianCost       string           `json:"comedian_cost" validation:"required,min=1"`
	StageTime          *string          `json:"stage_time"`
	Host               *string          `json:"host"`
	Instagram          *string          `json:"instagram"`
	SMS                *string          `json:"sms"`
	ComedianLineup     []string         `json:"comedian_lineup" validate:"required"`
	StartDateAndTime   time.Time        `json:"start_date_and_time"`
	EndDateAndTime     time.Time        `json:"end_date_and_time"`
}
