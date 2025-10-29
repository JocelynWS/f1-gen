package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type EUTRAPRACHConfiguration struct {
	RootSequenceIndex    int64  `lb:0,ub:837,mandatory`
	ZeroCorrelationIndex int64  `lb:0,ub:15,mandatory`
	HighSpeedFlag        bool   `mandatory`
	PrachFreqOffset      int64  `lb:0,ub:94,mandatory`
	PrachConfigIndex     *int64 `lb:0,ub:63,optional`
}

func (ie *EUTRAPRACHConfiguration) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.PrachConfigIndex != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)

	tmpRoot := NewINTEGER(ie.RootSequenceIndex, aper.Constraint{Lb: 0, Ub: 837}, false)
	if err = tmpRoot.Encode(w); err != nil {
		return utils.WrapError("Encode RootSequenceIndex", err)
	}

	tmpZero := NewINTEGER(ie.ZeroCorrelationIndex, aper.Constraint{Lb: 0, Ub: 15}, false)
	if err = tmpZero.Encode(w); err != nil {
		return utils.WrapError("Encode ZeroCorrelationIndex", err)
	}

	if err = w.WriteBool(ie.HighSpeedFlag); err != nil {
		return utils.WrapError("Encode HighSpeedFlag", err)
	}

	tmpFreq := NewINTEGER(ie.PrachFreqOffset, aper.Constraint{Lb: 0, Ub: 94}, false)
	if err = tmpFreq.Encode(w); err != nil {
		return utils.WrapError("Encode PrachFreqOffset", err)
	}

	if ie.PrachConfigIndex != nil {
		tmpConfig := NewINTEGER(*ie.PrachConfigIndex, aper.Constraint{Lb: 0, Ub: 63}, false)
		if err = tmpConfig.Encode(w); err != nil {
			return utils.WrapError("Encode PrachConfigIndex", err)
		}
	}

	return
}

func (ie *EUTRAPRACHConfiguration) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	optionals, err := r.ReadBits(2)
	if err != nil {
		return
	}

	tmpRoot := INTEGER{c: aper.Constraint{Lb: 0, Ub: 837}, ext: false}
	if err = tmpRoot.Decode(r); err != nil {
		return utils.WrapError("Read RootSequenceIndex", err)
	}
	ie.RootSequenceIndex = int64(tmpRoot.Value)

	tmpZero := INTEGER{c: aper.Constraint{Lb: 0, Ub: 15}, ext: false}
	if err = tmpZero.Decode(r); err != nil {
		return utils.WrapError("Read ZeroCorrelationIndex", err)
	}
	ie.ZeroCorrelationIndex = int64(tmpZero.Value)

	if ie.HighSpeedFlag, err = r.ReadBool(); err != nil {
		return utils.WrapError("Read HighSpeedFlag", err)
	}

	tmpFreq := INTEGER{c: aper.Constraint{Lb: 0, Ub: 94}, ext: false}
	if err = tmpFreq.Decode(r); err != nil {
		return utils.WrapError("Read PrachFreqOffset", err)
	}
	ie.PrachFreqOffset = int64(tmpFreq.Value)

	if aper.IsBitSet(optionals, 1) {
		tmpConfig := INTEGER{c: aper.Constraint{Lb: 0, Ub: 63}, ext: false}
		if err = tmpConfig.Decode(r); err != nil {
			return utils.WrapError("Read PrachConfigIndex", err)
		}
		ie.PrachConfigIndex = (*int64)(&tmpConfig.Value)
	}

	return
}
