package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type NRPRACHConfigItem struct {
	NRSCS                     NRSCS              `mandatory`
	PrachFreqStartfromCarrier int64              `lb:0,ub:maxNoOfPhysicalResourceBlocks-1,mandatory,valExt`
	Msg1FDM                   Msg1FDM            `madatory,valExt`
	ParchConfigIndex          int64              `lb:0,ub:255,mandatory,valExt`
	SsbPerRACHOccasion        SsbPerRACHOccasion `madatory,valExt`
	FreqDomainLength          FreqDomainLength   `mandatory`
	ZeroCorrelZoneConfig      int64              `lb:0,ub:15,mandatory`
	// IEExtension *NRPRACHConfigItemExtIEs `optional`
}

// Encode APER
func (ie *NRPRACHConfigItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0} // currently no optional
	w.WriteBits(optionals, 1)

	if err = ie.NRSCS.Encode(w); err != nil {
		err = utils.WrapError("Encode NRSCS", err)
		return
	}

	tmp_PrachFreqStartfromCarrier := NewINTEGER(
		ie.PrachFreqStartfromCarrier,
		aper.Constraint{Lb: 0, Ub: maxNoOfPhysicalResourceBlocks - 1},
		true,
	)
	if err = tmp_PrachFreqStartfromCarrier.Encode(w); err != nil {
		err = utils.WrapError("Encode PrachFreqStartfromCarrier", err)
		return
	}

	tmp_Msg1FDM := NewENUMERATED(int64(ie.Msg1FDM), aper.Constraint{Lb: 0, Ub: 3}, false)
	if err = tmp_Msg1FDM.Encode(w); err != nil {
		err = utils.WrapError("Encode Msg1FDM", err)
		return
	}

	tmp_ParchConfigIndex := NewINTEGER(ie.ParchConfigIndex, aper.Constraint{Lb: 0, Ub: 255}, true)
	if err = tmp_ParchConfigIndex.Encode(w); err != nil {
		err = utils.WrapError("Encode ParchConfigIndex", err)
		return
	}

	tmp_SsbPerRACHOccasion := NewENUMERATED(int64(ie.SsbPerRACHOccasion), aper.Constraint{Lb: 0, Ub: 7}, false)
	if err = tmp_SsbPerRACHOccasion.Encode(w); err != nil {
		err = utils.WrapError("Encode SsbPerRACHOccasion", err)
		return
	}

	if err = ie.FreqDomainLength.Encode(w); err != nil {
		err = utils.WrapError("Encode FreqDomainLength", err)
		return
	}

	tmp_ZeroCorrelZoneConfig := NewINTEGER(ie.ZeroCorrelZoneConfig, aper.Constraint{Lb: 0, Ub: 15}, false)
	if err = tmp_ZeroCorrelZoneConfig.Encode(w); err != nil {
		err = utils.WrapError("Encode ZeroCorrelZoneConfig", err)
		return
	}

	return
}

// Decode APER
func (ie *NRPRACHConfigItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil { // no optional
		return
	}

	if err = ie.NRSCS.Decode(r); err != nil {
		err = utils.WrapError("Read NRSCS", err)
		return
	}

	tmp_PrachFreqStartfromCarrier := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: maxNoOfPhysicalResourceBlocks - 1},
		ext: true,
	}
	if err = tmp_PrachFreqStartfromCarrier.Decode(r); err != nil {
		err = utils.WrapError("Read PrachFreqStartfromCarrier", err)
		return
	}
	ie.PrachFreqStartfromCarrier = int64(tmp_PrachFreqStartfromCarrier.Value)

	tmp_Msg1FDM := ENUMERATED{
		c:   aper.Constraint{Lb: 0, Ub: 3},
		ext: false,
	}
	if err = tmp_Msg1FDM.Decode(r); err != nil {
		err = utils.WrapError("Read Msg1FDM", err)
		return
	}
	ie.Msg1FDM = Msg1FDM(tmp_Msg1FDM.Value)

	tmp_ParchConfigIndex := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 255},
		ext: true,
	}
	if err = tmp_ParchConfigIndex.Decode(r); err != nil {
		err = utils.WrapError("Read ParchConfigIndex", err)
		return
	}
	ie.ParchConfigIndex = int64(tmp_ParchConfigIndex.Value)

	tmp_SsbPerRACHOccasion := ENUMERATED{
		c:   aper.Constraint{Lb: 0, Ub: 7},
		ext: false,
	}
	if err = tmp_SsbPerRACHOccasion.Decode(r); err != nil {
		err = utils.WrapError("Read SsbPerRACHOccasion", err)
		return
	}
	ie.SsbPerRACHOccasion = SsbPerRACHOccasion(tmp_SsbPerRACHOccasion.Value)

	if err = ie.FreqDomainLength.Decode(r); err != nil {
		err = utils.WrapError("Read FreqDomainLength", err)
		return
	}

	tmp_ZeroCorrelZoneConfig := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 15},
		ext: false,
	}
	if err = tmp_ZeroCorrelZoneConfig.Decode(r); err != nil {
		err = utils.WrapError("Read ZeroCorrelZoneConfig", err)
		return
	}
	ie.ZeroCorrelZoneConfig = int64(tmp_ZeroCorrelZoneConfig.Value)

	return
}
