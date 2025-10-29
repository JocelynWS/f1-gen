package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type BHChannelsToBeSetupItem struct {
	BHRLCChannelID     aper.BitString      `lb:16,ub:16,mandatory`
	BHQoSInformation   BHQoSInformation    `mandatory`
	RLCmode            RLCMode             `mandatory`
	BAPCtrlPDUChannel  *BAPCtrlPDUChannel  `optional`
	TrafficMappingInfo *TrafficMappingInfo `optional`
	// IEExtensions * `optional`
}

func (ie *BHChannelsToBeSetupItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.BAPCtrlPDUChannel != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.TrafficMappingInfo != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	tmp_BHRLCChannelID := NewBITSTRING(ie.BHRLCChannelID, aper.Constraint{Lb: 16, Ub: 16}, false)
	if err = tmp_BHRLCChannelID.Encode(w); err != nil {
		err = utils.WrapError("Encode BHRLCChannelID", err)
		return
	}
	if err = ie.BHQoSInformation.Encode(w); err != nil {
		err = utils.WrapError("Encode BHQoSInformation", err)
		return
	}
	if err = ie.RLCmode.Encode(w); err != nil {
		err = utils.WrapError("Encode RLCmode", err)
		return
	}
	if ie.BAPCtrlPDUChannel != nil {
		if err = ie.BAPCtrlPDUChannel.Encode(w); err != nil {
			err = utils.WrapError("Encode BAPCtrlPDUChannel", err)
			return
		}
	}
	if ie.TrafficMappingInfo != nil {
		if err = ie.TrafficMappingInfo.Encode(w); err != nil {
			err = utils.WrapError("Encode TrafficMappingInfo", err)
			return
		}
	}
	return
}
func (ie *BHChannelsToBeSetupItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
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
	if err = ie.BHQoSInformation.Decode(r); err != nil {
		err = utils.WrapError("Read BHQoSInformation", err)
		return
	}
	if err = ie.RLCmode.Decode(r); err != nil {
		err = utils.WrapError("Read RLCmode", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(BAPCtrlPDUChannel)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read BAPCtrlPDUChannel", err)
			return
		}
		ie.BAPCtrlPDUChannel = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp := new(TrafficMappingInfo)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read TrafficMappingInfo", err)
			return
		}
		ie.TrafficMappingInfo = tmp
	}
	return
}
