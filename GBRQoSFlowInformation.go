package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GBRQoSFlowInformation struct {
	MaxFlowBitRateDownlink        int64  `lb:0,ub:4000000000000,mandatory,valExt`
	MaxFlowBitRateUplink          int64  `lb:0,ub:4000000000000,mandatory,valExt`
	GuaranteedFlowBitRateDownlink int64  `lb:0,ub:4000000000000,mandatory,valExt`
	GuaranteedFlowBitRateUplink   int64  `lb:0,ub:4000000000000,mandatory,valExt`
	MaxPacketLossRateDownlink     *int64 `lb:0,ub:1000,optional`
	MaxPacketLossRateUplink       *int64 `lb:0,ub:1000,optional`
}

func (ie *GBRQoSFlowInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.MaxPacketLossRateDownlink != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.MaxPacketLossRateUplink != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)

	// Encode mandatory fields
	tmp := []INTEGER{
		NewINTEGER(ie.MaxFlowBitRateDownlink, aper.Constraint{Lb: 0, Ub: 4000000000000}, true),
		NewINTEGER(ie.MaxFlowBitRateUplink, aper.Constraint{Lb: 0, Ub: 4000000000000}, true),
		NewINTEGER(ie.GuaranteedFlowBitRateDownlink, aper.Constraint{Lb: 0, Ub: 4000000000000}, true),
		NewINTEGER(ie.GuaranteedFlowBitRateUplink, aper.Constraint{Lb: 0, Ub: 4000000000000}, true),
	}
	for i, t := range tmp {
		if err = t.Encode(w); err != nil {
			return utils.WrapError([]string{
				"MaxFlowBitRateDownlink", "MaxFlowBitRateUplink",
				"GuaranteedFlowBitRateDownlink", "GuaranteedFlowBitRateUplink",
			}[i], err)
		}
	}

	// Optional fields
	if ie.MaxPacketLossRateDownlink != nil {
		t := NewINTEGER(*ie.MaxPacketLossRateDownlink, aper.Constraint{Lb: 0, Ub: 1000}, false)
		if err = t.Encode(w); err != nil {
			return utils.WrapError("MaxPacketLossRateDownlink", err)
		}
	}
	if ie.MaxPacketLossRateUplink != nil {
		t := NewINTEGER(*ie.MaxPacketLossRateUplink, aper.Constraint{Lb: 0, Ub: 1000}, false)
		if err = t.Encode(w); err != nil {
			return utils.WrapError("MaxPacketLossRateUplink", err)
		}
	}

	return
}

func (ie *GBRQoSFlowInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	optionals, err := r.ReadBits(3)
	if err != nil {
		return
	}

	tmp := []struct {
		field  *int64
		lb, ub int64
		ext    bool
	}{
		{&ie.MaxFlowBitRateDownlink, 0, 4000000000000, true},
		{&ie.MaxFlowBitRateUplink, 0, 4000000000000, true},
		{&ie.GuaranteedFlowBitRateDownlink, 0, 4000000000000, true},
		{&ie.GuaranteedFlowBitRateUplink, 0, 4000000000000, true},
	}

	for i, t := range tmp {
		val := INTEGER{c: aper.Constraint{Lb: t.lb, Ub: t.ub}, ext: t.ext}
		if err = val.Decode(r); err != nil {
			return utils.WrapError([]string{
				"MaxFlowBitRateDownlink", "MaxFlowBitRateUplink",
				"GuaranteedFlowBitRateDownlink", "GuaranteedFlowBitRateUplink",
			}[i], err)
		}
		*t.field = int64(val.Value)
	}

	if aper.IsBitSet(optionals, 1) {
		val := INTEGER{c: aper.Constraint{Lb: 0, Ub: 1000}, ext: false}
		if err = val.Decode(r); err != nil {
			return utils.WrapError("MaxPacketLossRateDownlink", err)
		}
		ie.MaxPacketLossRateDownlink = (*int64)(&val.Value)
	}
	if aper.IsBitSet(optionals, 2) {
		val := INTEGER{c: aper.Constraint{Lb: 0, Ub: 1000}, ext: false}
		if err = val.Decode(r); err != nil {
			return utils.WrapError("MaxPacketLossRateUplink", err)
		}
		ie.MaxPacketLossRateUplink = (*int64)(&val.Value)
	}

	return
}
