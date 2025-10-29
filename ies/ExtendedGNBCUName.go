package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ExtendedGNBCUName struct {
	GNBCUNameVisibleString *GNBCUNameVisibleString `optional`
	GNBCUNameUTF8String    *GNBCUNameUTF8String    `optional`
	// IEExtensions * `optional`
}

func (ie *ExtendedGNBCUName) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.GNBCUNameVisibleString != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.GNBCUNameUTF8String != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if ie.GNBCUNameVisibleString != nil {
		if err = ie.GNBCUNameVisibleString.Encode(w); err != nil {
			err = utils.WrapError("Encode GNBCUNameVisibleString", err)
			return
		}
	}
	if ie.GNBCUNameUTF8String != nil {
		if err = ie.GNBCUNameUTF8String.Encode(w); err != nil {
			err = utils.WrapError("Encode GNBCUNameUTF8String", err)
			return
		}
	}
	return
}
func (ie *ExtendedGNBCUName) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(GNBCUNameVisibleString)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read GNBCUNameVisibleString", err)
			return
		}
		ie.GNBCUNameVisibleString = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp := new(GNBCUNameUTF8String)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read GNBCUNameUTF8String", err)
			return
		}
		ie.GNBCUNameUTF8String = tmp
	}
	return
}
