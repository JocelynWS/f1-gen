package ies

import "github.com/lvdund/ngap/aper"

const (
	PRSResourceQCLInfoPresentNothing uint64 = iota
	PRSResourceQCLInfoPresentQCLSourceSSB
	PRSResourceQCLInfoPresentQCLSourcePRS
)

type PRSResourceQCLInfo struct {
	Choice       uint64
	QCLSourceSSB *PRSResourceQCLSourceSSB
	QCLSourcePRS *PRSResourceQCLSourcePRS
	// ChoiceExtension // ChoiceExtensions
}

func (ie *PRSResourceQCLInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case PRSResourceQCLInfoPresentQCLSourceSSB:
		err = ie.QCLSourceSSB.Encode(w)
	case PRSResourceQCLInfoPresentQCLSourcePRS:
		err = ie.QCLSourcePRS.Encode(w)
	}
	return
}

func (ie *PRSResourceQCLInfo) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case PRSResourceQCLInfoPresentQCLSourceSSB:
		var tmp PRSResourceQCLSourceSSB
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.QCLSourceSSB = &tmp
	case PRSResourceQCLInfoPresentQCLSourcePRS:
		var tmp PRSResourceQCLSourcePRS
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.QCLSourcePRS = &tmp
	}
	return
}
