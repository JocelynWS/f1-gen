package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type EgressBHRLCCHItem struct {
	NextHopBAPAddress aper.BitString `lb:10,ub:10,mandatory`
	BHRLCChannelID    aper.BitString `lb:16,ub:16,mandatory`
	// IEExtensions *EgressBHRLCCHItemExtIEs `optional`
}

func (ie *EgressBHRLCCHItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_NextHopBAPAddress := NewBITSTRING(ie.NextHopBAPAddress, aper.Constraint{Lb: 10, Ub: 10}, false)
	if err = tmp_NextHopBAPAddress.Encode(w); err != nil {
		err = utils.WrapError("Encode NextHopBAPAddress", err)
		return
	}
	tmp_BHRLCChannelID := NewBITSTRING(ie.BHRLCChannelID, aper.Constraint{Lb: 16, Ub: 16}, false)
	if err = tmp_BHRLCChannelID.Encode(w); err != nil {
		err = utils.WrapError("Encode BHRLCChannelID", err)
		return
	}
	return
}
func (ie *EgressBHRLCCHItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_NextHopBAPAddress := BITSTRING{
		c:   aper.Constraint{Lb: 10, Ub: 10},
		ext: false,
	}
	if err = tmp_NextHopBAPAddress.Decode(r); err != nil {
		err = utils.WrapError("Read NextHopBAPAddress", err)
		return
	}
	ie.NextHopBAPAddress = aper.BitString{Bytes: tmp_NextHopBAPAddress.Value.Bytes, NumBits: tmp_NextHopBAPAddress.Value.NumBits}
	tmp_BHRLCChannelID := BITSTRING{
		c:   aper.Constraint{Lb: 16, Ub: 16},
		ext: false,
	}
	if err = tmp_BHRLCChannelID.Decode(r); err != nil {
		err = utils.WrapError("Read BHRLCChannelID", err)
		return
	}
	ie.BHRLCChannelID = aper.BitString{Bytes: tmp_BHRLCChannelID.Value.Bytes, NumBits: tmp_BHRLCChannelID.Value.NumBits}
	return
}
