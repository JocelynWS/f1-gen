package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PRSResourceItem struct {
	PRSResourceID        int64               `lb:0,ub:63,mandatory`
	SequenceID           int64               `lb:0,ub:4095,mandatory`
	REOffset             int64               `lb:0,ub:11,mandatory`
	ResourceSlotOffset   int64               `lb:0,ub:511,mandatory`
	ResourceSymbolOffset int64               `lb:0,ub:12,mandatory`
	QCLInfo              *PRSResourceQCLInfo `optional`
}

func (ie *PRSResourceItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.QCLInfo != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)

	tmp := aper.NewINTEGER(ie.PRSResourceID, aper.Constraint{Lb: 0, Ub: 63}, false)
	if err = tmp.Encode(w); err != nil {
		err = utils.WrapError("Encode PRSResourceID", err)
		return
	}

	tmp = aper.NewINTEGER(ie.SequenceID, aper.Constraint{Lb: 0, Ub: 4095}, false)
	if err = tmp.Encode(w); err != nil {
		err = utils.WrapError("Encode SequenceID", err)
		return
	}

	tmp = aper.NewINTEGER(ie.REOffset, aper.Constraint{Lb: 0, Ub: 11}, false)
	if err = tmp.Encode(w); err != nil {
		err = utils.WrapError("Encode REOffset", err)
		return
	}

	tmp = aper.NewINTEGER(ie.ResourceSlotOffset, aper.Constraint{Lb: 0, Ub: 511}, false)
	if err = tmp.Encode(w); err != nil {
		err = utils.WrapError("Encode ResourceSlotOffset", err)
		return
	}

	tmp = aper.NewINTEGER(ie.ResourceSymbolOffset, aper.Constraint{Lb: 0, Ub: 12}, false)
	if err = tmp.Encode(w); err != nil {
		err = utils.WrapError("Encode ResourceSymbolOffset", err)
		return
	}

	if ie.QCLInfo != nil {
		if err = ie.QCLInfo.Encode(w); err != nil {
			err = utils.WrapError("Encode QCLInfo", err)
			return
		}
	}

	return
}

func (ie *PRSResourceItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}

	tmp := aper.INTEGER{c: aper.Constraint{Lb: 0, Ub: 63}, ext: false}
	if err = tmp.Decode(r); err != nil {
		err = utils.WrapError("Read PRSResourceID", err)
		return
	}
	ie.PRSResourceID = int64(tmp.Value)

	tmp = aper.INTEGER{c: aper.Constraint{Lb: 0, Ub: 4095}, ext: false}
	if err = tmp.Decode(r); err != nil {
		err = utils.WrapError("Read SequenceID", err)
		return
	}
	ie.SequenceID = int64(tmp.Value)

	tmp = aper.INTEGER{c: aper.Constraint{Lb: 0, Ub: 11}, ext: false}
	if err = tmp.Decode(r); err != nil {
		err = utils.WrapError("Read REOffset", err)
		return
	}
	ie.REOffset = int64(tmp.Value)

	tmp = aper.INTEGER{c: aper.Constraint{Lb: 0, Ub: 511}, ext: false}
	if err = tmp.Decode(r); err != nil {
		err = utils.WrapError("Read ResourceSlotOffset", err)
		return
	}
	ie.ResourceSlotOffset = int64(tmp.Value)

	tmp = aper.INTEGER{c: aper.Constraint{Lb: 0, Ub: 12}, ext: false}
	if err = tmp.Decode(r); err != nil {
		err = utils.WrapError("Read ResourceSymbolOffset", err)
		return
	}
	ie.ResourceSymbolOffset = int64(tmp.Value)

	if aper.IsBitSet(optionals, 1) {
		tmpQCL := new(PRSResourceQCLInfo)
		if err = tmpQCL.Decode(r); err != nil {
			err = utils.WrapError("Read QCLInfo", err)
			return
		}
		ie.QCLInfo = tmpQCL
	} else {
		ie.QCLInfo = nil
	}

	return
}
