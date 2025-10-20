package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UACAssistanceInfo struct {
	UACPLMNList []UACPLMNItem `lb:1,ub:maxnoofUACPLMNs,mandatory`
	// IEExtensions * `optional`
}

func (ie *UACAssistanceInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if len(ie.UACPLMNList) > 0 {
		tmp := Sequence[*UACPLMNItem]{
			Value: []*UACPLMNItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofUACPLMNs},
			ext:   false,
		}
		for _, i := range ie.UACPLMNList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode UACPLMNList", err)
			return
		}
	} else {
		err = utils.WrapError("UACPLMNList is nil", err)
		return
	}
	return
}
func (ie *UACAssistanceInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_UACPLMNList := Sequence[*UACPLMNItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofUACPLMNs},
		ext: false,
	}
	fn := func() *UACPLMNItem { return new(UACPLMNItem) }
	if err = tmp_UACPLMNList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read UACPLMNList", err)
		return
	}
	ie.UACPLMNList = []UACPLMNItem{}
	for _, i := range tmp_UACPLMNList.Value {
		ie.UACPLMNList = append(ie.UACPLMNList, *i)
	}
	return
}
