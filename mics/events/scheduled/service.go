package scheduled

// Create Creates a mic and returns its ID.
func Create(mic *Mic) (int32, error) {
	return create(mic)
}

func DeleteByDefinitionId(definitionId int32) error {
	return deleteByDefinitionID(definitionId)
}
