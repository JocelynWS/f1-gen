package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PotentialSpCellItem struct {
	PotentialSpCellID NRCGI `mandatory`
	// IEExtensions * `optional`
}

func (ie *PotentialSpCellItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.PotentialSpCellID.Encode(w); err != nil {
		err = utils.WrapError("Encode PotentialSpCellID", err)
		return
	}
	return
}
func (ie *PotentialSpCellItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.PotentialSpCellID.Decode(r); err != nil {
		err = utils.WrapError("Read PotentialSpCellID", err)
		return
	}
	return
}
