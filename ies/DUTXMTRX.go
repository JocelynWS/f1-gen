package ies

import "github.com/lvdund/ngap/aper"

const (
	DUTXMTRXSupported    aper.Enumerated = 0
	DUTXMTRXNotSupported aper.Enumerated = 1
)

type DUTXMTRX struct {
	Value aper.Enumerated
}

func (ie *DUTXMTRX) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false)
	return
}

func (ie *DUTXMTRX) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false)
	ie.Value = aper.Enumerated(v)
	return
}
