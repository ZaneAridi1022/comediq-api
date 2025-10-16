package venues

type Venue struct {
	Name          string `json:"name"`
	Address       string `json:"address"`
	City          string `json:"city"`
	Neighbourhood string `json:"neighbourhood"`
	Borough       string `json:"borough"`
	Type          string `json:"type"`
	Contact       string `json:"contact"`
}

type VenueRoom struct {
	Venue Venue  `json:"venue"`
	Name  string `json:"name"`
}
