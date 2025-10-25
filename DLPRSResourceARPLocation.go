package ies

import "github.com/lvdund/ngap/aper"

const (
	DLPRSResourceARPLocationPresentNothing uint64 = iota
	DLPRSResourceARPLocationPresentRelativeGeodeticLocation
	DLPRSResourceARPLocationPresentRelativeCartesianLocation
)

type DLPRSResourceARPLocation struct {
	Choice                    uint64
	RelativeGeodeticLocation  *RelativeGeodeticLocation
	RelativeCartesianLocation *RelativeCartesianLocation
	// ChoiceExtension
}

func (ie *DLPRSResourceARPLocation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case DLPRSResourceARPLocationPresentRelativeGeodeticLocation:
		err = ie.RelativeGeodeticLocation.Encode(w)
	case DLPRSResourceARPLocationPresentRelativeCartesianLocation:
		err = ie.RelativeCartesianLocation.Encode(w)
	}
	return
}

func (ie *DLPRSResourceARPLocation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case DLPRSResourceARPLocationPresentRelativeGeodeticLocation:
		var tmp RelativeGeodeticLocation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.RelativeGeodeticLocation = &tmp
	case DLPRSResourceARPLocationPresentRelativeCartesianLocation:
		var tmp RelativeCartesianLocation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.RelativeCartesianLocation = &tmp
	}
	return
}
