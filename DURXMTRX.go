package ies

import "github.com/lvdund/ngap/aper"

const (
	DURXMTRXSupported    aper.Enumerated = 0
	DURXMTRXNotSupported aper.Enumerated = 1
)

type DURXMTRX struct {
	Value aper.Enumerated
}

func (ie *DURXMTRX) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false)
	return
}

func (ie *DURXMTRX) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false)
	ie.Value = aper.Enumerated(v)
	return
}
