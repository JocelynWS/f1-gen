package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SRBsModifiedItem struct {
	SRBID int64 `lb:0,ub:3,mandatory`
	LCID  int64 `lb:1,ub:32,mandatory`
	// IEExtensions * `optional`
}

func (ie *SRBsModifiedItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_SRBID := NewINTEGER(ie.SRBID, aper.Constraint{Lb: 0, Ub: 3}, false)
	if err = tmp_SRBID.Encode(w); err != nil {
		err = utils.WrapError("Encode SRBID", err)
		return
	}
	tmp_LCID := NewINTEGER(ie.LCID, aper.Constraint{Lb: 1, Ub: 32}, false)
	if err = tmp_LCID.Encode(w); err != nil {
		err = utils.WrapError("Encode LCID", err)
		return
	}
	return
}
func (ie *SRBsModifiedItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_SRBID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 3},
		ext: false,
	}
	if err = tmp_SRBID.Decode(r); err != nil {
		err = utils.WrapError("Read SRBID", err)
		return
	}
	ie.SRBID = int64(tmp_SRBID.Value)
	tmp_LCID := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 32},
		ext: false,
	}
	if err = tmp_LCID.Decode(r); err != nil {
		err = utils.WrapError("Read LCID", err)
		return
	}
	ie.LCID = int64(tmp_LCID.Value)
	return
}
