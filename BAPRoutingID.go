package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type BAPRoutingID struct {
	BAPAddress aper.BitString `lb:10,ub:10,mandatory`
	BAPPathID  aper.BitString `lb:10,ub:10,mandatory`
	// IEExtensions * `optional`
}

func (ie *BAPRoutingID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_BAPAddress := NewBITSTRING(ie.BAPAddress, aper.Constraint{Lb: 10, Ub: 10}, false)
	if err = tmp_BAPAddress.Encode(w); err != nil {
		err = utils.WrapError("Encode BAPAddress", err)
		return
	}
	tmp_BAPPathID := NewBITSTRING(ie.BAPPathID, aper.Constraint{Lb: 10, Ub: 10}, false)
	if err = tmp_BAPPathID.Encode(w); err != nil {
		err = utils.WrapError("Encode BAPPathID", err)
		return
	}
	return
}
func (ie *BAPRoutingID) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_BAPAddress := BITSTRING{
		c:   aper.Constraint{Lb: 10, Ub: 10},
		ext: false,
	}
	if err = tmp_BAPAddress.Decode(r); err != nil {
		err = utils.WrapError("Read BAPAddress", err)
		return
	}
	ie.BAPAddress = aper.BitString{Bytes: tmp_BAPAddress.Value.Bytes, NumBits: tmp_BAPAddress.Value.NumBits}
	tmp_BAPPathID := BITSTRING{
		c:   aper.Constraint{Lb: 10, Ub: 10},
		ext: false,
	}
	if err = tmp_BAPPathID.Decode(r); err != nil {
		err = utils.WrapError("Read BAPPathID", err)
		return
	}
	ie.BAPPathID = aper.BitString{Bytes: tmp_BAPPathID.Value.Bytes, NumBits: tmp_BAPPathID.Value.NumBits}
	return
}
