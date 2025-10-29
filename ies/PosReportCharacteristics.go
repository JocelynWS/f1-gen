package ies

import "github.com/lvdund/ngap/aper"

const (
	PosReportCharacteristicsOndemand aper.Enumerated = 0
	PosReportCharacteristicsPeriodic aper.Enumerated = 1
)

type PosReportCharacteristics struct {
	Value aper.Enumerated
}

func (ie *PosReportCharacteristics) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *PosReportCharacteristics) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
