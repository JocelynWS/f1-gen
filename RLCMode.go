package ies

import "github.com/lvdund/ngap/aper"

const (
	RLCModeRlcam                 aper.Enumerated = 0
	RLCModeRlcumbidirectional    aper.Enumerated = 1
	RLCModeRlcumunidirectionalul aper.Enumerated = 2
	RLCModeRlcumunidirectionaldl aper.Enumerated = 3
)

type RLCMode struct {
	Value aper.Enumerated
}

func (ie *RLCMode) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}

func (ie *RLCMode) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
