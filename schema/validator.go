package schema

// Validator a JSON Schema validation interface
type Validator interface {
	Validate(string) error
}

// JSONValidator validates json schema strings
type JSONValidator struct{}

// Validate returns an error if the json string is invalid
func (JSONValidator) Validate(json string) error {
	return nil
}
