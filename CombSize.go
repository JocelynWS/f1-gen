package f1ap

import "github.com/lvdund/ngap/aper"

const (
	CombSizeN2  aper.Enumerated = 0
	CombSizeN4  aper.Enumerated = 1
	CombSizeN6  aper.Enumerated = 2
	CombSizeN12 aper.Enumerated = 3
)

type CombSize struct {
	Value aper.Enumerated
}

func (ie *CombSize) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}

func (ie *CombSize) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
