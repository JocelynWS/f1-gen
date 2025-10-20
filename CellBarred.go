package ies

import "github.com/lvdund/ngap/aper"

const (
	CellBarredBarred    aper.Enumerated = 0
	CellBarredNotBarred aper.Enumerated = 1
)

type CellBarred struct {
	Value aper.Enumerated
}

func (ie *CellBarred) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *CellBarred) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
