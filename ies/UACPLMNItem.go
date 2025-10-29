package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UACPLMNItem struct {
	PLMNIdentity []byte      `lb:3,ub:3,mandatory`
	UACTypeItem  UACTypeItem `lb:1,ub:maxnoofUACperPLMN,mandatory`
}

func (ie *UACPLMNItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	if err = w.WriteBits([]byte{0x0}, 1); err != nil {
		return
	}

	tmp_PLMNIdentity := NewOCTETSTRING(ie.PLMNIdentity, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_PLMNIdentity.Encode(w); err != nil {
		err = utils.WrapError("Encode PLMNIdentity", err)
		return
	}

	if err = ie.UACTypeItem.Encode(w); err != nil {
		err = utils.WrapError("Encode UACTypeItem", err)
		return
	}

	return
}

func (ie *UACPLMNItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	if _, err = r.ReadBits(1); err != nil {
		return
	}

	tmp_PLMNIdentity := OCTETSTRING{
		c:   aper.Constraint{Lb: 3, Ub: 3},
		ext: false,
	}
	if err = tmp_PLMNIdentity.Decode(r); err != nil {
		err = utils.WrapError("Decode PLMNIdentity", err)
		return
	}
	ie.PLMNIdentity = tmp_PLMNIdentity.Value

	if err = ie.UACTypeItem.Decode(r); err != nil {
		err = utils.WrapError("Decode UACTypeItem", err)
		return
	}

	return
}
