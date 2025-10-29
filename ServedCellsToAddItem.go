package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ServedCellsToAddItem struct {
	ServedCellInformation  ServedCellInformation   `mandatory`
	GNBDUSystemInformation *GNBDUSystemInformation `optional`
	// IEExtensions * `optional`
}

func (ie *ServedCellsToAddItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.GNBDUSystemInformation != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
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
func (ie *ServedCellsToAddItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
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
