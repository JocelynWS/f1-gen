package ies

import "github.com/lvdund/ngap/aper"

const (
	DLPRSResourceSetARPLocationPresentNothing uint64 = iota
	DLPRSResourceSetARPLocationPresentRelativeGeodeticLocation
	DLPRSResourceSetARPLocationPresentRelativeCartesianLocation
)

type DLPRSResourceSetARPLocation struct {
	Choice                    uint64
	RelativeGeodeticLocation  *RelativeGeodeticLocation
	RelativeCartesianLocation *RelativeCartesianLocation
	// ChoiceExtension // ChoiceExtensions
}

func (ie *DLPRSResourceSetARPLocation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case DLPRSResourceSetARPLocationPresentRelativeGeodeticLocation:
		err = ie.RelativeGeodeticLocation.Encode(w)
	case DLPRSResourceSetARPLocationPresentRelativeCartesianLocation:
		err = ie.RelativeCartesianLocation.Encode(w)
	}
	return
}

func (ie *DLPRSResourceSetARPLocation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case DLPRSResourceSetARPLocationPresentRelativeGeodeticLocation:
		var tmp RelativeGeodeticLocation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.RelativeGeodeticLocation = &tmp
	case DLPRSResourceSetARPLocationPresentRelativeCartesianLocation:
		var tmp RelativeCartesianLocation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.RelativeCartesianLocation = &tmp
	}
	return
}
