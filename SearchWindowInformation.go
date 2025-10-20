package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SearchWindowInformation struct {
	ExpectedPropagationDelay int64 `lb:3841,ub:3841,mandatory,valExt`
	DelayUncertainty         int64 `lb:1,ub:246,mandatory,valExt`
	// IEExtensions * `optional`
}

func (ie *SearchWindowInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_ExpectedPropagationDelay := NewINTEGER(ie.ExpectedPropagationDelay, aper.Constraint{Lb: 3841, Ub: 3841}, true)
	if err = tmp_ExpectedPropagationDelay.Encode(w); err != nil {
		err = utils.WrapError("Encode ExpectedPropagationDelay", err)
		return
	}
	tmp_DelayUncertainty := NewINTEGER(ie.DelayUncertainty, aper.Constraint{Lb: 1, Ub: 246}, true)
	if err = tmp_DelayUncertainty.Encode(w); err != nil {
		err = utils.WrapError("Encode DelayUncertainty", err)
		return
	}
	return
}
func (ie *SearchWindowInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_ExpectedPropagationDelay := INTEGER{
		c:   aper.Constraint{Lb: 3841, Ub: 3841},
		ext: true,
	}
	if err = tmp_ExpectedPropagationDelay.Decode(r); err != nil {
		err = utils.WrapError("Read ExpectedPropagationDelay", err)
		return
	}
	ie.ExpectedPropagationDelay = int64(tmp_ExpectedPropagationDelay.Value)
	tmp_DelayUncertainty := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 246},
		ext: true,
	}
	if err = tmp_DelayUncertainty.Decode(r); err != nil {
		err = utils.WrapError("Read DelayUncertainty", err)
		return
	}
	ie.DelayUncertainty = int64(tmp_DelayUncertainty.Value)
	return
}
