package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/weconnect/grogg-tdd/schema"
)

// Schema holds both the JSON Scehma and a UUID from the database
type Schema struct {
	JSONSchema string
	ID         uuid.UUID
}

// DataRepository is basic data persistence abstraction
type DataRepository interface {
	SaveSchema(string) (*Schema, error)
}

func main() {
	inputSchemaString := os.Args[1]
	schema, err := RegisterSchema(inputSchemaString, schema.JSONValidator{}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Saved as: " + schema.ID.String())
	fmt.Println("Schema that was saved: " + schema.JSONSchema)
}

// RegisterSchema validates and stores a new JSON Schema
func RegisterSchema(schemaStr string, validator schema.Validator, repo DataRepository) (*Schema, error) {

	if err := validator.Validate(schemaStr); err != nil {
		return nil, err
	}

	// TODO: implement me!
	schema := &Schema{
		JSONSchema: schemaStr,
	}

	schema.ID = uuid.New()

	return schema, nil
}
