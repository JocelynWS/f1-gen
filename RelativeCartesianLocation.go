package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type RelativeCartesianLocation struct {
	XYZUnit             XYZUnit             `madatory,valExt`
	XValue              int64               `lb:-65536,ub:65535,madatory`
	YValue              int64               `lb:-65536,ub:65535,madatory`
	ZValue              int64               `lb:-32768,ub:32767,madatory`
	LocationUncertainty LocationUncertainty `madatory`
	// IEExtensions *RelativeCartesianLocationExtIEs `optional`
}

func (ie *RelativeCartesianLocation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if err = w.WriteBits(optionals, 1); err != nil {
		return
	}
	if err = ie.XYZUnit.Encode(w); err != nil {
		err = utils.WrapError("Encode XYZUnit", err)
		return
	}
	tmp_XValue := NewINTEGER(ie.XValue, aper.Constraint{Lb: -65536, Ub: 65535}, false)
	if err = tmp_XValue.Encode(w); err != nil {
		err = utils.WrapError("Encode XValue", err)
		return
	}
	tmp_YValue := NewINTEGER(ie.YValue, aper.Constraint{Lb: -65536, Ub: 65535}, false)
	if err = tmp_YValue.Encode(w); err != nil {
		err = utils.WrapError("Encode YValue", err)
		return
	}
	tmp_ZValue := NewINTEGER(ie.ZValue, aper.Constraint{Lb: -32768, Ub: 32767}, false)
	if err = tmp_ZValue.Encode(w); err != nil {
		err = utils.WrapError("Encode ZValue", err)
		return
	}
	if err = ie.LocationUncertainty.Encode(w); err != nil {
		err = utils.WrapError("Encode LocationUncertainty", err)
		return
	}
	return
}

func (ie *RelativeCartesianLocation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.XYZUnit.Decode(r); err != nil {
		err = utils.WrapError("Read XYZUnit", err)
		return
	}
	tmp_XValue := INTEGER{c: aper.Constraint{Lb: -65536, Ub: 65535}}
	if err = tmp_XValue.Decode(r); err != nil {
		err = utils.WrapError("Read XValue", err)
		return
	}
	ie.XValue = int64(tmp_XValue.Value)
	tmp_YValue := INTEGER{c: aper.Constraint{Lb: -65536, Ub: 65535}}
	if err = tmp_YValue.Decode(r); err != nil {
		err = utils.WrapError("Read YValue", err)
		return
	}
	ie.YValue = int64(tmp_YValue.Value)
	tmp_ZValue := INTEGER{c: aper.Constraint{Lb: -32768, Ub: 32767}}
	if err = tmp_ZValue.Decode(r); err != nil {
		err = utils.WrapError("Read ZValue", err)
		return
	}
	ie.ZValue = int64(tmp_ZValue.Value)
	if err = ie.LocationUncertainty.Decode(r); err != nil {
		err = utils.WrapError("Read LocationUncertainty", err)
		return
	}
	return
}
