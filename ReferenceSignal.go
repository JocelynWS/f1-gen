package ies

import "github.com/lvdund/ngap/aper"

const (
	ReferenceSignalPresentNothing uint64 = iota
	ReferenceSignalPresentNZPCSIRS
	ReferenceSignalPresentSSB
	ReferenceSignalPresentSRS
	ReferenceSignalPresentPositioningSRS
	ReferenceSignalPresentDLPRS
)

type ReferenceSignal struct {
	Choice         uint64
	NZPCSIRS       *NZPCSIRSResourceID
	SSB            *SSB
	SRS            *int64
	PositioningSRS *int64
	DLPRS          *DLPRS
	// ChoiceExtension // ChoiceExtensions
}

func (ie *ReferenceSignal) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 5, false); err != nil {
		return
	}
	switch ie.Choice {
	case ReferenceSignalPresentNZPCSIRS:
		err = ie.NZPCSIRS.Encode(w)
	case ReferenceSignalPresentSSB:
		err = ie.SSB.Encode(w)
	case ReferenceSignalPresentSRS:
		tmp := NewINTEGER(*ie.SRS, aper.Constraint{Lb: 0, Ub: 63}, false)
		err = tmp.Encode(w)
	case ReferenceSignalPresentPositioningSRS:
		tmp := NewINTEGER(*ie.PositioningSRS, aper.Constraint{Lb: 0, Ub: 63}, false)
		err = tmp.Encode(w)
	case ReferenceSignalPresentDLPRS:
		err = ie.DLPRS.Encode(w)
	}
	return
}

func (ie *ReferenceSignal) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(5, false); err != nil {
		return
	}
	switch ie.Choice {
	case ReferenceSignalPresentNZPCSIRS:
		var tmp NZPCSIRSResourceID
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NZPCSIRS = &tmp
	case ReferenceSignalPresentSSB:
		var tmp SSB
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SSB = &tmp
	case ReferenceSignalPresentSRS:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 63}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SRS = (*int64)(&tmp.Value)
	case ReferenceSignalPresentPositioningSRS:
		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 63}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.PositioningSRS = (*int64)(&tmp.Value)
	case ReferenceSignalPresentDLPRS:
		var tmp DLPRS
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.DLPRS = &tmp
	}
	return
}
