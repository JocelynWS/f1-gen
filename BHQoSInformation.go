package ies

import "github.com/lvdund/ngap/aper"

const (
	BHQoSInformationPresentNothing uint64 = iota
	BHQoSInformationPresentBHRLCCHQoS
	BHQoSInformationPresentEUTRANBHRLCCHQoS
	BHQoSInformationPresentCPTrafficType
)

type BHQoSInformation struct {
	Choice           uint64
	BHRLCCHQoS       *QoSFlowLevelQoSParameters
	EUTRANBHRLCCHQoS *EUTRANQoS
	CPTrafficType    *int64
	// ChoiceExtension // ChoiceExtensions
}

func (ie *BHQoSInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case BHQoSInformationPresentBHRLCCHQoS:
		err = ie.BHRLCCHQoS.Encode(w)
	case BHQoSInformationPresentEUTRANBHRLCCHQoS:
		err = ie.EUTRANBHRLCCHQoS.Encode(w)
	case BHQoSInformationPresentCPTrafficType:
		tmp := NewINTEGER(*ie.CPTrafficType, aper.Constraint{Lb: 0, Ub: 3}, false)
		err = tmp.Encode(w)
	}
	return
}

func (ie *BHQoSInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case BHQoSInformationPresentBHRLCCHQoS:
		var tmp QoSFlowLevelQoSParameters
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.BHRLCCHQoS = &tmp
	case BHQoSInformationPresentEUTRANBHRLCCHQoS:
		var tmp EUTRANQoS
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRANBHRLCCHQoS = &tmp
	case BHQoSInformationPresentCPTrafficType:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 3}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.CPTrafficType = (*int64)(&tmp.Value)
	}
	return
}
