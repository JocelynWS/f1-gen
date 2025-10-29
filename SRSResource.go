package f1ap

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

	tmp := NewINTEGER(ie.SRSResourceID, aper.Constraint{Lb: 0, Ub: 63}, false)
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SRSResourceID", err)
	}

	if err = ie.NrofSRSPorts.Encode(w); err != nil {
		return utils.WrapError("Encode NrofSRSPorts", err)
	}

	if err = ie.TransmissionComb.Encode(w); err != nil {
		return utils.WrapError("Encode TransmissionComb", err)
	}

	tmp3 := NewINTEGER(ie.StartPosition, aper.Constraint{Lb: 0, Ub: 13}, false)
	if err = tmp3.Encode(w); err != nil {
		return utils.WrapError("Encode StartPosition", err)
	}

	if err = ie.NrofSymbols.Encode(w); err != nil {
		return utils.WrapError("Encode NrofSymbols", err)
	}

	if err = ie.RepetitionFactor.Encode(w); err != nil {
		return utils.WrapError("Encode RepetitionFactor", err)
	}

	tmp6 := NewINTEGER(ie.FreqDomainPosition, aper.Constraint{Lb: 0, Ub: 67}, false)
	if err = tmp6.Encode(w); err != nil {
		return utils.WrapError("Encode FreqDomainPosition", err)
	}

	tmp7 := NewINTEGER(ie.FreqDomainShift, aper.Constraint{Lb: 0, Ub: 268}, false)
	if err = tmp7.Encode(w); err != nil {
		return utils.WrapError("Encode FreqDomainShift", err)
	}

	tmp8 := NewINTEGER(ie.CSRS, aper.Constraint{Lb: 0, Ub: 63}, false)
	if err = tmp8.Encode(w); err != nil {
		return utils.WrapError("Encode CSRS", err)
	}

	tmp9 := NewINTEGER(ie.BSRS, aper.Constraint{Lb: 0, Ub: 3}, false)
	if err = tmp9.Encode(w); err != nil {
		return utils.WrapError("Encode BSRS", err)
	}

	tmp10 := NewINTEGER(ie.BHop, aper.Constraint{Lb: 0, Ub: 3}, false)
	if err = tmp10.Encode(w); err != nil {
		return utils.WrapError("Encode BHop", err)
	}

	if err = ie.GroupOrSequenceHopping.Encode(w); err != nil {
		return utils.WrapError("Encode GroupOrSequenceHopping", err)
	}

	if err = ie.ResourceType.Encode(w); err != nil {
		return utils.WrapError("Encode ResourceType", err)
	}

	tmp12 := NewINTEGER(ie.SequenceId, aper.Constraint{Lb: 0, Ub: 1023}, false)
	if err = tmp12.Encode(w); err != nil {
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

	var tmp_SRSResourceID INTEGER
	tmp_SRSResourceID.c = aper.Constraint{Lb: 0, Ub: 63}
	if err = tmp_SRSResourceID.Decode(r); err != nil {
		return utils.WrapError("Read SRSResourceID", err)
	}
	ie.SRSResourceID = int64(tmp_SRSResourceID.Value)

	if err = ie.NrofSRSPorts.Decode(r); err != nil {
		return utils.WrapError("Read NrofSRSPorts", err)
	}

	if err = ie.TransmissionComb.Decode(r); err != nil {
		return utils.WrapError("Read TransmissionComb", err)
	}

	var tmp_StartPosition INTEGER
	tmp_StartPosition.c = aper.Constraint{Lb: 0, Ub: 13}
	if err = tmp_StartPosition.Decode(r); err != nil {
		return utils.WrapError("Read StartPosition", err)
	}
	ie.StartPosition = int64(tmp_StartPosition.Value)

	if err = ie.NrofSymbols.Decode(r); err != nil {
		return utils.WrapError("Read NrofSymbols", err)
	}

	if err = ie.RepetitionFactor.Decode(r); err != nil {
		return utils.WrapError("Read RepetitionFactor", err)
	}

	var tmp_FreqDomainPosition INTEGER
	tmp_FreqDomainPosition.c = aper.Constraint{Lb: 0, Ub: 67}
	if err = tmp_FreqDomainPosition.Decode(r); err != nil {
		return utils.WrapError("Read FreqDomainPosition", err)
	}
	ie.FreqDomainPosition = int64(tmp_FreqDomainPosition.Value)

	var tmp_FreqDomainShift INTEGER
	tmp_FreqDomainShift.c = aper.Constraint{Lb: 0, Ub: 268}
	if err = tmp_FreqDomainShift.Decode(r); err != nil {
		return utils.WrapError("Read FreqDomainShift", err)
	}
	ie.FreqDomainShift = int64(tmp_FreqDomainShift.Value)

	var tmp_CSRS INTEGER
	tmp_CSRS.c = aper.Constraint{Lb: 0, Ub: 63}
	if err = tmp_CSRS.Decode(r); err != nil {
		return utils.WrapError("Read CSRS", err)
	}
	ie.CSRS = int64(tmp_CSRS.Value)

	var tmp_BSRS INTEGER
	tmp_BSRS.c = aper.Constraint{Lb: 0, Ub: 3}
	if err = tmp_BSRS.Decode(r); err != nil {
		return utils.WrapError("Read BSRS", err)
	}
	ie.BSRS = int64(tmp_BSRS.Value)

	var tmp_BHop INTEGER
	tmp_BHop.c = aper.Constraint{Lb: 0, Ub: 3}
	if err = tmp_BHop.Decode(r); err != nil {
		return utils.WrapError("Read BHop", err)
	}
	ie.BHop = int64(tmp_BHop.Value)

	if err = ie.GroupOrSequenceHopping.Decode(r); err != nil {
		return utils.WrapError("Read GroupOrSequenceHopping", err)
	}

	if err = ie.ResourceType.Decode(r); err != nil {
		return utils.WrapError("Read ResourceType", err)
	}

	var tmp_SequenceId INTEGER
	tmp_SequenceId.c = aper.Constraint{Lb: 0, Ub: 1023}
	if err = tmp_SequenceId.Decode(r); err != nil {
		return utils.WrapError("Read SequenceId", err)
	}
	ie.SequenceId = int64(tmp_SequenceId.Value)

	return
}
