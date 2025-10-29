package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type IABv4AddressesRequested struct {
	IABv4AddressesRequested IABTNLAddressesRequested `mandatory`
	// IEExtensions * `optional`
}

func (ie *IABv4AddressesRequested) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.IABv4AddressesRequested.Encode(w); err != nil {
		err = utils.WrapError("Encode IABv4AddressesRequested", err)
		return
	}
	return
}
func (ie *IABv4AddressesRequested) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.IABv4AddressesRequested.Decode(r); err != nil {
		err = utils.WrapError("Read IABv4AddressesRequested", err)
		return
	}
	return
}
