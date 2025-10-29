package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AssociatedSCellItem struct {
	SCellID NRCGI `mandatory`
	// IEExtensions * `optional`
}

func (ie *AssociatedSCellItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.SCellID.Encode(w); err != nil {
		err = utils.WrapError("Encode SCellID", err)
		return
	}
	return
}
func (ie *AssociatedSCellItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.SCellID.Decode(r); err != nil {
		err = utils.WrapError("Read SCellID", err)
		return
	}
	return
}
