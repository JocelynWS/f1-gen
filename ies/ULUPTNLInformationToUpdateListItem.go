package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ULUPTNLInformationtoUpdateListItem struct {
	ULUPTNLInformation    UPTransportLayerInformation  `mandatory`
	NewULUPTNLInformation *UPTransportLayerInformation `optional`
	BHInfo                BHInfo                       `mandatory`
	// IEExtensions * `optional`
}

func (ie *ULUPTNLInformationtoUpdateListItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.NewULUPTNLInformation != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.ULUPTNLInformation.Encode(w); err != nil {
		err = utils.WrapError("Encode ULUPTNLInformation", err)
		return
	}
	if ie.NewULUPTNLInformation != nil {
		if err = ie.NewULUPTNLInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode NewULUPTNLInformation", err)
			return
		}
	}
	if err = ie.BHInfo.Encode(w); err != nil {
		err = utils.WrapError("Encode BHInfo", err)
		return
	}
	return
}
func (ie *ULUPTNLInformationtoUpdateListItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.ULUPTNLInformation.Decode(r); err != nil {
		err = utils.WrapError("Read ULUPTNLInformation", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(UPTransportLayerInformation)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read NewULUPTNLInformation", err)
			return
		}
		ie.NewULUPTNLInformation = tmp
	}
	if err = ie.BHInfo.Decode(r); err != nil {
		err = utils.WrapError("Read BHInfo", err)
		return
	}
	return
}
