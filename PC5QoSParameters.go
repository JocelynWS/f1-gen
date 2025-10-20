package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PC5QoSParameters struct {
	PC5QoSCharacteristics PC5QoSCharacteristics `mandatory`
	PC5QoSFlowBitRates    *PC5FlowBitRates      `optional`
	// IEExtensions * `optional`
}

func (ie *PC5QoSParameters) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.PC5QoSFlowBitRates != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.PC5QoSCharacteristics.Encode(w); err != nil {
		err = utils.WrapError("Encode PC5QoSCharacteristics", err)
		return
	}
	if ie.PC5QoSFlowBitRates != nil {
		if err = ie.PC5QoSFlowBitRates.Encode(w); err != nil {
			err = utils.WrapError("Encode PC5QoSFlowBitRates", err)
			return
		}
	}
	return
}
func (ie *PC5QoSParameters) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.PC5QoSCharacteristics.Decode(r); err != nil {
		err = utils.WrapError("Read PC5QoSCharacteristics", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(PC5FlowBitRates)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read PC5QoSFlowBitRates", err)
			return
		}
		ie.PC5QoSFlowBitRates = tmp
	}
	return
}
