package ies

import "github.com/lvdund/ngap/aper"

const (
	EventTypeOndemand aper.Enumerated = 0
	EventTypePeriodic aper.Enumerated = 1
	EventTypeStop     aper.Enumerated = 2
)

type EventType struct {
	Value aper.Enumerated
}

func (ie *EventType) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
	return
}

func (ie *EventType) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	ie.Value = aper.Enumerated(v)
	return
}
