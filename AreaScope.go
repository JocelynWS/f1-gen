package f1ap

import "github.com/lvdund/ngap/aper"

const (
	AreaScopeTrue aper.Enumerated = 0
)

type AreaScope struct {
	Value aper.Enumerated
}

func (ie *AreaScope) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, true)
	return
}

func (ie *AreaScope) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true)
	ie.Value = aper.Enumerated(v)
	return
}
