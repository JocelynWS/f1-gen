package ies

import "github.com/lvdund/ngap/aper"

const (
	PermutationDfu aper.Enumerated = 0
	PermutationUfd aper.Enumerated = 1
)

type Permutation struct {
	Value aper.Enumerated
}

func (ie *Permutation) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *Permutation) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
