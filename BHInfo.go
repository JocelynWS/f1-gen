package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type BHInfo struct {
	BAProutingID      *BAPRoutingID       `optional`
	EgressBHRLCCHList []EgressBHRLCCHItem `lb:1,ub:maxnoofEgressLinks,optional`
	// IEExtensions * `optional`
}

func (ie *BHInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.BAProutingID != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.EgressBHRLCCHList != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if ie.BAProutingID != nil {
		if err = ie.BAProutingID.Encode(w); err != nil {
			err = utils.WrapError("Encode BAProutingID", err)
			return
		}
	}
	if len(ie.EgressBHRLCCHList) > 0 {
		tmp := Sequence[*EgressBHRLCCHItem]{
			Value: []*EgressBHRLCCHItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofEgressLinks},
			ext:   false,
		}
		for _, i := range ie.EgressBHRLCCHList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode EgressBHRLCCHList", err)
			return
		}
	}
	return
}
func (ie *BHInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(BAPRoutingID)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read BAProutingID", err)
			return
		}
		ie.BAProutingID = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_EgressBHRLCCHList := Sequence[*EgressBHRLCCHItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEgressLinks},
			ext: false,
		}
		fn := func() *EgressBHRLCCHItem { return new(EgressBHRLCCHItem) }
		if err = tmp_EgressBHRLCCHList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read EgressBHRLCCHList", err)
			return
		}
		ie.EgressBHRLCCHList = []EgressBHRLCCHItem{}
		for _, i := range tmp_EgressBHRLCCHList.Value {
			ie.EgressBHRLCCHList = append(ie.EgressBHRLCCHList, *i)
		}
	}
	return
}
