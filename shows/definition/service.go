package definition

// Create Creates a show and returns its ID.
func Create(show *Show) (int32, error) {
	return create(show)
}
