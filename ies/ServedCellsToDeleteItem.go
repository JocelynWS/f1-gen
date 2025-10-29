package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ServedCellsToDeleteItem struct {
	OldNRCGI NRCGI `mandatory`
	// IEExtensions * `optional`
}

func (ie *ServedCellsToDeleteItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.OldNRCGI.Encode(w); err != nil {
		err = utils.WrapError("Encode OldNRCGI", err)
		return
	}
	return
}
func (ie *ServedCellsToDeleteItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.OldNRCGI.Decode(r); err != nil {
		err = utils.WrapError("Read OldNRCGI", err)
		return
	}
	return
}
