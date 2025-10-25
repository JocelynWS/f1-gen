package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type IPHeaderInformation struct {
	DestinationIABTNLAddress IABTNLAddress   `mandatory`
	DsInformationList        []DSCP          `lb:0,ub:maxnoofDSInfo,optional`
	IPv6FlowLabel            *aper.BitString `lb:20,ub:20,optional`
	// IEExtensions * `optional`
}

func (ie *IPHeaderInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.DsInformationList != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.IPv6FlowLabel != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if err = ie.DestinationIABTNLAddress.Encode(w); err != nil {
		err = utils.WrapError("Encode DestinationIABTNLAddress", err)
		return
	}
	if len(ie.DsInformationList) > 0 {
		tmp := Sequence[*DSCP]{
			Value: []*DSCP{},
			c:     aper.Constraint{Lb: 0, Ub: maxnoofDSInfo},
			ext:   false,
		}
		for _, i := range ie.DsInformationList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode DsInformationList", err)
			return
		}
	}
	if ie.IPv6FlowLabel != nil {
		tmp_IPv6FlowLabel := NewBITSTRING(*ie.IPv6FlowLabel, aper.Constraint{Lb: 20, Ub: 20}, false)
		if err = tmp_IPv6FlowLabel.Encode(w); err != nil {
			err = utils.WrapError("Encode IPv6FlowLabel", err)
			return
		}
	}

	return
}
func (ie *IPHeaderInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if err = ie.DestinationIABTNLAddress.Decode(r); err != nil {
		err = utils.WrapError("Read DestinationIABTNLAddress", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_DsInformationList := Sequence[*DSCP]{
			c:   aper.Constraint{Lb: 0, Ub: maxnoofDSInfo},
			ext: false,
		}
		fn := func() *DSCP { return new(DSCP) }
		if err = tmp_DsInformationList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read DsInformationList", err)
			return
		}
		ie.DsInformationList = []DSCP{}
		for _, i := range tmp_DsInformationList.Value {
			ie.DsInformationList = append(ie.DsInformationList, *i)
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_IPv6FlowLabel := BITSTRING{
			c:   aper.Constraint{Lb: 20, Ub: 20},
			ext: false,
		}
		if err = tmp_IPv6FlowLabel.Decode(r); err != nil {
			err = utils.WrapError("Read IPv6FlowLabel", err)
			return
		}
		ie.IPv6FlowLabel = &aper.BitString{Bytes: tmp_IPv6FlowLabel.Value.Bytes, NumBits: tmp_IPv6FlowLabel.Value.NumBits}
	}
	return
}
