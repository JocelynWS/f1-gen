package ies

import "github.com/lvdund/ngap/aper"

const (
	TimeStampSlotIndexPresentNothing uint64 = iota
	TimeStampSlotIndexPresentSCS15
	TimeStampSlotIndexPresentSCS30
	TimeStampSlotIndexPresentSCS60
	TimeStampSlotIndexPresentSCS120
)

type TimeStampSlotIndex struct {
	Choice uint64
	SCS15  *int64 
	SCS30  *int64 
	SCS60  *int64 
	SCS120 *int64 
    // ChoiceExtension
}

func (ie *TimeStampSlotIndex) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return
	}
	switch ie.Choice {
	case TimeStampSlotIndexPresentSCS15:
		tmp := NewINTEGER(*ie.SCS15, aper.Constraint{Lb: 0, Ub: 9}, false)
		err = tmp.Encode(w)
	case TimeStampSlotIndexPresentSCS30:
		tmp := NewINTEGER(*ie.SCS30, aper.Constraint{Lb: 0, Ub: 19}, false)
		err = tmp.Encode(w)
	case TimeStampSlotIndexPresentSCS60:
		tmp := NewINTEGER(*ie.SCS60, aper.Constraint{Lb: 0, Ub: 39}, false)
		err = tmp.Encode(w)
	case TimeStampSlotIndexPresentSCS120:
		tmp := NewINTEGER(*ie.SCS120, aper.Constraint{Lb: 0, Ub: 79}, false)
		err = tmp.Encode(w)
	}
	return
}

func (ie *TimeStampSlotIndex) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return
	}
	switch ie.Choice {
	case TimeStampSlotIndexPresentSCS15:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 9}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SCS15 = (*int64)(&tmp.Value)
	case TimeStampSlotIndexPresentSCS30:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 19}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SCS30 = (*int64)(&tmp.Value)
	case TimeStampSlotIndexPresentSCS60:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 39}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SCS60 = (*int64)(&tmp.Value)
	case TimeStampSlotIndexPresentSCS120:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 79}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SCS120 = (*int64)(&tmp.Value)
	}
	return
}
