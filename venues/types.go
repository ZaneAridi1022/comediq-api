package venues

type Venue struct {
	Name          string  `json:"name"          validate:"required,min=1"`
	Address       string  `json:"address"       validate:"required,min=1"`
	City          string  `json:"city"          validate:"required,min=1"`
	Neighbourhood *string `json:"neighbourhood"`
	Borough       *string `json:"borough"`
	Type          *string `json:"type"`
	Contact       *string `json:"contact"`
}

type VenueRoom struct {
	Venue Venue   `json:"venue" validate:"required"`
	Name  *string `json:"name"`
}
