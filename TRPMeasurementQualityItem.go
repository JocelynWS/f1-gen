package ies

import "github.com/lvdund/ngap/aper"

const (
	TRPMeasurementQualityItemPresentNothing uint64 = iota
	TRPMeasurementQualityItemPresentTimingMeasurementQuality
	TRPMeasurementQualityItemPresentAngleMeasurementQuality
)

type TRPMeasurementQualityItem struct {
	Choice                   uint64
	TimingMeasurementQuality *TimingMeasurementQuality
	AngleMeasurementQuality  *AngleMeasurementQuality
	// ChoiceExtension
}

func (ie *TRPMeasurementQualityItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case TRPMeasurementQualityItemPresentTimingMeasurementQuality:
		err = ie.TimingMeasurementQuality.Encode(w)
	case TRPMeasurementQualityItemPresentAngleMeasurementQuality:
		err = ie.AngleMeasurementQuality.Encode(w)
	}
	return
}

func (ie *TRPMeasurementQualityItem) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case TRPMeasurementQualityItemPresentTimingMeasurementQuality:
		var tmp TimingMeasurementQuality
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TimingMeasurementQuality = &tmp
	case TRPMeasurementQualityItemPresentAngleMeasurementQuality:
		var tmp AngleMeasurementQuality
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.AngleMeasurementQuality = &tmp
	}
	return
}
