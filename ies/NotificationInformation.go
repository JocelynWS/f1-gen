package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type NotificationInformation struct {
	MessageIdentifier aper.BitString `lb:16,ub:16,mandatory`
	SerialNumber      aper.BitString `lb:16,ub:16,mandatory`
	// IEExtensions * `optional`
}

func (ie *NotificationInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_MessageIdentifier := NewBITSTRING(ie.MessageIdentifier, aper.Constraint{Lb: 16, Ub: 16}, false)
	if err = tmp_MessageIdentifier.Encode(w); err != nil {
		err = utils.WrapError("Encode MessageIdentifier", err)
		return
	}
	tmp_SerialNumber := NewBITSTRING(ie.SerialNumber, aper.Constraint{Lb: 16, Ub: 16}, false)
	if err = tmp_SerialNumber.Encode(w); err != nil {
		err = utils.WrapError("Encode SerialNumber", err)
		return
	}
	return
}
func (ie *NotificationInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_MessageIdentifier := BITSTRING{
		c:   aper.Constraint{Lb: 16, Ub: 16},
		ext: false,
	}
	if err = tmp_MessageIdentifier.Decode(r); err != nil {
		err = utils.WrapError("Read MessageIdentifier", err)
		return
	}
	ie.MessageIdentifier = aper.BitString{Bytes: tmp_MessageIdentifier.Value.Bytes, NumBits: tmp_MessageIdentifier.Value.NumBits}
	tmp_SerialNumber := BITSTRING{
		c:   aper.Constraint{Lb: 16, Ub: 16},
		ext: false,
	}
	if err = tmp_SerialNumber.Decode(r); err != nil {
		err = utils.WrapError("Read SerialNumber", err)
		return
	}
	ie.SerialNumber = aper.BitString{Bytes: tmp_SerialNumber.Value.Bytes, NumBits: tmp_SerialNumber.Value.NumBits}
	return
}
