package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ULUPTNLAddressToUpdateListItem struct {
	OldIPAdress aper.BitString `lb:1,ub:160,mandatory,valExt`
	NewIPAdress aper.BitString `lb:1,ub:160,mandatory,valExt`
	// IEExtensions * `optional`
}

func (ie *ULUPTNLAddressToUpdateListItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_OldIPAdress := NewBITSTRING(ie.OldIPAdress, aper.Constraint{Lb: 1, Ub: 160}, true)
	if err = tmp_OldIPAdress.Encode(w); err != nil {
		err = utils.WrapError("Encode OldIPAdress", err)
		return
	}
	tmp_NewIPAdress := NewBITSTRING(ie.NewIPAdress, aper.Constraint{Lb: 1, Ub: 160}, true)
	if err = tmp_NewIPAdress.Encode(w); err != nil {
		err = utils.WrapError("Encode NewIPAdress", err)
		return
	}
	return
}
func (ie *ULUPTNLAddressToUpdateListItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_OldIPAdress := BITSTRING{
		c:   aper.Constraint{Lb: 1, Ub: 160},
		ext: true,
	}
	if err = tmp_OldIPAdress.Decode(r); err != nil {
		err = utils.WrapError("Read OldIPAdress", err)
		return
	}
	ie.OldIPAdress = aper.BitString{Bytes: tmp_OldIPAdress.Value.Bytes, NumBits: tmp_OldIPAdress.Value.NumBits}
	tmp_NewIPAdress := BITSTRING{
		c:   aper.Constraint{Lb: 1, Ub: 160},
		ext: true,
	}
	if err = tmp_NewIPAdress.Decode(r); err != nil {
		err = utils.WrapError("Read NewIPAdress", err)
		return
	}
	ie.NewIPAdress = aper.BitString{Bytes: tmp_NewIPAdress.Value.Bytes, NumBits: tmp_NewIPAdress.Value.NumBits}
	return
}
