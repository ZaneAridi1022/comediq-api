package helpers

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/comediq-api/database"
	"github.com/comediq-api/types"
	"github.com/google/uuid"
)

func ReadCSV() {
	filePath := "micData" + "/masterCopy.csv"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	micSet := make(map[string]uuid.UUID)
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.TrimLeadingSpace = true
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records:", err)
		return
	}

	for j, record := range records {
		if j == 0 {
			continue // Skip header row
		}
		// fmt.Printf("Record %d: %v\n", j, record)
		var verified string = "Unverified"
		verifiedExpanded := strings.Split(record[15], " ")
		if len(verifiedExpanded) > 1 {
			verified = strings.TrimSpace(verifiedExpanded[1])
		}
		active := (verified != "Unverified")
		newYork := "New York"
		uuid := uuid.New()
		openMic := types.HistoricalMic{
			UniqueIdentifier:   uuid,
			OpenMic:            ptr(record[1]),
			Day:                ptr(record[2]),
			StartTime:          ptr(record[3]),
			LatestEndTime:      ptr(record[4]),
			VenueName:          ptr(record[5]),
			Borough:            ptr(record[6]),
			Neighborhood:       ptr(record[7]),
			Location:           ptr(record[8]),
			VenueType:          ptr(record[9]),
			Cost:               ptr(record[10]),
			StageTime:          ptr(record[11]),
			SignUpInstructions: ptr(record[12]),
			HostsOrganizers:    ptr(record[13]),
			ChangesUpdates:     ptr(record[14]),
			LastVerified:       &verified,
			Active:             &active,
			SMS:                ptr(strings.TrimSpace(record[16])),
			City:               &newYork,
		}

		fmt.Printf("Mic details %d: %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v\n", j,
			openMic.UniqueIdentifier, *openMic.OpenMic, *openMic.Day,
			*openMic.StartTime, *openMic.LatestEndTime,
			*openMic.VenueName, *openMic.Borough,
			*openMic.Neighborhood, *openMic.Location,
			*openMic.VenueType, *openMic.Cost,
			*openMic.StageTime, *openMic.SignUpInstructions,
			*openMic.HostsOrganizers, *openMic.ChangesUpdates,
			*openMic.LastVerified)
		database.UpsertData(openMic, "open_mics_historical")
		micSet[record[0]] = openMic.UniqueIdentifier
	}

	fmt.Println("Mic Set:", micSet)

	time.Sleep(5 * time.Second)
	filePath2 := "micData" + "/profile_open_mics_rows_updated_retry.csv"
	file2, err := os.Open(filePath2)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file2.Close()

	reader2 := csv.NewReader(file2)
	reader2.FieldsPerRecord = -1
	reader2.TrimLeadingSpace = true
	records2, err := reader2.ReadAll()
	if err != nil {
		fmt.Println("Error reading records:", err)
		return
	}

	for i, record2 := range records2 {
		if i == 0 {
			continue // Skip header row
		}

		profileMic, err := types.NewProfileOpenMic(strings.TrimSpace(record2[5]),
			strings.TrimSpace(record2[0]),
			strings.TrimSpace(record2[1]),
			micSet[strings.TrimSpace(record2[2])],
			strings.TrimSpace(record2[3]),
			ptr(record2[4]),
			strings.TrimSpace(record2[6]),
		)

		if err != nil {
			fmt.Println("Error creating profile mic for record:", err)
			return
		}
		fmt.Printf("Profile Mic %d: %v, %v, %v, %v, %v, %v, %v\n", i,
			profileMic.ID, profileMic.ProfileID, profileMic.OpenMicID,
			profileMic.ScheduleType, profileMic.Notes,
			profileMic.CreatedAt, profileMic.LastModified)
		database.UpsertData(profileMic, "profile_open_mics")
	}
}

func ptr(s string) *string {
	trimmed := strings.TrimSpace(s)
	return &trimmed
}
