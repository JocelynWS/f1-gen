package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type IABTNLAddressesRequested struct {
	TNLAddressesOrPrefixesRequestedAllTraffic *int64 `lb:1,ub:256,optional`
	TNLAddressesOrPrefixesRequestedF1C        *int64 `lb:1,ub:256,optional`
	TNLAddressesOrPrefixesRequestedF1U        *int64 `lb:1,ub:256,optional`
	TNLAddressesOrPrefixesRequestedNoNF1      *int64 `lb:1,ub:256,optional`
	// IEExtensions * `optional`
}

func (ie *IABTNLAddressesRequested) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.TNLAddressesOrPrefixesRequestedAllTraffic != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.TNLAddressesOrPrefixesRequestedF1C != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.TNLAddressesOrPrefixesRequestedF1U != nil {
		aper.SetBit(optionals, 3)
	}
	if ie.TNLAddressesOrPrefixesRequestedNoNF1 != nil {
		aper.SetBit(optionals, 4)
	}
	w.WriteBits(optionals, 5)
	if ie.TNLAddressesOrPrefixesRequestedAllTraffic != nil {
		tmp_TNLAddressesOrPrefixesRequestedAllTraffic := NewINTEGER(*ie.TNLAddressesOrPrefixesRequestedAllTraffic, aper.Constraint{Lb: 1, Ub: 256}, false)
		if err = tmp_TNLAddressesOrPrefixesRequestedAllTraffic.Encode(w); err != nil {
			err = utils.WrapError("Encode TNLAddressesOrPrefixesRequestedAllTraffic", err)
			return
		}
	}
	if ie.TNLAddressesOrPrefixesRequestedF1C != nil {
		tmp_TNLAddressesOrPrefixesRequestedF1C := NewINTEGER(*ie.TNLAddressesOrPrefixesRequestedF1C, aper.Constraint{Lb: 1, Ub: 256}, false)
		if err = tmp_TNLAddressesOrPrefixesRequestedF1C.Encode(w); err != nil {
			err = utils.WrapError("Encode TNLAddressesOrPrefixesRequestedF1C", err)
			return
		}
	}
	if ie.TNLAddressesOrPrefixesRequestedF1U != nil {
		tmp_TNLAddressesOrPrefixesRequestedF1U := NewINTEGER(*ie.TNLAddressesOrPrefixesRequestedF1U, aper.Constraint{Lb: 1, Ub: 256}, false)
		if err = tmp_TNLAddressesOrPrefixesRequestedF1U.Encode(w); err != nil {
			err = utils.WrapError("Encode TNLAddressesOrPrefixesRequestedF1U", err)
			return
		}
	}
	if ie.TNLAddressesOrPrefixesRequestedNoNF1 != nil {
		tmp_TNLAddressesOrPrefixesRequestedNoNF1 := NewINTEGER(*ie.TNLAddressesOrPrefixesRequestedNoNF1, aper.Constraint{Lb: 1, Ub: 256}, false)
		if err = tmp_TNLAddressesOrPrefixesRequestedNoNF1.Encode(w); err != nil {
			err = utils.WrapError("Encode TNLAddressesOrPrefixesRequestedNoNF1", err)
			return
		}
	}
	return
}
func (ie *IABTNLAddressesRequested) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(5); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_TNLAddressesOrPrefixesRequestedAllTraffic := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 256},
			ext: false,
		}
		if err = tmp_TNLAddressesOrPrefixesRequestedAllTraffic.Decode(r); err != nil {
			err = utils.WrapError("Read TNLAddressesOrPrefixesRequestedAllTraffic", err)
			return
		}
		ie.TNLAddressesOrPrefixesRequestedAllTraffic = (*int64)(&tmp_TNLAddressesOrPrefixesRequestedAllTraffic.Value)
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_TNLAddressesOrPrefixesRequestedF1C := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 256},
			ext: false,
		}
		if err = tmp_TNLAddressesOrPrefixesRequestedF1C.Decode(r); err != nil {
			err = utils.WrapError("Read TNLAddressesOrPrefixesRequestedF1C", err)
			return
		}
		ie.TNLAddressesOrPrefixesRequestedF1C = (*int64)(&tmp_TNLAddressesOrPrefixesRequestedF1C.Value)
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_TNLAddressesOrPrefixesRequestedF1U := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 256},
			ext: false,
		}
		if err = tmp_TNLAddressesOrPrefixesRequestedF1U.Decode(r); err != nil {
			err = utils.WrapError("Read TNLAddressesOrPrefixesRequestedF1U", err)
			return
		}
		ie.TNLAddressesOrPrefixesRequestedF1U = (*int64)(&tmp_TNLAddressesOrPrefixesRequestedF1U.Value)
	}
	if aper.IsBitSet(optionals, 4) {
		tmp_TNLAddressesOrPrefixesRequestedNoNF1 := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 256},
			ext: false,
		}
		if err = tmp_TNLAddressesOrPrefixesRequestedNoNF1.Decode(r); err != nil {
			err = utils.WrapError("Read TNLAddressesOrPrefixesRequestedNoNF1", err)
			return
		}
		ie.TNLAddressesOrPrefixesRequestedNoNF1 = (*int64)(&tmp_TNLAddressesOrPrefixesRequestedNoNF1.Value)
	}
	return
}
