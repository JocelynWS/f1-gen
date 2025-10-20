package ies

import "github.com/lvdund/ngap/aper"

const (
	MeasurementPeriodicityPosmeasurementquantities         aper.Enumerated = 0
	MeasurementPeriodicityPosmeasurementquantitiesitem     aper.Enumerated = 1
	MeasurementPeriodicityPosmeasurementtype               aper.Enumerated = 2
	MeasurementPeriodicityTimingreportinggranularityfactor aper.Enumerated = 3
	MeasurementPeriodicityIeextensions                     aper.Enumerated = 4
)

type MeasurementPeriodicity struct {
	Value aper.Enumerated
}

func (ie *MeasurementPeriodicity) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, true)
	return
}

func (ie *MeasurementPeriodicity) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, true)
	ie.Value = aper.Enumerated(v)
	return
}
