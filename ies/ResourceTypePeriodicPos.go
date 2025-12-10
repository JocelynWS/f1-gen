package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ResourceTypePeriodicPos struct {
	Periodicity ResourceTypePeriodicPosPeriodicity `madatory,valueExt`
	Offset      int64                              `lb:0,ub:81919,madatory,valueExt`
	// IEExtensions *ResourceTypePeriodicPosExtIEs `optional`
}

func (ie *ResourceTypePeriodicPos) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if err = w.WriteBits(optionals, 1); err != nil {
		return
	}

	if err = ie.Periodicity.Encode(w); err != nil {
		err = utils.WrapError("Encode Periodicity", err)
		return
	}
	tmp_Offset := NewINTEGER(ie.Offset, aper.Constraint{Lb: 0, Ub: 81919}, true)
	if err = tmp_Offset.Encode(w); err != nil {
		err = utils.WrapError("Encode Offset", err)
		return
	}
	return
}

func (ie *ResourceTypePeriodicPos) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}

	if err = ie.Periodicity.Decode(r); err != nil {
		err = utils.WrapError("Read Periodicity", err)
		return
	}
	tmp_Offset := INTEGER{c: aper.Constraint{Lb: 0, Ub: 81919}, ext: true}
	if err = tmp_Offset.Decode(r); err != nil {
		err = utils.WrapError("Read Offset", err)
		return
	}
	ie.Offset = int64(tmp_Offset.Value)

	return
}
