package f1ap

import "github.com/lvdund/ngap/aper"

const (
	TRPPositionDirectAccuracyPresentNothing uint64 = iota
	TRPPositionDirectAccuracyPresentTRPPosition
	TRPPositionDirectAccuracyPresentTRPHAposition
)

type TRPPositionDirectAccuracy struct {
	Choice        uint64
	TRPPosition   *AccessPointPosition
	TRPHAposition *NGRANHighAccuracyAccessPointPosition
	// ChoiceExtension
}

func (ie *TRPPositionDirectAccuracy) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case TRPPositionDirectAccuracyPresentTRPPosition:
		err = ie.TRPPosition.Encode(w)
	case TRPPositionDirectAccuracyPresentTRPHAposition:
		err = ie.TRPHAposition.Encode(w)
	}
	return
}

func (ie *TRPPositionDirectAccuracy) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case TRPPositionDirectAccuracyPresentTRPPosition:
		var tmp AccessPointPosition
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TRPPosition = &tmp
	case TRPPositionDirectAccuracyPresentTRPHAposition:
		var tmp NGRANHighAccuracyAccessPointPosition
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TRPHAposition = &tmp
	}
	return
}
