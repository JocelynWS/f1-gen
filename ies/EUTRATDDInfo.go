package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type EUTRATDDInfo struct {
	OffsetToPointA int64 `lb:0,ub:2199,mandatory,valExt`
	// IEExtensions * `optional`
}

func (ie *EUTRATDDInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_OffsetToPointA := NewINTEGER(ie.OffsetToPointA, aper.Constraint{Lb: 0, Ub: 2199}, true)
	if err = tmp_OffsetToPointA.Encode(w); err != nil {
		err = utils.WrapError("Encode OffsetToPointA", err)
		return
	}
	return
}
func (ie *EUTRATDDInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_OffsetToPointA := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 2199},
		ext: true,
	}
	if err = tmp_OffsetToPointA.Decode(r); err != nil {
		err = utils.WrapError("Read OffsetToPointA", err)
		return
	}
	ie.OffsetToPointA = int64(tmp_OffsetToPointA.Value)
	return
}
