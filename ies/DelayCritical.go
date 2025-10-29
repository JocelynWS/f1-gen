package ies

import "github.com/lvdund/ngap/aper"

const (
	DelayCriticalDelayCritical    aper.Enumerated = 0
	DelayCriticalNonDelayCritical aper.Enumerated = 1
)

type DelayCritical struct {
	Value aper.Enumerated
}

func (ie *DelayCritical) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false)
	return
}

func (ie *DelayCritical) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false)
	ie.Value = aper.Enumerated(v)
	return
}
