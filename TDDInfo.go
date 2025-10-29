package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TDDInfo struct {
	NRFreqInfo            NRFreqInfo            `mandatory`
	TransmissionBandwidth TransmissionBandwidth `mandatory`
	// IEExtensions * `optional`
}

func (ie *TDDInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.NRFreqInfo.Encode(w); err != nil {
		err = utils.WrapError("Encode NRFreqInfo", err)
		return
	}
	if err = ie.TransmissionBandwidth.Encode(w); err != nil {
		err = utils.WrapError("Encode TransmissionBandwidth", err)
		return
	}
	return
}
func (ie *TDDInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.NRFreqInfo.Decode(r); err != nil {
		err = utils.WrapError("Read NRFreqInfo", err)
		return
	}
	if err = ie.TransmissionBandwidth.Decode(r); err != nil {
		err = utils.WrapError("Read TransmissionBandwidth", err)
		return
	}
	return
}
