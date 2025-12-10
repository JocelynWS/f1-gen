package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type CellMeasurementResultItem struct {
	CellID                          NRCGI                            `mandatory`
	RadioResourceStatus             *RadioResourceStatus             `optional`
	CompositeAvailableCapacityGroup *CompositeAvailableCapacityGroup `optional`
	SliceAvailableCapacity          *SliceAvailableCapacity          `optional`
	NumberofActiveUEs               *int64                           `lb:0,ub:16777215,optional,valueExt`
	// IEExtensions * `optional`
}

func (ie *CellMeasurementResultItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.RadioResourceStatus != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.CompositeAvailableCapacityGroup != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.SliceAvailableCapacity != nil {
		aper.SetBit(optionals, 3)
	}
	if ie.NumberofActiveUEs != nil {
		aper.SetBit(optionals, 4)
	}
	w.WriteBits(optionals, 5)
	if err = ie.CellID.Encode(w); err != nil {
		err = utils.WrapError("Encode CellID", err)
		return
	}
	if ie.RadioResourceStatus != nil {
		if err = ie.RadioResourceStatus.Encode(w); err != nil {
			err = utils.WrapError("Encode RadioResourceStatus", err)
			return
		}
	}
	if ie.CompositeAvailableCapacityGroup != nil {
		if err = ie.CompositeAvailableCapacityGroup.Encode(w); err != nil {
			err = utils.WrapError("Encode CompositeAvailableCapacityGroup", err)
			return
		}
	}
	if ie.SliceAvailableCapacity != nil {
		if err = ie.SliceAvailableCapacity.Encode(w); err != nil {
			err = utils.WrapError("Encode SliceAvailableCapacity", err)
			return
		}
	}
	if ie.NumberofActiveUEs != nil {
		tmp_NumberofActiveUEs := NewINTEGER(*ie.NumberofActiveUEs, aper.Constraint{Lb: 0, Ub: 16777215}, true)
		if err = tmp_NumberofActiveUEs.Encode(w); err != nil {
			err = utils.WrapError("Encode NumberofActiveUEs", err)
			return
		}
	}
	return
}
func (ie *CellMeasurementResultItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(5); err != nil {
		return
	}
	if err = ie.CellID.Decode(r); err != nil {
		err = utils.WrapError("Read CellID", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(RadioResourceStatus)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read RadioResourceStatus", err)
			return
		}
		ie.RadioResourceStatus = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp := new(CompositeAvailableCapacityGroup)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read CompositeAvailableCapacityGroup", err)
			return
		}
		ie.CompositeAvailableCapacityGroup = tmp
	}
	if aper.IsBitSet(optionals, 3) {
		tmp := new(SliceAvailableCapacity)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read SliceAvailableCapacity", err)
			return
		}
		ie.SliceAvailableCapacity = tmp
	}
	if aper.IsBitSet(optionals, 4) {
		tmp_NumberofActiveUEs := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 16777215},
			ext: true,
		}
		if err = tmp_NumberofActiveUEs.Decode(r); err != nil {
			err = utils.WrapError("Read NumberofActiveUEs", err)
			return
		}
		ie.NumberofActiveUEs = (*int64)(&tmp_NumberofActiveUEs.Value)
	}
	return
}
