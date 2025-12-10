package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ResourceSetTypeSemiPersistent struct {
	SemiPersistentSet SemiPersistentSet `madatory,valueExt`
	// IEExtensions *ResourceSetTypeSemiPersistentExtIEs `optional`
}

func (ie *ResourceSetTypeSemiPersistent) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if err = w.WriteBits(optionals, 1); err != nil {
		return
	}
	if err = ie.SemiPersistentSet.Encode(w); err != nil {
		err = utils.WrapError("Encode SemiPersistentSet", err)
		return
	}
	return
}

func (ie *ResourceSetTypeSemiPersistent) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.SemiPersistentSet.Decode(r); err != nil {
		err = utils.WrapError("Read SemiPersistentSet", err)
		return
	}
	return
}
