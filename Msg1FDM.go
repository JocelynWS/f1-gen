package ies

import "github.com/lvdund/ngap/aper"

const (
	Msg1FDMOne   aper.Enumerated = 0
	Msg1FDMTwo   aper.Enumerated = 1
	Msg1FDMFour  aper.Enumerated = 2
	Msg1FDMEight aper.Enumerated = 3
)

type Msg1FDM struct {
	Value aper.Enumerated
}

func (ie *Msg1FDM) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}

func (ie *Msg1FDM) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
