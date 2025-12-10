package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PWSSystemInformation struct {
	SIBtype    int64  `lb:6,ub:8,mandatory,valueExt`
	SIBmessage []byte `lb:0,ub:0,mandatory`
	// IEExtensions * `optional`
}

func (ie *PWSSystemInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_SIBtype := NewINTEGER(ie.SIBtype, aper.Constraint{Lb: 6, Ub: 8}, true)
	if err = tmp_SIBtype.Encode(w); err != nil {
		err = utils.WrapError("Encode SIBtype", err)
		return
	}
	tmp_SIBmessage := NewOCTETSTRING(ie.SIBmessage, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_SIBmessage.Encode(w); err != nil {
		err = utils.WrapError("Encode SIBmessage", err)
		return
	}
	return
}
func (ie *PWSSystemInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_SIBtype := INTEGER{
		c:   aper.Constraint{Lb: 6, Ub: 8},
		ext: true,
	}
	if err = tmp_SIBtype.Decode(r); err != nil {
		err = utils.WrapError("Read SIBtype", err)
		return
	}
	ie.SIBtype = int64(tmp_SIBtype.Value)
	tmp_SIBmessage := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_SIBmessage.Decode(r); err != nil {
		err = utils.WrapError("Read SIBmessage", err)
		return
	}
	ie.SIBmessage = tmp_SIBmessage.Value
	return
}
