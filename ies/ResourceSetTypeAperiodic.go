package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ResourceSetTypeAperiodic struct {
	SRSResourceTriggerList int64 `lb:1,ub:3,mandatory`
	SlotOffset             int64 `lb:0,ub:32,mandatory`
	// IEExtensions *ProtocolExtensionContainer `optional`
}

func (ie *ResourceSetTypeAperiodic) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)

	tmp_SRSResourceTriggerList := NewINTEGER(ie.SRSResourceTriggerList, aper.Constraint{Lb: 1, Ub: 3}, false)
	if err = tmp_SRSResourceTriggerList.Encode(w); err != nil {
		err = utils.WrapError("Encode SRSResourceTriggerList", err)
		return
	}

	tmp_SlotOffset := NewINTEGER(ie.SlotOffset, aper.Constraint{Lb: 0, Ub: 32}, false)
	if err = tmp_SlotOffset.Encode(w); err != nil {
		err = utils.WrapError("Encode SlotOffset", err)
		return
	}
	return
}

func (ie *ResourceSetTypeAperiodic) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}

	tmp_SRSResourceTriggerList := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 3},
		ext: false,
	}
	if err = tmp_SRSResourceTriggerList.Decode(r); err != nil {
		err = utils.WrapError("Read SRSResourceTriggerList", err)
		return
	}
	ie.SRSResourceTriggerList = int64(tmp_SRSResourceTriggerList.Value)
	tmp_SlotOffset := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 32},
		ext: false,
	}
	if err = tmp_SlotOffset.Decode(r); err != nil {
		err = utils.WrapError("Read SlotOffset", err)
		return
	}
	ie.SlotOffset = int64(tmp_SlotOffset.Value)

	return
}
