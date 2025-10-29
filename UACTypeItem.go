package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UACTypeItem struct {
	UACReductionIndication int64           `lb:0,ub:100,mandatory`
	UACCategoryType        UACCategoryType `mandatory`
	// IEExtensions * `optional`
}

func (ie *UACTypeItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_UACReductionIndication := NewINTEGER(ie.UACReductionIndication, aper.Constraint{Lb: 0, Ub: 100}, false)
	if err = tmp_UACReductionIndication.Encode(w); err != nil {
		err = utils.WrapError("Encode UACReductionIndication", err)
		return
	}
	if err = ie.UACCategoryType.Encode(w); err != nil {
		err = utils.WrapError("Encode UACCategoryType", err)
		return
	}
	return
}
func (ie *UACTypeItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_UACReductionIndication := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 100},
		ext: false,
	}
	if err = tmp_UACReductionIndication.Decode(r); err != nil {
		err = utils.WrapError("Read UACReductionIndication", err)
		return
	}
	ie.UACReductionIndication = int64(tmp_UACReductionIndication.Value)
	if err = ie.UACCategoryType.Decode(r); err != nil {
		err = utils.WrapError("Read UACCategoryType", err)
		return
	}
	return
}
