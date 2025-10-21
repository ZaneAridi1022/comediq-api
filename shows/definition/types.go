package definition

import (
	"github.com/comediq-api/venues"
)

type Show struct {
	ID                 *int32           `json:"id"`
	VenueRoom          venues.VenueRoom `json:"venue_room"`
	Name               string           `json:"name"`
	SignupInstructions string           `json:"signup_instructions"`
	Notes              *string          `json:"notes"`
	AudienceCost       *string          `json:"audience_cost"`
	ComedianCost       *string          `json:"comedian_cost"`
	StageTime          *string          `json:"stage_time"`
	Host               *string          `json:"host"`
	SMS                *string          `json:"sms"`
	Verified           string           `json:"verified"`
	// OccurrenceRules     [7]OneMonthDayRule `json:"occurrence_rules"`
}
