package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type BHChannelsFailedToBeSetupModItem struct {
	BHRLCChannelID aper.BitString `lb:16,ub:16,mandatory`
	Cause          *Cause         `optional`
	// IEExtensions * `optional`
}

func (ie *BHChannelsFailedToBeSetupModItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.Cause != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_BHRLCChannelID := NewBITSTRING(ie.BHRLCChannelID, aper.Constraint{Lb: 16, Ub: 16}, false)
	if err = tmp_BHRLCChannelID.Encode(w); err != nil {
		err = utils.WrapError("Encode BHRLCChannelID", err)
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
func (ie *BHChannelsFailedToBeSetupModItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_BHRLCChannelID := BITSTRING{
		c:   aper.Constraint{Lb: 16, Ub: 16},
		ext: false,
	}
	if err = tmp_BHRLCChannelID.Decode(r); err != nil {
		err = utils.WrapError("Read BHRLCChannelID", err)
		return
	}
	ie.BHRLCChannelID = aper.BitString{Bytes: tmp_BHRLCChannelID.Value.Bytes, NumBits: tmp_BHRLCChannelID.Value.NumBits}
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
