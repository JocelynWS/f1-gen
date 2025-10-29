package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ServedCellsToModifyItem struct {
	OldNRCGI               NRCGI                   `mandatory`
	ServedCellInformation  ServedCellInformation   `mandatory`
	GNBDUSystemInformation *GNBDUSystemInformation `optional`
	// IEExtensions * `optional`
}

func (ie *ServedCellsToModifyItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.GNBDUSystemInformation != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.OldNRCGI.Encode(w); err != nil {
		err = utils.WrapError("Encode OldNRCGI", err)
		return
	}
	if err = ie.ServedCellInformation.Encode(w); err != nil {
		err = utils.WrapError("Encode ServedCellInformation", err)
		return
	}
	if ie.GNBDUSystemInformation != nil {
		if err = ie.GNBDUSystemInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode GNBDUSystemInformation", err)
			return
		}
	}
	return
}
func (ie *ServedCellsToModifyItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.OldNRCGI.Decode(r); err != nil {
		err = utils.WrapError("Read OldNRCGI", err)
		return
	}
	if err = ie.ServedCellInformation.Decode(r); err != nil {
		err = utils.WrapError("Read ServedCellInformation", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(GNBDUSystemInformation)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read GNBDUSystemInformation", err)
			return
		}
		ie.GNBDUSystemInformation = tmp
	}
	return
}
