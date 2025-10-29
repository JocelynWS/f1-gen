package ies

import "github.com/lvdund/ngap/aper"

const (
	HSNAFlexibleHard         aper.Enumerated = 0
	HSNAFlexibleSoft         aper.Enumerated = 1
	HSNAFlexibleNotavailable aper.Enumerated = 2
)

type HSNAFlexible struct {
	Value aper.Enumerated
}

func (ie *HSNAFlexible) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false)
	return
}

func (ie *HSNAFlexible) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	ie.Value = aper.Enumerated(v)
	return
}
