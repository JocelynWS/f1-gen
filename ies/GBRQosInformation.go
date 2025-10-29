package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GBRQosInformation struct {
	ERABMaximumBitrateDL    int64 `lb:0,ub:4000000000000,mandatory,valExt`
	ERABMaximumBitrateUL    int64 `lb:0,ub:4000000000000,mandatory,valExt`
	ERABGuaranteedBitrateDL int64 `lb:0,ub:4000000000000,mandatory,valExt`
	ERABGuaranteedBitrateUL int64 `lb:0,ub:4000000000000,mandatory,valExt`
	// IEExtensions * `optional`
}

func (ie *GBRQosInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_ERABMaximumBitrateDL := NewINTEGER(ie.ERABMaximumBitrateDL, aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
	if err = tmp_ERABMaximumBitrateDL.Encode(w); err != nil {
		err = utils.WrapError("Encode ERABMaximumBitrateDL", err)
		return
	}
	tmp_ERABMaximumBitrateUL := NewINTEGER(ie.ERABMaximumBitrateUL, aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
	if err = tmp_ERABMaximumBitrateUL.Encode(w); err != nil {
		err = utils.WrapError("Encode ERABMaximumBitrateUL", err)
		return
	}
	tmp_ERABGuaranteedBitrateDL := NewINTEGER(ie.ERABGuaranteedBitrateDL, aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
	if err = tmp_ERABGuaranteedBitrateDL.Encode(w); err != nil {
		err = utils.WrapError("Encode ERABGuaranteedBitrateDL", err)
		return
	}
	tmp_ERABGuaranteedBitrateUL := NewINTEGER(ie.ERABGuaranteedBitrateUL, aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
	if err = tmp_ERABGuaranteedBitrateUL.Encode(w); err != nil {
		err = utils.WrapError("Encode ERABGuaranteedBitrateUL", err)
		return
	}
	return
}
func (ie *GBRQosInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_ERABMaximumBitrateDL := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: true,
	}
	if err = tmp_ERABMaximumBitrateDL.Decode(r); err != nil {
		err = utils.WrapError("Read ERABMaximumBitrateDL", err)
		return
	}
	ie.ERABMaximumBitrateDL = int64(tmp_ERABMaximumBitrateDL.Value)
	tmp_ERABMaximumBitrateUL := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: true,
	}
	if err = tmp_ERABMaximumBitrateUL.Decode(r); err != nil {
		err = utils.WrapError("Read ERABMaximumBitrateUL", err)
		return
	}
	ie.ERABMaximumBitrateUL = int64(tmp_ERABMaximumBitrateUL.Value)
	tmp_ERABGuaranteedBitrateDL := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: true,
	}
	if err = tmp_ERABGuaranteedBitrateDL.Decode(r); err != nil {
		err = utils.WrapError("Read ERABGuaranteedBitrateDL", err)
		return
	}
	ie.ERABGuaranteedBitrateDL = int64(tmp_ERABGuaranteedBitrateDL.Value)
	tmp_ERABGuaranteedBitrateUL := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: true,
	}
	if err = tmp_ERABGuaranteedBitrateUL.Decode(r); err != nil {
		err = utils.WrapError("Read ERABGuaranteedBitrateUL", err)
		return
	}
	ie.ERABGuaranteedBitrateUL = int64(tmp_ERABGuaranteedBitrateUL.Value)
	return
}
