package ical_test

import (
	"github.com/koshatul/ical-go/src/ical"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ICal Library", func() {
	It("Parse", func() {
		f, err := afs.Open("vcalendar_1.0.ics")
		Expect(err).NotTo(HaveOccurred())

		ical, err := ical.ParseCalendar(f)
		Expect(err).NotTo(HaveOccurred())

		events := ical.GetEvents()
		event := events[0]
		Expect(event.Description()).To(Equal("Steve and John to review newest proposal material"))
		// for _, event := range events {
		// 	log.Printf("Event: %#v", event)
		// }
	})

})
