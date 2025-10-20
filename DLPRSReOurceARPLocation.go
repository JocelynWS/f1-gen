package ies

import "github.com/lvdund/ngap/aper"

const (
	DLPRSReOurceARPLocationPresentNothing uint64 = iota
	DLPRSReOurceARPLocationPresentRelativeGeodeticLocation
	DLPRSReOurceARPLocationPresentRelativeCarteIanLocation
	DLPRSReOurceARPLocationPresentChoiceExtension
)

type DLPRSReOurceARPLocation struct {
	Choice                   uint64
	RelativeGeodeticLocation *RelativeGeodeticLocation
	RelativeCarteIanLocation *RelativeCarteIanLocation
	// ChoiceExtension *DLPRSReOurceARPLocationExtIEs
}

func (ie *DLPRSReOurceARPLocation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case DLPRSReOurceARPLocationPresentRelativeGeodeticLocation:
		err = ie.RelativeGeodeticLocation.Encode(w)
	case DLPRSReOurceARPLocationPresentRelativeCarteIanLocation:
		err = ie.RelativeCarteIanLocation.Encode(w)
	}
	return
}

func (ie *DLPRSReOurceARPLocation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case DLPRSReOurceARPLocationPresentRelativeGeodeticLocation:
		var tmp RelativeGeodeticLocation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.RelativeGeodeticLocation = &tmp
	case DLPRSReOurceARPLocationPresentRelativeCarteIanLocation:
		var tmp RelativeCarteIanLocation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.RelativeCarteIanLocation = &tmp
	}
	return
}
