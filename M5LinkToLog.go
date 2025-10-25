package ies

import "github.com/lvdund/ngap/aper"

const (
	M5LinkstologUplink                aper.Enumerated = 0
	M5LinkstologDownlink              aper.Enumerated = 1
	M5LinkstologBothUplinkAndDownlink aper.Enumerated = 2
)

type M5LinksToLog struct {
	Value aper.Enumerated
}

func (ie *M5LinksToLog) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
	return
}

func (ie *M5LinksToLog) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	ie.Value = aper.Enumerated(v)
	return
}
