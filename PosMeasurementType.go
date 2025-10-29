package f1ap

import "github.com/lvdund/ngap/aper"

const (
	PosMeasurementTypeGnbrxtx   aper.Enumerated = 0
	PosMeasurementTypeUlsrsrsrp aper.Enumerated = 1
	PosMeasurementTypeUlaoa     aper.Enumerated = 2
	PosMeasurementTypeUlrtoa    aper.Enumerated = 3
)

type PosMeasurementType struct {
	Value aper.Enumerated
}

func (ie *PosMeasurementType) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}

func (ie *PosMeasurementType) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
