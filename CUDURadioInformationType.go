package ies

import "github.com/lvdund/ngap/aper"

const (
	CUDURadioInformationTypePresentNothing uint64 = iota
	CUDURadioInformationTypePresentRIM
)

type CUDURadioInformationType struct {
	Choice uint64
	RIM    *CUDURIMInformation
	// ChoiceExtension // ChoiceExtensions
}

func (ie *CUDURadioInformationType) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case CUDURadioInformationTypePresentRIM:
		err = ie.RIM.Encode(w)
	}
	return
}

func (ie *CUDURadioInformationType) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return
	}
	switch ie.Choice {
	case CUDURadioInformationTypePresentRIM:
		var tmp CUDURIMInformation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.RIM = &tmp
	}
	return
}
