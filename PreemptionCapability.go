package ies

import "github.com/lvdund/ngap/aper"

const (
	PreEmptionCapabilityShallnottriggerpreemption aper.Enumerated = 0
	PreEmptionCapabilityMaytriggerpreemption      aper.Enumerated = 1
)

type PreEmptionCapability struct {
	Value aper.Enumerated
}

func (ie *PreEmptionCapability) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false)
	return
}

func (ie *PreEmptionCapability) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false)
	ie.Value = aper.Enumerated(v)
	return
}
