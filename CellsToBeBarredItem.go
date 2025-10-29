package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type CellsToBeBarredItem struct {
	NRCGI      NRCGI      `mandatory`
	CellBarred CellBarred `mandatory`
	// IEExtensions * `optional`
}

func (ie *CellsToBeBarredItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.NRCGI.Encode(w); err != nil {
		err = utils.WrapError("Encode NRCGI", err)
		return
	}
	if err = ie.CellBarred.Encode(w); err != nil {
		err = utils.WrapError("Encode CellBarred", err)
		return
	}
	return
}
func (ie *CellsToBeBarredItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.NRCGI.Decode(r); err != nil {
		err = utils.WrapError("Read NRCGI", err)
		return
	}
	if err = ie.CellBarred.Decode(r); err != nil {
		err = utils.WrapError("Read CellBarred", err)
		return
	}
	return
}
