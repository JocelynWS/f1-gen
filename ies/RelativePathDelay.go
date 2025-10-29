package ies

import "github.com/lvdund/ngap/aper"

const (
	RelativePathDelayPresentNothing uint64 = iota
	RelativePathDelayPresentK0
	RelativePathDelayPresentK1
	RelativePathDelayPresentK2
	RelativePathDelayPresentK3
	RelativePathDelayPresentK4
	RelativePathDelayPresentK5
)

type RelativePathDelay struct {
	Choice uint64
	K0     *int64
	K1     *int64
	K2     *int64
	K3     *int64
	K4     *int64
	K5     *int64
	// ChoiceExtension
}

func (ie *RelativePathDelay) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 6, false); err != nil {
		return
	}
	switch ie.Choice {
	case RelativePathDelayPresentK0:
		tmp := NewINTEGER(*ie.K0, aper.Constraint{Lb: 0, Ub: 16351}, false)
		err = tmp.Encode(w)
	case RelativePathDelayPresentK1:
		tmp := NewINTEGER(*ie.K1, aper.Constraint{Lb: 0, Ub: 8176}, false)
		err = tmp.Encode(w)
	case RelativePathDelayPresentK2:
		tmp := NewINTEGER(*ie.K2, aper.Constraint{Lb: 0, Ub: 4088}, false)
		err = tmp.Encode(w)
	case RelativePathDelayPresentK3:
		tmp := NewINTEGER(*ie.K3, aper.Constraint{Lb: 0, Ub: 2044}, false)
		err = tmp.Encode(w)
	case RelativePathDelayPresentK4:
		tmp := NewINTEGER(*ie.K4, aper.Constraint{Lb: 0, Ub: 1022}, false)
		err = tmp.Encode(w)
	case RelativePathDelayPresentK5:
		tmp := NewINTEGER(*ie.K5, aper.Constraint{Lb: 0, Ub: 511}, false)
		err = tmp.Encode(w)
	}
	return
}

func (ie *RelativePathDelay) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(6, false); err != nil {
		return
	}
	switch ie.Choice {
	case RelativePathDelayPresentK0:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 16351}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.K0 = (*int64)(&tmp.Value)
	case RelativePathDelayPresentK1:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 8176}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.K1 = (*int64)(&tmp.Value)
	case RelativePathDelayPresentK2:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 4088}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.K2 = (*int64)(&tmp.Value)
	case RelativePathDelayPresentK3:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 2044}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.K3 = (*int64)(&tmp.Value)
	case RelativePathDelayPresentK4:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 1022}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.K4 = (*int64)(&tmp.Value)
	case RelativePathDelayPresentK5:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 511}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.K5 = (*int64)(&tmp.Value)
	}
	return
}
