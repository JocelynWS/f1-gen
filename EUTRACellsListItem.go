package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type EUTRACellsListItem struct {
	EUTRACellID                 aper.BitString              `lb:28,ub:28,mandatory`
	ServedEUTRACellsInformation ServedEUTRACellsInformation `mandatory`
	// IEExtensions * `optional`
}

func (ie *EUTRACellsListItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_EUTRACellID := NewBITSTRING(ie.EUTRACellID, aper.Constraint{Lb: 28, Ub: 28}, false)
	if err = tmp_EUTRACellID.Encode(w); err != nil {
		err = utils.WrapError("Encode EUTRACellID", err)
		return
	}
	if err = ie.ServedEUTRACellsInformation.Encode(w); err != nil {
		err = utils.WrapError("Encode ServedEUTRACellsInformation", err)
		return
	}
	return
}
func (ie *EUTRACellsListItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_EUTRACellID := BITSTRING{
		c:   aper.Constraint{Lb: 28, Ub: 28},
		ext: false,
	}
	if err = tmp_EUTRACellID.Decode(r); err != nil {
		err = utils.WrapError("Read EUTRACellID", err)
		return
	}
	ie.EUTRACellID = aper.BitString{Bytes: tmp_EUTRACellID.Value.Bytes, NumBits: tmp_EUTRACellID.Value.NumBits}
	if err = ie.ServedEUTRACellsInformation.Decode(r); err != nil {
		err = utils.WrapError("Read ServedEUTRACellsInformation", err)
		return
	}
	return
}
