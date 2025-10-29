package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	IABTNLAddressPresentNothing uint64 = iota
	IABTNLAddressPresentIpv4Address
	IABTNLAddressPresentIpv6Address
	IABTNLAddressPresentIpv6Prefix
	IABTNLAddressPresentChoiceExtension
)

type IABTNLAddress struct {
	Choice      uint64
	IPv4Address *aper.BitString
	IPv6Address *aper.BitString
	IPv6Prefix  *aper.BitString
	//ChoiceExtension
}

func (ie *IABTNLAddress) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case IABTNLAddressPresentIpv4Address:
		tmp := NewBITSTRING(*ie.IPv4Address, aper.Constraint{Lb: 32, Ub: 32}, false)
		err = tmp.Encode(w)
	case IABTNLAddressPresentIpv6Address:
		tmp := NewBITSTRING(*ie.IPv6Address, aper.Constraint{Lb: 128, Ub: 128}, false)
		err = tmp.Encode(w)
	case IABTNLAddressPresentIpv6Prefix:
		tmp := NewBITSTRING(*ie.IPv6Prefix, aper.Constraint{Lb: 64, Ub: 64}, false)
		err = tmp.Encode(w)
	}
	return
}

func (ie *IABTNLAddress) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case IABTNLAddressPresentIpv4Address:
		tmp := BITSTRING{c: aper.Constraint{Lb: 32, Ub: 32}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read IPv4Address", err)
			return
		}
		ie.IPv4Address = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	case IABTNLAddressPresentIpv6Address:
		tmp := BITSTRING{c: aper.Constraint{Lb: 128, Ub: 128}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read IPv6Address", err)
			return
		}
		ie.IPv6Address = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	case IABTNLAddressPresentIpv6Prefix:
		tmp := BITSTRING{c: aper.Constraint{Lb: 64, Ub: 64}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read IPv6Prefix", err)
			return
		}
		ie.IPv6Prefix = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	}
	return
}
