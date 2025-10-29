package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type IABAllocatedTNLAddressItem struct {
	IABTNLAddress      IABTNLAddress       `mandatory`
	IABTNLAddressUsage *IABTNLAddressUsage `optional`
	// IEExtensions * `optional`
}

func (ie *IABAllocatedTNLAddressItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.IABTNLAddressUsage != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.IABTNLAddress.Encode(w); err != nil {
		err = utils.WrapError("Encode IABTNLAddress", err)
		return
	}
	if ie.IABTNLAddressUsage != nil {
		if err = ie.IABTNLAddressUsage.Encode(w); err != nil {
			err = utils.WrapError("Encode IABTNLAddressUsage", err)
			return
		}
	}
	return
}
func (ie *IABAllocatedTNLAddressItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.IABTNLAddress.Decode(r); err != nil {
		err = utils.WrapError("Read IABTNLAddress", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(IABTNLAddressUsage)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read IABTNLAddressUsage", err)
			return
		}
		ie.IABTNLAddressUsage = tmp
	}
	return
}
