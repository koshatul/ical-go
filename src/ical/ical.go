package ical

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type ICal struct {
	events []*VEvent
}

type itemType uint32

const (
	itemNone itemType = iota
	itemEvent
	itemAlarm
)

func ParseCalendar(r io.Reader) (*ICal, error) {
	ical := &ICal{}

	event := &VEvent{}
	scanner := bufio.NewScanner(r)

	current := itemNone
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.SplitN(line, ":", 2)
		lineSplit[0] = strings.ToUpper(lineSplit[0])
		switch lineSplit[0] {
		case "BEGIN":
			// Starting item
			switch lineSplit[1] {
			case "VEVENT":
				event = &VEvent{}
				current = itemEvent
			}
		case "END":
			switch lineSplit[1] {
			case "VEVENT":
				ical.AddEvents(event)
			}
		case "DESCRIPTION", "DTSTART", "DTSTAMP", "DTEND", "DURATION", "UID":
			switch current {
			case itemEvent:
				event.AddFields(VEventField{
					name:  lineSplit[0],
					value: lineSplit[1],
				})
			default:
				return nil, fmt.Errorf("Unexpected line: %s", line)
			}
		}

		// ical.ParseLine(scanner.Text())
	}
	err := scanner.Err()
	return ical, err
}

func (v *ICal) AddEvents(events ...*VEvent) {
	v.events = append(v.events, events...)
}

func (v *ICal) GetEvents() []*VEvent {
	return v.events
}

func (v *ICal) ParseLine(line string) {

}
