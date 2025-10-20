package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type CUDURIMInformation struct {
	VictimgNBSetID       aper.BitString       `lb:22,ub:22,mandatory`
	RIMRSDetectionStatus RIMRSDetectionStatus `mandatory`
	// IEExtensions * `optional`
}

func (ie *CUDURIMInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_VictimgNBSetID := NewBITSTRING(ie.VictimgNBSetID, aper.Constraint{Lb: 22, Ub: 22}, false)
	if err = tmp_VictimgNBSetID.Encode(w); err != nil {
		err = utils.WrapError("Encode VictimgNBSetID", err)
		return
	}
	if err = ie.RIMRSDetectionStatus.Encode(w); err != nil {
		err = utils.WrapError("Encode RIMRSDetectionStatus", err)
		return
	}
	return
}
func (ie *CUDURIMInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_VictimgNBSetID := BITSTRING{
		c:   aper.Constraint{Lb: 22, Ub: 22},
		ext: false,
	}
	if err = tmp_VictimgNBSetID.Decode(r); err != nil {
		err = utils.WrapError("Read VictimgNBSetID", err)
		return
	}
	ie.VictimgNBSetID = aper.BitString{Bytes: tmp_VictimgNBSetID.Value.Bytes, NumBits: tmp_VictimgNBSetID.Value.NumBits}
	if err = ie.RIMRSDetectionStatus.Decode(r); err != nil {
		err = utils.WrapError("Read RIMRSDetectionStatus", err)
		return
	}
	return
}
