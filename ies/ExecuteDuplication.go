package ies

import "github.com/lvdund/ngap/aper"

const (
	ExecuteDuplicationTrue aper.Enumerated = 0
)

type ExecuteDuplication struct {
	Value aper.Enumerated
}

func (ie *ExecuteDuplication) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, true)
	return
}

func (ie *ExecuteDuplication) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true)
	ie.Value = aper.Enumerated(v)
	return
}
