package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SliceAvailableCapacityItem struct {
	PLMNIdentity                []byte                        `lb:3,ub:3,mandatory`
	SNSSAIAvailableCapacityList []SNSSAIAvailableCapacityItem `lb:1,ub:maxnoofSliceItems,mandatory,valueExt`
	// IEExtensions * `optional`
}

func (ie *SliceAvailableCapacityItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PLMNIdentity := NewOCTETSTRING(ie.PLMNIdentity, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_PLMNIdentity.Encode(w); err != nil {
		err = utils.WrapError("Encode PLMNIdentity", err)
		return
	}
	if len(ie.SNSSAIAvailableCapacityList) > 0 {
		tmp := Sequence[*SNSSAIAvailableCapacityItem]{
			Value: []*SNSSAIAvailableCapacityItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofSliceItems},
			ext:   true,
		}
		for _, i := range ie.SNSSAIAvailableCapacityList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode SNSSAIAvailableCapacityList", err)
			return
		}
	} else {
		err = utils.WrapError("SNSSAIAvailableCapacityList is nil", err)
		return
	}
	return
}
func (ie *SliceAvailableCapacityItem) Decode(r *aper.AperReader) (err error) {
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
		err = utils.WrapError("Read PLMNIdentity", err)
		return
	}
	ie.PLMNIdentity = tmp_PLMNIdentity.Value
	tmp_SNSSAIAvailableCapacityList := Sequence[*SNSSAIAvailableCapacityItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofSliceItems},
		ext: true,
	}
	fn := func() *SNSSAIAvailableCapacityItem { return new(SNSSAIAvailableCapacityItem) }
	if err = tmp_SNSSAIAvailableCapacityList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read SNSSAIAvailableCapacityList", err)
		return
	}
	ie.SNSSAIAvailableCapacityList = []SNSSAIAvailableCapacityItem{}
	for _, i := range tmp_SNSSAIAvailableCapacityList.Value {
		ie.SNSSAIAvailableCapacityList = append(ie.SNSSAIAvailableCapacityList, *i)
	}
	return
}
