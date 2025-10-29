package f1ap

import "github.com/lvdund/ngap/aper"

const (
	NrofSymbolsSRSN1 aper.Enumerated = 0
	NrofSymbolsSRSN2 aper.Enumerated = 1
	NrofSymbolsSRSN4 aper.Enumerated = 2
)

type NrofSymbolsSRS struct {
	Value aper.Enumerated
}

func (ie *NrofSymbolsSRS) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false)
	return
}

func (ie *NrofSymbolsSRS) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	ie.Value = aper.Enumerated(v)
	return
}
