package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type CellstobeActivatedListItem struct {
	NRCGI NRCGI  `mandatory`
	NRPCI *NRPCI `optional`
	// IEExtensions * `optional`
}

func (ie *CellstobeActivatedListItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.NRPCI != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.NRCGI.Encode(w); err != nil {
		err = utils.WrapError("Encode NRCGI", err)
		return
	}
	if ie.NRPCI != nil {
		if err = ie.NRPCI.Encode(w); err != nil {
			err = utils.WrapError("Encode NRPCI", err)
			return
		}
	}
	return
}
func (ie *CellstobeActivatedListItem) Decode(r *aper.AperReader) (err error) {
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
		tmp := new(NRPCI)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read NRPCI", err)
			return
		}
		ie.NRPCI = tmp
	}
	return
}
