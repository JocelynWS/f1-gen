package ies

import "github.com/lvdund/ngap/aper"

const (
	RRCReconfigurationCompleteIndicatorTrue    aper.Enumerated = 0
	RRCReconfigurationCompleteIndicatorFailure aper.Enumerated = 1
)

type RRCReconfigurationCompleteIndicator struct {
	Value aper.Enumerated
}

func (ie *RRCReconfigurationCompleteIndicator) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *RRCReconfigurationCompleteIndicator) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
