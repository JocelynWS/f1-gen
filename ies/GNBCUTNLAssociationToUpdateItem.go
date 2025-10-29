package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GNBCUTNLAssociationToUpdateItem struct {
	TNLAssociationTransportLayerAddress CPTransportLayerAddress `mandatory`
	TNLAssociationUsage                 *TNLAssociationUsage    `optional`
	// IEExtensions * `optional`
}

func (ie *GNBCUTNLAssociationToUpdateItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.TNLAssociationUsage != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.TNLAssociationTransportLayerAddress.Encode(w); err != nil {
		err = utils.WrapError("Encode TNLAssociationTransportLayerAddress", err)
		return
	}
	if ie.TNLAssociationUsage != nil {
		if err = ie.TNLAssociationUsage.Encode(w); err != nil {
			err = utils.WrapError("Encode TNLAssociationUsage", err)
			return
		}
	}
	return
}
func (ie *GNBCUTNLAssociationToUpdateItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.TNLAssociationTransportLayerAddress.Decode(r); err != nil {
		err = utils.WrapError("Read TNLAssociationTransportLayerAddress", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(TNLAssociationUsage)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read TNLAssociationUsage", err)
			return
		}
		ie.TNLAssociationUsage = tmp
	}
	return
}
