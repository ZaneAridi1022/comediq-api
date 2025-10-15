package types

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ScheduleType string

const (
	Upcoming  ScheduleType = "upcomming"
	Completed ScheduleType = "completed"
	Cancelled ScheduleType = "cancelled"
)

type ProfileOpenMic struct {
	ID           uuid.UUID    `json:"id"`
	CreatedAt    time.Time    `json:"created_at"`
	ProfileID    uuid.UUID    `json:"profile_id"`
	OpenMicID    uuid.UUID    `json:"open_mic_id"`
	ScheduleType ScheduleType `json:"schedule_type"`
	Notes        *string      `json:"notes"`
	LastModified time.Time    `json:"last_modified"`
}

func NewProfileOpenMic(
	id string, createdAtStr string, profileID string,
	openMicID uuid.UUID, scheduleType string, notes *string,
	lastModifiedStr string,
) (ProfileOpenMic, error) {
	const pgTimestampLayout = "2006-01-02 15:04:05.999999-07"
	createdAt, err := time.Parse(pgTimestampLayout, createdAtStr)
	if err != nil {
		return ProfileOpenMic{}, fmt.Errorf("invalid created_at: %w", err)
	}

	var lastModified time.Time
	if lastModifiedStr == "" {
		lastModified = time.Now()
	} else {
		lastModified, err = time.Parse(pgTimestampLayout, lastModifiedStr)
		if err != nil {
			return ProfileOpenMic{}, fmt.Errorf("invalid last_modified: %w", err)
		}
	}
	ID, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("Invalid UUID for ProfileOpenMic ID:", err)
		return ProfileOpenMic{}, err
	}

	profileIDParsed, err := uuid.Parse(profileID)
	if err != nil {
		fmt.Println("Invalid UUID for ProfileOpenMic ProfileID:", err)
		return ProfileOpenMic{}, err
	}

	return ProfileOpenMic{
		ID:           ID,
		CreatedAt:    createdAt,
		ProfileID:    profileIDParsed,
		OpenMicID:    openMicID,
		ScheduleType: ScheduleType(scheduleType),
		Notes:        notes,
		LastModified: lastModified,
	}, nil
}
