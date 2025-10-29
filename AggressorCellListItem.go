package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AggressorCellListItem struct {
	AggressorCellID NRCGI `mandatory`
	// IEExtensions * `optional`
}

func (ie *AggressorCellListItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.AggressorCellID.Encode(w); err != nil {
		err = utils.WrapError("Encode AggressorCellID", err)
		return
	}
	return
}
func (ie *AggressorCellListItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.AggressorCellID.Decode(r); err != nil {
		err = utils.WrapError("Read AggressorCellID", err)
		return
	}
	return
}
