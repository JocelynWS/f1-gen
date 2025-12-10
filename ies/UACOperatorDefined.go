package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UACOperatorDefined struct {
	AccessCategory int64          `lb:32,ub:63,mandatory,valueExt`
	AccessIdentity aper.BitString `lb:7,ub:7,mandatory`
	// IEExtensions * `optional`
}

func (ie *UACOperatorDefined) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_AccessCategory := NewINTEGER(ie.AccessCategory, aper.Constraint{Lb: 32, Ub: 63}, true)
	if err = tmp_AccessCategory.Encode(w); err != nil {
		err = utils.WrapError("Encode AccessCategory", err)
		return
	}
	tmp_AccessIdentity := NewBITSTRING(ie.AccessIdentity, aper.Constraint{Lb: 7, Ub: 7}, false)
	if err = tmp_AccessIdentity.Encode(w); err != nil {
		err = utils.WrapError("Encode AccessIdentity", err)
		return
	}
	return
}
func (ie *UACOperatorDefined) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_AccessCategory := INTEGER{
		c:   aper.Constraint{Lb: 32, Ub: 63},
		ext: true,
	}
	if err = tmp_AccessCategory.Decode(r); err != nil {
		err = utils.WrapError("Read AccessCategory", err)
		return
	}
	ie.AccessCategory = int64(tmp_AccessCategory.Value)
	tmp_AccessIdentity := BITSTRING{
		c:   aper.Constraint{Lb: 7, Ub: 7},
		ext: false,
	}
	if err = tmp_AccessIdentity.Decode(r); err != nil {
		err = utils.WrapError("Read AccessIdentity", err)
		return
	}
	ie.AccessIdentity = aper.BitString{Bytes: tmp_AccessIdentity.Value.Bytes, NumBits: tmp_AccessIdentity.Value.NumBits}
	return
}
