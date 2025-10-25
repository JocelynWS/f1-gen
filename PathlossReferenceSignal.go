package ies

import "github.com/lvdund/ngap/aper"

const (
	PathlossReferenceSignalPresentNothing uint64 = iota
	PathlossReferenceSignalPresentSSB
	PathlossReferenceSignalPresentDLPRS
)

type PathlossReferenceSignal struct {
	Choice uint64
	SSB    *SSB
	DLPRS  *DLPRS
	// ChoiceExtension
}

func (ie *PathlossReferenceSignal) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case PathlossReferenceSignalPresentSSB:
		err = ie.SSB.Encode(w)
	case PathlossReferenceSignalPresentDLPRS:
		err = ie.DLPRS.Encode(w)
	}
	return
}

func (ie *PathlossReferenceSignal) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case PathlossReferenceSignalPresentSSB:
		var tmp SSB
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SSB = &tmp
	case PathlossReferenceSignalPresentDLPRS:
		var tmp DLPRS
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.DLPRS = &tmp
	}
	return
}
