package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weconnect/grogg-tdd"
)

var _ = Describe("Main", func() {
	Describe("#RegisterSchema", func() {
		It("should work", func() {
			schema, err := RegisterSchema("{}")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(schema.JSONSchema).Should(Equal("{}"))
			Ω(schema.ID).ShouldNot(BeNil())
		})
	})
})
