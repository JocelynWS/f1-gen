package f1ap

import "github.com/lvdund/ngap/aper"

const (
	TRPInformationTypeResponseItemPresentNothing uint64 = iota
	TRPInformationTypeResponseItemPresentPCINR
	TRPInformationTypeResponseItemPresentNGRANCGI
	TRPInformationTypeResponseItemPresentNRARFCN
	TRPInformationTypeResponseItemPresentPRSConfiguration
	TRPInformationTypeResponseItemPresentSSBinformation
	TRPInformationTypeResponseItemPresentSFNInitialisationTime
	TRPInformationTypeResponseItemPresentSpatialDirectionInformation
	TRPInformationTypeResponseItemPresentGeographicalCoordinates
)

type TRPInformationTypeResponseItem struct {
	Choice                      uint64
	PCINR                       *int64
	NGRANCGI                    *NRCGI
	NRARFCN                     *int64
	PRSConfiguration            *PRSConfiguration
	SSBinformation              *SSBInformation
	SFNInitialisationTime       *RelativeTime1900
	SpatialDirectionInformation *SpatialDirectionInformation
	GeographicalCoordinates     *GeographicalCoordinates
	// ChoiceExtension
}

func (ie *TRPInformationTypeResponseItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 8, false); err != nil {
		return
	}
	switch ie.Choice {
	case TRPInformationTypeResponseItemPresentPCINR:
		tmp := NewINTEGER(*ie.PCINR, aper.Constraint{Lb: 0, Ub: 1007}, false)
		err = tmp.Encode(w)
	case TRPInformationTypeResponseItemPresentNGRANCGI:
		err = ie.NGRANCGI.Encode(w)
	case TRPInformationTypeResponseItemPresentNRARFCN:
		tmp := NewINTEGER(*ie.NRARFCN, aper.Constraint{Lb: 0, Ub: maxNRARFCN}, false)
		err = tmp.Encode(w)
	case TRPInformationTypeResponseItemPresentPRSConfiguration:
		err = ie.PRSConfiguration.Encode(w)
	case TRPInformationTypeResponseItemPresentSSBinformation:
		err = ie.SSBinformation.Encode(w)
	case TRPInformationTypeResponseItemPresentSFNInitialisationTime:
		err = ie.SFNInitialisationTime.Encode(w)
	case TRPInformationTypeResponseItemPresentSpatialDirectionInformation:
		err = ie.SpatialDirectionInformation.Encode(w)
	case TRPInformationTypeResponseItemPresentGeographicalCoordinates:
		err = ie.GeographicalCoordinates.Encode(w)
	}
	return
}

func (ie *TRPInformationTypeResponseItem) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(8, false); err != nil {
		return
	}
	switch ie.Choice {
	case TRPInformationTypeResponseItemPresentPCINR:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 1007}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.PCINR = (*int64)(&tmp.Value)
	case TRPInformationTypeResponseItemPresentNGRANCGI:
		var tmp NRCGI
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NGRANCGI = &tmp
	case TRPInformationTypeResponseItemPresentNRARFCN:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxNRARFCN}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NRARFCN = (*int64)(&tmp.Value)
	case TRPInformationTypeResponseItemPresentPRSConfiguration:
		var tmp PRSConfiguration
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.PRSConfiguration = &tmp
	case TRPInformationTypeResponseItemPresentSSBinformation:
		var tmp SSBInformation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SSBinformation = &tmp
	case TRPInformationTypeResponseItemPresentSFNInitialisationTime:
		var tmp RelativeTime1900
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SFNInitialisationTime = &tmp
	case TRPInformationTypeResponseItemPresentSpatialDirectionInformation:
		var tmp SpatialDirectionInformation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SpatialDirectionInformation = &tmp
	case TRPInformationTypeResponseItemPresentGeographicalCoordinates:
		var tmp GeographicalCoordinates
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GeographicalCoordinates = &tmp
	}
	return
}
