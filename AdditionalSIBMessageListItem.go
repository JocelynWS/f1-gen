package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AdditionalSIBMessageListItem struct {
	AdditionalSIB []byte `lb:0,ub:0,mandatory`
	// IEExtensions * `optional`
}

func (ie *AdditionalSIBMessageListItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_AdditionalSIB := NewOCTETSTRING(ie.AdditionalSIB, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_AdditionalSIB.Encode(w); err != nil {
		err = utils.WrapError("Encode AdditionalSIB", err)
		return
	}
	return
}
func (ie *AdditionalSIBMessageListItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_AdditionalSIB := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_AdditionalSIB.Decode(r); err != nil {
		err = utils.WrapError("Read AdditionalSIB", err)
		return
	}
	ie.AdditionalSIB = tmp_AdditionalSIB.Value
	return
}
