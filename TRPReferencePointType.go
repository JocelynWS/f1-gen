package ies

import "github.com/lvdund/ngap/aper"

const (
	TRPReferencePointTypePresentNothing uint64 = iota
	TRPReferencePointTypePresentTRPPositionRelativeGeodetic
	TRPReferencePointTypePresentTRPPositionRelativeCartesian
)

type TRPReferencePointType struct {
	Choice                       uint64
	TRPPositionRelativeGeodetic  *RelativeGeodeticLocation
	TRPPositionRelativeCartesian *RelativeCartesianLocation
	// ChoiceExtension // ChoiceExtensions
}

func (ie *TRPReferencePointType) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case TRPReferencePointTypePresentTRPPositionRelativeGeodetic:
		err = ie.TRPPositionRelativeGeodetic.Encode(w)
	case TRPReferencePointTypePresentTRPPositionRelativeCartesian:
		err = ie.TRPPositionRelativeCartesian.Encode(w)
	}
	return
}

func (ie *TRPReferencePointType) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case TRPReferencePointTypePresentTRPPositionRelativeGeodetic:
		var tmp RelativeGeodeticLocation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TRPPositionRelativeGeodetic = &tmp
	case TRPReferencePointTypePresentTRPPositionRelativeCartesian:
		var tmp RelativeCartesianLocation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TRPPositionRelativeCartesian = &tmp
	}
	return
}
