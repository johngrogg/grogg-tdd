package schema_test

import (
	"errors"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/johngrogg/grogg-tdd/schema"
)

type invalidMockValidator struct{}

func (invalidMockValidator) Validate(str string) error {
	return errors.New("invalid")
}

type successMockValidator struct{}

func (successMockValidator) Validate(str string) error {
	return nil
}

type successMockRepository struct{}

func (successMockRepository) Save(schemaStr string, orm ORM) (*Schema, error) {
	return &Schema{
		JSONSchema: schemaStr,
		ID:         uuid.New(),
	}, nil
}

type failureMockRepository struct{}

func (failureMockRepository) Save(str string, orm ORM) (*Schema, error) {
	return nil, errors.New("this means nothing asdjfio")
}

var _ = Describe("Register", func() {
	var (
		validator  Validator
		repository DataRepository

		schema *Schema
		err    error
	)

	BeforeEach(func() {
		validator = successMockValidator{}
		repository = successMockRepository{}
	})

	JustBeforeEach(func() {
		schema, err = Register("{asdfgh}", validator, repository)
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
			validator = invalidMockValidator{}
		})

		It("should return an error", func() {
			Ω(err).Should(MatchError("invalid"))
		})
	})

	Context("when the repository fails to save", func() {
		BeforeEach(func() {
			repository = failureMockRepository{}
		})

		It("should return the error", func() {
			Ω(err).Should(MatchError("this means nothing asdjfio"))
		})
	})
})
