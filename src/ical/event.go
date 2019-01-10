package ical

import "strings"

type VEvent struct {
	fields []VEventField
}

func NewVEvent() *VEvent {
	return &VEvent{}
}

func (v *VEvent) AddFields(fields ...VEventField) {
	v.fields = append(v.fields, fields...)
}

func (v *VEvent) getString(name string) string {
	for _, field := range v.fields {
		if strings.EqualFold(field.Name(), name) {
			return field.Value()
		}
	}
	return ""
}

func (v *VEvent) Description() string {
	return v.getString("DESCRIPTION")
}
