package ies

import "github.com/lvdund/ngap/aper"

const (
	ResourceTypePosPresentNothing uint64 = iota
	ResourceTypePosPresentPeriodic
	ResourceTypePosPresentSemiPersistent
	ResourceTypePosPresentAperiodic
)

type ResourceTypePos struct {
	Choice         uint64
	Periodic       *ResourceTypePeriodicPos
	SemiPersistent *ResourceTypeSemiPersistentPos
	Aperiodic      *ResourceTypeAperiodicPos
	// ChoiceExtension
}

func (ie *ResourceTypePos) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case ResourceTypePosPresentPeriodic:
		err = ie.Periodic.Encode(w)
	case ResourceTypePosPresentSemiPersistent:
		err = ie.SemiPersistent.Encode(w)
	case ResourceTypePosPresentAperiodic:
		err = ie.Aperiodic.Encode(w)
	}
	return
}

func (ie *ResourceTypePos) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case ResourceTypePosPresentPeriodic:
		var tmp ResourceTypePeriodicPos
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.Periodic = &tmp
	case ResourceTypePosPresentSemiPersistent:
		var tmp ResourceTypeSemiPersistentPos
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SemiPersistent = &tmp
	case ResourceTypePosPresentAperiodic:
		var tmp ResourceTypeAperiodicPos
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.Aperiodic = &tmp
	}
	return
}
