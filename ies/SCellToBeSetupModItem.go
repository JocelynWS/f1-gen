package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SCellToBeSetupModItem struct {
	SCellID           NRCGI             `mandatory`
	SCellIndex        int64             `lb:1,ub:31,mandatory,valueExt`
	SCellULConfigured *CellULConfigured `optional`
	// IEExtensions * `optional`
}

func (ie *SCellToBeSetupModItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.SCellULConfigured != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.SCellID.Encode(w); err != nil {
		err = utils.WrapError("Encode SCellID", err)
		return
	}
	tmp_SCellIndex := NewINTEGER(ie.SCellIndex, aper.Constraint{Lb: 1, Ub: 31}, true)
	if err = tmp_SCellIndex.Encode(w); err != nil {
		err = utils.WrapError("Encode SCellIndex", err)
		return
	}
	if ie.SCellULConfigured != nil {
		if err = ie.SCellULConfigured.Encode(w); err != nil {
			err = utils.WrapError("Encode SCellULConfigured", err)
			return
		}
	}
	return
}
func (ie *SCellToBeSetupModItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.SCellID.Decode(r); err != nil {
		err = utils.WrapError("Read SCellID", err)
		return
	}
	tmp_SCellIndex := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 31},
		ext: true,
	}
	if err = tmp_SCellIndex.Decode(r); err != nil {
		err = utils.WrapError("Read SCellIndex", err)
		return
	}
	ie.SCellIndex = int64(tmp_SCellIndex.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp := new(CellULConfigured)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read SCellULConfigured", err)
			return
		}
		ie.SCellULConfigured = tmp
	}
	return
}
