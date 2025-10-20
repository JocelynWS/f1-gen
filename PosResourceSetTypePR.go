package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PosResourceSetTypePR struct {
	PosPeriodicSet PosPeriodicSet `madatory,valExt`
	// IEExtensions *PosResourceSetTypePRExtIEs `optional`
}

func (ie *PosResourceSetTypePR) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if err = w.WriteBits(optionals, 1); err != nil {
		return
	}
	if err = ie.PosPeriodicSet.Encode(w); err != nil {
		err = utils.WrapError("Encode PosPeriodicSet", err)
		return
	}
	return
}

func (ie *PosResourceSetTypePR) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.PosPeriodicSet.Decode(r); err != nil {
		err = utils.WrapError("Read PosPeriodicSet", err)
		return
	}
	return
}
