package ies

import "github.com/lvdund/ngap/aper"

const (
	FrequencyShift7p5khzFalse aper.Enumerated = 0
	FrequencyShift7p5khzTrue  aper.Enumerated = 1
)

type FrequencyShift7p5khz struct {
	Value aper.Enumerated
}

func (ie *FrequencyShift7p5khz) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *FrequencyShift7p5khz) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
