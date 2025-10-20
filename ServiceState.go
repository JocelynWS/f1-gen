package ies

import "github.com/lvdund/ngap/aper"

const (
	ServiceStateInservice    aper.Enumerated = 0
	ServiceStateOutofservice aper.Enumerated = 1
)

type ServiceState struct {
	Value aper.Enumerated
}

func (ie *ServiceState) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *ServiceState) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
