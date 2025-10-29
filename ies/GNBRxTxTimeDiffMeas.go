package ies

import "github.com/lvdund/ngap/aper"

const (
	GNBRxTxTimeDiffMeasPresentNothing uint64 = iota
	GNBRxTxTimeDiffMeasPresentK0
	GNBRxTxTimeDiffMeasPresentK1
	GNBRxTxTimeDiffMeasPresentK2
	GNBRxTxTimeDiffMeasPresentK3
	GNBRxTxTimeDiffMeasPresentK4
	GNBRxTxTimeDiffMeasPresentK5
)

type GNBRxTxTimeDiffMeas struct {
	Choice uint64
	K0     *int64
	K1     *int64
	K2     *int64
	K3     *int64
	K4     *int64
	K5     *int64
	// ChoiceExtension
}

func (ie *GNBRxTxTimeDiffMeas) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 6, false); err != nil {
		return
	}
	switch ie.Choice {
	case GNBRxTxTimeDiffMeasPresentK0:
		tmp := NewINTEGER(*ie.K0, aper.Constraint{Lb: 0, Ub: 1970049}, false)
		err = tmp.Encode(w)
	case GNBRxTxTimeDiffMeasPresentK1:
		tmp := NewINTEGER(*ie.K1, aper.Constraint{Lb: 0, Ub: 985025}, false)
		err = tmp.Encode(w)
	case GNBRxTxTimeDiffMeasPresentK2:
		tmp := NewINTEGER(*ie.K2, aper.Constraint{Lb: 0, Ub: 492513}, false)
		err = tmp.Encode(w)
	case GNBRxTxTimeDiffMeasPresentK3:
		tmp := NewINTEGER(*ie.K3, aper.Constraint{Lb: 0, Ub: 246257}, false)
		err = tmp.Encode(w)
	case GNBRxTxTimeDiffMeasPresentK4:
		tmp := NewINTEGER(*ie.K4, aper.Constraint{Lb: 0, Ub: 123129}, false)
		err = tmp.Encode(w)
	case GNBRxTxTimeDiffMeasPresentK5:
		tmp := NewINTEGER(*ie.K5, aper.Constraint{Lb: 0, Ub: 61565}, false)
		err = tmp.Encode(w)
	}
	return
}

func (ie *GNBRxTxTimeDiffMeas) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(6, false); err != nil {
		return
	}
	switch ie.Choice {
	case GNBRxTxTimeDiffMeasPresentK0:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 1970049}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.K0 = (*int64)(&tmp.Value)
	case GNBRxTxTimeDiffMeasPresentK1:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 985025}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.K1 = (*int64)(&tmp.Value)
	case GNBRxTxTimeDiffMeasPresentK2:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 492513}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.K2 = (*int64)(&tmp.Value)
	case GNBRxTxTimeDiffMeasPresentK3:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 246257}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.K3 = (*int64)(&tmp.Value)
	case GNBRxTxTimeDiffMeasPresentK4:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 123129}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.K4 = (*int64)(&tmp.Value)
	case GNBRxTxTimeDiffMeasPresentK5:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 61565}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.K5 = (*int64)(&tmp.Value)
	}
	return
}
