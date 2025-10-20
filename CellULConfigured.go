package ies

import "github.com/lvdund/ngap/aper"

const (
	CellULConfiguredNone     aper.Enumerated = 0
	CellULConfiguredUl       aper.Enumerated = 1
	CellULConfiguredSul      aper.Enumerated = 2
	CellULConfiguredUlAndSul aper.Enumerated = 3
)

type CellULConfigured struct {
	Value aper.Enumerated
}

func (ie *CellULConfigured) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}

func (ie *CellULConfigured) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
