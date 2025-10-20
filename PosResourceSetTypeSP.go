package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PosResourceSetTypeSP struct {
	PosSemiPersistentSet PosSemiPersistentSet `madatory,valExt`
	// IEExtensions *PosResourceSetTypeSPExtIEs `optional`
}

func (ie *PosResourceSetTypeSP) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if err = w.WriteBits(optionals, 1); err != nil {
		return
	}
	if err = ie.PosSemiPersistentSet.Encode(w); err != nil {
		err = utils.WrapError("Encode PosSemiPersistentSet", err)
		return
	}
	return
}

func (ie *PosResourceSetTypeSP) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.PosSemiPersistentSet.Decode(r); err != nil {
		err = utils.WrapError("Read PosSemiPersistentSet", err)
		return
	}
	return
}
