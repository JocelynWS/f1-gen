package ies

import "github.com/lvdund/ngap/aper"

const (
	EUTRACoexModeInfoPresentNothing uint64 = iota
	EUTRACoexModeInfoPresentFDD
	EUTRACoexModeInfoPresentTDD
)

type EUTRACoexModeInfo struct {
	Choice uint64
	FDD    *EUTRACoexFDDInfo
	TDD    *EUTRACoexTDDInfo
}

func (ie *EUTRACoexModeInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case EUTRACoexModeInfoPresentFDD:
		err = ie.FDD.Encode(w)
	case EUTRACoexModeInfoPresentTDD:
		err = ie.TDD.Encode(w)
	}
	return
}

func (ie *EUTRACoexModeInfo) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case EUTRACoexModeInfoPresentFDD:
		var tmp EUTRACoexFDDInfo
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.FDD = &tmp
	case EUTRACoexModeInfoPresentTDD:
		var tmp EUTRACoexTDDInfo
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TDD = &tmp
	}
	return
}
