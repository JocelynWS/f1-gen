package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DRBActivityItem struct {
	DRBID       int64        `lb:1,ub:32,mandatory,valueExt`
	DRBActivity *DRBActivity `optional`
	// IEExtensions * `optional`
}

func (ie *DRBActivityItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.DRBActivity != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_DRBID := NewINTEGER(ie.DRBID, aper.Constraint{Lb: 1, Ub: 32}, true)
	if err = tmp_DRBID.Encode(w); err != nil {
		err = utils.WrapError("Encode DRBID", err)
		return
	}
	if ie.DRBActivity != nil {
		if err = ie.DRBActivity.Encode(w); err != nil {
			err = utils.WrapError("Encode DRBActivity", err)
			return
		}
	}
	return
}
func (ie *DRBActivityItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_DRBID := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 32},
		ext: true,
	}
	if err = tmp_DRBID.Decode(r); err != nil {
		err = utils.WrapError("Read DRBID", err)
		return
	}
	ie.DRBID = int64(tmp_DRBID.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp := new(DRBActivity)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read DRBActivity", err)
			return
		}
		ie.DRBActivity = tmp
	}
	return
}
