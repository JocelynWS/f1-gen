package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	SSBTransmissionBitmapPresentNothing uint64 = iota
	SSBTransmissionBitmapPresentShortbitmap
	SSBTransmissionBitmapPresentMediumbitmap
	SSBTransmissionBitmapPresentLongbitmap
	SSBTransmissionBitmapPresentChoiceExtension
)

type SSBTransmissionBitmap struct {
	Choice       uint64
	ShortBitmap  *aper.BitString
	MediumBitmap *aper.BitString
	LongBitmap   *aper.BitString
	//ChoiceExtension
}

func (ie *SSBTransmissionBitmap) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case SSBTransmissionBitmapPresentShortbitmap:
		tmp := NewBITSTRING(*ie.ShortBitmap, aper.Constraint{Lb: 4, Ub: 4}, false)
		err = tmp.Encode(w)
	case SSBTransmissionBitmapPresentMediumbitmap:
		tmp := NewBITSTRING(*ie.MediumBitmap, aper.Constraint{Lb: 8, Ub: 8}, false)
		err = tmp.Encode(w)
	case SSBTransmissionBitmapPresentLongbitmap:
		tmp := NewBITSTRING(*ie.LongBitmap, aper.Constraint{Lb: 64, Ub: 64}, false)
		err = tmp.Encode(w)
	}
	return
}

func (ie *SSBTransmissionBitmap) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case SSBTransmissionBitmapPresentShortbitmap:
		tmp := BITSTRING{c: aper.Constraint{Lb: 4, Ub: 4}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read ShortBitmap", err)
			return
		}
		ie.ShortBitmap = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	case SSBTransmissionBitmapPresentMediumbitmap:
		tmp := BITSTRING{c: aper.Constraint{Lb: 8, Ub: 8}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read MediumBitmap", err)
			return
		}
		ie.MediumBitmap = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	case SSBTransmissionBitmapPresentLongbitmap:
		tmp := BITSTRING{c: aper.Constraint{Lb: 64, Ub: 64}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read LongBitmap", err)
			return
		}
		ie.LongBitmap = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	}
	return
}
