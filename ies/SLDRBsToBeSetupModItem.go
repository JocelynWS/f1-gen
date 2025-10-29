package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SLDRBsToBeSetupModItem struct {
	SLDRBID          int64            `lb:1,ub:512,mandatory`
	SLDRBInformation SLDRBInformation `mandatory`
	RLCMode          *RLCMode         `optional`
	// IEExtensions * `optional`
}

func (ie *SLDRBsToBeSetupModItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.RLCMode != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_SLDRBID := NewINTEGER(ie.SLDRBID, aper.Constraint{Lb: 1, Ub: 512}, false)
	if err = tmp_SLDRBID.Encode(w); err != nil {
		err = utils.WrapError("Encode SLDRBID", err)
		return
	}
	if err = ie.SLDRBInformation.Encode(w); err != nil {
		err = utils.WrapError("Encode SLDRBInformation", err)
		return
	}
	if ie.RLCMode != nil {
		if err = ie.RLCMode.Encode(w); err != nil {
			err = utils.WrapError("Encode RLCMode", err)
			return
		}
	}
	return
}
func (ie *SLDRBsToBeSetupModItem) Decode(r *aper.AperReader) (err error) {
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
	if err = ie.SLDRBInformation.Decode(r); err != nil {
		err = utils.WrapError("Read SLDRBInformation", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(RLCMode)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read RLCMode", err)
			return
		}
		ie.RLCMode = tmp
	}
	return
}
