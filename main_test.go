package main_test

import (
	"errors"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/johngrogg/grogg-tdd"
	"github.com/johngrogg/grogg-tdd/schema"
)

type invalidValidator struct{}

func (invalidValidator) Validate(str string) error {
	return errors.New("invalid")
}

type successValidator struct{}

func (successValidator) Validate(str string) error {
	return nil
}

type successRepository struct{}

func (successRepository) Save(schemaStr string) (*schema.Schema, error) {
	return &schema.Schema{
		JSONSchema: schemaStr,
		ID:         uuid.New(),
	}, nil
}

type failureRepository struct{}

func (failureRepository) Save(str string) (*schema.Schema, error) {
	return nil, errors.New("this means nothing asdjfio")
}

var _ = Describe("Main", func() {
	Describe("#RegisterSchema", func() {
		var (
			validator  schema.Validator
			repository schema.DataRepository

			schema *schema.Schema
			err    error
		)

		BeforeEach(func() {
			validator = successValidator{}
			repository = successRepository{}
		})

		JustBeforeEach(func() {
			schema, err = RegisterSchema("{asdfgh}", validator, repository)
		})

		Context("when there are no errors", func() {
			It("should return the new schema", func() {
				Ω(err).ShouldNot(HaveOccurred())
				Ω(schema.JSONSchema).Should(Equal("{asdfgh}"))
				Ω(schema.ID).ShouldNot(BeNil())
			})
		})

		Context("when the schema string is invalid", func() {
			BeforeEach(func() {
				validator = invalidValidator{}
			})

			It("should return an error", func() {
				Ω(err).Should(MatchError("invalid"))
			})
		})

		Context("when the repository fails to save", func() {
			BeforeEach(func() {
				repository = failureRepository{}
			})

			It("should return the error", func() {
				Ω(err).Should(MatchError("this means nothing asdjfio"))
			})
		})
	})
})
