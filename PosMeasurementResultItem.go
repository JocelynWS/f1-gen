package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PosMeasurementResultItem struct {
	MeasuredResultsValue MeasuredResultsValue   `mandatory`
	TimeStamp            TimeStamp              `mandatory`
	MeasurementQuality   *TRPMeasurementQuality `optional`
	MeasurementBeamInfo  *MeasurementBeamInfo   `optional`
	// IEExtensions * `optional`
}

func (ie *PosMeasurementResultItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.MeasurementQuality != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.MeasurementBeamInfo != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if err = ie.MeasuredResultsValue.Encode(w); err != nil {
		err = utils.WrapError("Encode MeasuredResultsValue", err)
		return
	}
	if err = ie.TimeStamp.Encode(w); err != nil {
		err = utils.WrapError("Encode TimeStamp", err)
		return
	}
	if ie.MeasurementQuality != nil {
		if err = ie.MeasurementQuality.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasurementQuality", err)
			return
		}
	}
	if ie.MeasurementBeamInfo != nil {
		if err = ie.MeasurementBeamInfo.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasurementBeamInfo", err)
			return
		}
	}
	return
}
func (ie *PosMeasurementResultItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if err = ie.MeasuredResultsValue.Decode(r); err != nil {
		err = utils.WrapError("Read MeasuredResultsValue", err)
		return
	}
	if err = ie.TimeStamp.Decode(r); err != nil {
		err = utils.WrapError("Read TimeStamp", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(TRPMeasurementQuality)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read MeasurementQuality", err)
			return
		}
		ie.MeasurementQuality = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp := new(MeasurementBeamInfo)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read MeasurementBeamInfo", err)
			return
		}
		ie.MeasurementBeamInfo = tmp
	}
	return
}
