package ies

import "github.com/lvdund/ngap/aper"

const (
	Shift7dot5kHzTrue aper.Enumerated = 0
)

type Shift7dot5kHz struct {
	Value aper.Enumerated
}

func (ie *Shift7dot5kHz) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, true)
	return
}

func (ie *Shift7dot5kHz) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true)
	ie.Value = aper.Enumerated(v)
	return
}
