package ies

import "github.com/lvdund/ngap/aper"

const (
	NrofSymbolsN1  aper.Enumerated = 0
	NrofSymbolsN2  aper.Enumerated = 1
	NrofSymbolsN4  aper.Enumerated = 2
	NrofSymbolsN8  aper.Enumerated = 3
	NrofSymbolsN12 aper.Enumerated = 4
)

type NrofSymbols struct {
	Value aper.Enumerated
}

func (ie *NrofSymbols) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, false)
	return
}

func (ie *NrofSymbols) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, false)
	ie.Value = aper.Enumerated(v)
	return
}
