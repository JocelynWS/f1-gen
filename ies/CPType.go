package ies

import "github.com/lvdund/ngap/aper"

const (
	CPTypeNormal   aper.Enumerated = 0
	CPTypeExtended aper.Enumerated = 1
)

type CPType struct {
	Value aper.Enumerated
}

func (ie *CPType) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *CPType) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
