package ies

import "github.com/lvdund/ngap/aper"

const (
	DUFSlotConfigItemPresentNothing uint64 = iota
	DUFSlotConfigItemPresentExplicitFormat
	DUFSlotConfigItemPresentImplicitFormat
)

type DUFSlotConfigItem struct {
	Choice         uint64
	ExplicitFormat *ExplicitFormat
	ImplicitFormat *ImplicitFormat
	// ChoiceExtension // ChoiceExtensions
}

func (ie *DUFSlotConfigItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case DUFSlotConfigItemPresentExplicitFormat:
		err = ie.ExplicitFormat.Encode(w)
	case DUFSlotConfigItemPresentImplicitFormat:
		err = ie.ImplicitFormat.Encode(w)
	}
	return
}

func (ie *DUFSlotConfigItem) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case DUFSlotConfigItemPresentExplicitFormat:
		var tmp ExplicitFormat
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.ExplicitFormat = &tmp
	case DUFSlotConfigItemPresentImplicitFormat:
		var tmp ImplicitFormat
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.ImplicitFormat = &tmp
	}
	return
}
