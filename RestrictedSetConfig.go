package f1ap

import "github.com/lvdund/ngap/aper"

const (
	RestrictedSetConfigUnrestrictedSet    aper.Enumerated = 0
	RestrictedSetConfigRestrictedSetTypeA aper.Enumerated = 1
	RestrictedSetConfigRestrictedSetTypeB aper.Enumerated = 2
)

type RestrictedSetConfig struct {
	Value aper.Enumerated
}

func (ie *RestrictedSetConfig) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
	return
}

func (ie *RestrictedSetConfig) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	ie.Value = aper.Enumerated(v)
	return
}
