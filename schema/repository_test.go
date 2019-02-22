package schema_test

import (
	"errors"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weconnect/grogg-tdd/schema"
)

type successfulMockInsert struct{}

func (successfulMockInsert) Insert(str string) (uuid.UUID, error) {
	return uuid.New(), nil
}

type invalidMockInsert struct{}

func (invalidMockInsert) Insert(str string) (uuid.UUID, error) {
	return uuid.New(), errors.New("db insert fail")
}

var _ = Describe("Repository", func() {
	var (
		orm ORM

		repository Repository
	)

	BeforeEach(func() {
		orm = successfulMockInsert{}
		repository = Repository{}
	})

	Describe("#Save", func() {
		var (
			schema *Schema
			err    error
		)

		JustBeforeEach(func() {
			schema, err = repository.Save("{asdfgh}", orm)
		})

		Context("when there are no errors", func() {
			It("should return the new schema", func() {
				Ω(err).ShouldNot(HaveOccurred())
				Ω(schema.JSONSchema).Should(Equal("{asdfgh}"))
				Ω(schema.ID).ShouldNot(BeNil())
			})
		})

		Context("when the ORM fails to save", func() {
			BeforeEach(func() {
				orm = invalidMockInsert{}
			})

			It("should return the error", func() {
				Ω(err).Should(MatchError("db insert fail"))
			})
		})
	})
})
