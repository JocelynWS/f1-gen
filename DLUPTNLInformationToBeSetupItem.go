package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DLUPTNLInformationToBeSetupItem struct {
	DLUPTNLInformation UPTransportLayerInformation `mandatory`
	// IEExtensions * `optional`
}

func (ie *DLUPTNLInformationToBeSetupItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.DLUPTNLInformation.Encode(w); err != nil {
		err = utils.WrapError("Encode DLUPTNLInformation", err)
		return
	}
	return
}
func (ie *DLUPTNLInformationToBeSetupItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.DLUPTNLInformation.Decode(r); err != nil {
		err = utils.WrapError("Read DLUPTNLInformation", err)
		return
	}
	return
}
