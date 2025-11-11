package past

// Create Creates a mic and returns its ID.
func Create(mic *Mic) (int32, error) {
	return create(mic)
}
