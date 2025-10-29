package ies

import "github.com/lvdund/ngap/aper"

const (
	BandwidthSRSPresentNothing uint64 = iota
	BandwidthSRSPresentFR1
	BandwidthSRSPresentFR2
)

type BandwidthSRS struct {
	Choice uint64
	FR1    *FR1Bandwidth
	FR2    *FR2Bandwidth
	// ChoiceExtension
}

func (ie *BandwidthSRS) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case BandwidthSRSPresentFR1:
		err = ie.FR1.Encode(w)
	case BandwidthSRSPresentFR2:
		err = ie.FR2.Encode(w)
	}
	return
}

func (ie *BandwidthSRS) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case BandwidthSRSPresentFR1:
		var tmp FR1Bandwidth
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.FR1 = &tmp
	case BandwidthSRSPresentFR2:
		var tmp FR2Bandwidth
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.FR2 = &tmp
	}
	return
}
