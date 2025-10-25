package ies

import "github.com/lvdund/ngap/aper"

type UEAssociatedLogicalF1ConnectionItemResAck struct {
	UEAssociatedLogicalF1ConnectionItem UEAssociatedLogicalF1ConnectionItem `mandatory`
}

func (ie *UEAssociatedLogicalF1ConnectionItemResAck) Encode(w *aper.AperWriter) (err error) {
	if err = ie.UEAssociatedLogicalF1ConnectionItem.Encode(w); err != nil {
		return
	}
	return
}

func (ie *UEAssociatedLogicalF1ConnectionItemResAck) Decode(r *aper.AperReader) (err error) {
	if err = ie.UEAssociatedLogicalF1ConnectionItem.Decode(r); err != nil {
		return
	}
	return
}
