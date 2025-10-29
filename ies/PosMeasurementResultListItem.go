package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PosMeasurementResultListItem struct {
	PosMeasurementResult []PosMeasurementResultItem `lb:1,ub:maxnoofPosMeas,mandatory,valExt`
	TRPID                int64                      `lb:0,ub:4095,mandatory,valExt`
	// IEExtensions * `optional`
}

func (ie *PosMeasurementResultListItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if len(ie.PosMeasurementResult) > 0 {
		tmp := Sequence[*PosMeasurementResultItem]{
			Value: []*PosMeasurementResultItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofPosMeas},
			ext:   true,
		}
		for _, i := range ie.PosMeasurementResult {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode PosMeasurementResult", err)
			return
		}
	} else {
		err = utils.WrapError("PosMeasurementResult is nil", err)
		return
	}
	tmp_TRPID := NewINTEGER(ie.TRPID, aper.Constraint{Lb: 0, Ub: 4095}, true)
	if err = tmp_TRPID.Encode(w); err != nil {
		err = utils.WrapError("Encode TRPID", err)
		return
	}
	return
}
func (ie *PosMeasurementResultListItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_PosMeasurementResult := Sequence[*PosMeasurementResultItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofPosMeas},
		ext: true,
	}
	fn := func() *PosMeasurementResultItem { return new(PosMeasurementResultItem) }
	if err = tmp_PosMeasurementResult.Decode(r, fn); err != nil {
		err = utils.WrapError("Read PosMeasurementResult", err)
		return
	}
	ie.PosMeasurementResult = []PosMeasurementResultItem{}
	for _, i := range tmp_PosMeasurementResult.Value {
		ie.PosMeasurementResult = append(ie.PosMeasurementResult, *i)
	}
	tmp_TRPID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4095},
		ext: true,
	}
	if err = tmp_TRPID.Decode(r); err != nil {
		err = utils.WrapError("Read TRPID", err)
		return
	}
	ie.TRPID = int64(tmp_TRPID.Value)
	return
}
