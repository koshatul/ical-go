package ical_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIcal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "iCal Suite")
}
