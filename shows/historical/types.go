package historical

import (
	"time"

	"github.com/comediq-api/venues"
)

type Show struct {
	ID             *int32           `json:"id"`
	VenueRoom      venues.VenueRoom `json:"venue_room"`
	Name           string           `json:"name"`
	Host           *string          `json:"host"`
	AudienceCost   *string          `json:"audience_cost"`
	ComedianCost   *string          `json:"comedian_cost"`
	StageTime      *string          `json:"stage_time"`
	ComedianLineup []string         `json:"comedian_lineup"`
	StartTime      time.Time        `json:"start_time"`
	EndTime        time.Time        `json:"end_time"`
}
