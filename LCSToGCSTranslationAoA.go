package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type LCSToGCSTranslationAoA struct {
	Alpha int64 `lb:0,ub:3599,mandatory`
	Beta  int64 `lb:0,ub:3599,mandatory`
	Gamma int64 `lb:0,ub:3599,mandatory`
	// IEExtensions * `optional`
}

func (ie *LCSToGCSTranslationAoA) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_Alpha := NewINTEGER(ie.Alpha, aper.Constraint{Lb: 0, Ub: 3599}, false)
	if err = tmp_Alpha.Encode(w); err != nil {
		err = utils.WrapError("Encode Alpha", err)
		return
	}
	tmp_Beta := NewINTEGER(ie.Beta, aper.Constraint{Lb: 0, Ub: 3599}, false)
	if err = tmp_Beta.Encode(w); err != nil {
		err = utils.WrapError("Encode Beta", err)
		return
	}
	tmp_Gamma := NewINTEGER(ie.Gamma, aper.Constraint{Lb: 0, Ub: 3599}, false)
	if err = tmp_Gamma.Encode(w); err != nil {
		err = utils.WrapError("Encode Gamma", err)
		return
	}
	return
}
func (ie *LCSToGCSTranslationAoA) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_Alpha := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 3599},
		ext: false,
	}
	if err = tmp_Alpha.Decode(r); err != nil {
		err = utils.WrapError("Read Alpha", err)
		return
	}
	ie.Alpha = int64(tmp_Alpha.Value)
	tmp_Beta := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 3599},
		ext: false,
	}
	if err = tmp_Beta.Decode(r); err != nil {
		err = utils.WrapError("Read Beta", err)
		return
	}
	ie.Beta = int64(tmp_Beta.Value)
	tmp_Gamma := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 3599},
		ext: false,
	}
	if err = tmp_Gamma.Decode(r); err != nil {
		err = utils.WrapError("Read Gamma", err)
		return
	}
	ie.Gamma = int64(tmp_Gamma.Value)
	return
}
