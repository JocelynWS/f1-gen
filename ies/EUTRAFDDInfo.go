package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type EUTRAFDDInfo struct {
	ULOffsetToPointA int64 `lb:0,ub:2199,mandatory,valueExt`
	DLOffsetToPointA int64 `lb:0,ub:2199,mandatory,valueExt`
	// IEExtensions * `optional`
}

func (ie *EUTRAFDDInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_ULOffsetToPointA := NewINTEGER(ie.ULOffsetToPointA, aper.Constraint{Lb: 0, Ub: 2199}, true)
	if err = tmp_ULOffsetToPointA.Encode(w); err != nil {
		err = utils.WrapError("Encode ULOffsetToPointA", err)
		return
	}
	tmp_DLOffsetToPointA := NewINTEGER(ie.DLOffsetToPointA, aper.Constraint{Lb: 0, Ub: 2199}, true)
	if err = tmp_DLOffsetToPointA.Encode(w); err != nil {
		err = utils.WrapError("Encode DLOffsetToPointA", err)
		return
	}
	return
}
func (ie *EUTRAFDDInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_ULOffsetToPointA := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 2199},
		ext: true,
	}
	if err = tmp_ULOffsetToPointA.Decode(r); err != nil {
		err = utils.WrapError("Read ULOffsetToPointA", err)
		return
	}
	ie.ULOffsetToPointA = int64(tmp_ULOffsetToPointA.Value)
	tmp_DLOffsetToPointA := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 2199},
		ext: true,
	}
	if err = tmp_DLOffsetToPointA.Decode(r); err != nil {
		err = utils.WrapError("Read DLOffsetToPointA", err)
		return
	}
	ie.DLOffsetToPointA = int64(tmp_DLOffsetToPointA.Value)
	return
}
