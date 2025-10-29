package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ULUPTNLInformationToBeSetupItem struct {
	ULUPTNLInformation UPTransportLayerInformation
}

func (ie *ULUPTNLInformationToBeSetupItem) Encode(w *aper.AperWriter) (err error) {
	if err = ie.ULUPTNLInformation.Encode(w); err != nil {
		return utils.WrapError("Encode ULUPTNLInformation", err)
	}

	return nil
}

func (ie *ULUPTNLInformationToBeSetupItem) Decode(r *aper.AperReader) (err error) {
	if err = ie.ULUPTNLInformation.Decode(r); err != nil {
		return utils.WrapError("Decode ULUPTNLInformation", err)
	}

	return nil
}
