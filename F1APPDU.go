package ies

import "github.com/lvdund/ngap/aper"

const (
	F1apPduPresentNothing uint64 = iota
	F1apPduPresentInitiatingMessage
	F1apPduPresentSuccessfulOutcome
	F1apPduPresentUnsuccessfulOutcome
)

type F1apPdu struct {
	Choice              uint64
	InitiatingMessage   *InitiatingMessage
	SuccessfulOutcome   *SuccessfulOutcome
	UnsuccessfulOutcome *UnsuccessfulOutcome
}

func (ie *F1apPdu) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case F1apPduPresentInitiatingMessage:
		err = ie.InitiatingMessage.Encode(w)
	case F1apPduPresentSuccessfulOutcome:
		err = ie.SuccessfulOutcome.Encode(w)
	case F1apPduPresentUnsuccessfulOutcome:
		err = ie.UnsuccessfulOutcome.Encode(w)
	}
	return
}

func (ie *F1apPdu) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case F1apPduPresentInitiatingMessage:
		var tmp InitiatingMessage
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.InitiatingMessage = &tmp
	case F1apPduPresentSuccessfulOutcome:
		var tmp SuccessfulOutcome
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SuccessfulOutcome = &tmp
	case F1apPduPresentUnsuccessfulOutcome:
		var tmp UnsuccessfulOutcome
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.UnsuccessfulOutcome = &tmp
	}
	return
}




