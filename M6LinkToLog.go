package ies

import "github.com/lvdund/ngap/aper"

const (
	M6LinkstologUplink                aper.Enumerated = 0
	M6LinkstologDownlink              aper.Enumerated = 1
	M6LinkstologBothUplinkAndDownlink aper.Enumerated = 2
)

type M6LinksToLog struct {
	Value aper.Enumerated
}

func (ie *M6LinksToLog) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
	return
}

func (ie *M6LinksToLog) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	ie.Value = aper.Enumerated(v)
	return
}
