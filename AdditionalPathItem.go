package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AdditionalPathItem struct {
	RelativePathDelay RelativePathDelay
	PathQuality       *TRPMeasurementQuality
}

func (ie *AdditionalPathItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.PathQuality != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)

	if err = ie.RelativePathDelay.Encode(w); err != nil {
		err = utils.WrapError("Encode RelativePathDelay", err)
		return
	}

	if aper.IsBitSet(optionals, 1) {
		if err = ie.PathQuality.Encode(w); err != nil {
			err = utils.WrapError("Encode PathQuality", err)
			return
		}
	}
	return
}

func (ie *AdditionalPathItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}

	if err = ie.RelativePathDelay.Decode(r); err != nil {
		err = utils.WrapError("Decode RelativePathDelay", err)
		return
	}

	if aper.IsBitSet(optionals, 1) {
		ie.PathQuality = new(TRPMeasurementQuality)
		if err = ie.PathQuality.Decode(r); err != nil {
			err = utils.WrapError("Decode PathQuality", err)
			return
		}
	}
	return
}
