package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	DLPRSMutingPatternPresentNothing uint64 = iota
	DLPRSMutingPatternPresentTwo
	DLPRSMutingPatternPresentFour
	DLPRSMutingPatternPresentSix
	DLPRSMutingPatternPresentEight
	DLPRSMutingPatternPresentSixteen
	DLPRSMutingPatternPresentThirtyTwo
	DLPRSMutingPatternPresentChoiceExtension
)

type DLPRSMutingPattern struct {
	Choice    uint64
	Two       *aper.BitString
	Four      *aper.BitString
	Six       *aper.BitString
	Eight     *aper.BitString
	Sixteen   *aper.BitString
	ThirtyTwo *aper.BitString
	//ChoiceExtension
}

func (ie *DLPRSMutingPattern) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 6, false); err != nil {
		return
	}
	switch ie.Choice {
	case DLPRSMutingPatternPresentTwo:
		tmp := NewBITSTRING(*ie.Two, aper.Constraint{Lb: 2, Ub: 2}, false)
		err = tmp.Encode(w)
	case DLPRSMutingPatternPresentFour:
		tmp := NewBITSTRING(*ie.Four, aper.Constraint{Lb: 4, Ub: 4}, false)
		err = tmp.Encode(w)
	case DLPRSMutingPatternPresentSix:
		tmp := NewBITSTRING(*ie.Six, aper.Constraint{Lb: 6, Ub: 6}, false)
		err = tmp.Encode(w)
	case DLPRSMutingPatternPresentEight:
		tmp := NewBITSTRING(*ie.Eight, aper.Constraint{Lb: 8, Ub: 8}, false)
		err = tmp.Encode(w)
	case DLPRSMutingPatternPresentSixteen:
		tmp := NewBITSTRING(*ie.Sixteen, aper.Constraint{Lb: 16, Ub: 16}, false)
		err = tmp.Encode(w)
	case DLPRSMutingPatternPresentThirtyTwo:
		tmp := NewBITSTRING(*ie.ThirtyTwo, aper.Constraint{Lb: 32, Ub: 32}, false)
		err = tmp.Encode(w)
	}
	return
}

func (ie *DLPRSMutingPattern) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(6, false); err != nil {
		return
	}
	switch ie.Choice {
	case DLPRSMutingPatternPresentTwo:
		tmp := BITSTRING{c: aper.Constraint{Lb: 2, Ub: 2}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read Two", err)
			return
		}
		ie.Two = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	case DLPRSMutingPatternPresentFour:
		tmp := BITSTRING{c: aper.Constraint{Lb: 4, Ub: 4}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read Four", err)
			return
		}
		ie.Four = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	case DLPRSMutingPatternPresentSix:
		tmp := BITSTRING{c: aper.Constraint{Lb: 6, Ub: 6}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read Six", err)
			return
		}
		ie.Six = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	case DLPRSMutingPatternPresentEight:
		tmp := BITSTRING{c: aper.Constraint{Lb: 8, Ub: 8}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read Eight", err)
			return
		}
		ie.Eight = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	case DLPRSMutingPatternPresentSixteen:
		tmp := BITSTRING{c: aper.Constraint{Lb: 16, Ub: 16}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read Sixteen", err)
			return
		}
		ie.Sixteen = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	case DLPRSMutingPatternPresentThirtyTwo:
		tmp := BITSTRING{c: aper.Constraint{Lb: 32, Ub: 32}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read ThirtyTwo", err)
			return
		}
		ie.ThirtyTwo = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	}
	return
}
