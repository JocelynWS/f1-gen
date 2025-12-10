package ies

import "github.com/lvdund/ngap/aper"

const (
	AbortTransmissionPresentNothing uint64 = iota
	AbortTransmissionPresentSRSResourceSetID
	AbortTransmissionPresentReleaseALL
)

type AbortTransmission struct {
	Choice           uint64
	SRSResourceSetID *int64
	ReleaseALL       *NULL
	// ChoiceExtension
}

func (ie *AbortTransmission) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case AbortTransmissionPresentSRSResourceSetID:
		tmp := NewINTEGER(*ie.SRSResourceSetID, aper.Constraint{Lb: 0, Ub: 15}, true)
		err = tmp.Encode(w)
	case AbortTransmissionPresentReleaseALL:
		err = ie.ReleaseALL.Encode(w)
	}
	return
}

func (ie *AbortTransmission) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case AbortTransmissionPresentSRSResourceSetID:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 15}, true)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SRSResourceSetID = (*int64)(&tmp.Value)
	case AbortTransmissionPresentReleaseALL:
		var tmp NULL
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.ReleaseALL = &tmp
	}
	return
}
