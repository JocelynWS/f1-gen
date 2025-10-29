package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type Dynamic5QIDescriptor struct {
	QoSPriorityLevel   int64             `lb:1,ub:127,mandatory`
	PacketDelayBudget  PacketDelayBudget `mandatory`
	PacketErrorRate    PacketErrorRate   `mandatory`
	FiveQI             *int64            `lb:0,ub:255,optional,valExt`
	DelayCritical      *DelayCritical    `optional`
	AveragingWindow    *int64            `lb:0,ub:4095,optional`
	MaxDataBurstVolume *int64            `lb:0,ub:4095,optional`
	// IEExtensions *Dynamic5QIDescriptorExtIEs `optional`
}

func (ie *Dynamic5QIDescriptor) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.FiveQI != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.DelayCritical != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.AveragingWindow != nil {
		aper.SetBit(optionals, 3)
	}
	if ie.MaxDataBurstVolume != nil {
		aper.SetBit(optionals, 4)
	}
	w.WriteBits(optionals, 5)

	// Encode mandatory fields
	tmpQoS := NewINTEGER(ie.QoSPriorityLevel, aper.Constraint{Lb: 1, Ub: 127}, false)
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
	if ie.FiveQI != nil {
		tmp := NewINTEGER(*ie.FiveQI, aper.Constraint{Lb: 0, Ub: 255}, true)
		if err = tmp.Encode(w); err != nil {
			return utils.WrapError("Encode FiveQI", err)
		}
	}

	if ie.DelayCritical != nil {
		tmp := NewENUMERATED(int64(ie.DelayCritical.Value), aper.Constraint{Lb: 0, Ub: 1}, false)
		if err = tmp.Encode(w); err != nil {
			return utils.WrapError("Encode DelayCritical", err)
		}
	}

	if ie.AveragingWindow != nil {
		tmp := NewINTEGER(*ie.AveragingWindow, aper.Constraint{Lb: 0, Ub: 4095}, false)
		if err = tmp.Encode(w); err != nil {
			return utils.WrapError("Encode AveragingWindow", err)
		}
	}

	if ie.MaxDataBurstVolume != nil {
		tmp := NewINTEGER(*ie.MaxDataBurstVolume, aper.Constraint{Lb: 0, Ub: 4095}, false)
		if err = tmp.Encode(w); err != nil {
			return utils.WrapError("Encode MaxDataBurstVolume", err)
		}
	}

	return
}

func (ie *Dynamic5QIDescriptor) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	optionals, err := r.ReadBits(5)
	if err != nil {
		return
	}

	// Decode mandatory fields
	tmpQoS := INTEGER{c: aper.Constraint{Lb: 1, Ub: 127}, ext: false}
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
	if aper.IsBitSet(optionals, 1) {
		tmp := INTEGER{c: aper.Constraint{Lb: 0, Ub: 255}, ext: true}
		if err = tmp.Decode(r); err != nil {
			return utils.WrapError("Read FiveQI", err)
		}
		ie.FiveQI = (*int64)(&tmp.Value)
	}

	if aper.IsBitSet(optionals, 2) {
		tmp := ENUMERATED{c: aper.Constraint{Lb: 0, Ub: 1}, ext: false}
		if err = tmp.Decode(r); err != nil {
			return utils.WrapError("Read DelayCritical", err)
		}
		ie.DelayCritical = &DelayCritical{Value: aper.Enumerated(tmp.Value)}
	}

	if aper.IsBitSet(optionals, 3) {
		tmp := INTEGER{c: aper.Constraint{Lb: 0, Ub: 4095}, ext: false}
		if err = tmp.Decode(r); err != nil {
			return utils.WrapError("Read AveragingWindow", err)
		}
		ie.AveragingWindow = (*int64)(&tmp.Value)
	}

	if aper.IsBitSet(optionals, 4) {
		tmp := INTEGER{c: aper.Constraint{Lb: 0, Ub: 4095}, ext: false}
		if err = tmp.Decode(r); err != nil {
			return utils.WrapError("Read MaxDataBurstVolume", err)
		}
		ie.MaxDataBurstVolume = (*int64)(&tmp.Value)
	}

	return
}
