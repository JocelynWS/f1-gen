package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TransportUPLayerAddressInfoToAddItem struct {
	IPSecTransportLayerAddress    aper.BitString `lb:1,ub:160,mandatory,valExt`
	GTPTransportLayerAddressToAdd []GTPTLAItem   `lb:1,ub:maxnoofGTPTLAs,optional,valExt`
	// IEExtensions * `optional`
}

func (ie *TransportUPLayerAddressInfoToAddItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.GTPTransportLayerAddressToAdd != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_IPSecTransportLayerAddress := NewBITSTRING(ie.IPSecTransportLayerAddress, aper.Constraint{Lb: 1, Ub: 160}, true)
	if err = tmp_IPSecTransportLayerAddress.Encode(w); err != nil {
		err = utils.WrapError("Encode IPSecTransportLayerAddress", err)
		return
	}
	if len(ie.GTPTransportLayerAddressToAdd) > 0 {
		tmp := Sequence[*GTPTLAItem]{
			Value: []*GTPTLAItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofGTPTLAs},
			ext:   true,
		}
		for _, i := range ie.GTPTransportLayerAddressToAdd {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode GTPTransportLayerAddressToAdd", err)
			return
		}
	}
	return
}
func (ie *TransportUPLayerAddressInfoToAddItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_IPSecTransportLayerAddress := BITSTRING{
		c:   aper.Constraint{Lb: 1, Ub: 160},
		ext: true,
	}
	if err = tmp_IPSecTransportLayerAddress.Decode(r); err != nil {
		err = utils.WrapError("Read IPSecTransportLayerAddress", err)
		return
	}
	ie.IPSecTransportLayerAddress = aper.BitString{Bytes: tmp_IPSecTransportLayerAddress.Value.Bytes, NumBits: tmp_IPSecTransportLayerAddress.Value.NumBits}
	if aper.IsBitSet(optionals, 1) {
		tmp_GTPTransportLayerAddressToAdd := Sequence[*GTPTLAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofGTPTLAs},
			ext: true,
		}
		fn := func() *GTPTLAItem { return new(GTPTLAItem) }
		if err = tmp_GTPTransportLayerAddressToAdd.Decode(r, fn); err != nil {
			err = utils.WrapError("Read GTPTransportLayerAddressToAdd", err)
			return
		}
		ie.GTPTransportLayerAddressToAdd = []GTPTLAItem{}
		for _, i := range tmp_GTPTransportLayerAddressToAdd.Value {
			ie.GTPTransportLayerAddressToAdd = append(ie.GTPTransportLayerAddressToAdd, *i)
		}
	}
	return
}
