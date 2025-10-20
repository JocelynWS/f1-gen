package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DRXCycle struct {
	LongDRXCycleLength  LongDRXCycleLength   `mandatory`
	ShortDRXCycleLength *ShortDRXCycleLength `optional`
	ShortDRXCycleTimer  *int64               `lb:1,ub:16,optional`
	// IEExtensions * `optional`
}

func (ie *DRXCycle) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.ShortDRXCycleLength != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.ShortDRXCycleTimer != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if err = ie.LongDRXCycleLength.Encode(w); err != nil {
		err = utils.WrapError("Encode LongDRXCycleLength", err)
		return
	}
	if ie.ShortDRXCycleLength != nil {
		if err = ie.ShortDRXCycleLength.Encode(w); err != nil {
			err = utils.WrapError("Encode ShortDRXCycleLength", err)
			return
		}
	}
	if ie.ShortDRXCycleTimer != nil {
		tmp_ShortDRXCycleTimer := NewINTEGER(*ie.ShortDRXCycleTimer, aper.Constraint{Lb: 1, Ub: 16}, false)
		if err = tmp_ShortDRXCycleTimer.Encode(w); err != nil {
			err = utils.WrapError("Encode ShortDRXCycleTimer", err)
			return
		}
	}
	return
}
func (ie *DRXCycle) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if err = ie.LongDRXCycleLength.Decode(r); err != nil {
		err = utils.WrapError("Read LongDRXCycleLength", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(ShortDRXCycleLength)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read ShortDRXCycleLength", err)
			return
		}
		ie.ShortDRXCycleLength = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_ShortDRXCycleTimer := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 16},
			ext: false,
		}
		if err = tmp_ShortDRXCycleTimer.Decode(r); err != nil {
			err = utils.WrapError("Read ShortDRXCycleTimer", err)
			return
		}
		ie.ShortDRXCycleTimer = (*int64)(&tmp_ShortDRXCycleTimer.Value)
	}
	return
}
