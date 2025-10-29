package f1ap

import "github.com/lvdund/ngap/aper"

const (
	DRBActivityActive    aper.Enumerated = 0
	DRBActivityNotActive aper.Enumerated = 1
)

type DRBActivity struct {
	Value aper.Enumerated
}

func (ie *DRBActivity) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false)
	return
}

func (ie *DRBActivity) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false)
	ie.Value = aper.Enumerated(v)
	return
}
