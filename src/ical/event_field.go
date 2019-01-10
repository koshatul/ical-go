package ical

type VEventField struct {
	name  string
	value string
}

func (v *VEventField) Name() string {
	return v.name
}

func (v *VEventField) Value() string {
	return v.value
}
