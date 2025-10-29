package f1ap

import "github.com/lvdund/ngap/aper"

const (
	ULUEConfigurationNoData aper.Enumerated = 0
	ULUEConfigurationShared aper.Enumerated = 1
	ULUEConfigurationOnly   aper.Enumerated = 2
)

type ULUEConfiguration struct {
	Value aper.Enumerated
}

func (ie *ULUEConfiguration) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
	return
}

func (ie *ULUEConfiguration) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	ie.Value = aper.Enumerated(v)
	return
}
