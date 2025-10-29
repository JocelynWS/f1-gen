package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TSCTrafficCharacteristics struct {
	TSCAssistanceInformationDL *TSCAssistanceInformation `optional`
	TSCAssistanceInformationUL *TSCAssistanceInformation `optional`
	// IEExtensions * `optional`
}

func (ie *TSCTrafficCharacteristics) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.TSCAssistanceInformationDL != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.TSCAssistanceInformationUL != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if ie.TSCAssistanceInformationDL != nil {
		if err = ie.TSCAssistanceInformationDL.Encode(w); err != nil {
			err = utils.WrapError("Encode TSCAssistanceInformationDL", err)
			return
		}
	}
	if ie.TSCAssistanceInformationUL != nil {
		if err = ie.TSCAssistanceInformationUL.Encode(w); err != nil {
			err = utils.WrapError("Encode TSCAssistanceInformationUL", err)
			return
		}
	}
	return
}
func (ie *TSCTrafficCharacteristics) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(TSCAssistanceInformation)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read TSCAssistanceInformationDL", err)
			return
		}
		ie.TSCAssistanceInformationDL = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp := new(TSCAssistanceInformation)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read TSCAssistanceInformationUL", err)
			return
		}
		ie.TSCAssistanceInformationUL = tmp
	}
	return
}
