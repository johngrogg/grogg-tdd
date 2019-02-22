package main

import (
	"fmt"
	"os"

	"github.com/johngrogg/grogg-tdd/schema"
	"github.com/johngrogg/grogg-tdd/schema/validators"
)

func main() {
	inputSchemaString := os.Args[1]
	schema, err := schema.Register(inputSchemaString, validators.JSONSchema{}, schema.Repository{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Saved as: " + schema.ID.String())
	fmt.Println("Schema that was saved: " + schema.JSONSchema)
}
