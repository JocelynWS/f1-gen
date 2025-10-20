package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SSBTFConfiguration struct {
	SSBFrequency          int64                `lb:0,ub:3279165,mandatory`
	SSBSubcarrierSpacing  SSBSubcarrierSpacing `mandatory`
	SSBTransmitPower      int64                `lb:-60,ub:50,mandatory`
	SSBPeriodicity        SSBPeriodicity       `mandatory`
	SSBHalfFrameOffset    int64                `lb:0,ub:1,mandatory`
	SSBSFNOffset          int64                `lb:0,ub:15,mandatory`
	SSBPositionInBurst    *SSBPositionsInBurst `optional`
	SFNInitialisationTime *aper.BitString      `lb:64,ub:64,optional`
	// IEExtensions * `optional`
}

func (ie *SSBTFConfiguration) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.SSBPositionInBurst != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.SFNInitialisationTime != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)

	if err = NewINTEGER(ie.SSBFrequency, aper.Constraint{Lb: 0, Ub: 3279165}, false).Encode(w); err != nil {
		return utils.WrapError("Encode SSBFrequency", err)
	}

	if err = NewENUMERATED(int64(ie.SSBSubcarrierSpacing), aper.Constraint{Lb: 0, Ub: 4}, false).Encode(w); err != nil {
		return utils.WrapError("Encode SSBSubcarrierSpacing", err)
	}

	if err = NewINTEGER(ie.SSBTransmitPower, aper.Constraint{Lb: -60, Ub: 50}, false).Encode(w); err != nil {
		return utils.WrapError("Encode SSBTransmitPower", err)
	}

	if err = NewENUMERATED(int64(ie.SSBPeriodicity), aper.Constraint{Lb: 0, Ub: 5}, false).Encode(w); err != nil {
		return utils.WrapError("Encode SSBPeriodicity", err)
	}

	if err = NewINTEGER(ie.SSBHalfFrameOffset, aper.Constraint{Lb: 0, Ub: 1}, false).Encode(w); err != nil {
		return utils.WrapError("Encode SSBHalfFrameOffset", err)
	}

	if err = NewINTEGER(ie.SSBSFNOffset, aper.Constraint{Lb: 0, Ub: 15}, false).Encode(w); err != nil {
		return utils.WrapError("Encode SSBSFNOffset", err)
	}

	if ie.SSBPositionInBurst != nil {
		if err = ie.SSBPositionInBurst.Encode(w); err != nil {
			return utils.WrapError("Encode SSBPositionInBurst", err)
		}
	}

	if ie.SFNInitialisationTime != nil {
		if err = NewBITSTRING(ie.SFNInitialisationTime, aper.Constraint{Lb: 64, Ub: 64}, false).Encode(w); err != nil {
			return utils.WrapError("Encode SFNInitialisationTime", err)
		}
	}

	return
}

func (ie *SSBTFConfiguration) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}

	tmp_SSBFrequency := INTEGER{c: aper.Constraint{Lb: 0, Ub: 3279165}}
	if err = tmp_SSBFrequency.Decode(r); err != nil {
		return utils.WrapError("Read SSBFrequency", err)
	}
	ie.SSBFrequency = int64(tmp_SSBFrequency.Value)

	tmp_SSBSubcarrierSpacing := ENUMERATED{c: aper.Constraint{Lb: 0, Ub: 4}}
	if err = tmp_SSBSubcarrierSpacing.Decode(r); err != nil {
		return utils.WrapError("Read SSBSubcarrierSpacing", err)
	}
	ie.SSBSubcarrierSpacing = SSBSubcarrierSpacing(tmp_SSBSubcarrierSpacing.Value)

	tmp_SSBTransmitPower := INTEGER{c: aper.Constraint{Lb: -60, Ub: 50}}
	if err = tmp_SSBTransmitPower.Decode(r); err != nil {
		return utils.WrapError("Read SSBTransmitPower", err)
	}
	ie.SSBTransmitPower = int64(tmp_SSBTransmitPower.Value)

	tmp_SSBPeriodicity := ENUMERATED{c: aper.Constraint{Lb: 0, Ub: 5}}
	if err = tmp_SSBPeriodicity.Decode(r); err != nil {
		return utils.WrapError("Read SSBPeriodicity", err)
	}
	ie.SSBPeriodicity = SSBPeriodicity(tmp_SSBPeriodicity.Value)

	tmp_SSBHalfFrameOffset := INTEGER{c: aper.Constraint{Lb: 0, Ub: 1}}
	if err = tmp_SSBHalfFrameOffset.Decode(r); err != nil {
		return utils.WrapError("Read SSBHalfFrameOffset", err)
	}
	ie.SSBHalfFrameOffset = int64(tmp_SSBHalfFrameOffset.Value)

	tmp_SSBSFNOffset := INTEGER{c: aper.Constraint{Lb: 0, Ub: 15}}
	if err = tmp_SSBSFNOffset.Decode(r); err != nil {
		return utils.WrapError("Read SSBSFNOffset", err)
	}
	ie.SSBSFNOffset = int64(tmp_SSBSFNOffset.Value)

	if aper.IsBitSet(optionals, 1) {
		tmp := new(SSBPositionsInBurst)
		if err = tmp.Decode(r); err != nil {
			return utils.WrapError("Read SSBPositionInBurst", err)
		}
		ie.SSBPositionInBurst = tmp
	}

	if aper.IsBitSet(optionals, 2) {
		tmp_SFNInitialisationTime := BITSTRING{c: aper.Constraint{Lb: 64, Ub: 64}}
		if err = tmp_SFNInitialisationTime.Decode(r); err != nil {
			return utils.WrapError("Read SFNInitialisationTime", err)
		}
		ie.SFNInitialisationTime = &aper.BitString{
			Bytes:   tmp_SFNInitialisationTime.Value.Bytes,
			NumBits: tmp_SFNInitialisationTime.Value.NumBits,
		}
	}

	return
}
