package ies

import "github.com/lvdund/ngap/aper"

const (
	ECIDMeasuredResultsValuePresentNothing uint64 = iota
	ECIDMeasuredResultsValuePresentValueAngleofArrivalNR
)

type ECIDMeasuredResultsValue struct {
	Choice                uint64
	ValueAngleofArrivalNR *ULAoA
	// ChoiceExtension
}

func (ie *ECIDMeasuredResultsValue) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case ECIDMeasuredResultsValuePresentValueAngleofArrivalNR:
		err = ie.ValueAngleofArrivalNR.Encode(w)
	}
	return
}

func (ie *ECIDMeasuredResultsValue) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return
	}
	switch ie.Choice {
	case ECIDMeasuredResultsValuePresentValueAngleofArrivalNR:
		var tmp ULAoA
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.ValueAngleofArrivalNR = &tmp
	}
	return
}
