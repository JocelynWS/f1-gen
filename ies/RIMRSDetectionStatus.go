package ies

import "github.com/lvdund/ngap/aper"

const (
	RIMRSDetectionStatusRsDetected    aper.Enumerated = 0
	RIMRSDetectionStatusRsDisappeared aper.Enumerated = 1
)

type RIMRSDetectionStatus struct {
	Value aper.Enumerated
}

func (ie *RIMRSDetectionStatus) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *RIMRSDetectionStatus) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
