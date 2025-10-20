package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ResourceCoordinationTransferInformation struct {
	MeNBCellID                        aper.BitString                     `lb:28,ub:28,mandatory`
	ResourceCoordinationEUTRACellInfo *ResourceCoordinationEUTRACellInfo `optional`
	// IEExtensions * `optional`
}

func (ie *ResourceCoordinationTransferInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.ResourceCoordinationEUTRACellInfo != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_MeNBCellID := NewBITSTRING(ie.MeNBCellID, aper.Constraint{Lb: 28, Ub: 28}, false)
	if err = tmp_MeNBCellID.Encode(w); err != nil {
		err = utils.WrapError("Encode MeNBCellID", err)
		return
	}
	if ie.ResourceCoordinationEUTRACellInfo != nil {
		if err = ie.ResourceCoordinationEUTRACellInfo.Encode(w); err != nil {
			err = utils.WrapError("Encode ResourceCoordinationEUTRACellInfo", err)
			return
		}
	}
	return
}
func (ie *ResourceCoordinationTransferInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_MeNBCellID := BITSTRING{
		c:   aper.Constraint{Lb: 28, Ub: 28},
		ext: false,
	}
	if err = tmp_MeNBCellID.Decode(r); err != nil {
		err = utils.WrapError("Read MeNBCellID", err)
		return
	}
	ie.MeNBCellID = aper.BitString{Bytes: tmp_MeNBCellID.Value.Bytes, NumBits: tmp_MeNBCellID.Value.NumBits}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(ResourceCoordinationEUTRACellInfo)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read ResourceCoordinationEUTRACellInfo", err)
			return
		}
		ie.ResourceCoordinationEUTRACellInfo = tmp
	}
	return
}
