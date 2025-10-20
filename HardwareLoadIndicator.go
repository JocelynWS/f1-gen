package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type HardwareLoadIndicator struct {
	DLHardwareLoadIndicator int64 `lb:0,ub:100,mandatory,valExt`
	ULHardwareLoadIndicator int64 `lb:0,ub:100,mandatory,valExt`
	// IEExtensions * `optional`
}

func (ie *HardwareLoadIndicator) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_DLHardwareLoadIndicator := NewINTEGER(ie.DLHardwareLoadIndicator, aper.Constraint{Lb: 0, Ub: 100}, true)
	if err = tmp_DLHardwareLoadIndicator.Encode(w); err != nil {
		err = utils.WrapError("Encode DLHardwareLoadIndicator", err)
		return
	}
	tmp_ULHardwareLoadIndicator := NewINTEGER(ie.ULHardwareLoadIndicator, aper.Constraint{Lb: 0, Ub: 100}, true)
	if err = tmp_ULHardwareLoadIndicator.Encode(w); err != nil {
		err = utils.WrapError("Encode ULHardwareLoadIndicator", err)
		return
	}
	return
}
func (ie *HardwareLoadIndicator) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_DLHardwareLoadIndicator := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 100},
		ext: true,
	}
	if err = tmp_DLHardwareLoadIndicator.Decode(r); err != nil {
		err = utils.WrapError("Read DLHardwareLoadIndicator", err)
		return
	}
	ie.DLHardwareLoadIndicator = int64(tmp_DLHardwareLoadIndicator.Value)
	tmp_ULHardwareLoadIndicator := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 100},
		ext: true,
	}
	if err = tmp_ULHardwareLoadIndicator.Decode(r); err != nil {
		err = utils.WrapError("Read ULHardwareLoadIndicator", err)
		return
	}
	ie.ULHardwareLoadIndicator = int64(tmp_ULHardwareLoadIndicator.Value)
	return
}
