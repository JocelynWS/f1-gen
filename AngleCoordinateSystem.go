package ies

import "github.com/lvdund/ngap/aper"

const (
	AngleCoordinateSystem_LCS aper.Enumerated = 0
	AngleCoordinateSystem_GCS aper.Enumerated = 1
)

type AngleCoordinateSystem struct {
	Value aper.Enumerated
}

func (ie *AngleCoordinateSystem) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *AngleCoordinateSystem) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	if err != nil {
		return
	}
	ie.Value = aper.Enumerated(v)
	return
}
