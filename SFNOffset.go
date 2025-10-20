package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SFNOffset struct {
	SFNTimeOffset aper.BitString `lb:24,ub:24,mandatory`
	// IEExtensions * `optional`
}

func (ie *SFNOffset) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_SFNTimeOffset := NewBITSTRING(ie.SFNTimeOffset, aper.Constraint{Lb: 24, Ub: 24}, false)
	if err = tmp_SFNTimeOffset.Encode(w); err != nil {
		err = utils.WrapError("Encode SFNTimeOffset", err)
		return
	}
	return
}
func (ie *SFNOffset) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_SFNTimeOffset := BITSTRING{
		c:   aper.Constraint{Lb: 24, Ub: 24},
		ext: false,
	}
	if err = tmp_SFNTimeOffset.Decode(r); err != nil {
		err = utils.WrapError("Read SFNTimeOffset", err)
		return
	}
	ie.SFNTimeOffset = aper.BitString{Bytes: tmp_SFNTimeOffset.Value.Bytes, NumBits: tmp_SFNTimeOffset.Value.NumBits}
	return
}
