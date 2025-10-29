package f1ap

import "github.com/lvdund/ngap/aper"

const (
	ConfiguredTACIndicationTrue aper.Enumerated = 0
)

type ConfiguredTACIndication struct {
	Value aper.Enumerated
}

func (ie *ConfiguredTACIndication) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, true)
	return
}

func (ie *ConfiguredTACIndication) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true)
	ie.Value = aper.Enumerated(v)
	return
}
