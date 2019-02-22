package schema

// Register validates and stores a new JSON Schema
func Register(schemaStr string, validator Validator, repo DataRepository) (*Schema, error) {

	if err := validator.Validate(schemaStr); err != nil {
		return nil, err
	}

	schema, err := repo.Save(schemaStr, nil)
	if err != nil {
		return nil, err
	}

	return schema, nil
}
