package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ExtendedGNBDUName struct {
	GNBDUNameVisibleString *GNBDUNameVisibleString `optional`
	GNBDUNameUTF8String    *GNBDUNameUTF8String    `optional`
	// IEExtensions * `optional`
}

func (ie *ExtendedGNBDUName) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.GNBDUNameVisibleString != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.GNBDUNameUTF8String != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if ie.GNBDUNameVisibleString != nil {
		if err = ie.GNBDUNameVisibleString.Encode(w); err != nil {
			err = utils.WrapError("Encode GNBDUNameVisibleString", err)
			return
		}
	}
	if ie.GNBDUNameUTF8String != nil {
		if err = ie.GNBDUNameUTF8String.Encode(w); err != nil {
			err = utils.WrapError("Encode GNBDUNameUTF8String", err)
			return
		}
	}
	return
}
func (ie *ExtendedGNBDUName) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(GNBDUNameVisibleString)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read GNBDUNameVisibleString", err)
			return
		}
		ie.GNBDUNameVisibleString = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp := new(GNBDUNameUTF8String)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read GNBDUNameUTF8String", err)
			return
		}
		ie.GNBDUNameUTF8String = tmp
	}
	return
}
