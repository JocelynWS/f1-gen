package f1ap

import "github.com/lvdund/ngap/aper"

const (
	EUTRAModeInfoPresentNothing uint64 = iota
	EUTRAModeInfoPresentEUTRAFDD
	EUTRAModeInfoPresentEUTRATDD
)

type EUTRAModeInfo struct {
	Choice   uint64
	EUTRAFDD *EUTRAFDDInfo
	EUTRATDD *EUTRATDDInfo
	// ChoiceExtension
}

func (ie *EUTRAModeInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case EUTRAModeInfoPresentEUTRAFDD:
		err = ie.EUTRAFDD.Encode(w)
	case EUTRAModeInfoPresentEUTRATDD:
		err = ie.EUTRATDD.Encode(w)
	}
	return
}

func (ie *EUTRAModeInfo) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case EUTRAModeInfoPresentEUTRAFDD:
		var tmp EUTRAFDDInfo
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRAFDD = &tmp
	case EUTRAModeInfoPresentEUTRATDD:
		var tmp EUTRATDDInfo
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRATDD = &tmp
	}
	return
}
