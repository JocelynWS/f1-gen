package f1ap

import "github.com/lvdund/ngap/aper"

const (
	MilliArcSecondUnitsZerodot03 aper.Enumerated = 0
	MilliArcSecondUnitsZerodot3  aper.Enumerated = 1
	MilliArcSecondUnitsThree     aper.Enumerated = 2
)

type MilliArcSecondUnits struct {
	Value aper.Enumerated
}

func (ie *MilliArcSecondUnits) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
	return
}

func (ie *MilliArcSecondUnits) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	ie.Value = aper.Enumerated(v)
	return
}
