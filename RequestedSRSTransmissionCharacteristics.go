package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type RequestedSRSTransmissionCharacteristics struct {
	NumberOfTransmissions *int64               `lb:0,ub:500,optional,valExt`
	ResourceType          SRBResourceType      `mandatory,valExt`
	BandwidthSRS          BandwidthSRS         `mandatory`
	SRSResourceSetList    []SRSResourceSetItem `lb:1,ub:maxnoofSRSResourceSets,optional,valExt`
	SSBInformation        *SSBInformation      `optional`
	// IEExtensions *ProtocolExtensionContainer `optional`
}

func (ie *RequestedSRSTransmissionCharacteristics) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.NumberOfTransmissions != nil {
		aper.SetBit(optionals, 1)
	}
	if len(ie.SRSResourceSetList) > 0 {
		aper.SetBit(optionals, 2)
	}
	if ie.SSBInformation != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)

	if ie.NumberOfTransmissions != nil {
		tmp_NumberOfTransmissions := NewINTEGER(*ie.NumberOfTransmissions, aper.Constraint{Lb: 0, Ub: 500}, true)
		if err = tmp_NumberOfTransmissions.Encode(w); err != nil {
			err = utils.WrapError("Encode NumberOfTransmissions", err)
			return
		}
	}

	tmp_ResourceType := NewENUMERATED(int64(ie.ResourceType.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
	if err = tmp_ResourceType.Encode(w); err != nil {
		err = utils.WrapError("Encode ResourceType", err)
		return
	}

	if err = ie.BandwidthSRS.Encode(w); err != nil {
		err = utils.WrapError("Encode BandwidthSRS", err)
		return
	}

	if len(ie.SRSResourceSetList) > 0 {
		tmp := Sequence[*SRSResourceSetItem]{
			Value: []*SRSResourceSetItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoSRSResourceSets},
			ext:   true,
		}
		for _, i := range ie.SRSResourceSetList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode SRSResourceSetList", err)
			return
		}
	}

	if ie.SSBInformation != nil {
		if err = ie.SSBInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode SSBInformation", err)
			return
		}
	}
	return
}

func (ie *RequestedSRSTransmissionCharacteristics) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}

	if aper.IsBitSet(optionals, 1) {
		tmp_NumberOfTransmissions := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 500},
			ext: true,
		}
		if err = tmp_NumberOfTransmissions.Decode(r); err != nil {
			err = utils.WrapError("Read NumberOfTransmissions", err)
			return
		}
		ie.NumberOfTransmissions = (*int64)(&tmp_NumberOfTransmissions.Value)
	}

	tmp_ResourceType := ENUMERATED{
		c:   aper.Constraint{Lb: 0, Ub: 2},
		ext: true,
	}
	if err = tmp_ResourceType.Decode(r); err != nil {
		err = utils.WrapError("Read ResourceType", err)
		return
	}
	ie.ResourceType = SRBResourceType{Value: aper.Enumerated(tmp_ResourceType.Value)}

	if err = ie.BandwidthSRS.Decode(r); err != nil {
		err = utils.WrapError("Read BandwidthSRS", err)
		return
	}

	if aper.IsBitSet(optionals, 2) {
		tmp_SRSResourceSetList := Sequence[*SRSResourceSetItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoSRSResourceSets},
			ext: true,
		}
		fn := func() *SRSResourceSetItem { return new(SRSResourceSetItem) }
		if err = tmp_SRSResourceSetList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read SRSResourceSetList", err)
			return
		}
		ie.SRSResourceSetList = []SRSResourceSetItem{}
		for _, i := range tmp_SRSResourceSetList.Value {
			ie.SRSResourceSetList = append(ie.SRSResourceSetList, *i)
		}
	}

	if aper.IsBitSet(optionals, 3) {
		tmp := new(SSBInformation)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read SSBInformation", err)
			return
		}
		ie.SSBInformation = tmp
	}
	return
}
