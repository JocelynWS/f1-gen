package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ResourceSetTypePeriodic struct {
	PeriodicSet PeriodicSet `madatory,valueExt`
	// IEExtensions *ResourceSetTypePeriodicExtIEs `optional`
}

func (ie *ResourceSetTypePeriodic) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if err = w.WriteBits(optionals, 1); err != nil {
		return
	}
	if err = ie.PeriodicSet.Encode(w); err != nil {
		err = utils.WrapError("Encode PeriodicSet", err)
		return
	}
	return
}

func (ie *ResourceSetTypePeriodic) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.PeriodicSet.Decode(r); err != nil {
		err = utils.WrapError("Read PeriodicSet", err)
		return
	}
	return
}
