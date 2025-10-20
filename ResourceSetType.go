package ies

import "github.com/lvdund/ngap/aper"

const (
	ResourceSetTypePresentNothing uint64 = iota
	ResourceSetTypePresentPeriodic
	ResourceSetTypePresentSemiPersistent
	ResourceSetTypePresentAperiodic
)

type ResourceSetType struct {
	Choice         uint64
	Periodic       *ResourceSetTypePeriodic
	SemiPersistent *ResourceSetTypeSemiPersistent
	Aperiodic      *ResourceSetTypeAperiodic
	// ChoiceExtension // ChoiceExtensions
}

func (ie *ResourceSetType) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case ResourceSetTypePresentPeriodic:
		err = ie.Periodic.Encode(w)
	case ResourceSetTypePresentSemiPersistent:
		err = ie.SemiPersistent.Encode(w)
	case ResourceSetTypePresentAperiodic:
		err = ie.Aperiodic.Encode(w)
	}
	return
}

func (ie *ResourceSetType) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case ResourceSetTypePresentPeriodic:
		var tmp ResourceSetTypePeriodic
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.Periodic = &tmp
	case ResourceSetTypePresentSemiPersistent:
		var tmp ResourceSetTypeSemiPersistent
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SemiPersistent = &tmp
	case ResourceSetTypePresentAperiodic:
		var tmp ResourceSetTypeAperiodic
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.Aperiodic = &tmp
	}
	return
}
