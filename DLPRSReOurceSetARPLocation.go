package f1ap

import "github.com/lvdund/ngap/aper"

const (
	DLPRSReOurceSetARPLocationPresentNothing uint64 = iota
	DLPRSReOurceSetARPLocationPresentRelativeGeodeticLocation
	DLPRSReOurceSetARPLocationPresentRelativeCartesianLocation
	DLPRSReOurceSetARPLocationPresentChoiceExtension
)

type DLPRSReOurceSetARPLocation struct {
	Choice                    uint64
	RelativeGeodeticLocation  *RelativeGeodeticLocation
	RelativeCartesianLocation *RelativeCartesianLocation
	// ChoiceExtension *DLPRSReOurceSetARPLocationExtIEs
}

func (ie *DLPRSReOurceSetARPLocation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case DLPRSReOurceSetARPLocationPresentRelativeGeodeticLocation:
		err = ie.RelativeGeodeticLocation.Encode(w)
	case DLPRSReOurceSetARPLocationPresentRelativeCartesianLocation:
		err = ie.RelativeCartesianLocation.Encode(w)
	}
	return
}

func (ie *DLPRSReOurceSetARPLocation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case DLPRSReOurceSetARPLocationPresentRelativeGeodeticLocation:
		var tmp RelativeGeodeticLocation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.RelativeGeodeticLocation = &tmp
	case DLPRSReOurceSetARPLocationPresentRelativeCartesianLocation:
		var tmp RelativeCartesianLocation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.RelativeCartesianLocation = &tmp
	}
	return
}
