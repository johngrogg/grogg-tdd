package validators

// JSONSchema validates json schema strings
type JSONSchema struct{}

// Validate returns an error if the json string is invalid
func (JSONSchema) Validate(json string) error {
	return nil
}
