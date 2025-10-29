package f1ap

import "github.com/lvdund/ngap/aper"

const (
	PeriodicitySRSPeriodicitylist     aper.Enumerated = 0
	PeriodicitySRSPeriodicitylistitem aper.Enumerated = 1
	PeriodicitySRSPeriodicitysrs      aper.Enumerated = 2
	PeriodicitySRSIeextensions        aper.Enumerated = 3
)

type PeriodicitySRS struct {
	Value aper.Enumerated
}

func (ie *PeriodicitySRS) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}

func (ie *PeriodicitySRS) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
