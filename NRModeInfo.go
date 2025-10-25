package ies

import "github.com/lvdund/ngap/aper"

const (
	NRModeInfoPresentNothing uint64 = iota
	NRModeInfoPresentFDD
	NRModeInfoPresentTDD
)

type NRModeInfo struct {
	Choice uint64
	FDD    *FDDInfo
	TDD    *TDDInfo
	// ChoiceExtension
}

func (ie *NRModeInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case NRModeInfoPresentFDD:
		err = ie.FDD.Encode(w)
	case NRModeInfoPresentTDD:
		err = ie.TDD.Encode(w)
	}
	return
}

func (ie *NRModeInfo) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case NRModeInfoPresentFDD:
		var tmp FDDInfo
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.FDD = &tmp
	case NRModeInfoPresentTDD:
		var tmp TDDInfo
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TDD = &tmp
	}
	return
}
