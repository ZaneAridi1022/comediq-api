package definition

import (
	"strings"
	"time"

	"github.com/comediq-api/shows/upcoming"
	"github.com/comediq-api/utils"
)

// Create Creates a show and returns its ID.
func Create(show *Show) (int32, error) {
	return create(show)
}

// AutoSchedule creates all upcoming.Show from now based off occurrence rules
func AutoSchedule(show *Show) error {
	dateNow := time.Now()
	for _, rule := range show.OccurrenceRules {
		for _, date := range rule.GetDatesFrom(dateNow) {
			_, err := upcoming.Create(&upcoming.Show{
				VenueRoom:          show.VenueRoom,
				Name:               show.Name,
				SignupInstructions: show.SignupInstructions,
				Notes:              show.Notes,
				AudienceCost:       show.AudienceCost,
				ComedianCost:       show.ComedianCost,
				StageTime:          show.StageTime,
				Host:               show.Host,
				SMS:                show.SMS,
				Verified:           show.Verified,
				ComedianLineup:     []string{},
				StartDateAndTime:   utils.CombineDateAndTime(date, rule.StartTime),
				EndDateAndTime:     utils.CombineDateAndTime(date, rule.EndTime),
				Open:               show.Open,
			})
			// Skip if duplicates are made. (Questionable error handling)
			if err != nil && !strings.Contains(err.Error(), "violates unique constraint") {
				return err
			}
		}
	}
	return nil
}
