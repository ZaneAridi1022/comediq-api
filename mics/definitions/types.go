package definitions

import (
	"time"

	"github.com/comediq-api/utils"
	"github.com/comediq-api/venues"
)

type Mic struct {
	ID                 *int32                `json:"id"`
	VenueRoom          venues.VenueRoom      `json:"venue_room"          validation:"required"`
	Name               string                `json:"name"                validation:"required,min=1"`
	SignupInstructions *string               `json:"signup_instructions" validation:"required,min=1"`
	Notes              *string               `json:"notes"`
	AudienceCost       string                `json:"audience_cost"       validation:"required,min=1"`
	ComedianCost       string                `json:"comedian_cost"       validation:"required,min=1"`
	StageTime          *string               `json:"stage_time"`
	Host               *string               `json:"host"`
	Instagram          *string               `json:"instagram"`
	SMS                *string               `json:"sms"`
	Verified           *string               `json:"verified"`
	Active             bool                  `json:"active"`
	OccurrenceRules    []WeekDayInAMonthRule `json:"occurrence_rules" validation:"dive"`
}

type WeekDayInAMonthRule struct {
	WeeksInAdvance     int32     `json:"weeks_in_advice"`
	WeeksOfMonthPolicy [5]bool   `json:"weeks_of_month_policy"`
	WeekDay            int32     `json:"week_day"`
	StartTime          time.Time `json:"start_time"`
	EndTime            time.Time `json:"end_time"`
}

func (rule *WeekDayInAMonthRule) GetDatesFrom(fromDate time.Time) []time.Time {
	oneDayDuration := time.Hour * 24
	var dates []time.Time
	for i := 0; i < int(rule.WeeksInAdvance)*7; i++ {
		occurringDate := fromDate.Add(time.Duration(i) * oneDayDuration)
		validWeekday := int32(occurringDate.Weekday()) == rule.WeekDay
		validWeekOfMonth := rule.WeeksOfMonthPolicy[utils.GetWeekOfMonth(occurringDate)]
		if validWeekday && validWeekOfMonth {
			dates = append(dates, occurringDate)
		}
	}
	return dates
}
