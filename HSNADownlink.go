package ies

import "github.com/lvdund/ngap/aper"

const (
	HSNADownlinkHard         aper.Enumerated = 0
	HSNADownlinkSoft         aper.Enumerated = 1
	HSNADownlinkNotavailable aper.Enumerated = 2
)

type HSNADownlink struct {
	Value aper.Enumerated
}

func (ie *HSNADownlink) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false)
	return
}

func (ie *HSNADownlink) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	ie.Value = aper.Enumerated(v)
	return
}
