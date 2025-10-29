package ies

import "github.com/lvdund/ngap/aper"

const (
	LatitudeSignNorth aper.Enumerated = 0
	LatitudeSignSouth aper.Enumerated = 1
)

type LatitudeSign struct {
	Value aper.Enumerated
}

func (ie *LatitudeSign) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false)
	return
}

func (ie *LatitudeSign) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false)
	ie.Value = aper.Enumerated(v)
	return
}
