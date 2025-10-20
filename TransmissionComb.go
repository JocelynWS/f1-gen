package ies

import "github.com/lvdund/ngap/aper"

const (
	TransmissionCombPresentNothing uint64 = iota
	TransmissionCombPresentN2
	TransmissionCombPresentN4
	TransmissionCombPresentChoiceExt
)

type TransmissionCombN2 struct {
	CombOffset  int64 `lb:0,ub:1,mandatory"`
	CyclicShift int64 `lb:0,ub:7,mandatory"`
}

type TransmissionCombN4 struct {
	CombOffset  int64 `lb:0,ub:3,mandatory"`
	CyclicShift int64 `lb:0,ub:11,mandatory"`
}

type TransmissionComb struct {
	Choice uint64
	N2     *TransmissionCombN2
	N4     *TransmissionCombN4
	// ChoiceExtension *TransmissionCombExtIEs
}

func (ie *TransmissionComb) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case TransmissionCombPresentN2:
		err = ie.N2.Encode(w)
	case TransmissionCombPresentN4:
		err = ie.N4.Encode(w)
	case TransmissionCombPresentChoiceExt:
	}
	return
}

func (ie *TransmissionComb) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case TransmissionCombPresentN2:
		var tmp TransmissionCombN2
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.N2 = &tmp
	case TransmissionCombPresentN4:
		var tmp TransmissionCombN4
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.N4 = &tmp
	}
	return
}

func (n2 *TransmissionCombN2) Encode(w *aper.AperWriter) (err error) {
	tmp1 := aper.NewINTEGER(n2.CombOffset, aper.Constraint{Lb: 0, Ub: 1}, false)
	if err = tmp1.Encode(w); err != nil {
		return
	}
	tmp2 := aper.NewINTEGER(n2.CyclicShift, aper.Constraint{Lb: 0, Ub: 7}, false)
	if err = tmp2.Encode(w); err != nil {
		return
	}

	return
}

func (n2 *TransmissionCombN2) Decode(r *aper.AperReader) (err error) {
	tmp1 := aper.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 1}, false)
	if err = tmp1.Decode(r); err != nil {
		return
	}
	n2.CombOffset = tmp1.Value

	tmp2 := aper.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 7}, false)
	if err = tmp2.Decode(r); err != nil {
		return
	}
	n2.CyclicShift = tmp2.Value

	return
}

func (n4 *TransmissionCombN4) Encode(w *aper.AperWriter) (err error) {
	tmp1 := aper.NewINTEGER(n4.CombOffset, aper.Constraint{Lb: 0, Ub: 3}, false)
	if err = tmp1.Encode(w); err != nil {
		return
	}
	tmp2 := aper.NewINTEGER(n4.CyclicShift, aper.Constraint{Lb: 0, Ub: 11}, false)
	if err = tmp2.Encode(w); err != nil {
		return
	}
	return
}

func (n4 *TransmissionCombN4) Decode(r *aper.AperReader) (err error) {
	tmp1 := aper.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 3}, false)
	if err = tmp1.Decode(r); err != nil {
		return
	}
	n4.CombOffset = tmp1.Value

	tmp2 := aper.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 11}, false)
	if err = tmp2.Decode(r); err != nil {
		return
	}
	n4.CyclicShift = tmp2.Value

	return
}
