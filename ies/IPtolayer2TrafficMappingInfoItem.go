package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type IPtolayer2TrafficMappingInfoItem struct {
	MappingInformationIndex aper.BitString      `lb:1,ub:65536,mandatory,valueExt`
	IPHeaderInformation     IPHeaderInformation `mandatory`
	BHInfo                  BHInfo              `mandatory`
	// IEExtensions *ProtocolExtensionContainerIPtolayer2TrafficMappingInfoItemExtIEs `optional`
}

func (ie *IPtolayer2TrafficMappingInfoItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)

	tmp_MappingInformationIndex := NewBITSTRING(ie.MappingInformationIndex, aper.Constraint{Lb: 1, Ub: 65536}, true)
	if err = tmp_MappingInformationIndex.Encode(w); err != nil {
		err = utils.WrapError("Encode MappingInformationIndex", err)
		return
	}

	if err = ie.IPHeaderInformation.Encode(w); err != nil {
		err = utils.WrapError("Encode IPHeaderInformation", err)
		return
	}

	if err = ie.BHInfo.Encode(w); err != nil {
		err = utils.WrapError("Encode BHInfo", err)
		return
	}

	return
}

func (ie *IPtolayer2TrafficMappingInfoItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	if _, err = r.ReadBits(1); err != nil {
		return
	}

	tmp_MappingInformationIndex := BITSTRING{
		c:   aper.Constraint{Lb: 1, Ub: 65536},
		ext: true,
	}
	if err = tmp_MappingInformationIndex.Decode(r); err != nil {
		err = utils.WrapError("Read MappingInformationIndex", err)
		return
	}
	ie.MappingInformationIndex = aper.BitString{
		Bytes:   tmp_MappingInformationIndex.Value.Bytes,
		NumBits: tmp_MappingInformationIndex.Value.NumBits,
	}

	if err = ie.IPHeaderInformation.Decode(r); err != nil {
		err = utils.WrapError("Read IPHeaderInformation", err)
		return
	}

	if err = ie.BHInfo.Decode(r); err != nil {
		err = utils.WrapError("Read BHInfo", err)
		return
	}

	return
}
