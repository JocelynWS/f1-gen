package f1ap

import "github.com/lvdund/ngap/aper"

const (
	XYZUnitMM aper.Enumerated = 0
	XYZUnitCM aper.Enumerated = 1
	XYZUnitDM aper.Enumerated = 2
)

type XYZUnit struct {
	Value aper.Enumerated
}

func (ie *XYZUnit) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
	return
}

func (ie *XYZUnit) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	ie.Value = aper.Enumerated(v)
	return
}
