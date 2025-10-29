package f1ap

import (
	"github.com/lvdund/ngap/aper"
)

type PosMeasurementPeriodicity int64

const (
	PosMeasurementPeriodicity_ms120   PosMeasurementPeriodicity = 0
	PosMeasurementPeriodicity_ms240   PosMeasurementPeriodicity = 1
	PosMeasurementPeriodicity_ms480   PosMeasurementPeriodicity = 2
	PosMeasurementPeriodicity_ms640   PosMeasurementPeriodicity = 3
	PosMeasurementPeriodicity_ms1024  PosMeasurementPeriodicity = 4
	PosMeasurementPeriodicity_ms2048  PosMeasurementPeriodicity = 5
	PosMeasurementPeriodicity_ms5120  PosMeasurementPeriodicity = 6
	PosMeasurementPeriodicity_ms10240 PosMeasurementPeriodicity = 7
	PosMeasurementPeriodicity_min1    PosMeasurementPeriodicity = 8
	PosMeasurementPeriodicity_min6    PosMeasurementPeriodicity = 9
	PosMeasurementPeriodicity_min12   PosMeasurementPeriodicity = 10
	PosMeasurementPeriodicity_min30   PosMeasurementPeriodicity = 11
	PosMeasurementPeriodicity_min60   PosMeasurementPeriodicity = 12
)

func (e *PosMeasurementPeriodicity) Encode(w *aper.AperWriter) (err error) {
	tmp := ENUMERATED{
		Value: aper.Enumerated(*e),
		c:     aper.Constraint{Lb: 0, Ub: 12},
		ext:   true,
	}
	if err = tmp.Encode(w); err != nil {
		return
	}
	return
}

func (e *PosMeasurementPeriodicity) Decode(r *aper.AperReader) (err error) {
	tmp := ENUMERATED{
		c:   aper.Constraint{Lb: 0, Ub: 12},
		ext: true,
	}
	if err = tmp.Decode(r); err != nil {
		return
	}
	*e = PosMeasurementPeriodicity(tmp.Value)
	return
}
