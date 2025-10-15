package venues

import (
	"encoding/json"
	"fmt"

	"github.com/comediq-api/database"
)

type venueDatabaseRow struct {
	ID            *int32 `json:"id"`
	Address       string `json:"address"`
	Name          string `json:"name"`
	City          string `json:"city"`
	Neighbourhood string `json:"neighbourhood"`
	Borough       string `json:"borough"`
	Type          string `json:"type"`
	Contact       string `json:"contact"`
}

type venueRoomDatabaseRow struct {
	ID       *int32 `json:"id"`
	VenueID  int32  `json:"venue_id"`
	RoomName string `json:"room_name"`
}

func UpsertVenueRoom(venueRoom *VenueRoom) (int32, error) {
	venueID, err := upsertVenue(venueRoom)
	if err != nil {
		return 0, err
	}

	venueRoomRow := venueRoomDatabaseRow{
		VenueID:  venueID,
		RoomName: venueRoom.RoomName,
	}

	data, _, err := database.Client.From("venue_rooms").
		Upsert(venueRoomRow, "venue_id,room_name", "id", "").
		Execute()
	if err != nil {
		return 0, err
	}

	return unmarshalID(data)
}

func upsertVenue(venueRoom *VenueRoom) (int32, error) {
	venueRow := venueDatabaseRow{
		Address:       venueRoom.Address,
		Name:          venueRoom.Name,
		City:          venueRoom.City,
		Neighbourhood: venueRoom.Neighbourhood,
		Borough:       venueRoom.Borough,
		Type:          venueRoom.Type,
		Contact:       venueRoom.Contact,
	}

	data, _, err := database.Client.From("venues").
		Upsert(venueRow, "address,room_name,city,neighbourhood,borough,type,contact", "id", "").
		Execute()
	if err != nil {
		return 0, err
	}

	return unmarshalID(data)
}

func unmarshalID(data []byte) (int32, error) {
	var result []struct {
		ID int32 `json:"id"`
	}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, fmt.Errorf("database operation returned nothing")
	}
	return result[0].ID, nil
}
