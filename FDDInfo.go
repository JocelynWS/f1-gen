package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type FDDInfo struct {
	ULNRFreqInfo            NRFreqInfo            `mandatory`
	DLNRFreqInfo            NRFreqInfo            `mandatory`
	ULTransmissionBandwidth TransmissionBandwidth `mandatory`
	DLTransmissionBandwidth TransmissionBandwidth `mandatory`
	// IEExtensions * `optional`
}

func (ie *FDDInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.ULNRFreqInfo.Encode(w); err != nil {
		err = utils.WrapError("Encode ULNRFreqInfo", err)
		return
	}
	if err = ie.DLNRFreqInfo.Encode(w); err != nil {
		err = utils.WrapError("Encode DLNRFreqInfo", err)
		return
	}
	if err = ie.ULTransmissionBandwidth.Encode(w); err != nil {
		err = utils.WrapError("Encode ULTransmissionBandwidth", err)
		return
	}
	if err = ie.DLTransmissionBandwidth.Encode(w); err != nil {
		err = utils.WrapError("Encode DLTransmissionBandwidth", err)
		return
	}
	return
}
func (ie *FDDInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.ULNRFreqInfo.Decode(r); err != nil {
		err = utils.WrapError("Read ULNRFreqInfo", err)
		return
	}
	if err = ie.DLNRFreqInfo.Decode(r); err != nil {
		err = utils.WrapError("Read DLNRFreqInfo", err)
		return
	}
	if err = ie.ULTransmissionBandwidth.Decode(r); err != nil {
		err = utils.WrapError("Read ULTransmissionBandwidth", err)
		return
	}
	if err = ie.DLTransmissionBandwidth.Decode(r); err != nil {
		err = utils.WrapError("Read DLTransmissionBandwidth", err)
		return
	}
	return
}
