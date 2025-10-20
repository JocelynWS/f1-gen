package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SItypeItem struct {
	SItype int64 `lb:1,ub:32,mandatory`
	// IEExtensions * `optional`
}

func (ie *SItypeItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_SItype := NewINTEGER(ie.SItype, aper.Constraint{Lb: 1, Ub: 32}, false)
	if err = tmp_SItype.Encode(w); err != nil {
		err = utils.WrapError("Encode SItype", err)
		return
	}
	return
}
func (ie *SItypeItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_SItype := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 32},
		ext: false,
	}
	if err = tmp_SItype.Decode(r); err != nil {
		err = utils.WrapError("Read SItype", err)
		return
	}
	ie.SItype = int64(tmp_SItype.Value)
	return
}
