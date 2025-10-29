package f1ap

import "github.com/lvdund/ngap/aper"

const (
	ReflectiveQoSAttributeSubjectTo aper.Enumerated = 0
)

type ReflectiveQoSAttribute struct {
	Value aper.Enumerated
}

func (ie *ReflectiveQoSAttribute) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, true)
	return
}

func (ie *ReflectiveQoSAttribute) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true)
	ie.Value = aper.Enumerated(v)
	return
}
