package ies

import "github.com/lvdund/ngap/aper"

const (
	ReportingPeriodicityMs500   aper.Enumerated = 0
	ReportingPeriodicityMs1000  aper.Enumerated = 1
	ReportingPeriodicityMs2000  aper.Enumerated = 2
	ReportingPeriodicityMs5000  aper.Enumerated = 3
	ReportingPeriodicityMs10000 aper.Enumerated = 4
)

type ReportingPeriodicity struct {
	Value aper.Enumerated
}

func (ie *ReportingPeriodicity) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, true)
	return
}

func (ie *ReportingPeriodicity) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, true)
	ie.Value = aper.Enumerated(v)
	return
}
