package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGroggTdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GroggTdd Suite")
}
