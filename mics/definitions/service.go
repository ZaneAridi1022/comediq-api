package definitions

import (
	"errors"
	"time"

	"github.com/comediq-api/mics/events/scheduled"
	"github.com/comediq-api/utils"
)

// Create Creates a mic and returns its ID.
func Create(mic *Mic) (int32, error) {
	return create(mic)
}

// AutoSchedule creates all scheduled.Mic from now based off occurrence rules.
// Deletes all upcoming mics associated with the provided mic definition and schedules new upcoming mics.
func AutoSchedule(micDefinition *Mic) error {
	if micDefinition.ID == nil {
		return errors.New("mic definition must have a registered id")
	}

	err := scheduled.DeleteByDefinitionId(*micDefinition.ID)
	if err != nil {
		return err
	}

	dateNow := time.Now()
	for _, rule := range micDefinition.OccurrenceRules {
		for _, date := range rule.GetDatesFrom(dateNow) {
			_, err := scheduled.Create(&scheduled.Mic{
				DefinitionID:       micDefinition.ID,
				VenueRoom:          micDefinition.VenueRoom,
				Name:               micDefinition.Name,
				SignupInstructions: micDefinition.SignupInstructions,
				Notes:              micDefinition.Notes,
				AudienceCost:       micDefinition.AudienceCost,
				ComedianCost:       micDefinition.ComedianCost,
				StageTime:          micDefinition.StageTime,
				Host:               micDefinition.Host,
				Instagram:          micDefinition.Instagram,
				SMS:                micDefinition.SMS,
				ComedianLineup:     []string{},
				StartDateAndTime:   utils.CombineDateAndTime(date, rule.StartTime),
				EndDateAndTime:     utils.CombineDateAndTime(date, rule.EndTime),
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
