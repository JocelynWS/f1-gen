package ies

import "github.com/lvdund/ngap/aper"

const (
	RepetitionFactorN1 aper.Enumerated = 0
	RepetitionFactorN2 aper.Enumerated = 1
	RepetitionFactorN4 aper.Enumerated = 2
)

type RepetitionFactor struct {
	Value aper.Enumerated
}

func (ie *RepetitionFactor) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false)
	return
}

func (ie *RepetitionFactor) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	ie.Value = aper.Enumerated(v)
	return
}
