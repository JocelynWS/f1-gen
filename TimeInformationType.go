package ies

import "github.com/lvdund/ngap/aper"

const (
	TimeInformationTypeLocalclock aper.Enumerated = 0
)

type TimeInformationType struct {
	Value aper.Enumerated
}

func (ie *TimeInformationType) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false)
	return
}

func (ie *TimeInformationType) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false)
	ie.Value = aper.Enumerated(v)
	return
}
