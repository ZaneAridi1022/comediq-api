package past

import (
	"time"

	"github.com/comediq-api/venues"
)

type Mic struct {
	ID               *int32           `json:"id"`
	VenueRoom        venues.VenueRoom `json:"venue_room"          validate:"required"`
	Name             string           `json:"name"                validate:"required,min=1"`
	Host             *string          `json:"host"`
	AudienceCost     string           `json:"audience_cost" validate:"required,min=1"`
	ComedianCost     string           `json:"comedian_cost" validate:"required,min=1"`
	StageTime        *string          `json:"stage_time"`
	ComedianLineup   []string         `json:"comedian_lineup" validate:"required,min=1,required"`
	StartDateAndTime time.Time        `json:"start_date_and_time" validate:"required"`
	EndDateAndTime   time.Time        `json:"end_date_and_time"   validate:"required"`
}
