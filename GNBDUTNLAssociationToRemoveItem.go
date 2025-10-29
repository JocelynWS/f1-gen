package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GNBDUTNLAssociationToRemoveItem struct {
	TNLAssociationTransportLayerAddress      CPTransportLayerAddress  `mandatory`
	TNLAssociationTransportLayerAddressgNBCU *CPTransportLayerAddress `optional`
	// IEExtensions * `optional`
}

func (ie *GNBDUTNLAssociationToRemoveItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.TNLAssociationTransportLayerAddressgNBCU != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.TNLAssociationTransportLayerAddress.Encode(w); err != nil {
		err = utils.WrapError("Encode TNLAssociationTransportLayerAddress", err)
		return
	}
	if ie.TNLAssociationTransportLayerAddressgNBCU != nil {
		if err = ie.TNLAssociationTransportLayerAddressgNBCU.Encode(w); err != nil {
			err = utils.WrapError("Encode TNLAssociationTransportLayerAddressgNBCU", err)
			return
		}
	}
	return
}
func (ie *GNBDUTNLAssociationToRemoveItem) Decode(r *aper.AperReader) (err error) {
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
		tmp := new(CPTransportLayerAddress)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read TNLAssociationTransportLayerAddressgNBCU", err)
			return
		}
		ie.TNLAssociationTransportLayerAddressgNBCU = tmp
	}
	return
}
