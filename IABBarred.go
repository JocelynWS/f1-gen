package f1ap

import "github.com/lvdund/ngap/aper"

const (
	IABBarredBarred    aper.Enumerated = 0
	IABBarredNotBarred aper.Enumerated = 1
)

type IABBarred struct {
	Value aper.Enumerated
}

func (ie *IABBarred) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *IABBarred) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
