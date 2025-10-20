package ies

import "github.com/lvdund/ngap/aper"

const (
	SpatialRelationPosPresentNothing uint64 = iota
	SpatialRelationPosPresentSSBPos
	SpatialRelationPosPresentPRSInformationPos
)

type SpatialRelationPos struct {
	Choice            uint64
	SSBPos            *SSB
	PRSInformationPos *PRSInformationPos
	// ChoiceExtension // ChoiceExtensions
}

func (ie *SpatialRelationPos) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case SpatialRelationPosPresentSSBPos:
		err = ie.SSBPos.Encode(w)
	case SpatialRelationPosPresentPRSInformationPos:
		err = ie.PRSInformationPos.Encode(w)
	}
	return
}

func (ie *SpatialRelationPos) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case SpatialRelationPosPresentSSBPos:
		var tmp SSB
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SSBPos = &tmp
	case SpatialRelationPosPresentPRSInformationPos:
		var tmp PRSInformationPos
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.PRSInformationPos = &tmp
	}
	return
}
