package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PRSResourceQCLSourceSSB struct {
	PCINR    int64     `lb:0,ub:1007,mandatory`
	SSBIndex *SSBIndex `optional`
}

func (ie *PRSResourceQCLSourceSSB) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.SSBIndex != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)

	tmp := aper.NewINTEGER(ie.PCINR, aper.Constraint{Lb: 0, Ub: 1007}, false)
	if err = tmp.Encode(w); err != nil {
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

func (ie *PRSResourceQCLSourceSSB) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}

	tmp := aper.INTEGER{c: aper.Constraint{Lb: 0, Ub: 1007}, ext: false}
	if err = tmp.Decode(r); err != nil {
		err = utils.WrapError("Read PCINR", err)
		return
	}
	ie.PCINR = int64(tmp.Value)

	if aper.IsBitSet(optionals, 1) {
		tmpSSB := new(SSBIndex)
		if err = tmpSSB.Decode(r); err != nil {
			err = utils.WrapError("Read SSBIndex", err)
			return
		}
		ie.SSBIndex = tmpSSB
	} else {
		ie.SSBIndex = nil
	}

	return
}
