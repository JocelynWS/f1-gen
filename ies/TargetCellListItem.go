package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TargetCellListItem struct {
	TargetCell NRCGI `mandatory`
	// IEExtensions * `optional`
}

func (ie *TargetCellListItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.TargetCell.Encode(w); err != nil {
		err = utils.WrapError("Encode TargetCell", err)
		return
	}
	return
}
func (ie *TargetCellListItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.TargetCell.Decode(r); err != nil {
		err = utils.WrapError("Read TargetCell", err)
		return
	}
	return
}
