package ies

import "github.com/lvdund/ngap/aper"

const (
	DCBasedDuplicationConfiguredTrue  aper.Enumerated = 0
	DCBasedDuplicationConfiguredFalse aper.Enumerated = 1
)

type DCBasedDuplicationConfigured struct {
	Value aper.Enumerated
}

func (ie *DCBasedDuplicationConfigured) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *DCBasedDuplicationConfigured) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
