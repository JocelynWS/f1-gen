package ies

import "github.com/lvdund/ngap/aper"

const (
	ResourceNumberofSymbolsN2  aper.Enumerated = 0
	ResourceNumberofSymbolsN4  aper.Enumerated = 1
	ResourceNumberofSymbolsN6  aper.Enumerated = 2
	ResourceNumberofSymbolsN12 aper.Enumerated = 3
)

type ResourceNumberofSymbols struct {
	Value aper.Enumerated
}

func (ie *ResourceNumberofSymbols) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}

func (ie *ResourceNumberofSymbols) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
