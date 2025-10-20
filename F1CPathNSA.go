package ies

import "github.com/lvdund/ngap/aper"

const (
	F1CPathNSALte  aper.Enumerated = 0
	F1CPathNSANr   aper.Enumerated = 1
	F1CPathNSABoth aper.Enumerated = 2
)

type F1CPathNSA struct {
	Value aper.Enumerated
}

func (ie *F1CPathNSA) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false)
	return
}

func (ie *F1CPathNSA) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	ie.Value = aper.Enumerated(v)
	return
}
