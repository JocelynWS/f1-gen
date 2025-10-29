package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type CompositeAvailableCapacity struct {
	CellCapacityClassValue *int64        `lb:1,ub:100,optional,valExt`
	CapacityValue          CapacityValue `mandatory`
	// IEExtensions * `optional`
}

func (ie *CompositeAvailableCapacity) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.CellCapacityClassValue != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.CellCapacityClassValue != nil {
		tmp_CellCapacityClassValue := NewINTEGER(*ie.CellCapacityClassValue, aper.Constraint{Lb: 1, Ub: 100}, true)
		if err = tmp_CellCapacityClassValue.Encode(w); err != nil {
			err = utils.WrapError("Encode CellCapacityClassValue", err)
			return
		}
	}
	if err = ie.CapacityValue.Encode(w); err != nil {
		err = utils.WrapError("Encode CapacityValue", err)
		return
	}
	return
}
func (ie *CompositeAvailableCapacity) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_CellCapacityClassValue := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 100},
			ext: true,
		}
		if err = tmp_CellCapacityClassValue.Decode(r); err != nil {
			err = utils.WrapError("Read CellCapacityClassValue", err)
			return
		}
		ie.CellCapacityClassValue = (*int64)(&tmp_CellCapacityClassValue.Value)
	}
	if err = ie.CapacityValue.Decode(r); err != nil {
		err = utils.WrapError("Read CapacityValue", err)
		return
	}
	return
}
