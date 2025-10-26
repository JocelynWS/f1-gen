package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SSB struct {
	PCINR    NRPCI     `mandatory`
	SSBIndex *SSBIndex `optional`
}

func (ie *SSB) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.SSBIndex != nil {
		aper.SetBit(optionals, 0)
	}
	w.WriteBits(optionals, 1)

	if err = ie.PCINR.Encode(w); err != nil {
		err = utils.WrapError("Encode PCINR", err)
		return
	}

	if ie.SSBIndex != nil {
		if err = ie.SSBIndex.Encode(w); err != nil {
			err = utils.WrapError("Encode SSBIndex", err)
			return
		}
	}
	return
}

func (ie *SSB) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(1); err != nil {
		return
	}

	if err = ie.PCINR.Decode(r); err != nil {
		err = utils.WrapError("Decode PCINR", err)
		return
	}

	if aper.IsBitSet(optionals, 0) {
		ie.SSBIndex = new(SSBIndex)
		if err = ie.SSBIndex.Decode(r); err != nil {
			err = utils.WrapError("Decode SSBIndex", err)
			return
		}
	}
	return
}
