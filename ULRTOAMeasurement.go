package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ULRTOAMeasurement struct {
	ULRTOAMeasurementItem ULRTOAMeasurementItem `mandatory`
	AdditionalPathList    *AdditionalPathItem   `optional`
	// IEExtensions * `optional`
}

func (ie *ULRTOAMeasurement) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.AdditionalPathList != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.ULRTOAMeasurementItem.Encode(w); err != nil {
		err = utils.WrapError("Encode ULRTOAMeasurementItem", err)
		return
	}
	if ie.AdditionalPathList != nil {
		if err = ie.AdditionalPathList.Encode(w); err != nil {
			err = utils.WrapError("Encode AdditionalPathList", err)
			return
		}
	}
	return
}
func (ie *ULRTOAMeasurement) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.ULRTOAMeasurementItem.Decode(r); err != nil {
		err = utils.WrapError("Read ULRTOAMeasurementItem", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(AdditionalPathItem)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read AdditionalPathList", err)
			return
		}
		ie.AdditionalPathList = tmp
	}
	return
}
