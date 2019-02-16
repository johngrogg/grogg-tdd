package schema

import "github.com/google/uuid"

// Schema holds both the JSON Scehma and a UUID from the database
type Schema struct {
	JSONSchema string
	ID         uuid.UUID
}

// DataRepository is basic data persistence abstraction
type DataRepository interface {
	Save(string, ORM) (*Schema, error)
}

// Repository validates json schema strings
type Repository struct{}

// Save persists the JSON Schema into our database
func (Repository) Save(json string, orm ORM) (*Schema, error) {
	schema := &Schema{
		JSONSchema: json,
		ID:         uuid.New(),
	}
	return schema, nil
}

// ORM works with the database
type ORM interface {
	Insert(string) (uuid.UUID, error)
}
