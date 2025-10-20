package ies

import "github.com/lvdund/ngap/aper"

const (
	FreqDomainLengthPresentNothing uint64 = iota
	FreqDomainLengthPresentL839
	FreqDomainLengthPresentL139
)

type FreqDomainLength struct {
	Choice uint64
	L839   *L839Info
	L139   *L139Info
	// ChoiceExtension // ChoiceExtensions
}

func (ie *FreqDomainLength) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case FreqDomainLengthPresentL839:
		err = ie.L839.Encode(w)
	case FreqDomainLengthPresentL139:
		err = ie.L139.Encode(w)
	}
	return
}

func (ie *FreqDomainLength) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case FreqDomainLengthPresentL839:
		var tmp L839Info
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.L839 = &tmp
	case FreqDomainLengthPresentL139:
		var tmp L139Info
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.L139 = &tmp
	}
	return
}
