package ies

import "github.com/lvdund/ngap/aper"

const (
	TransmissionActionIndicatorStop    aper.Enumerated = 0
	TransmissionActionIndicatorRestart aper.Enumerated = 1
)

type TransmissionActionIndicator struct {
	Value aper.Enumerated
}

func (ie *TransmissionActionIndicator) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *TransmissionActionIndicator) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
