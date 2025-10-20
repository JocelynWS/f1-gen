package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GNBRxTxTimeDiff struct {
	RxTxTimeDiff       GNBRxTxTimeDiffMeas `mandatory`
	AdditionalPathList *AdditionalPathList `optional`
	// IEExtensions * `optional`
}

func (ie *GNBRxTxTimeDiff) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.AdditionalPathList != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.RxTxTimeDiff.Encode(w); err != nil {
		err = utils.WrapError("Encode RxTxTimeDiff", err)
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
func (ie *GNBRxTxTimeDiff) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.RxTxTimeDiff.Decode(r); err != nil {
		err = utils.WrapError("Read RxTxTimeDiff", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(AdditionalPathList)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read AdditionalPathList", err)
			return
		}
		ie.AdditionalPathList = tmp
	}
	return
}
