package venues

import (
	"encoding/json"
	"fmt"

	"github.com/comediq-api/database"
)

const venuesTableName = "venues"
const venueRoomsTableName = "venue_rooms"

type venueDatabaseRow struct {
	ID            *int32 `json:"id,omitempty"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	City          string `json:"city"`
	Neighbourhood string `json:"neighbourhood"`
	Borough       string `json:"borough"`
	Type          string `json:"type"`
	Contact       string `json:"contact"`
}
type venueRoomDatabaseRow struct {
	ID      *int32 `json:"id,omitempty"`
	VenueID int32  `json:"venue_id"`
	Name    string `json:"name"`
}

func UpsertVenueRoom(venueRoom *VenueRoom) (int32, error) {
	venueID, err := UpsertVenue(&venueRoom.Venue)
	if err != nil {
		return 0, err
	}

	venueRoomRow := venueRoomDatabaseRow{
		VenueID: venueID,
		Name:    venueRoom.Name,
	}

	data, _, err := database.Client.From(venueRoomsTableName).
		Upsert(venueRoomRow, "venue_id,name", "representation", "").
		Execute()
	if err != nil {
		return 0, err
	}

	return unmarshalID(data)
}

func UpsertVenue(venue *Venue) (int32, error) {
	venueRow := venueDatabaseRow{
		Name:          venue.Name,
		Address:       venue.Address,
		City:          venue.City,
		Neighbourhood: venue.Neighbourhood,
		Borough:       venue.Borough,
		Type:          venue.Type,
		Contact:       venue.Contact,
	}

	data, _, err := database.Client.From(venuesTableName).
		Upsert(venueRow, "name,address,city,neighbourhood,borough,type,contact", "representation", "").
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
