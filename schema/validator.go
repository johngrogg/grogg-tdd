package schema

// Validator a JSON Schema validation interface
type Validator interface {
	Validate(string) error
}
