package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type BAPlayerBHRLCchannelMappingInfoItem struct {
	MappingInformationIndex aper.BitString  `lb:1,ub:65536,mandatory,valueExt`
	PriorHopBAPAddress      *aper.BitString `lb:10,ub:10,optional`
	IngressbHRLCChannelID   *aper.BitString `lb:16,ub:16,optional`
	NextHopBAPAddress       *aper.BitString `lb:10,ub:10,optional`
	EgressbHRLCChannelID    *aper.BitString `lb:16,ub:16,optional`
}

func (ie *BAPlayerBHRLCchannelMappingInfoItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.PriorHopBAPAddress != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.IngressbHRLCChannelID != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.NextHopBAPAddress != nil {
		aper.SetBit(optionals, 3)
	}
	if ie.EgressbHRLCChannelID != nil {
		aper.SetBit(optionals, 4)
	}
	w.WriteBits(optionals, 5)
	tmp_MappingInformationIndex := NewBITSTRING(ie.MappingInformationIndex, aper.Constraint{Lb: 1, Ub: 65536}, true)
	if err = tmp_MappingInformationIndex.Encode(w); err != nil {
		err = utils.WrapError("Encode MappingInformationIndex", err)
		return
	}
	if ie.PriorHopBAPAddress != nil {
		tmp_PriorHopBAPAddress := NewBITSTRING(*ie.PriorHopBAPAddress, aper.Constraint{Lb: 10, Ub: 10}, false)
		if err = tmp_PriorHopBAPAddress.Encode(w); err != nil {
			err = utils.WrapError("Encode PriorHopBAPAddress", err)
			return
		}
	}
	if ie.IngressbHRLCChannelID != nil {
		tmp_IngressbHRLCChannelID := NewBITSTRING(*ie.IngressbHRLCChannelID, aper.Constraint{Lb: 16, Ub: 16}, false)
		if err = tmp_IngressbHRLCChannelID.Encode(w); err != nil {
			err = utils.WrapError("Encode IngressbHRLCChannelID", err)
			return
		}
	}
	if ie.NextHopBAPAddress != nil {
		tmp_NextHopBAPAddress := NewBITSTRING(*ie.NextHopBAPAddress, aper.Constraint{Lb: 10, Ub: 10}, false)
		if err = tmp_NextHopBAPAddress.Encode(w); err != nil {
			err = utils.WrapError("Encode NextHopBAPAddress", err)
			return
		}
	}
	if ie.EgressbHRLCChannelID != nil {
		tmp_EgressbHRLCChannelID := NewBITSTRING(*ie.EgressbHRLCChannelID, aper.Constraint{Lb: 16, Ub: 16}, false)
		if err = tmp_EgressbHRLCChannelID.Encode(w); err != nil {
			err = utils.WrapError("Encode EgressbHRLCChannelID", err)
			return
		}
	}
	return
}
func (ie *BAPlayerBHRLCchannelMappingInfoItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(5); err != nil {
		return
	}
	tmp_MappingInformationIndex := BITSTRING{
		c:   aper.Constraint{Lb: 1, Ub: 65536},
		ext: true,
	}
	if err = tmp_MappingInformationIndex.Decode(r); err != nil {
		err = utils.WrapError("Read MappingInformationIndex", err)
		return
	}
	ie.MappingInformationIndex = aper.BitString{Bytes: tmp_MappingInformationIndex.Value.Bytes, NumBits: tmp_MappingInformationIndex.Value.NumBits}
	if aper.IsBitSet(optionals, 1) {
		tmp_PriorHopBAPAddress := BITSTRING{
			c:   aper.Constraint{Lb: 10, Ub: 10},
			ext: false,
		}
		if err = tmp_PriorHopBAPAddress.Decode(r); err != nil {
			err = utils.WrapError("Read PriorHopBAPAddress", err)
			return
		}
		ie.PriorHopBAPAddress = &aper.BitString{Bytes: tmp_PriorHopBAPAddress.Value.Bytes, NumBits: tmp_PriorHopBAPAddress.Value.NumBits}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_IngressbHRLCChannelID := BITSTRING{
			c:   aper.Constraint{Lb: 16, Ub: 16},
			ext: false,
		}
		if err = tmp_IngressbHRLCChannelID.Decode(r); err != nil {
			err = utils.WrapError("Read IngressbHRLCChannelID", err)
			return
		}
		ie.IngressbHRLCChannelID = &aper.BitString{Bytes: tmp_IngressbHRLCChannelID.Value.Bytes, NumBits: tmp_IngressbHRLCChannelID.Value.NumBits}
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_NextHopBAPAddress := BITSTRING{
			c:   aper.Constraint{Lb: 10, Ub: 10},
			ext: false,
		}
		if err = tmp_NextHopBAPAddress.Decode(r); err != nil {
			err = utils.WrapError("Read NextHopBAPAddress", err)
			return
		}
		ie.NextHopBAPAddress = &aper.BitString{Bytes: tmp_NextHopBAPAddress.Value.Bytes, NumBits: tmp_NextHopBAPAddress.Value.NumBits}
	}
	if aper.IsBitSet(optionals, 4) {
		tmp_EgressbHRLCChannelID := BITSTRING{
			c:   aper.Constraint{Lb: 16, Ub: 16},
			ext: false,
		}
		if err = tmp_EgressbHRLCChannelID.Decode(r); err != nil {
			err = utils.WrapError("Read EgressbHRLCChannelID", err)
			return
		}
		ie.EgressbHRLCChannelID = &aper.BitString{Bytes: tmp_EgressbHRLCChannelID.Value.Bytes, NumBits: tmp_EgressbHRLCChannelID.Value.NumBits}
	}
	return
}
