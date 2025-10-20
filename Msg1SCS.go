package ies

import "github.com/lvdund/ngap/aper"

const (
	Msg1SCSSCS15  aper.Enumerated = 0
	Msg1SCSSCS30  aper.Enumerated = 1
	Msg1SCSSCS60  aper.Enumerated = 2
	Msg1SCSSCS120 aper.Enumerated = 3
)

type Msg1SCS struct {
	Value aper.Enumerated
}

func (ie *Msg1SCS) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}

func (ie *Msg1SCS) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
