package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type NeighbourCellInformationItem struct {
	NRCGI                 NRCGI                  `mandatory`
	IntendedTDDDLULConfig *IntendedTDDDLULConfig `optional`
	// IEExtensions * `optional`
}

func (ie *NeighbourCellInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.IntendedTDDDLULConfig != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.NRCGI.Encode(w); err != nil {
		err = utils.WrapError("Encode NRCGI", err)
		return
	}
	if ie.IntendedTDDDLULConfig != nil {
		if err = ie.IntendedTDDDLULConfig.Encode(w); err != nil {
			err = utils.WrapError("Encode IntendedTDDDLULConfig", err)
			return
		}
	}
	return
}
func (ie *NeighbourCellInformationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.NRCGI.Decode(r); err != nil {
		err = utils.WrapError("Read NRCGI", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(IntendedTDDDLULConfig)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read IntendedTDDDLULConfig", err)
			return
		}
		ie.IntendedTDDDLULConfig = tmp
	}
	return
}
