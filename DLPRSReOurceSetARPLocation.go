package ies

import "github.com/lvdund/ngap/aper"

const (
	DLPRSReOurceSetARPLocationPresentNothing uint64 = iota
	DLPRSReOurceSetARPLocationPresentRelativeGeodeticLocation
	DLPRSReOurceSetARPLocationPresentRelativeCarteIanLocation
	DLPRSReOurceSetARPLocationPresentChoiceExtension
)

type DLPRSReOurceSetARPLocation struct {
	Choice                   uint64
	RelativeGeodeticLocation *RelativeGeodeticLocation
	RelativeCarteIanLocation *RelativeCarteIanLocation
	// ChoiceExtension *DLPRSReOurceSetARPLocationExtIEs
}

func (ie *DLPRSReOurceSetARPLocation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case DLPRSReOurceSetARPLocationPresentRelativeGeodeticLocation:
		err = ie.RelativeGeodeticLocation.Encode(w)
	case DLPRSReOurceSetARPLocationPresentRelativeCarteIanLocation:
		err = ie.RelativeCarteIanLocation.Encode(w)
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
	case DLPRSReOurceSetARPLocationPresentRelativeCarteIanLocation:
		var tmp RelativeCarteIanLocation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.RelativeCarteIanLocation = &tmp
	}
	return
}
