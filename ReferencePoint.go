package f1ap

import "github.com/lvdund/ngap/aper"

const (
	ReferencePointPresentNothing uint64 = iota
	ReferencePointPresentCoordinateID
	ReferencePointPresentReferencePointCoordinate
	ReferencePointPresentReferencePointCoordinateHA
)

type ReferencePoint struct {
	Choice                     uint64
	CoordinateID               *int64
	ReferencePointCoordinate   *AccessPointPosition
	ReferencePointCoordinateHA *NGRANHighAccuracyAccessPointPosition
	// ChoiceExtension
}

func (ie *ReferencePoint) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case ReferencePointPresentCoordinateID:
		tmp := NewINTEGER(*ie.CoordinateID, aper.Constraint{Lb: 0, Ub: 511}, false)
		err = tmp.Encode(w)
	case ReferencePointPresentReferencePointCoordinate:
		err = ie.ReferencePointCoordinate.Encode(w)
	case ReferencePointPresentReferencePointCoordinateHA:
		err = ie.ReferencePointCoordinateHA.Encode(w)
	}
	return
}

func (ie *ReferencePoint) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case ReferencePointPresentCoordinateID:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 511}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.CoordinateID = (*int64)(&tmp.Value)
	case ReferencePointPresentReferencePointCoordinate:
		var tmp AccessPointPosition
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.ReferencePointCoordinate = &tmp
	case ReferencePointPresentReferencePointCoordinateHA:
		var tmp NGRANHighAccuracyAccessPointPosition
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.ReferencePointCoordinateHA = &tmp
	}
	return
}
