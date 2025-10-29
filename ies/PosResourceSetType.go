package ies

import "github.com/lvdund/ngap/aper"

const (
	PosResourceSetTypePresentNothing uint64 = iota
	PosResourceSetTypePresentPeriodic
	PosResourceSetTypePresentSemiPersistent
	PosResourceSetTypePresentAperiodic
)

type PosResourceSetType struct {
	Choice         uint64
	Periodic       *PosResourceSetTypePR
	SemiPersistent *PosResourceSetTypeSP
	Aperiodic      *PosResourceSetTypeAP
	// ChoiceExtension
}

func (ie *PosResourceSetType) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case PosResourceSetTypePresentPeriodic:
		err = ie.Periodic.Encode(w)
	case PosResourceSetTypePresentSemiPersistent:
		err = ie.SemiPersistent.Encode(w)
	case PosResourceSetTypePresentAperiodic:
		err = ie.Aperiodic.Encode(w)
	}
	return
}

func (ie *PosResourceSetType) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case PosResourceSetTypePresentPeriodic:
		var tmp PosResourceSetTypePR
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.Periodic = &tmp
	case PosResourceSetTypePresentSemiPersistent:
		var tmp PosResourceSetTypeSP
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SemiPersistent = &tmp
	case PosResourceSetTypePresentAperiodic:
		var tmp PosResourceSetTypeAP
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.Aperiodic = &tmp
	}
	return
}
