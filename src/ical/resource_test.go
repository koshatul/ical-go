package ical_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/spf13/afero"
)

var (
	afs afero.Fs
)

var _ = BeforeSuite(func() {
	afs = afero.NewMemMapFs()
	afero.WriteFile(afs, "vcalendar_1.0.ics", []byte(vCalendar10), 0755)
})

const vCalendar10 = `BEGIN:VCALENDAR
VERSION:1.0
BEGIN:VEVENT
CATEGORIES:MEETING
STATUS:TENTATIVE
DTSTART:19960401T033000Z
DTEND:19960401T043000Z
SUMMARY:Your Proposal Review
DESCRIPTION:Steve and John to review newest proposal material
CLASS:PRIVATE
END:VEVENT
END:VCALENDAR`
