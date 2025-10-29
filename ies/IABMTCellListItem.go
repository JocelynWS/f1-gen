package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type IABMTCellListItem struct {
	NRCellIdentity aper.BitString `lb:36,ub:36,mandatory`
	DURXMTRX       DURXMTRX       `mandatory`
	DUTXMTTX       DUTXMTTX       `mandatory`
	DURXMTTX       DURXMTTX       `mandatory`
	DUTXMTRX       DUTXMTRX       `mandatory`
	// IEExtensions * `optional`
}

func (ie *IABMTCellListItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_NRCellIdentity := NewBITSTRING(ie.NRCellIdentity, aper.Constraint{Lb: 36, Ub: 36}, false)
	if err = tmp_NRCellIdentity.Encode(w); err != nil {
		err = utils.WrapError("Encode NRCellIdentity", err)
		return
	}
	if err = ie.DURXMTRX.Encode(w); err != nil {
		err = utils.WrapError("Encode DURXMTRX", err)
		return
	}
	if err = ie.DUTXMTTX.Encode(w); err != nil {
		err = utils.WrapError("Encode DUTXMTTX", err)
		return
	}
	if err = ie.DURXMTTX.Encode(w); err != nil {
		err = utils.WrapError("Encode DURXMTTX", err)
		return
	}
	if err = ie.DUTXMTRX.Encode(w); err != nil {
		err = utils.WrapError("Encode DUTXMTRX", err)
		return
	}
	return
}
func (ie *IABMTCellListItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_NRCellIdentity := BITSTRING{
		c:   aper.Constraint{Lb: 36, Ub: 36},
		ext: false,
	}
	if err = tmp_NRCellIdentity.Decode(r); err != nil {
		err = utils.WrapError("Read NRCellIdentity", err)
		return
	}
	ie.NRCellIdentity = aper.BitString{Bytes: tmp_NRCellIdentity.Value.Bytes, NumBits: tmp_NRCellIdentity.Value.NumBits}
	if err = ie.DURXMTRX.Decode(r); err != nil {
		err = utils.WrapError("Read DURXMTRX", err)
		return
	}
	if err = ie.DUTXMTTX.Decode(r); err != nil {
		err = utils.WrapError("Read DUTXMTTX", err)
		return
	}
	if err = ie.DURXMTTX.Decode(r); err != nil {
		err = utils.WrapError("Read DURXMTTX", err)
		return
	}
	if err = ie.DUTXMTRX.Decode(r); err != nil {
		err = utils.WrapError("Read DUTXMTRX", err)
		return
	}
	return
}
