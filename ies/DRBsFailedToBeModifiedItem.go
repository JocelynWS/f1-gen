package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DRBsFailedToBeModifiedItem struct {
	DRBID int64  `lb:1,ub:32,mandatory`
	Cause *Cause `optional`
	// IEExtensions * `optional`
}

func (ie *DRBsFailedToBeModifiedItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.Cause != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_DRBID := NewINTEGER(ie.DRBID, aper.Constraint{Lb: 1, Ub: 32}, false)
	if err = tmp_DRBID.Encode(w); err != nil {
		err = utils.WrapError("Encode DRBID", err)
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
func (ie *DRBsFailedToBeModifiedItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_DRBID := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 32},
		ext: false,
	}
	if err = tmp_DRBID.Decode(r); err != nil {
		err = utils.WrapError("Read DRBID", err)
		return
	}
	ie.DRBID = int64(tmp_DRBID.Value)
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
