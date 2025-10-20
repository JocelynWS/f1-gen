package ies

import "github.com/lvdund/ngap/aper"

const (
	RATFrequencyPriorityInformationPresentNothing uint64 = iota
	RATFrequencyPriorityInformationPresentENDC
	RATFrequencyPriorityInformationPresentNGRAN
)

type RATFrequencyPriorityInformation struct {
	Choice uint64
	ENDC   *int64
	NGRAN  *int64
	// ChoiceExtension // ChoiceExtensions
}

func (ie *RATFrequencyPriorityInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case RATFrequencyPriorityInformationPresentENDC:
		tmp := NewINTEGER(*ie.ENDC, aper.Constraint{Lb: 1, Ub: 256}, false)
		err = tmp.Encode(w)
	case RATFrequencyPriorityInformationPresentNGRAN:
		tmp := NewINTEGER(*ie.NGRAN, aper.Constraint{Lb: 1, Ub: 256}, false)
		err = tmp.Encode(w)
	}
	return
}

func (ie *RATFrequencyPriorityInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case RATFrequencyPriorityInformationPresentENDC:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 1, Ub: 256}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.ENDC = (*int64)(&tmp.Value)
	case RATFrequencyPriorityInformationPresentNGRAN:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 1, Ub: 256}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NGRAN = (*int64)(&tmp.Value)
	}
	return
}
