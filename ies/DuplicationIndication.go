package ies

import "github.com/lvdund/ngap/aper"

const (
	DuplicationIndicationTrue  aper.Enumerated = 0
	DuplicationIndicationFalse aper.Enumerated = 1
)

type DuplicationIndication struct {
	Value aper.Enumerated
}

func (ie *DuplicationIndication) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *DuplicationIndication) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
