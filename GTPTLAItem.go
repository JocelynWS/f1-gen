package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GTPTLAItem struct {
	GTPTransportLayerAddress aper.BitString `lb:1,ub:160,mandatory,valExt`
	// IEExtensions * `optional`
}

func (ie *GTPTLAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_GTPTransportLayerAddress := NewBITSTRING(ie.GTPTransportLayerAddress, aper.Constraint{Lb: 1, Ub: 160}, true)
	if err = tmp_GTPTransportLayerAddress.Encode(w); err != nil {
		err = utils.WrapError("Encode GTPTransportLayerAddress", err)
		return
	}
	return
}
func (ie *GTPTLAItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_GTPTransportLayerAddress := BITSTRING{
		c:   aper.Constraint{Lb: 1, Ub: 160},
		ext: true,
	}
	if err = tmp_GTPTransportLayerAddress.Decode(r); err != nil {
		err = utils.WrapError("Read GTPTransportLayerAddress", err)
		return
	}
	ie.GTPTransportLayerAddress = aper.BitString{Bytes: tmp_GTPTransportLayerAddress.Value.Bytes, NumBits: tmp_GTPTransportLayerAddress.Value.NumBits}
	return
}
