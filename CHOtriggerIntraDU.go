package ies

import "github.com/lvdund/ngap/aper"

const (
	CHOtriggerIntraDUChoinitiation aper.Enumerated = 0
	CHOtriggerIntraDUChoreplace    aper.Enumerated = 1
	CHOtriggerIntraDUChocancel     aper.Enumerated = 2
)

type CHOtriggerIntraDU struct {
	Value aper.Enumerated
}

func (ie *CHOtriggerIntraDU) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
	return
}

func (ie *CHOtriggerIntraDU) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	ie.Value = aper.Enumerated(v)
	return
}
