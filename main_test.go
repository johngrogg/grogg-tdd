package main_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weconnect/grogg-tdd"
	"github.com/weconnect/grogg-tdd/schema"
)

type invalidValidator struct{}

func (invalidValidator) Validate(str string) error {
	return errors.New("invalid")
}

type successValidator struct{}

func (successValidator) Validate(str string) error {
	return nil
}

var _ = Describe("Main", func() {
	Describe("#RegisterSchema", func() {
		var (
			validator schema.Validator

			schema *Schema
			err    error
		)

		BeforeEach(func() {
			validator = successValidator{}
		})

		JustBeforeEach(func() {
			schema, err = RegisterSchema("{}", validator)
		})

		It("should work", func() {
			Ω(err).ShouldNot(HaveOccurred())
			Ω(schema.JSONSchema).Should(Equal("{}"))
			Ω(schema.ID).ShouldNot(BeNil())
		})

		Context("when the schema string is invalid", func() {

			BeforeEach(func() {
				validator = invalidValidator{}
			})

			It("should return an error", func() {
				Ω(err).Should(HaveOccurred())
			})
		})
	})
})
