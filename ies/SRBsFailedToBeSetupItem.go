package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SRBsFailedToBeSetupItem struct {
	SRBID int64  `lb:0,ub:3,mandatory,valueExt`
	Cause *Cause `optional`
	// IEExtensions * `optional`
}

func (ie *SRBsFailedToBeSetupItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.Cause != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_SRBID := NewINTEGER(ie.SRBID, aper.Constraint{Lb: 0, Ub: 3}, true)
	if err = tmp_SRBID.Encode(w); err != nil {
		err = utils.WrapError("Encode SRBID", err)
		return
	}
	if ie.Cause != nil {
		if err = ie.Cause.Encode(w); err != nil {
			err = utils.WrapError("Encode Cause", err)
			return
		}
	}
	return
}
func (ie *SRBsFailedToBeSetupItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_SRBID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 3},
		ext: true,
	}
	if err = tmp_SRBID.Decode(r); err != nil {
		err = utils.WrapError("Read SRBID", err)
		return
	}
	ie.SRBID = int64(tmp_SRBID.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp := new(Cause)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read Cause", err)
			return
		}
		ie.Cause = tmp
	}
	return
}
