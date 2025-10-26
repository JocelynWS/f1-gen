package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ConditionalIntraDUMobilityInformation struct {
	ChoTriggerIntraDU   CHOtriggerIntraDU   `mandatory`
	TargetCellsTocancel *TargetCellListItem `optional`
	//IEExtensions 		*ProtocolExtensionContainer `optional` // IEExtensions
}

func (ie *ConditionalIntraDUMobilityInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.TargetCellsTocancel != nil {
		aper.SetBit(optionals, 0)
	}
	w.WriteBits(optionals, 2)

	if err = ie.ChoTriggerIntraDU.Encode(w); err != nil {
		err = utils.WrapError("Encode ChoTriggerIntraDU", err)
		return
	}

	if ie.TargetCellsTocancel != nil {
		if err = ie.TargetCellsTocancel.Encode(w); err != nil {
			err = utils.WrapError("Encode TargetCellsTocancel", err)
			return
		}
	}

	return
}

func (ie *ConditionalIntraDUMobilityInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}

	if err = ie.ChoTriggerIntraDU.Decode(r); err != nil {
		err = utils.WrapError("Decode ChoTriggerIntraDU", err)
		return
	}

	if aper.IsBitSet(optionals, 0) {
		ie.TargetCellsTocancel = new(TargetCellListItem)
		if err = ie.TargetCellsTocancel.Decode(r); err != nil {
			err = utils.WrapError("Decode TargetCellsTocancel", err)
			return
		}
	}

	return
}
