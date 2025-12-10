package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AlternativeQoSParaSetItem struct {
	AlternativeQoSParaSetIndex int64            `lb:1,ub:8,mandatory`
	GuaranteedFlowBitRateDL    *int64           `lb:0,ub:4000000000000,optional,valueExt`
	GuaranteedFlowBitRateUL    *int64           `lb:0,ub:4000000000000,optional,valueExt`
	PacketDelayBudget          *int64           `lb:0,ub:1023,optional`
	PacketErrorRate            *PacketErrorRate `optional`
	// IEExtensions * `optional`
}

func (ie *AlternativeQoSParaSetItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.GuaranteedFlowBitRateDL != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.GuaranteedFlowBitRateUL != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.PacketDelayBudget != nil {
		aper.SetBit(optionals, 3)
	}
	if ie.PacketErrorRate != nil {
		aper.SetBit(optionals, 4)
	}
	w.WriteBits(optionals, 5)
	tmp_AlternativeQoSParaSetIndex := NewINTEGER(ie.AlternativeQoSParaSetIndex, aper.Constraint{Lb: 1, Ub: 8}, false)
	if err = tmp_AlternativeQoSParaSetIndex.Encode(w); err != nil {
		err = utils.WrapError("Encode AlternativeQoSParaSetIndex", err)
		return
	}
	if ie.GuaranteedFlowBitRateDL != nil {
		tmp_GuaranteedFlowBitRateDL := NewINTEGER(*ie.GuaranteedFlowBitRateDL, aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
		if err = tmp_GuaranteedFlowBitRateDL.Encode(w); err != nil {
			err = utils.WrapError("Encode GuaranteedFlowBitRateDL", err)
			return
		}
	}
	if ie.GuaranteedFlowBitRateUL != nil {
		tmp_GuaranteedFlowBitRateUL := NewINTEGER(*ie.GuaranteedFlowBitRateUL, aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
		if err = tmp_GuaranteedFlowBitRateUL.Encode(w); err != nil {
			err = utils.WrapError("Encode GuaranteedFlowBitRateUL", err)
			return
		}
	}
	if ie.PacketDelayBudget != nil {
		tmp_PacketDelayBudget := NewINTEGER(*ie.PacketDelayBudget, aper.Constraint{Lb: 0, Ub: 1023}, false)
		if err = tmp_PacketDelayBudget.Encode(w); err != nil {
			err = utils.WrapError("Encode PacketDelayBudget", err)
			return
		}
	}
	if ie.PacketErrorRate != nil {
		if err = ie.PacketErrorRate.Encode(w); err != nil {
			err = utils.WrapError("Encode PacketErrorRate", err)
			return
		}
	}
	return
}
func (ie *AlternativeQoSParaSetItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(5); err != nil {
		return
	}
	tmp_AlternativeQoSParaSetIndex := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 8},
		ext: false,
	}
	if err = tmp_AlternativeQoSParaSetIndex.Decode(r); err != nil {
		err = utils.WrapError("Read AlternativeQoSParaSetIndex", err)
		return
	}
	ie.AlternativeQoSParaSetIndex = int64(tmp_AlternativeQoSParaSetIndex.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp_GuaranteedFlowBitRateDL := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
			ext: true,
		}
		if err = tmp_GuaranteedFlowBitRateDL.Decode(r); err != nil {
			err = utils.WrapError("Read GuaranteedFlowBitRateDL", err)
			return
		}
		ie.GuaranteedFlowBitRateDL = (*int64)(&tmp_GuaranteedFlowBitRateDL.Value)
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_GuaranteedFlowBitRateUL := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
			ext: true,
		}
		if err = tmp_GuaranteedFlowBitRateUL.Decode(r); err != nil {
			err = utils.WrapError("Read GuaranteedFlowBitRateUL", err)
			return
		}
		ie.GuaranteedFlowBitRateUL = (*int64)(&tmp_GuaranteedFlowBitRateUL.Value)
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_PacketDelayBudget := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 1023},
			ext: false,
		}
		if err = tmp_PacketDelayBudget.Decode(r); err != nil {
			err = utils.WrapError("Read PacketDelayBudget", err)
			return
		}
		ie.PacketDelayBudget = (*int64)(&tmp_PacketDelayBudget.Value)
	}
	if aper.IsBitSet(optionals, 4) {
		tmp := new(PacketErrorRate)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read PacketErrorRate", err)
			return
		}
		ie.PacketErrorRate = tmp
	}
	return
}
