package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ExplicitFormat struct {
	Permutation         Permutation `mandatory`
	NoofDownlinkSymbols *int64      `lb:0,ub:14,optional`
	NoofUplinkSymbols   *int64      `lb:0,ub:14,optional`
	// IEExtensions * `optional`
}

func (ie *ExplicitFormat) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.NoofDownlinkSymbols != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.NoofUplinkSymbols != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if err = ie.Permutation.Encode(w); err != nil {
		err = utils.WrapError("Encode Permutation", err)
		return
	}
	if ie.NoofDownlinkSymbols != nil {
		tmp_NoofDownlinkSymbols := NewINTEGER(*ie.NoofDownlinkSymbols, aper.Constraint{Lb: 0, Ub: 14}, false)
		if err = tmp_NoofDownlinkSymbols.Encode(w); err != nil {
			err = utils.WrapError("Encode NoofDownlinkSymbols", err)
			return
		}
	}
	if ie.NoofUplinkSymbols != nil {
		tmp_NoofUplinkSymbols := NewINTEGER(*ie.NoofUplinkSymbols, aper.Constraint{Lb: 0, Ub: 14}, false)
		if err = tmp_NoofUplinkSymbols.Encode(w); err != nil {
			err = utils.WrapError("Encode NoofUplinkSymbols", err)
			return
		}
	}
	return
}
func (ie *ExplicitFormat) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if err = ie.Permutation.Decode(r); err != nil {
		err = utils.WrapError("Read Permutation", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_NoofDownlinkSymbols := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 14},
			ext: false,
		}
		if err = tmp_NoofDownlinkSymbols.Decode(r); err != nil {
			err = utils.WrapError("Read NoofDownlinkSymbols", err)
			return
		}
		ie.NoofDownlinkSymbols = (*int64)(&tmp_NoofDownlinkSymbols.Value)
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_NoofUplinkSymbols := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 14},
			ext: false,
		}
		if err = tmp_NoofUplinkSymbols.Decode(r); err != nil {
			err = utils.WrapError("Read NoofUplinkSymbols", err)
			return
		}
		ie.NoofUplinkSymbols = (*int64)(&tmp_NoofUplinkSymbols.Value)
	}
	return
}
