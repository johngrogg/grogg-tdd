package main

import (
	"fmt"
	"os"

	"github.com/weconnect/grogg-tdd/schema"
	"github.com/weconnect/grogg-tdd/schema/validators"
)

func main() {
	inputSchemaString := os.Args[1]
	registeredSchema, err := schema.Register(inputSchemaString, validators.JSONValidator{}, schema.Repository{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Saved as: " + registeredSchema.ID.String())
	fmt.Println("Schema that was saved: " + registeredSchema.JSONSchema)
}
