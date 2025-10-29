package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SLDRBsFailedToBeModifiedItem struct {
	SLDRBID int64  `lb:1,ub:512,mandatory`
	Cause   *Cause `optional`
	// IEExtensions * `optional`
}

func (ie *SLDRBsFailedToBeModifiedItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.Cause != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_SLDRBID := NewINTEGER(ie.SLDRBID, aper.Constraint{Lb: 1, Ub: 512}, false)
	if err = tmp_SLDRBID.Encode(w); err != nil {
		err = utils.WrapError("Encode SLDRBID", err)
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
func (ie *SLDRBsFailedToBeModifiedItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_SLDRBID := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 512},
		ext: false,
	}
	if err = tmp_SLDRBID.Decode(r); err != nil {
		err = utils.WrapError("Read SLDRBID", err)
		return
	}
	ie.SLDRBID = int64(tmp_SLDRBID.Value)
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
