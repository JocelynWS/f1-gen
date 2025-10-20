package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type BHRoutingInformationAddedListItem struct {
	BAPRoutingID      BAPRoutingID   `mandatory`
	NextHopBAPAddress aper.BitString `lb:10,ub:10,mandatory`
	// IEExtensions * `optional`
}

func (ie *BHRoutingInformationAddedListItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.BAPRoutingID.Encode(w); err != nil {
		err = utils.WrapError("Encode BAPRoutingID", err)
		return
	}
	tmp_NextHopBAPAddress := NewBITSTRING(ie.NextHopBAPAddress, aper.Constraint{Lb: 10, Ub: 10}, false)
	if err = tmp_NextHopBAPAddress.Encode(w); err != nil {
		err = utils.WrapError("Encode NextHopBAPAddress", err)
		return
	}
	return
}
func (ie *BHRoutingInformationAddedListItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.BAPRoutingID.Decode(r); err != nil {
		err = utils.WrapError("Read BAPRoutingID", err)
		return
	}
	tmp_NextHopBAPAddress := BITSTRING{
		c:   aper.Constraint{Lb: 10, Ub: 10},
		ext: false,
	}
	if err = tmp_NextHopBAPAddress.Decode(r); err != nil {
		err = utils.WrapError("Read NextHopBAPAddress", err)
		return
	}
	ie.NextHopBAPAddress = aper.BitString{Bytes: tmp_NextHopBAPAddress.Value.Bytes, NumBits: tmp_NextHopBAPAddress.Value.NumBits}
	return
}
