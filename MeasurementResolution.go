package ies

import "github.com/lvdund/ngap/aper"

const (
	MeasurementResolutionM0dot1 aper.Enumerated = 0
	MeasurementResolutionM1     aper.Enumerated = 1
	MeasurementResolutionM10    aper.Enumerated = 2
	MeasurementResolutionM30    aper.Enumerated = 3
)

type MeasurementResolution struct {
	Value aper.Enumerated
}

func (ie *MeasurementResolution) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}

func (ie *MeasurementResolution) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
