package definition

import (
	"time"

	"github.com/comediq-api/utils"
	"github.com/comediq-api/venues"
)

type Show struct {
	ID                 *int32                `json:"id"`
	VenueRoom          venues.VenueRoom      `json:"venue_room"`
	Name               string                `json:"name"`
	SignupInstructions string                `json:"signup_instructions"`
	Notes              *string               `json:"notes"`
	AudienceCost       *string               `json:"audience_cost"`
	ComedianCost       *string               `json:"comedian_cost"`
	StageTime          *string               `json:"stage_time"`
	Host               *string               `json:"host"`
	SMS                *string               `json:"sms"`
	Verified           string                `json:"verified"`
	Open               bool                  `json:"open"`
	OccurrenceRules    []WeekDayInAMonthRule `json:"occurrence_rules"`
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
