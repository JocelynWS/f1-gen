package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PRSInformationPos struct {
	PRSIDPos            *int64 `lb:0,ub:255,optional`
	PRSResourceSetIDPos int64  `lb:0,ub:7,mandatory`
	PRSResourceIDPos    int64  `lb:0,ub:63,mandatory`
}

func (ie *PRSInformationPos) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.PRSIDPos != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)

	if ie.PRSIDPos != nil {
		tmp := NewINTEGER(*ie.PRSIDPos, aper.Constraint{Lb: 0, Ub: 255}, false)
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode PRSIDPos", err)
			return
		}
	}

	tmp_PRSResourceSetIDPos := NewINTEGER(ie.PRSResourceSetIDPos, aper.Constraint{Lb: 0, Ub: 7}, false)
	if err = tmp_PRSResourceSetIDPos.Encode(w); err != nil {
		err = utils.WrapError("Encode PRSResourceSetIDPos", err)
		return
	}

	tmp_PRSResourceIDPos := NewINTEGER(ie.PRSResourceIDPos, aper.Constraint{Lb: 0, Ub: 63}, false)
	if err = tmp_PRSResourceIDPos.Encode(w); err != nil {
		err = utils.WrapError("Encode PRSResourceIDPos", err)
		return
	}

	return
}

func (ie *PRSInformationPos) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}

	if aper.IsBitSet(optionals, 1) {
		tmp := INTEGER{c: aper.Constraint{Lb: 0, Ub: 255}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read PRSIDPos", err)
			return
		}
		val := int64(tmp.Value)
		ie.PRSIDPos = &val
	} else {
		ie.PRSIDPos = nil
	}

	tmp_PRSResourceSetIDPos := INTEGER{c: aper.Constraint{Lb: 0, Ub: 7}, ext: false}
	if err = tmp_PRSResourceSetIDPos.Decode(r); err != nil {
		err = utils.WrapError("Read PRSResourceSetIDPos", err)
		return
	}
	ie.PRSResourceSetIDPos = int64(tmp_PRSResourceSetIDPos.Value)

	tmp_PRSResourceIDPos := INTEGER{c: aper.Constraint{Lb: 0, Ub: 63}, ext: false}
	if err = tmp_PRSResourceIDPos.Decode(r); err != nil {
		err = utils.WrapError("Read PRSResourceIDPos", err)
		return
	}
	ie.PRSResourceIDPos = int64(tmp_PRSResourceIDPos.Value)

	return
}