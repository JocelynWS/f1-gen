package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SlotConfigurationItem struct {
	SlotIndex         int64             `lb:0,ub:5119,mandatory,valExt`
	SymbolAllocInSlot SymbolAllocInSlot `mandatory`
	// IEExtensions * `optional`
}

func (ie *SlotConfigurationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_SlotIndex := NewINTEGER(ie.SlotIndex, aper.Constraint{Lb: 0, Ub: 5119}, true)
	if err = tmp_SlotIndex.Encode(w); err != nil {
		err = utils.WrapError("Encode SlotIndex", err)
		return
	}
	if err = ie.SymbolAllocInSlot.Encode(w); err != nil {
		err = utils.WrapError("Encode SymbolAllocInSlot", err)
		return
	}
	return
}
func (ie *SlotConfigurationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_SlotIndex := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 5119},
		ext: true,
	}
	if err = tmp_SlotIndex.Decode(r); err != nil {
		err = utils.WrapError("Read SlotIndex", err)
		return
	}
	ie.SlotIndex = int64(tmp_SlotIndex.Value)
	if err = ie.SymbolAllocInSlot.Decode(r); err != nil {
		err = utils.WrapError("Read SymbolAllocInSlot", err)
		return
	}
	return
}
