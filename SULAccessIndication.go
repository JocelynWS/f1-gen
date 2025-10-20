package ies

import "github.com/lvdund/ngap/aper"

const (
	SULAccessIndicationTrue aper.Enumerated = 0
)

type SULAccessIndication struct {
	Value aper.Enumerated
}

func (ie *SULAccessIndication) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, true)
	return
}

func (ie *SULAccessIndication) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true)
	ie.Value = aper.Enumerated(v)
	return
}
