package ies

import "github.com/lvdund/ngap/aper"

const (
	ResourceRepetitionFactorRf1  aper.Enumerated = 0
	ResourceRepetitionFactorRf2  aper.Enumerated = 1
	ResourceRepetitionFactorRf4  aper.Enumerated = 2
	ResourceRepetitionFactorRf6  aper.Enumerated = 3
	ResourceRepetitionFactorRf8  aper.Enumerated = 4
	ResourceRepetitionFactorRf16 aper.Enumerated = 5
	ResourceRepetitionFactorRf32 aper.Enumerated = 6
)

type ResourceRepetitionFactor struct {
	Value aper.Enumerated
}

func (ie *ResourceRepetitionFactor) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, true)
	return
}

func (ie *ResourceRepetitionFactor) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, true)
	ie.Value = aper.Enumerated(v)
	return
}
