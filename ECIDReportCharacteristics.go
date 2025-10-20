package ies

import "github.com/lvdund/ngap/aper"

const (
	ECIDReportCharacteristicsOndemand aper.Enumerated = 0
	ECIDReportCharacteristicsPeriodic aper.Enumerated = 1
)

type ECIDReportCharacteristics struct {
	Value aper.Enumerated
}

func (ie *ECIDReportCharacteristics) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *ECIDReportCharacteristics) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
