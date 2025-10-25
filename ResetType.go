package ies

import "github.com/lvdund/ngap/aper"

const (
	ResetTypePresentNothing uint64 = iota
	ResetTypePresentF1Interface
	ResetTypePresentPartOfF1Interface
)

type ResetType struct {
	Choice            uint64
	F1Interface       *ResetAll
	PartOfF1Interface []UEAssociatedLogicalF1ConnectionItemRes
	// ChoiceExtension
}

func (ie *ResetType) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case ResetTypePresentF1Interface:
		err = ie.F1Interface.Encode(w)
	case ResetTypePresentPartOfF1Interface:
		tmp := Sequence[*UEAssociatedLogicalF1ConnectionItemRes]{}
		for _, i := range ie.PartOfF1Interface {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	}
	return
}

func (ie *ResetType) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case ResetTypePresentF1Interface:
		var tmp ResetAll
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.F1Interface = &tmp
	case ResetTypePresentPartOfF1Interface:
		tmp := NewSequence[*UEAssociatedLogicalF1ConnectionItemRes](nil, aper.Constraint{Lb: 0, Ub: 65535}, false)
		fn := func() *UEAssociatedLogicalF1ConnectionItemRes { return new(UEAssociatedLogicalF1ConnectionItemRes) }
		if err = tmp.Decode(r, fn); err != nil {
			return
		}
		for _, i := range tmp.Value {
			ie.PartOfF1Interface = append(ie.PartOfF1Interface, *i)
		}
	}
	return
}
