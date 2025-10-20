package ies

import "github.com/lvdund/ngap/aper"

const (
	ResourceTimeGapTg1  aper.Enumerated = 0
	ResourceTimeGapTg2  aper.Enumerated = 1
	ResourceTimeGapTg4  aper.Enumerated = 2
	ResourceTimeGapTg8  aper.Enumerated = 3
	ResourceTimeGapTg16 aper.Enumerated = 4
	ResourceTimeGapTg32 aper.Enumerated = 5
)

type ResourceTimeGap struct {
	Value aper.Enumerated
}

func (ie *ResourceTimeGap) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, true)
	return
}

func (ie *ResourceTimeGap) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, true)
	ie.Value = aper.Enumerated(v)
	return
}
