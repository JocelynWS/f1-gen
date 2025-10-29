package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ResourceCoordinationEUTRACellInfo struct {
	EUTRAModeInfo           EUTRACoexModeInfo       `mandatory`
	EUTRAPRACHConfiguration EUTRAPRACHConfiguration `mandatory`
	// IEExtensions * `optional`
}

func (ie *ResourceCoordinationEUTRACellInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.EUTRAModeInfo.Encode(w); err != nil {
		err = utils.WrapError("Encode EUTRAModeInfo", err)
		return
	}
	if err = ie.EUTRAPRACHConfiguration.Encode(w); err != nil {
		err = utils.WrapError("Encode EUTRAPRACHConfiguration", err)
		return
	}
	return
}
func (ie *ResourceCoordinationEUTRACellInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.EUTRAModeInfo.Decode(r); err != nil {
		err = utils.WrapError("Read EUTRAModeInfo", err)
		return
	}
	if err = ie.EUTRAPRACHConfiguration.Decode(r); err != nil {
		err = utils.WrapError("Read EUTRAPRACHConfiguration", err)
		return
	}
	return
}
