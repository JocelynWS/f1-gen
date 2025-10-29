package f1ap

import "github.com/lvdund/ngap/aper"

const (
	DUCURadioInformationTypePresentNothing uint64 = iota
	DUCURadioInformationTypePresentRIM
)

type DUCURadioInformationType struct {
	Choice uint64
	RIM    *DUCURIMInformation
	// ChoiceExtension
}

func (ie *DUCURadioInformationType) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case DUCURadioInformationTypePresentRIM:
		err = ie.RIM.Encode(w)
	}
	return
}

func (ie *DUCURadioInformationType) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return
	}
	switch ie.Choice {
	case DUCURadioInformationTypePresentRIM:
		var tmp DUCURIMInformation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.RIM = &tmp
	}
	return
}
