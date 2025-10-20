package ies

import "github.com/lvdund/ngap/aper"

const (
	MeasuredResultsValuePresentNothing uint64 = iota
	MeasuredResultsValuePresentULAngleOfArrival
	MeasuredResultsValuePresentULSRSRSRP
	MeasuredResultsValuePresentULRTOA
	MeasuredResultsValuePresentGNBRxTxTimeDiff
)

type MeasuredResultsValue struct {
	Choice           uint64
	ULAngleOfArrival *ULAoA
	ULSRSRSRP        *int64
	ULRTOA           *ULRTOAMeasurement
	GNBRxTxTimeDiff  *GNBRxTxTimeDiff
	// ChoiceExtension // ChoiceExtensions
}

func (ie *MeasuredResultsValue) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return
	}
	switch ie.Choice {
	case MeasuredResultsValuePresentULAngleOfArrival:
		err = ie.ULAngleOfArrival.Encode(w)
	case MeasuredResultsValuePresentULSRSRSRP:
		tmp := NewINTEGER(*ie.ULSRSRSRP, aper.Constraint{Lb: 0, Ub: 126}, false)
		err = tmp.Encode(w)
	case MeasuredResultsValuePresentULRTOA:
		err = ie.ULRTOA.Encode(w)
	case MeasuredResultsValuePresentGNBRxTxTimeDiff:
		err = ie.GNBRxTxTimeDiff.Encode(w)
	}
	return
}

func (ie *MeasuredResultsValue) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return
	}
	switch ie.Choice {
	case MeasuredResultsValuePresentULAngleOfArrival:
		var tmp ULAoA
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.ULAngleOfArrival = &tmp
	case MeasuredResultsValuePresentULSRSRSRP:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 126}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.ULSRSRSRP = (*int64)(&tmp.Value)
	case MeasuredResultsValuePresentULRTOA:
		var tmp ULRTOAMeasurement
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.ULRTOA = &tmp
	case MeasuredResultsValuePresentGNBRxTxTimeDiff:
		var tmp GNBRxTxTimeDiff
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GNBRxTxTimeDiff = &tmp
	}
	return
}
