package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GNBCUTNLAssociationToAddItem struct {
	TNLAssociationTransportLayerAddress CPTransportLayerAddress `mandatory`
	TNLAssociationUsage                 TNLAssociationUsage     `mandatory`
	// IEExtensions * `optional`
}

func (ie *GNBCUTNLAssociationToAddItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.TNLAssociationTransportLayerAddress.Encode(w); err != nil {
		err = utils.WrapError("Encode TNLAssociationTransportLayerAddress", err)
		return
	}
	if err = ie.TNLAssociationUsage.Encode(w); err != nil {
		err = utils.WrapError("Encode TNLAssociationUsage", err)
		return
	}
	return
}
func (ie *GNBCUTNLAssociationToAddItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.TNLAssociationTransportLayerAddress.Decode(r); err != nil {
		err = utils.WrapError("Read TNLAssociationTransportLayerAddress", err)
		return
	}
	if err = ie.TNLAssociationUsage.Decode(r); err != nil {
		err = utils.WrapError("Read TNLAssociationUsage", err)
		return
	}
	return
}
