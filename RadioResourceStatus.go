package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type RadioResourceStatus struct {
	SSBAreaRadioResourceStatusList SSBAreaRadioResourceStatusItem `mandatory`
	// IEExtensions * `optional`
}

func (ie *RadioResourceStatus) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.SSBAreaRadioResourceStatusList.Encode(w); err != nil {
		err = utils.WrapError("Encode SSBAreaRadioResourceStatusList", err)
		return
	}
	return
}
func (ie *RadioResourceStatus) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.SSBAreaRadioResourceStatusList.Decode(r); err != nil {
		err = utils.WrapError("Read SSBAreaRadioResourceStatusList", err)
		return
	}
	return
}
