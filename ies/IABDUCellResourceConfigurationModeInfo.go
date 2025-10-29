package ies

import "github.com/lvdund/ngap/aper"

const (
	IABDUCellResourceConfigurationModeInfoPresentNothing uint64 = iota
	IABDUCellResourceConfigurationModeInfoPresentFDD
	IABDUCellResourceConfigurationModeInfoPresentTDD
)

type IABDUCellResourceConfigurationModeInfo struct {
	Choice uint64
	FDD    *IABDUCellResourceConfigurationFDDInfo
	TDD    *IABDUCellResourceConfigurationTDDInfo
	// ChoiceExtension
}

func (ie *IABDUCellResourceConfigurationModeInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case IABDUCellResourceConfigurationModeInfoPresentFDD:
		err = ie.FDD.Encode(w)
	case IABDUCellResourceConfigurationModeInfoPresentTDD:
		err = ie.TDD.Encode(w)
	}
	return
}

func (ie *IABDUCellResourceConfigurationModeInfo) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case IABDUCellResourceConfigurationModeInfoPresentFDD:
		var tmp IABDUCellResourceConfigurationFDDInfo
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.FDD = &tmp
	case IABDUCellResourceConfigurationModeInfoPresentTDD:
		var tmp IABDUCellResourceConfigurationTDDInfo
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TDD = &tmp
	}
	return
}
