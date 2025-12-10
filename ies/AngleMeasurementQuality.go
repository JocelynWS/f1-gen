package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AngleMeasurementQuality struct {
	AzimuthQuality int64      `lb:0,ub:255,madatory`
	ZenithQuality  *int64     `optional,lb:0,ub:255`
	Resolution     Resolution `madatory,valueExt`
	// IEExtensions *AngleMeasurementQualityExtIEs `optional`
}

func (ie *AngleMeasurementQuality) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.ZenithQuality != nil {
		aper.SetBit(optionals, 0)
	}
	if err = w.WriteBits(optionals, 2); err != nil {
		return
	}
	tmp_AzimuthQuality := NewINTEGER(ie.AzimuthQuality, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_AzimuthQuality.Encode(w); err != nil {
		err = utils.WrapError("Encode AzimuthQuality", err)
		return
	}
	if ie.ZenithQuality != nil {
		tmp_ZenithQuality := NewINTEGER(*ie.ZenithQuality, aper.Constraint{Lb: 0, Ub: 255}, false)
		if err = tmp_ZenithQuality.Encode(w); err != nil {
			err = utils.WrapError("Encode ZenithQuality", err)
			return
		}
	}
	if err = ie.Resolution.Encode(w); err != nil {
		err = utils.WrapError("Encode Resolution", err)
		return
	}
	return
}

func (ie *AngleMeasurementQuality) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_AzimuthQuality := INTEGER{c: aper.Constraint{Lb: 0, Ub: 255}}
	if err = tmp_AzimuthQuality.Decode(r); err != nil {
		err = utils.WrapError("Read AzimuthQuality", err)
		return
	}
	ie.AzimuthQuality = int64(tmp_AzimuthQuality.Value)
	if aper.IsBitSet(optionals, 0) {
		tmp_ZenithQuality := INTEGER{c: aper.Constraint{Lb: 0, Ub: 255}}
		if err = tmp_ZenithQuality.Decode(r); err != nil {
			err = utils.WrapError("Read ZenithQuality", err)
			return
		}
		val := int64(tmp_ZenithQuality.Value)
		ie.ZenithQuality = &val
	}
	if err = ie.Resolution.Decode(r); err != nil {
		err = utils.WrapError("Read Resolution", err)
		return
	}
	return
}
