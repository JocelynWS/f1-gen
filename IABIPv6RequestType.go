package ies

import "github.com/lvdund/ngap/aper"

const (
	IABIPv6RequestTypePresentNothing uint64 = iota
	IABIPv6RequestTypePresentIPv6Address
	IABIPv6RequestTypePresentIPv6Prefix
)

type IABIPv6RequestType struct {
	Choice      uint64
	IPv6Address *IABTNLAddressesRequested
	IPv6Prefix  *IABTNLAddressesRequested
	// ChoiceExtension
}

func (ie *IABIPv6RequestType) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case IABIPv6RequestTypePresentIPv6Address:
		err = ie.IPv6Address.Encode(w)
	case IABIPv6RequestTypePresentIPv6Prefix:
		err = ie.IPv6Prefix.Encode(w)
	}
	return
}

func (ie *IABIPv6RequestType) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case IABIPv6RequestTypePresentIPv6Address:
		var tmp IABTNLAddressesRequested
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.IPv6Address = &tmp
	case IABIPv6RequestTypePresentIPv6Prefix:
		var tmp IABTNLAddressesRequested
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.IPv6Prefix = &tmp
	}
	return
}
