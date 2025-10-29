package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SNSSAIAvailableCapacityItem struct {
	SNSSAI                              SNSSAI `mandatory`
	SliceAvailableCapacityValueDownlink *int64 `lb:0,ub:100,optional`
	SliceAvailableCapacityValueUplink   *int64 `lb:0,ub:100,optional`
	// IEExtensions * `optional`
}

func (ie *SNSSAIAvailableCapacityItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.SliceAvailableCapacityValueDownlink != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.SliceAvailableCapacityValueUplink != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if err = ie.SNSSAI.Encode(w); err != nil {
		err = utils.WrapError("Encode SNSSAI", err)
		return
	}
	if ie.SliceAvailableCapacityValueDownlink != nil {
		tmp_SliceAvailableCapacityValueDownlink := NewINTEGER(*ie.SliceAvailableCapacityValueDownlink, aper.Constraint{Lb: 0, Ub: 100}, false)
		if err = tmp_SliceAvailableCapacityValueDownlink.Encode(w); err != nil {
			err = utils.WrapError("Encode SliceAvailableCapacityValueDownlink", err)
			return
		}
	}
	if ie.SliceAvailableCapacityValueUplink != nil {
		tmp_SliceAvailableCapacityValueUplink := NewINTEGER(*ie.SliceAvailableCapacityValueUplink, aper.Constraint{Lb: 0, Ub: 100}, false)
		if err = tmp_SliceAvailableCapacityValueUplink.Encode(w); err != nil {
			err = utils.WrapError("Encode SliceAvailableCapacityValueUplink", err)
			return
		}
	}
	return
}
func (ie *SNSSAIAvailableCapacityItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if err = ie.SNSSAI.Decode(r); err != nil {
		err = utils.WrapError("Read SNSSAI", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_SliceAvailableCapacityValueDownlink := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 100},
			ext: false,
		}
		if err = tmp_SliceAvailableCapacityValueDownlink.Decode(r); err != nil {
			err = utils.WrapError("Read SliceAvailableCapacityValueDownlink", err)
			return
		}
		ie.SliceAvailableCapacityValueDownlink = (*int64)(&tmp_SliceAvailableCapacityValueDownlink.Value)
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_SliceAvailableCapacityValueUplink := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 100},
			ext: false,
		}
		if err = tmp_SliceAvailableCapacityValueUplink.Decode(r); err != nil {
			err = utils.WrapError("Read SliceAvailableCapacityValueUplink", err)
			return
		}
		ie.SliceAvailableCapacityValueUplink = (*int64)(&tmp_SliceAvailableCapacityValueUplink.Value)
	}
	return
}
