package historical

import (
	"time"

	"github.com/comediq-api/venues"
)

type Show struct {
	ID               *int32           `json:"id"`
	VenueRoom        venues.VenueRoom `json:"venue_room"`
	Name             string           `json:"name"`
	Host             *string          `json:"host"`
	AudienceCost     *string          `json:"audience_cost"`
	ComedianCost     *string          `json:"comedian_cost"`
	StageTime        *string          `json:"stage_time"`
	ComedianLineup   []string         `json:"comedian_lineup"`
	StartDateAndTime time.Time        `json:"start_date_and_time"`
	EndDateAndTime   time.Time        `json:"end_date_and_time"`
	Open             bool             `json:"open"`
}
