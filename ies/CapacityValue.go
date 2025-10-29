package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type CapacityValue struct {
	CapacityValue            int64                      `lb:0,ub:100,mandatory`
	SSBAreaCapacityValueList []SSBAreaCapacityValueItem `lb:1,ub:maxnoofSSBs,optional,valExt`
	// IEExtensions * `optional`
}

func (ie *CapacityValue) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.SSBAreaCapacityValueList != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_CapacityValue := NewINTEGER(ie.CapacityValue, aper.Constraint{Lb: 0, Ub: 100}, false)
	if err = tmp_CapacityValue.Encode(w); err != nil {
		err = utils.WrapError("Encode CapacityValue", err)
		return
	}
	if len(ie.SSBAreaCapacityValueList) > 0 {
		tmp := Sequence[*SSBAreaCapacityValueItem]{
			Value: []*SSBAreaCapacityValueItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofSSBs},
			ext:   true,
		}
		for _, i := range ie.SSBAreaCapacityValueList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode SSBAreaCapacityValueList", err)
			return
		}
	}
	return
}
func (ie *CapacityValue) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_CapacityValue := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 100},
		ext: false,
	}
	if err = tmp_CapacityValue.Decode(r); err != nil {
		err = utils.WrapError("Read CapacityValue", err)
		return
	}
	ie.CapacityValue = int64(tmp_CapacityValue.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp_SSBAreaCapacityValueList := Sequence[*SSBAreaCapacityValueItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSSBs},
			ext: true,
		}
		fn := func() *SSBAreaCapacityValueItem { return new(SSBAreaCapacityValueItem) }
		if err = tmp_SSBAreaCapacityValueList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read SSBAreaCapacityValueList", err)
			return
		}
		ie.SSBAreaCapacityValueList = []SSBAreaCapacityValueItem{}
		for _, i := range tmp_SSBAreaCapacityValueList.Value {
			ie.SSBAreaCapacityValueList = append(ie.SSBAreaCapacityValueList, *i)
		}
	}
	return
}
