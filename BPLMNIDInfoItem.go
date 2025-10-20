package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type BPLMNIDInfoItem struct {
	PLMNIdentityList         []AvailablePLMNListItem     `lb:1,ub:maxnoofBPLMNs,mandatory`
	ExtendedPLMNIdentityList []ExtendedAvailablePLMNItem `lb:1,ub:maxnoofExtendedBPLMNs,optional`
	FiveGSTAC                []byte                      `lb:3,ub:3,optional`
	NrCellID                 aper.BitString              `lb:36,ub:36,mandatory`
	Ranac                    *int64                      `lb:0,ub:255,optional`
	// IEExtensions * `optional`
}

func (ie *BPLMNIDInfoItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.ExtendedPLMNIdentityList != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.FiveGSTAC != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.Ranac != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)
	if len(ie.PLMNIdentityList) > 0 {
		tmp := Sequence[*AvailablePLMNListItem]{
			Value: []*AvailablePLMNListItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofBPLMNs},
			ext:   false,
		}
		for _, i := range ie.PLMNIdentityList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode PLMNIdentityList", err)
			return
		}
	} else {
		err = utils.WrapError("PLMNIdentityList is nil", err)
		return
	}
	if len(ie.ExtendedPLMNIdentityList) > 0 {
		tmp := Sequence[*ExtendedAvailablePLMNItem]{
			Value: []*ExtendedAvailablePLMNItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofExtendedBPLMNs},
			ext:   false,
		}
		for _, i := range ie.ExtendedPLMNIdentityList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode ExtendedPLMNIdentityList", err)
			return
		}
	}
	if ie.FiveGSTAC != nil {
		tmp_FiveGSTAC := NewOCTETSTRING(ie.FiveGSTAC, aper.Constraint{Lb: 3, Ub: 3}, false)
		if err = tmp_FiveGSTAC.Encode(w); err != nil {
			err = utils.WrapError("Encode FiveGSTAC", err)
			return
		}
	}
	tmp_NrCellID := NewBITSTRING(ie.NrCellID, aper.Constraint{Lb: 36, Ub: 36}, false)
	if err = tmp_NrCellID.Encode(w); err != nil {
		err = utils.WrapError("Encode NrCellID", err)
		return
	}
	if ie.Ranac != nil {
		tmp_Ranac := NewINTEGER(*ie.Ranac, aper.Constraint{Lb: 0, Ub: 255}, false)
		if err = tmp_Ranac.Encode(w); err != nil {
			err = utils.WrapError("Encode Ranac", err)
			return
		}
	}
	return
}
func (ie *BPLMNIDInfoItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	tmp_PLMNIdentityList := Sequence[*AvailablePLMNListItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofBPLMNs},
		ext: false,
	}
	fn := func() *AvailablePLMNListItem { return new(AvailablePLMNListItem) }
	if err = tmp_PLMNIdentityList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read PLMNIdentityList", err)
		return
	}
	ie.PLMNIdentityList = []AvailablePLMNListItem{}
	for _, i := range tmp_PLMNIdentityList.Value {
		ie.PLMNIdentityList = append(ie.PLMNIdentityList, *i)
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_ExtendedPLMNIdentityList := Sequence[*ExtendedAvailablePLMNItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofExtendedBPLMNs},
			ext: false,
		}
		fn := func() *ExtendedAvailablePLMNItem { return new(ExtendedAvailablePLMNItem) }
		if err = tmp_ExtendedPLMNIdentityList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read ExtendedPLMNIdentityList", err)
			return
		}
		ie.ExtendedPLMNIdentityList = []ExtendedAvailablePLMNItem{}
		for _, i := range tmp_ExtendedPLMNIdentityList.Value {
			ie.ExtendedPLMNIdentityList = append(ie.ExtendedPLMNIdentityList, *i)
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_FiveGSTAC := OCTETSTRING{
			c:   aper.Constraint{Lb: 3, Ub: 3},
			ext: false,
		}
		if err = tmp_FiveGSTAC.Decode(r); err != nil {
			err = utils.WrapError("Read FiveGSTAC", err)
			return
		}
		ie.FiveGSTAC = tmp_FiveGSTAC.Value
	}
	tmp_NrCellID := BITSTRING{
		c:   aper.Constraint{Lb: 36, Ub: 36},
		ext: false,
	}
	if err = tmp_NrCellID.Decode(r); err != nil {
		err = utils.WrapError("Read NrCellID", err)
		return
	}
	ie.NrCellID = aper.BitString{Bytes: tmp_NrCellID.Value.Bytes, NumBits: tmp_NrCellID.Value.NumBits}
	if aper.IsBitSet(optionals, 3) {
		tmp_Ranac := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 255},
			ext: false,
		}
		if err = tmp_Ranac.Decode(r); err != nil {
			err = utils.WrapError("Read Ranac", err)
			return
		}
		ie.Ranac = (*int64)(&tmp_Ranac.Value)
	}
	return
}
