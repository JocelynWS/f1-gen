package ies

import (
	"github.com/lvdund/ngap/aper"
)

type ECIDMeasurementPeriodicity int64

const (
	ECIDMeasurementPeriodicity_ms120   ECIDMeasurementPeriodicity = 0
	ECIDMeasurementPeriodicity_ms240   ECIDMeasurementPeriodicity = 1
	ECIDMeasurementPeriodicity_ms480   ECIDMeasurementPeriodicity = 2
	ECIDMeasurementPeriodicity_ms640   ECIDMeasurementPeriodicity = 3
	ECIDMeasurementPeriodicity_ms1024  ECIDMeasurementPeriodicity = 4
	ECIDMeasurementPeriodicity_ms2048  ECIDMeasurementPeriodicity = 5
	ECIDMeasurementPeriodicity_ms5120  ECIDMeasurementPeriodicity = 6
	ECIDMeasurementPeriodicity_ms10240 ECIDMeasurementPeriodicity = 7
	ECIDMeasurementPeriodicity_min1    ECIDMeasurementPeriodicity = 8
	ECIDMeasurementPeriodicity_min6    ECIDMeasurementPeriodicity = 9
	ECIDMeasurementPeriodicity_min12   ECIDMeasurementPeriodicity = 10
	ECIDMeasurementPeriodicity_min30   ECIDMeasurementPeriodicity = 11
	ECIDMeasurementPeriodicity_min60   ECIDMeasurementPeriodicity = 12
)

func (e *ECIDMeasurementPeriodicity) Encode(w *aper.AperWriter) (err error) {
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

func (e *ECIDMeasurementPeriodicity) Decode(r *aper.AperReader) (err error) {
	tmp := ENUMERATED{
		c:   aper.Constraint{Lb: 0, Ub: 12},
		ext: true,
	}
	if err = tmp.Decode(r); err != nil {
		return
	}
	*e = ECIDMeasurementPeriodicity(tmp.Value)
	return
}
