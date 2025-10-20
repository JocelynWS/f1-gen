package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type NumDLULSymbols struct {
	NumDLSymbols int64 `lb:0,ub:13,mandatory,valExt`
	NumULSymbols int64 `lb:0,ub:13,mandatory,valExt`
	// IEExtensions * `optional`
}

func (ie *NumDLULSymbols) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_NumDLSymbols := NewINTEGER(ie.NumDLSymbols, aper.Constraint{Lb: 0, Ub: 13}, true)
	if err = tmp_NumDLSymbols.Encode(w); err != nil {
		err = utils.WrapError("Encode NumDLSymbols", err)
		return
	}
	tmp_NumULSymbols := NewINTEGER(ie.NumULSymbols, aper.Constraint{Lb: 0, Ub: 13}, true)
	if err = tmp_NumULSymbols.Encode(w); err != nil {
		err = utils.WrapError("Encode NumULSymbols", err)
		return
	}
	return
}
func (ie *NumDLULSymbols) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_NumDLSymbols := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 13},
		ext: true,
	}
	if err = tmp_NumDLSymbols.Decode(r); err != nil {
		err = utils.WrapError("Read NumDLSymbols", err)
		return
	}
	ie.NumDLSymbols = int64(tmp_NumDLSymbols.Value)
	tmp_NumULSymbols := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 13},
		ext: true,
	}
	if err = tmp_NumULSymbols.Decode(r); err != nil {
		err = utils.WrapError("Read NumULSymbols", err)
		return
	}
	ie.NumULSymbols = int64(tmp_NumULSymbols.Value)
	return
}
