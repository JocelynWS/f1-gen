package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ULAoA struct {
	AzimuthAoA             int64                   `lb:0,ub:3599,mandatory`
	ZenithAoA              *int64                  `lb:0,ub:1799,optional`
	LCSToGCSTranslationAoA *LCSToGCSTranslationAoA `optional`
	//IEExtensions ProtocolExtensionContainer mandatory
}

func (ie *ULAoA) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.ZenithAoA != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.LCSToGCSTranslationAoA != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 2)

	tmp_AzimuthAoA := NewINTEGER(ie.AzimuthAoA, aper.Constraint{Lb: 0, Ub: 3599}, false)
	if err = tmp_AzimuthAoA.Encode(w); err != nil {
		err = utils.WrapError("Encode AzimuthAoA", err)
		return
	}

	if ie.ZenithAoA != nil {
		tmp_ZenithAoA := NewINTEGER(*ie.ZenithAoA, aper.Constraint{Lb: 0, Ub: 1799}, false)
		if err = tmp_ZenithAoA.Encode(w); err != nil {
			err = utils.WrapError("Encode ZenithAoA", err)
			return
		}
	}

	if ie.LCSToGCSTranslationAoA != nil {
		if err = ie.LCSToGCSTranslationAoA.Encode(w); err != nil {
			err = utils.WrapError("Encode LCSToGCSTranslationAoA", err)
			return
		}
	}

	return
}

func (ie *ULAoA) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}

	tmp_AzimuthAoA := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 3599},
		ext: false,
	}
	if err = tmp_AzimuthAoA.Decode(r); err != nil {
		err = utils.WrapError("Read AzimuthAoA", err)
		return
	}
	ie.AzimuthAoA = int64(tmp_AzimuthAoA.Value)

	if aper.IsBitSet(optionals, 1) {
		tmp_ZenithAoA := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 1799},
			ext: false,
		}
		if err = tmp_ZenithAoA.Decode(r); err != nil {
			err = utils.WrapError("Read ZenithAoA", err)
			return
		}
		ie.ZenithAoA = (*int64)(&tmp_ZenithAoA.Value)
	}

	if aper.IsBitSet(optionals, 2) {
		tmp := new(LCSToGCSTranslationAoA)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read LCSToGCSTranslationAoA", err)
			return
		}
		ie.LCSToGCSTranslationAoA = tmp
	}

	return
}
