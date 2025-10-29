package ies

import "github.com/lvdund/ngap/aper"

const (
	SSBPeriodicityMs5   aper.Enumerated = 0
	SSBPeriodicityMs10  aper.Enumerated = 1
	SSBPeriodicityMs20  aper.Enumerated = 2
	SSBPeriodicityMs40  aper.Enumerated = 3
	SSBPeriodicityMs80  aper.Enumerated = 4
	SSBPeriodicityMs160 aper.Enumerated = 5
)

type SSBPeriodicity struct {
	Value aper.Enumerated
}

func (ie *SSBPeriodicity) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, true)
	return
}

func (ie *SSBPeriodicity) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, true)
	ie.Value = aper.Enumerated(v)
	return
}
