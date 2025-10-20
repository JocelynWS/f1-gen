package ies

import "github.com/lvdund/ngap/aper"

const (
	UACCategoryTypePresentNothing uint64 = iota
	UACCategoryTypePresentUACstandardized
	UACCategoryTypePresentUACOperatorDefined
)

type UACCategoryType struct {
	Choice             uint64
	UACstandardized    *UACAction
	UACOperatorDefined *UACOperatorDefined
	// ChoiceExtension // ChoiceExtensions
}

func (ie *UACCategoryType) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case UACCategoryTypePresentUACstandardized:
		err = ie.UACstandardized.Encode(w)
	case UACCategoryTypePresentUACOperatorDefined:
		err = ie.UACOperatorDefined.Encode(w)
	}
	return
}

func (ie *UACCategoryType) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case UACCategoryTypePresentUACstandardized:
		var tmp UACAction
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.UACstandardized = &tmp
	case UACCategoryTypePresentUACOperatorDefined:
		var tmp UACOperatorDefined
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.UACOperatorDefined = &tmp
	}
	return
}
