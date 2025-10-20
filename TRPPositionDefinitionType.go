package ies

import "github.com/lvdund/ngap/aper"

const (
	TRPPositionDefinitionTypePresentNothing uint64 = iota
	TRPPositionDefinitionTypePresentDirect
	TRPPositionDefinitionTypePresentReferenced
)

type TRPPositionDefinitionType struct {
	Choice     uint64
	Direct     *TRPPositionDirect
	Referenced *TRPPositionReferenced
	// ChoiceExtension // ChoiceExtensions
}

func (ie *TRPPositionDefinitionType) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case TRPPositionDefinitionTypePresentDirect:
		err = ie.Direct.Encode(w)
	case TRPPositionDefinitionTypePresentReferenced:
		err = ie.Referenced.Encode(w)
	}
	return
}

func (ie *TRPPositionDefinitionType) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case TRPPositionDefinitionTypePresentDirect:
		var tmp TRPPositionDirect
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.Direct = &tmp
	case TRPPositionDefinitionTypePresentReferenced:
		var tmp TRPPositionReferenced
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.Referenced = &tmp
	}
	return
}
