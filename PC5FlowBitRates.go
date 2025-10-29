package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PC5FlowBitRates struct {
	GuaranteedFlowBitRate int64 `lb:0,ub:4000000000000,mandatory,valExt`
	MaximumFlowBitRate    int64 `lb:0,ub:4000000000000,mandatory,valExt`
	// IEExtensions * `optional`
}

func (ie *PC5FlowBitRates) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_GuaranteedFlowBitRate := NewINTEGER(ie.GuaranteedFlowBitRate, aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
	if err = tmp_GuaranteedFlowBitRate.Encode(w); err != nil {
		err = utils.WrapError("Encode GuaranteedFlowBitRate", err)
		return
	}
	tmp_MaximumFlowBitRate := NewINTEGER(ie.MaximumFlowBitRate, aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
	if err = tmp_MaximumFlowBitRate.Encode(w); err != nil {
		err = utils.WrapError("Encode MaximumFlowBitRate", err)
		return
	}
	return
}
func (ie *PC5FlowBitRates) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_GuaranteedFlowBitRate := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: true,
	}
	if err = tmp_GuaranteedFlowBitRate.Decode(r); err != nil {
		err = utils.WrapError("Read GuaranteedFlowBitRate", err)
		return
	}
	ie.GuaranteedFlowBitRate = int64(tmp_GuaranteedFlowBitRate.Value)
	tmp_MaximumFlowBitRate := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: true,
	}
	if err = tmp_MaximumFlowBitRate.Decode(r); err != nil {
		err = utils.WrapError("Read MaximumFlowBitRate", err)
		return
	}
	ie.MaximumFlowBitRate = int64(tmp_MaximumFlowBitRate.Value)
	return
}
