package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SRSResource struct {
	SRSResourceID          int64                  `lb:0,ub:63,mandatory`
	NrofSRSPorts           NrofSRSPorts           `mandatory`
	TransmissionComb       TransmissionComb       `mandatory`
	StartPosition          int64                  `lb:0,ub:13,mandatory`
	NrofSymbols            NrofSymbols            `mandatory`
	RepetitionFactor       RepetitionFactor       `mandatory`
	FreqDomainPosition     int64                  `lb:0,ub:67,mandatory`
	FreqDomainShift        int64                  `lb:0,ub:268,mandatory`
	CSRS                   int64                  `lb:0,ub:63,mandatory`
	BSRS                   int64                  `lb:0,ub:3,mandatory`
	BHop                   int64                  `lb:0,ub:3,mandatory`
	GroupOrSequenceHopping GroupOrSequenceHopping `mandatory`
	ResourceType           ResourceType           `mandatory`
	SequenceId             int64                  `lb:0,ub:1023,mandatory`
	// IEExtensions *SRSResourceExtIEs `optional`
}

func (ie *SRSResource) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)

	if err = NewINTEGER(ie.SRSResourceID, aper.Constraint{Lb: 0, Ub: 63}, false).Encode(w); err != nil {
		return utils.WrapError("Encode SRSResourceID", err)
	}

	if err = NewENUMERATED(int64(ie.NrofSRSPorts), aper.Constraint{Lb: 0, Ub: 2}, false).Encode(w); err != nil {
		return utils.WrapError("Encode NrofSRSPorts", err)
	}

	if err = ie.TransmissionComb.Encode(w); err != nil {
		return utils.WrapError("Encode TransmissionComb", err)
	}

	if err = NewINTEGER(ie.StartPosition, aper.Constraint{Lb: 0, Ub: 13}, false).Encode(w); err != nil {
		return utils.WrapError("Encode StartPosition", err)
	}

	if err = NewENUMERATED(int64(ie.NrofSymbols), aper.Constraint{Lb: 0, Ub: 2}, false).Encode(w); err != nil {
		return utils.WrapError("Encode NrofSymbols", err)
	}

	if err = NewENUMERATED(int64(ie.RepetitionFactor), aper.Constraint{Lb: 0, Ub: 2}, false).Encode(w); err != nil {
		return utils.WrapError("Encode RepetitionFactor", err)
	}

	if err = NewINTEGER(ie.FreqDomainPosition, aper.Constraint{Lb: 0, Ub: 67}, false).Encode(w); err != nil {
		return utils.WrapError("Encode FreqDomainPosition", err)
	}

	if err = NewINTEGER(ie.FreqDomainShift, aper.Constraint{Lb: 0, Ub: 268}, false).Encode(w); err != nil {
		return utils.WrapError("Encode FreqDomainShift", err)
	}

	if err = NewINTEGER(ie.CSRS, aper.Constraint{Lb: 0, Ub: 63}, false).Encode(w); err != nil {
		return utils.WrapError("Encode CSRS", err)
	}

	if err = NewINTEGER(ie.BSRS, aper.Constraint{Lb: 0, Ub: 3}, false).Encode(w); err != nil {
		return utils.WrapError("Encode BSRS", err)
	}

	if err = NewINTEGER(ie.BHop, aper.Constraint{Lb: 0, Ub: 3}, false).Encode(w); err != nil {
		return utils.WrapError("Encode BHop", err)
	}

	if err = NewENUMERATED(int64(ie.GroupOrSequenceHopping), aper.Constraint{Lb: 0, Ub: 2}, false).Encode(w); err != nil {
		return utils.WrapError("Encode GroupOrSequenceHopping", err)
	}

	if err = ie.ResourceType.Encode(w); err != nil {
		return utils.WrapError("Encode ResourceType", err)
	}

	if err = NewINTEGER(ie.SequenceId, aper.Constraint{Lb: 0, Ub: 1023}, false).Encode(w); err != nil {
		return utils.WrapError("Encode SequenceId", err)
	}

	return
}

func (ie *SRSResource) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}

	tmp_SRSResourceID := INTEGER{c: aper.Constraint{Lb: 0, Ub: 63}}
	if err = tmp_SRSResourceID.Decode(r); err != nil {
		return utils.WrapError("Read SRSResourceID", err)
	}
	ie.SRSResourceID = int64(tmp_SRSResourceID.Value)

	tmp_NrofSRSPorts := ENUMERATED{c: aper.Constraint{Lb: 0, Ub: 2}}
	if err = tmp_NrofSRSPorts.Decode(r); err != nil {
		return utils.WrapError("Read NrofSRSPorts", err)
	}
	ie.NrofSRSPorts = NrofSRSPorts(tmp_NrofSRSPorts.Value)

	if err = ie.TransmissionComb.Decode(r); err != nil {
		return utils.WrapError("Read TransmissionComb", err)
	}

	tmp_StartPosition := INTEGER{c: aper.Constraint{Lb: 0, Ub: 13}}
	if err = tmp_StartPosition.Decode(r); err != nil {
		return utils.WrapError("Read StartPosition", err)
	}
	ie.StartPosition = int64(tmp_StartPosition.Value)

	tmp_NrofSymbols := ENUMERATED{c: aper.Constraint{Lb: 0, Ub: 2}}
	if err = tmp_NrofSymbols.Decode(r); err != nil {
		return utils.WrapError("Read NrofSymbols", err)
	}
	ie.NrofSymbols = NrofSymbols(tmp_NrofSymbols.Value)

	tmp_RepetitionFactor := ENUMERATED{c: aper.Constraint{Lb: 0, Ub: 2}}
	if err = tmp_RepetitionFactor.Decode(r); err != nil {
		return utils.WrapError("Read RepetitionFactor", err)
	}
	ie.RepetitionFactor = RepetitionFactor(tmp_RepetitionFactor.Value)

	tmp_FreqDomainPosition := INTEGER{c: aper.Constraint{Lb: 0, Ub: 67}}
	if err = tmp_FreqDomainPosition.Decode(r); err != nil {
		return utils.WrapError("Read FreqDomainPosition", err)
	}
	ie.FreqDomainPosition = int64(tmp_FreqDomainPosition.Value)

	tmp_FreqDomainShift := INTEGER{c: aper.Constraint{Lb: 0, Ub: 268}}
	if err = tmp_FreqDomainShift.Decode(r); err != nil {
		return utils.WrapError("Read FreqDomainShift", err)
	}
	ie.FreqDomainShift = int64(tmp_FreqDomainShift.Value)

	tmp_CSRS := INTEGER{c: aper.Constraint{Lb: 0, Ub: 63}}
	if err = tmp_CSRS.Decode(r); err != nil {
		return utils.WrapError("Read CSRS", err)
	}
	ie.CSRS = int64(tmp_CSRS.Value)

	tmp_BSRS := INTEGER{c: aper.Constraint{Lb: 0, Ub: 3}}
	if err = tmp_BSRS.Decode(r); err != nil {
		return utils.WrapError("Read BSRS", err)
	}
	ie.BSRS = int64(tmp_BSRS.Value)

	tmp_BHop := INTEGER{c: aper.Constraint{Lb: 0, Ub: 3}}
	if err = tmp_BHop.Decode(r); err != nil {
		return utils.WrapError("Read BHop", err)
	}
	ie.BHop = int64(tmp_BHop.Value)

	tmp_GroupOrSequenceHopping := ENUMERATED{c: aper.Constraint{Lb: 0, Ub: 2}}
	if err = tmp_GroupOrSequenceHopping.Decode(r); err != nil {
		return utils.WrapError("Read GroupOrSequenceHopping", err)
	}
	ie.GroupOrSequenceHopping = GroupOrSequenceHopping(tmp_GroupOrSequenceHopping.Value)

	if err = ie.ResourceType.Decode(r); err != nil {
		return utils.WrapError("Read ResourceType", err)
	}

	tmp_SequenceId := INTEGER{c: aper.Constraint{Lb: 0, Ub: 1023}}
	if err = tmp_SequenceId.Decode(r); err != nil {
		return utils.WrapError("Read SequenceId", err)
	}
	ie.SequenceId = int64(tmp_SequenceId.Value)

	return
}
