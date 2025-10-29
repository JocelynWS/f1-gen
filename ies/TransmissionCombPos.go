package ies

import "github.com/lvdund/ngap/aper"

// CHOICE indicator
const (
	TransmissionCombPosPresentNothing uint64 = iota
	TransmissionCombPosPresentN2
	TransmissionCombPosPresentN4
	TransmissionCombPosPresentN8
	TransmissionCombPosPresentChoiceExt
)

// Struct N2
type TransmissionCombPosN2 struct {
	CombOffset  int64 `lb:0,ub:1,mandatory"`
	CyclicShift int64 `lb:0,ub:7,mandatory"`
}

func (n2 *TransmissionCombPosN2) Encode(w *aper.AperWriter) error {
	tmp := NewINTEGER(n2.CombOffset, aper.Constraint{Lb: 0, Ub: 1}, false)
	if err := tmp.Encode(w); err != nil {
		return err
	}
	tmp2 := NewINTEGER(n2.CyclicShift, aper.Constraint{Lb: 0, Ub: 7}, false)
	if err := tmp2.Encode(w); err != nil {
		return err
	}
	return nil
}

func (n2 *TransmissionCombPosN2) Decode(r *aper.AperReader) error {
	var tmp INTEGER
	tmp.c = aper.Constraint{Lb: 0, Ub: 1}
	if err := tmp.Decode(r); err != nil {
		return err
	}
	n2.CombOffset = int64(tmp.Value)

	var tmp2 INTEGER
	tmp2.c = aper.Constraint{Lb: 0, Ub: 7}
	if err := tmp2.Decode(r); err != nil {
		return err
	}
	n2.CyclicShift = int64(tmp2.Value)

	return nil
}

// Struct N4
type TransmissionCombPosN4 struct {
	CombOffset  int64 `lb:0,ub:3,mandatory"`
	CyclicShift int64 `lb:0,ub:11,mandatory"`
}

func (n4 *TransmissionCombPosN4) Encode(w *aper.AperWriter) error {
	tmp := NewINTEGER(n4.CombOffset, aper.Constraint{Lb: 0, Ub: 3}, false)
	if err := tmp.Encode(w); err != nil {
		return err
	}
	tmp2 := NewINTEGER(n4.CyclicShift, aper.Constraint{Lb: 0, Ub: 11}, false)
	if err := tmp2.Encode(w); err != nil {
		return err
	}
	return nil
}

func (n4 *TransmissionCombPosN4) Decode(r *aper.AperReader) error {
	var tmp INTEGER
	tmp.c = aper.Constraint{Lb: 0, Ub: 3}
	if err := tmp.Decode(r); err != nil {
		return err
	}
	n4.CombOffset = int64(tmp.Value)

	var tmp2 INTEGER
	tmp2.c = aper.Constraint{Lb: 0, Ub: 11}
	if err := tmp2.Decode(r); err != nil {
		return err
	}
	n4.CyclicShift = int64(tmp2.Value)

	return nil
}

// Struct N8
type TransmissionCombPosN8 struct {
	CombOffset  int64 `lb:0,ub:7,mandatory"`
	CyclicShift int64 `lb:0,ub:5,mandatory"`
}

func (n8 *TransmissionCombPosN8) Encode(w *aper.AperWriter) error {
	tmp := NewINTEGER(n8.CombOffset, aper.Constraint{Lb: 0, Ub: 7}, false)
	if err := tmp.Encode(w); err != nil {
		return err
	}
	tmp2 := NewINTEGER(n8.CyclicShift, aper.Constraint{Lb: 0, Ub: 5}, false)
	if err := tmp2.Encode(w); err != nil {
		return err
	}
	return nil
}

func (n8 *TransmissionCombPosN8) Decode(r *aper.AperReader) error {
	var tmp INTEGER
	tmp.c = aper.Constraint{Lb: 0, Ub: 7}
	if err := tmp.Decode(r); err != nil {
		return err
	}
	n8.CombOffset = int64(tmp.Value)

	var tmp2 INTEGER
	tmp2.c = aper.Constraint{Lb: 0, Ub: 5}
	if err := tmp2.Decode(r); err != nil {
		return err
	}
	n8.CyclicShift = int64(tmp2.Value)

	return nil
}

// CHOICE chính
type TransmissionCombPos struct {
	Choice uint64
	N2     *TransmissionCombPosN2
	N4     *TransmissionCombPosN4
	N8     *TransmissionCombPosN8
	// ChoiceExtension *TransmissionCombPosExtIEs
}

func (ie *TransmissionCombPos) Encode(w *aper.AperWriter) error {
	if err := w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TransmissionCombPosPresentN2:
		return ie.N2.Encode(w)
	case TransmissionCombPosPresentN4:
		return ie.N4.Encode(w)
	case TransmissionCombPosPresentN8:
		return ie.N8.Encode(w)
	}
	return nil
}

func (ie *TransmissionCombPos) Decode(r *aper.AperReader) error {
	choice, err := r.ReadChoice(3, false)
	if err != nil {
		return err
	}
	ie.Choice = choice

	switch ie.Choice {
	case TransmissionCombPosPresentN2:
		n2 := &TransmissionCombPosN2{}
		if err := n2.Decode(r); err != nil {
			return err
		}
		ie.N2 = n2
	case TransmissionCombPosPresentN4:
		n4 := &TransmissionCombPosN4{}
		if err := n4.Decode(r); err != nil {
			return err
		}
		ie.N4 = n4
	case TransmissionCombPosPresentN8:
		n8 := &TransmissionCombPosN8{}
		if err := n8.Decode(r); err != nil {
			return err
		}
		ie.N8 = n8
	}
	return nil
}
