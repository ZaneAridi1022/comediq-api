package types

import (
    "github.com/google/uuid"
)
type HistoricalMic struct {
	UniqueIdentifier   uuid.UUID  `json:"unique_identifier"`
	OpenMic            *string `json:"open_mic"`
	Day                *string `json:"day"`
	StartTime          *string `json:"start_time"`
	LatestEndTime      *string `json:"latest_end_time"`
	VenueName          *string `json:"venue_name"`
	Borough            *string `json:"borough"`
	Neighborhood       *string `json:"neighborhood"`
	Location           *string `json:"location"`
	VenueType          *string `json:"venue_type"`
	Cost               *string `json:"cost"`
	StageTime          *string `json:"stage_time"`
	SignUpInstructions *string `json:"sign_up_instructions"`
	HostsOrganizers    *string `json:"hosts_organizers"`
	ChangesUpdates     *string `json:"changes_updates"`
	LastVerified       *string `json:"last_verified"`
	OtherRules         *string `json:"other_rules"`
	Active             *bool   `json:"active"`
	SMS                *string `json:"sms_response"`
	City               *string `json:"city"`
}

func NewOpenMic() *HistoricalMic {
    return &HistoricalMic{
        UniqueIdentifier:   uuid.New(),
        OpenMic:            nil,
        Day:                nil,
        StartTime:          nil,
        LatestEndTime:      nil,
        VenueName:          nil,
        Borough:            nil,
        Neighborhood:       nil,
        Location:           nil,
        VenueType:          nil,
        Cost:               nil,
        StageTime:          nil,
        SignUpInstructions: nil,
        HostsOrganizers:    nil,
        ChangesUpdates:     nil,
        LastVerified:       nil,
        OtherRules:         nil,
        Active:             nil,
        SMS:                nil,
        City:               nil,
    }
}



