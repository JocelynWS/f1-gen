package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DynamicPQIDescriptor struct {
	ResourceType       *ResourceType     `optional`
	QoSPriorityLevel   int64             `lb:1,ub:8,mandatory`
	PacketDelayBudget  PacketDelayBudget `mandatory`
	PacketErrorRate    PacketErrorRate   `mandatory`
	AveragingWindow    *int64            `lb:0,ub:4095,optional,valueExt`
	MaxDataBurstVolume *int64            `lb:0,ub:4095,optional,valueExt`
	// IEExtensions *DynamicPQIDescriptorExtIEs `optional`
}

func (ie *DynamicPQIDescriptor) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.ResourceType != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.AveragingWindow != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.MaxDataBurstVolume != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)

	// Encode optional ResourceType
	if ie.ResourceType != nil {
		tmp := NewENUMERATED(int64(ie.ResourceType.Value), aper.Constraint{Lb: 0, Ub: 2}, false)
		if err = tmp.Encode(w); err != nil {
			return utils.WrapError("Encode ResourceType", err)
		}
	}

	// Encode mandatory fields
	tmpQoS := NewINTEGER(ie.QoSPriorityLevel, aper.Constraint{Lb: 1, Ub: 8}, false)
	if err = tmpQoS.Encode(w); err != nil {
		return utils.WrapError("Encode QoSPriorityLevel", err)
	}

	if err = ie.PacketDelayBudget.Encode(w); err != nil {
		return utils.WrapError("Encode PacketDelayBudget", err)
	}

	if err = ie.PacketErrorRate.Encode(w); err != nil {
		return utils.WrapError("Encode PacketErrorRate", err)
	}

	// Encode optional fields
	if ie.AveragingWindow != nil {
		tmp := NewINTEGER(*ie.AveragingWindow, aper.Constraint{Lb: 0, Ub: 4095}, true)
		if err = tmp.Encode(w); err != nil {
			return utils.WrapError("Encode AveragingWindow", err)
		}
	}

	if ie.MaxDataBurstVolume != nil {
		tmp := NewINTEGER(*ie.MaxDataBurstVolume, aper.Constraint{Lb: 0, Ub: 4095}, true)
		if err = tmp.Encode(w); err != nil {
			return utils.WrapError("Encode MaxDataBurstVolume", err)
		}
	}

	return
}

func (ie *DynamicPQIDescriptor) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	optionals, err := r.ReadBits(4)
	if err != nil {
		return
	}

	// Decode optional ResourceType
	if aper.IsBitSet(optionals, 1) {
		tmp := ENUMERATED{c: aper.Constraint{Lb: 0, Ub: 2}, ext: false}
		if err = tmp.Decode(r); err != nil {
			return utils.WrapError("Read ResourceType", err)
		}
		ie.ResourceType = &ResourceType{Value: aper.Enumerated(tmp.Value)}
	}

	// Decode mandatory fields
	tmpQoS := INTEGER{c: aper.Constraint{Lb: 1, Ub: 8}, ext: false}
	if err = tmpQoS.Decode(r); err != nil {
		return utils.WrapError("Read QoSPriorityLevel", err)
	}
	ie.QoSPriorityLevel = int64(tmpQoS.Value)

	if err = ie.PacketDelayBudget.Decode(r); err != nil {
		return utils.WrapError("Read PacketDelayBudget", err)
	}

	if err = ie.PacketErrorRate.Decode(r); err != nil {
		return utils.WrapError("Read PacketErrorRate", err)
	}

	// Decode optional fields
	if aper.IsBitSet(optionals, 2) {
		tmp := INTEGER{c: aper.Constraint{Lb: 0, Ub: 4095}, ext: true}
		if err = tmp.Decode(r); err != nil {
			return utils.WrapError("Read AveragingWindow", err)
		}
		ie.AveragingWindow = (*int64)(&tmp.Value)
	}

	if aper.IsBitSet(optionals, 3) {
		tmp := INTEGER{c: aper.Constraint{Lb: 0, Ub: 4095}, ext: true}
		if err = tmp.Decode(r); err != nil {
			return utils.WrapError("Read MaxDataBurstVolume", err)
		}
		ie.MaxDataBurstVolume = (*int64)(&tmp.Value)
	}

	return
}
