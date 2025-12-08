package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type LCStoGCSTranslation struct {
	Alpha     int64 `lb:0,ub:359,mandatory`
	AlphaFine int64 `lb:0,ub:9,mandatory`
	Beta      int64 `lb:0,ub:359,mandatory`
	BetaFine  int64 `lb:0,ub:9,mandatory`
	Gamma     int64 `lb:0,ub:359,mandatory`
	GammaFine int64 `lb:0,ub:9,mandatory`
	// IEExtensions * `optional`
}

func (ie *LCStoGCSTranslation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)

	tmp_Alpha := NewINTEGER(ie.Alpha, aper.Constraint{Lb: 0, Ub: 359}, false)
	if err = tmp_Alpha.Encode(w); err != nil {
		err = utils.WrapError("Encode Alpha", err)
		return
	}

	tmp_AlphaFine := NewINTEGER(ie.AlphaFine, aper.Constraint{Lb: 0, Ub: 9}, false)
	if err = tmp_AlphaFine.Encode(w); err != nil {
		err = utils.WrapError("Encode AlphaFine", err)
		return
	}

	tmp_Beta := NewINTEGER(ie.Beta, aper.Constraint{Lb: 0, Ub: 359}, false)
	if err = tmp_Beta.Encode(w); err != nil {
		err = utils.WrapError("Encode Beta", err)
		return
	}

	tmp_BetaFine := NewINTEGER(ie.BetaFine, aper.Constraint{Lb: 0, Ub: 9}, false)
	if err = tmp_BetaFine.Encode(w); err != nil {
		err = utils.WrapError("Encode BetaFine", err)
		return
	}

	tmp_Gamma := NewINTEGER(ie.Gamma, aper.Constraint{Lb: 0, Ub: 359}, false)
	if err = tmp_Gamma.Encode(w); err != nil {
		err = utils.WrapError("Encode Gamma", err)
		return
	}

	tmp_GammaFine := NewINTEGER(ie.GammaFine, aper.Constraint{Lb: 0, Ub: 9}, false)
	if err = tmp_GammaFine.Encode(w); err != nil {
		err = utils.WrapError("Encode GammaFine", err)
		return
	}

	return
}

func (ie *LCStoGCSTranslation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	if _, err = r.ReadBits(1); err != nil {
		return
	}

	tmp_Alpha := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 359},
		ext: false,
	}
	if err = tmp_Alpha.Decode(r); err != nil {
		err = utils.WrapError("Read Alpha", err)
		return
	}
	ie.Alpha = int64(tmp_Alpha.Value)

	tmp_AlphaFine := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 9},
		ext: false,
	}
	if err = tmp_AlphaFine.Decode(r); err != nil {
		err = utils.WrapError("Read AlphaFine", err)
		return
	}
	ie.AlphaFine = int64(tmp_AlphaFine.Value)

	tmp_Beta := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 359},
		ext: false,
	}
	if err = tmp_Beta.Decode(r); err != nil {
		err = utils.WrapError("Read Beta", err)
		return
	}
	ie.Beta = int64(tmp_Beta.Value)

	tmp_BetaFine := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 9},
		ext: false,
	}
	if err = tmp_BetaFine.Decode(r); err != nil {
		err = utils.WrapError("Read BetaFine", err)
		return
	}
	ie.BetaFine = int64(tmp_BetaFine.Value)

	tmp_Gamma := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 359},
		ext: false,
	}
	if err = tmp_Gamma.Decode(r); err != nil {
		err = utils.WrapError("Read Gamma", err)
		return
	}
	ie.Gamma = int64(tmp_Gamma.Value)

	tmp_GammaFine := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 9},
		ext: false,
	}
	if err = tmp_GammaFine.Decode(r); err != nil {
		err = utils.WrapError("Read GammaFine", err)
		return
	}
	ie.GammaFine = int64(tmp_GammaFine.Value)

	return
}
