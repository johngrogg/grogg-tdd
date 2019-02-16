package main

import (
	"fmt"
	"os"

	"github.com/johngrogg/grogg-tdd/schema"
)

func main() {
	inputSchemaString := os.Args[1]
	schema, err := RegisterSchema(inputSchemaString, schema.JSONValidator{}, schema.Repository{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Saved as: " + schema.ID.String())
	fmt.Println("Schema that was saved: " + schema.JSONSchema)
}

// RegisterSchema validates and stores a new JSON Schema
func RegisterSchema(schemaStr string, validator schema.Validator, repo schema.DataRepository) (*schema.Schema, error) {

	if err := validator.Validate(schemaStr); err != nil {
		return nil, err
	}

	schema, err := repo.Save(schemaStr)
	if err != nil {
		return nil, err
	}

	return schema, nil
}
