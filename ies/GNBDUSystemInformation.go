package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GNBDUSystemInformation struct {
	MIBMessage  []byte `lb:0,ub:0,mandatory`
	SIB1Message []byte `lb:0,ub:0,mandatory`
	// IEExtensions * `optional`
}

func (ie *GNBDUSystemInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_MIBMessage := NewOCTETSTRING(ie.MIBMessage, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_MIBMessage.Encode(w); err != nil {
		err = utils.WrapError("Encode MIBMessage", err)
		return
	}
	tmp_SIB1Message := NewOCTETSTRING(ie.SIB1Message, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_SIB1Message.Encode(w); err != nil {
		err = utils.WrapError("Encode SIB1Message", err)
		return
	}
	return
}
func (ie *GNBDUSystemInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_MIBMessage := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_MIBMessage.Decode(r); err != nil {
		err = utils.WrapError("Read MIBMessage", err)
		return
	}
	ie.MIBMessage = tmp_MIBMessage.Value
	tmp_SIB1Message := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_SIB1Message.Decode(r); err != nil {
		err = utils.WrapError("Read SIB1Message", err)
		return
	}
	ie.SIB1Message = tmp_SIB1Message.Value
	return
}
