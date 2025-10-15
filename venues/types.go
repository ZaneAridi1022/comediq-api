package venues

type VenueRoom struct {
	ID            *int32 `json:"id"`
	Address       string `json:"address"`
	Name          string `json:"name"`
	City          string `json:"city"`
	Neighbourhood string `json:"neighbourhood"`
	Borough       string `json:"borough"`
	Type          string `json:"type"`
	Contact       string `json:"contact"`
	RoomName      string `json:"room"`
}
