package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PRSResourceSetItem struct {
	PRSResourceSetID         PRSResourceSetID         `madatory`
	SubcarrierSpacing        SubcarrierSpacingPRS     `madatory,valExt` // Đổi tên để tránh xung đột với SubcarrierSpacing trước đó
	PRSBandwidth             int64                    `lb:1,ub:63,madatory`
	StartPRB                 int64                    `lb:0,ub:2176,madatory`
	PointA                   int64                    `lb:0,ub:3279165,madatory`
	CombSize                 CombSize                 `madatory,valExt`
	CPType                   CPType                   `madatory,valExt`
	ResourceSetPeriodicity   ResourceSetPeriodicity   `madatory,valExt`
	ResourceSetSlotOffset    int64                    `lb:0,ub:81919,madatory,valExt`
	ResourceRepetitionFactor ResourceRepetitionFactor `madatory,valExt`
	ResourceTimeGap          ResourceTimeGap          `madatory,valExt`
	ResourceNumberofSymbols  ResourceNumberofSymbols  `madatory,valExt`
	PRSMuting                *PRSMuting               `optional`
	PRSResourceTransmitPower int64                    `lb:-60,ub:50,madatory`
	PRSResourceList          PRSResourceItem          `madatory`
	// IEExtensions *PRSResourceSetItemExtIEs `optional`
}

func (ie *PRSResourceSetItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.PRSMuting != nil {
		aper.SetBit(optionals, 0)
	}
	if err = w.WriteBits(optionals, 2); err != nil {
		return
	}

	if err = ie.PRSResourceSetID.Encode(w); err != nil {
		err = utils.WrapError("Encode PRSResourceSetID", err)
		return
	}
	if err = ie.SubcarrierSpacing.Encode(w); err != nil {
		err = utils.WrapError("Encode SubcarrierSpacing", err)
		return
	}
	tmp_PRSBandwidth := NewINTEGER(ie.PRSBandwidth, aper.Constraint{Lb: 1, Ub: 63}, false)
	if err = tmp_PRSBandwidth.Encode(w); err != nil {
		err = utils.WrapError("Encode PRSBandwidth", err)
		return
	}
	tmp_StartPRB := NewINTEGER(ie.StartPRB, aper.Constraint{Lb: 0, Ub: 2176}, false)
	if err = tmp_StartPRB.Encode(w); err != nil {
		err = utils.WrapError("Encode StartPRB", err)
		return
	}
	tmp_PointA := NewINTEGER(ie.PointA, aper.Constraint{Lb: 0, Ub: 3279165}, false)
	if err = tmp_PointA.Encode(w); err != nil {
		err = utils.WrapError("Encode PointA", err)
		return
	}
	if err = ie.CombSize.Encode(w); err != nil {
		err = utils.WrapError("Encode CombSize", err)
		return
	}
	if err = ie.CPType.Encode(w); err != nil {
		err = utils.WrapError("Encode CPType", err)
		return
	}
	if err = ie.ResourceSetPeriodicity.Encode(w); err != nil {
		err = utils.WrapError("Encode ResourceSetPeriodicity", err)
		return
	}
	tmp_ResourceSetSlotOffset := NewINTEGER(ie.ResourceSetSlotOffset, aper.Constraint{Lb: 0, Ub: 81919}, true)
	if err = tmp_ResourceSetSlotOffset.Encode(w); err != nil {
		err = utils.WrapError("Encode ResourceSetSlotOffset", err)
		return
	}
	if err = ie.ResourceRepetitionFactor.Encode(w); err != nil {
		err = utils.WrapError("Encode ResourceRepetitionFactor", err)
		return
	}
	if err = ie.ResourceTimeGap.Encode(w); err != nil {
		err = utils.WrapError("Encode ResourceTimeGap", err)
		return
	}
	if err = ie.ResourceNumberofSymbols.Encode(w); err != nil {
		err = utils.WrapError("Encode ResourceNumberofSymbols", err)
		return
	}
	if ie.PRSMuting != nil {
		if err = ie.PRSMuting.Encode(w); err != nil {
			err = utils.WrapError("Encode PRSMuting", err)
			return
		}
	}
	tmp_PRSResourceTransmitPower := NewINTEGER(ie.PRSResourceTransmitPower, aper.Constraint{Lb: -60, Ub: 50}, false)
	if err = tmp_PRSResourceTransmitPower.Encode(w); err != nil {
		err = utils.WrapError("Encode PRSResourceTransmitPower", err)
		return
	}
	if err = ie.PRSResourceList.Encode(w); err != nil {
		err = utils.WrapError("Encode PRSResourceList", err)
		return
	}

	return
}

func (ie *PRSResourceSetItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}

	if err = ie.PRSResourceSetID.Decode(r); err != nil {
		err = utils.WrapError("Read PRSResourceSetID", err)
		return
	}
	if err = ie.SubcarrierSpacing.Decode(r); err != nil {
		err = utils.WrapError("Read SubcarrierSpacing", err)
		return
	}
	tmp_PRSBandwidth := INTEGER{c: aper.Constraint{Lb: 1, Ub: 63}}
	if err = tmp_PRSBandwidth.Decode(r); err != nil {
		err = utils.WrapError("Read PRSBandwidth", err)
		return
	}
	ie.PRSBandwidth = int64(tmp_PRSBandwidth.Value)
	tmp_StartPRB := INTEGER{c: aper.Constraint{Lb: 0, Ub: 2176}}
	if err = tmp_StartPRB.Decode(r); err != nil {
		err = utils.WrapError("Read StartPRB", err)
		return
	}
	ie.StartPRB = int64(tmp_StartPRB.Value)
	tmp_PointA := INTEGER{c: aper.Constraint{Lb: 0, Ub: 3279165}}
	if err = tmp_PointA.Decode(r); err != nil {
		err = utils.WrapError("Read PointA", err)
		return
	}
	ie.PointA = int64(tmp_PointA.Value)
	if err = ie.CombSize.Decode(r); err != nil {
		err = utils.WrapError("Read CombSize", err)
		return
	}
	if err = ie.CPType.Decode(r); err != nil {
		err = utils.WrapError("Read CPType", err)
		return
	}
	if err = ie.ResourceSetPeriodicity.Decode(r); err != nil {
		err = utils.WrapError("Read ResourceSetPeriodicity", err)
		return
	}
	tmp_ResourceSetSlotOffset := INTEGER{c: aper.Constraint{Lb: 0, Ub: 81919}, ext: true}
	if err = tmp_ResourceSetSlotOffset.Decode(r); err != nil {
		err = utils.WrapError("Read ResourceSetSlotOffset", err)
		return
	}
	ie.ResourceSetSlotOffset = int64(tmp_ResourceSetSlotOffset.Value)
	if err = ie.ResourceRepetitionFactor.Decode(r); err != nil {
		err = utils.WrapError("Read ResourceRepetitionFactor", err)
		return
	}
	if err = ie.ResourceTimeGap.Decode(r); err != nil {
		err = utils.WrapError("Read ResourceTimeGap", err)
		return
	}
	if err = ie.ResourceNumberofSymbols.Decode(r); err != nil {
		err = utils.WrapError("Read ResourceNumberofSymbols", err)
		return
	}
	if aper.IsBitSet(optionals, 0) {
		tmp := new(PRSMuting)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read PRSMuting", err)
			return
		}
		ie.PRSMuting = tmp
	}
	tmp_PRSResourceTransmitPower := INTEGER{c: aper.Constraint{Lb: -60, Ub: 50}}
	if err = tmp_PRSResourceTransmitPower.Decode(r); err != nil {
		err = utils.WrapError("Read PRSResourceTransmitPower", err)
		return
	}
	ie.PRSResourceTransmitPower = int64(tmp_PRSResourceTransmitPower.Value)
	if err = ie.PRSResourceList.Decode(r); err != nil {
		err = utils.WrapError("Read PRSResourceList", err)
		return
	}

	return
}
