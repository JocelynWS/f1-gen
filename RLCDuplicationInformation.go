package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type RLCDuplicationInformation struct {
	RLCDuplicationStateList []RLCDuplicationStateItem `lb:1,ub:maxnoofRLCDuplicationState,mandatory`
	PrimaryPathIndication   *PrimaryPathIndication    `optional`
	// IEExtensions * `optional`
}

func (ie *RLCDuplicationInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.PrimaryPathIndication != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if len(ie.RLCDuplicationStateList) > 0 {
		tmp := Sequence[*RLCDuplicationStateItem]{
			Value: []*RLCDuplicationStateItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofRLCDuplicationState},
			ext:   false,
		}
		for _, i := range ie.RLCDuplicationStateList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode RLCDuplicationStateList", err)
			return
		}
	} else {
		err = utils.WrapError("RLCDuplicationStateList is nil", err)
		return
	}
	if ie.PrimaryPathIndication != nil {
		if err = ie.PrimaryPathIndication.Encode(w); err != nil {
			err = utils.WrapError("Encode PrimaryPathIndication", err)
			return
		}
	}
	return
}
func (ie *RLCDuplicationInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_RLCDuplicationStateList := Sequence[*RLCDuplicationStateItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofRLCDuplicationState},
		ext: false,
	}
	fn := func() *RLCDuplicationStateItem { return new(RLCDuplicationStateItem) }
	if err = tmp_RLCDuplicationStateList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read RLCDuplicationStateList", err)
		return
	}
	ie.RLCDuplicationStateList = []RLCDuplicationStateItem{}
	for _, i := range tmp_RLCDuplicationStateList.Value {
		ie.RLCDuplicationStateList = append(ie.RLCDuplicationStateList, *i)
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(PrimaryPathIndication)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read PrimaryPathIndication", err)
			return
		}
		ie.PrimaryPathIndication = tmp
	}
	return
}
