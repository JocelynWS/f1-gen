package f1ap

import "github.com/lvdund/ngap/aper"

const (
	SRSTypePresentNothing uint64 = iota
	SRSTypePresentSemipersistentSRS
	SRSTypePresentAperiodicSRS
)

type SRSType struct {
	Choice            uint64
	SemipersistentSRS *SemipersistentSRS
	AperiodicSRS      *AperiodicSRS
	// ChoiceExtension
}

func (ie *SRSType) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case SRSTypePresentSemipersistentSRS:
		err = ie.SemipersistentSRS.Encode(w)
	case SRSTypePresentAperiodicSRS:
		err = ie.AperiodicSRS.Encode(w)
	}
	return
}

func (ie *SRSType) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case SRSTypePresentSemipersistentSRS:
		var tmp SemipersistentSRS
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SemipersistentSRS = &tmp
	case SRSTypePresentAperiodicSRS:
		var tmp AperiodicSRS
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.AperiodicSRS = &tmp
	}
	return
}
