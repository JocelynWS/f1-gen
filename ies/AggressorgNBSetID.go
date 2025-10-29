package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AggressorgNBSetID struct {
	AggressorgNBSetID aper.BitString `lb:22,ub:22,mandatory`
	// IEExtensions * `optional`
}

func (ie *AggressorgNBSetID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_AggressorgNBSetID := NewBITSTRING(ie.AggressorgNBSetID, aper.Constraint{Lb: 22, Ub: 22}, false)
	if err = tmp_AggressorgNBSetID.Encode(w); err != nil {
		err = utils.WrapError("Encode AggressorgNBSetID", err)
		return
	}
	return
}
func (ie *AggressorgNBSetID) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_AggressorgNBSetID := BITSTRING{
		c:   aper.Constraint{Lb: 22, Ub: 22},
		ext: false,
	}
	if err = tmp_AggressorgNBSetID.Decode(r); err != nil {
		err = utils.WrapError("Read AggressorgNBSetID", err)
		return
	}
	ie.AggressorgNBSetID = aper.BitString{Bytes: tmp_AggressorgNBSetID.Value.Bytes, NumBits: tmp_AggressorgNBSetID.Value.NumBits}
	return
}
