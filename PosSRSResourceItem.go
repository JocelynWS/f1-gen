package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PosSRSResourceItem struct {
	SrsPosResourceId       int64                  `lb:0,ub:63,mandatory`
	TransmissionCombPos    TransmissionCombPos    `mandatory`
	StartPosition          int64                  `lb:0,ub:13,mandatory`
	NrofSymbols            NrofSymbols            `mandatory`
	FreqDomainShift        int64                  `lb:0,ub:268,mandatory`
	CSRS                   int64                  `lb:0,ub:63,mandatory`
	GroupOrSequenceHopping GroupOrSequenceHopping `mandatory`
	ResourceTypePos        ResourceTypePos        `mandatory`
	SequenceId             int64                  `lb:0,ub:65535,mandatory`
	SpatialRelationPos     *SpatialRelationPos    `optional`
	// IEExtensions *PosSRSResourceItemExtIEs `optional`
}

func (ie *PosSRSResourceItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.SpatialRelationPos != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)

	if err = NewINTEGER(ie.SrsPosResourceId, aper.Constraint{Lb: 0, Ub: 63}, false).Encode(w); err != nil {
		return utils.WrapError("Encode SrsPosResourceId", err)
	}
	if err = ie.TransmissionCombPos.Encode(w); err != nil {
		return utils.WrapError("Encode TransmissionCombPos", err)
	}
	if err = NewINTEGER(ie.StartPosition, aper.Constraint{Lb: 0, Ub: 13}, false).Encode(w); err != nil {
		return utils.WrapError("Encode StartPosition", err)
	}
	if err = NewENUMERATED(int64(ie.NrofSymbols), aper.Constraint{Lb: 0, Ub: 4}, false).Encode(w); err != nil {
		return utils.WrapError("Encode NrofSymbols", err)
	}
	if err = NewINTEGER(ie.FreqDomainShift, aper.Constraint{Lb: 0, Ub: 268}, false).Encode(w); err != nil {
		return utils.WrapError("Encode FreqDomainShift", err)
	}
	if err = NewINTEGER(ie.CSRS, aper.Constraint{Lb: 0, Ub: 63}, false).Encode(w); err != nil {
		return utils.WrapError("Encode CSRS", err)
	}
	if err = NewENUMERATED(int64(ie.GroupOrSequenceHopping), aper.Constraint{Lb: 0, Ub: 2}, false).Encode(w); err != nil {
		return utils.WrapError("Encode GroupOrSequenceHopping", err)
	}
	if err = ie.ResourceTypePos.Encode(w); err != nil {
		return utils.WrapError("Encode ResourceTypePos", err)
	}
	if err = NewINTEGER(ie.SequenceId, aper.Constraint{Lb: 0, Ub: 65535}, false).Encode(w); err != nil {
		return utils.WrapError("Encode SequenceId", err)
	}
	if ie.SpatialRelationPos != nil {
		if err = ie.SpatialRelationPos.Encode(w); err != nil {
			return utils.WrapError("Encode SpatialRelationPos", err)
		}
	}
	return
}

func (ie *PosSRSResourceItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}

	tmp := INTEGER{c: aper.Constraint{Lb: 0, Ub: 63}}
	if err = tmp.Decode(r); err != nil {
		return utils.WrapError("Read SrsPosResourceId", err)
	}
	ie.SrsPosResourceId = int64(tmp.Value)

	if err = ie.TransmissionCombPos.Decode(r); err != nil {
		return utils.WrapError("Read TransmissionCombPos", err)
	}
	tmp = INTEGER{c: aper.Constraint{Lb: 0, Ub: 13}}
	if err = tmp.Decode(r); err != nil {
		return utils.WrapError("Read StartPosition", err)
	}
	ie.StartPosition = int64(tmp.Value)

	tmpEnum := ENUMERATED{c: aper.Constraint{Lb: 0, Ub: 4}}
	if err = tmpEnum.Decode(r); err != nil {
		return utils.WrapError("Read NrofSymbols", err)
	}
	ie.NrofSymbols = NrofSymbolsEnum(tmpEnum.Value)

	tmp = INTEGER{c: aper.Constraint{Lb: 0, Ub: 268}}
	if err = tmp.Decode(r); err != nil {
		return utils.WrapError("Read FreqDomainShift", err)
	}
	ie.FreqDomainShift = int64(tmp.Value)

	tmp = INTEGER{c: aper.Constraint{Lb: 0, Ub: 63}}
	if err = tmp.Decode(r); err != nil {
		return utils.WrapError("Read CSRS", err)
	}
	ie.CSRS = int64(tmp.Value)

	tmpEnum = ENUMERATED{c: aper.Constraint{Lb: 0, Ub: 2}}
	if err = tmpEnum.Decode(r); err != nil {
		return utils.WrapError("Read GroupOrSequenceHopping", err)
	}
	ie.GroupOrSequenceHopping = GroupOrSequenceHoppingEnum(tmpEnum.Value)

	if err = ie.ResourceTypePos.Decode(r); err != nil {
		return utils.WrapError("Read ResourceTypePos", err)
	}

	tmp = INTEGER{c: aper.Constraint{Lb: 0, Ub: 65535}}
	if err = tmp.Decode(r); err != nil {
		return utils.WrapError("Read SequenceId", err)
	}
	ie.SequenceId = int64(tmp.Value)

	if aper.IsBitSet(optionals, 1) {
		tmpSR := new(SpatialRelationPos)
		if err = tmpSR.Decode(r); err != nil {
			return utils.WrapError("Read SpatialRelationPos", err)
		}
		ie.SpatialRelationPos = tmpSR
	}
	return
}
