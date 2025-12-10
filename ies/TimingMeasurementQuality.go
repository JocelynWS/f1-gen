package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TimingMeasurementQuality struct {
	MeasurementQuality int64                 `lb:0,ub:31,madatory`
	Resolution         MeasurementResolution `madatory,valueExt`
	// IEExtensions *TimingMeasurementQualityExtIEs `optional`
}

func (ie *TimingMeasurementQuality) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if err = w.WriteBits(optionals, 1); err != nil {
		return
	}

	tmp_MeasurementQuality := NewINTEGER(ie.MeasurementQuality, aper.Constraint{Lb: 0, Ub: 31}, false)
	if err = tmp_MeasurementQuality.Encode(w); err != nil {
		err = utils.WrapError("Encode MeasurementQuality", err)
		return
	}
	if err = ie.Resolution.Encode(w); err != nil {
		err = utils.WrapError("Encode Resolution", err)
		return
	}
	return
}

func (ie *TimingMeasurementQuality) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}

	tmp_MeasurementQuality := INTEGER{c: aper.Constraint{Lb: 0, Ub: 31}}
	if err = tmp_MeasurementQuality.Decode(r); err != nil {
		err = utils.WrapError("Read MeasurementQuality", err)
		return
	}
	ie.MeasurementQuality = int64(tmp_MeasurementQuality.Value)

	if err = ie.Resolution.Decode(r); err != nil {
		err = utils.WrapError("Read Resolution", err)
		return
	}
	return
}
