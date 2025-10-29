package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type LocationUncertainty struct {
	HorizontalUncertainty int64 `lb:0,ub:255,mandatory`
	HorizontalConfidence  int64 `lb:0,ub:100,mandatory`
	VerticalUncertainty   int64 `lb:0,ub:255,mandatory`
	VerticalConfidence    int64 `lb:0,ub:100,mandatory`
	// IEExtensions * `optional`
}

func (ie *LocationUncertainty) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_HorizontalUncertainty := NewINTEGER(ie.HorizontalUncertainty, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_HorizontalUncertainty.Encode(w); err != nil {
		err = utils.WrapError("Encode HorizontalUncertainty", err)
		return
	}
	tmp_HorizontalConfidence := NewINTEGER(ie.HorizontalConfidence, aper.Constraint{Lb: 0, Ub: 100}, false)
	if err = tmp_HorizontalConfidence.Encode(w); err != nil {
		err = utils.WrapError("Encode HorizontalConfidence", err)
		return
	}
	tmp_VerticalUncertainty := NewINTEGER(ie.VerticalUncertainty, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_VerticalUncertainty.Encode(w); err != nil {
		err = utils.WrapError("Encode VerticalUncertainty", err)
		return
	}
	tmp_VerticalConfidence := NewINTEGER(ie.VerticalConfidence, aper.Constraint{Lb: 0, Ub: 100}, false)
	if err = tmp_VerticalConfidence.Encode(w); err != nil {
		err = utils.WrapError("Encode VerticalConfidence", err)
		return
	}
	return
}
func (ie *LocationUncertainty) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_HorizontalUncertainty := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 255},
		ext: false,
	}
	if err = tmp_HorizontalUncertainty.Decode(r); err != nil {
		err = utils.WrapError("Read HorizontalUncertainty", err)
		return
	}
	ie.HorizontalUncertainty = int64(tmp_HorizontalUncertainty.Value)
	tmp_HorizontalConfidence := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 100},
		ext: false,
	}
	if err = tmp_HorizontalConfidence.Decode(r); err != nil {
		err = utils.WrapError("Read HorizontalConfidence", err)
		return
	}
	ie.HorizontalConfidence = int64(tmp_HorizontalConfidence.Value)
	tmp_VerticalUncertainty := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 255},
		ext: false,
	}
	if err = tmp_VerticalUncertainty.Decode(r); err != nil {
		err = utils.WrapError("Read VerticalUncertainty", err)
		return
	}
	ie.VerticalUncertainty = int64(tmp_VerticalUncertainty.Value)
	tmp_VerticalConfidence := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 100},
		ext: false,
	}
	if err = tmp_VerticalConfidence.Decode(r); err != nil {
		err = utils.WrapError("Read VerticalConfidence", err)
		return
	}
	ie.VerticalConfidence = int64(tmp_VerticalConfidence.Value)
	return
}
