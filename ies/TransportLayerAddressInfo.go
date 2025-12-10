package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TransportLayerAddressInfo struct {
	TransportUPLayerAddressInfoToAddList    []TransportUPLayerAddressInfoToAddItem    `lb:1,ub:maxnoofTLAIAB,optional,valueExt`
	TransportUPLayerAddressInfoToRemoveList []TransportUPLayerAddressInfoToRemoveItem `lb:1,ub:maxnoofTLAIAB,optional,valueExt`
	// IEExtensions * `optional`
}

func (ie *TransportLayerAddressInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.TransportUPLayerAddressInfoToAddList != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.TransportUPLayerAddressInfoToRemoveList != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if len(ie.TransportUPLayerAddressInfoToAddList) > 0 {
		tmp := Sequence[*TransportUPLayerAddressInfoToAddItem]{
			Value: []*TransportUPLayerAddressInfoToAddItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofTLAs},
			ext:   true,
		}
		for _, i := range ie.TransportUPLayerAddressInfoToAddList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode TransportUPLayerAddressInfoToAddList", err)
			return
		}
	}
	if len(ie.TransportUPLayerAddressInfoToRemoveList) > 0 {
		tmp := Sequence[*TransportUPLayerAddressInfoToRemoveItem]{
			Value: []*TransportUPLayerAddressInfoToRemoveItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofTLAs},
			ext:   true,
		}
		for _, i := range ie.TransportUPLayerAddressInfoToRemoveList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode TransportUPLayerAddressInfoToRemoveList", err)
			return
		}
	}
	return
}
func (ie *TransportLayerAddressInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_TransportUPLayerAddressInfoToAddList := Sequence[*TransportUPLayerAddressInfoToAddItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTLAs},
			ext: true,
		}
		fn := func() *TransportUPLayerAddressInfoToAddItem { return new(TransportUPLayerAddressInfoToAddItem) }
		if err = tmp_TransportUPLayerAddressInfoToAddList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read TransportUPLayerAddressInfoToAddList", err)
			return
		}
		ie.TransportUPLayerAddressInfoToAddList = []TransportUPLayerAddressInfoToAddItem{}
		for _, i := range tmp_TransportUPLayerAddressInfoToAddList.Value {
			ie.TransportUPLayerAddressInfoToAddList = append(ie.TransportUPLayerAddressInfoToAddList, *i)
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_TransportUPLayerAddressInfoToRemoveList := Sequence[*TransportUPLayerAddressInfoToRemoveItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTLAs},
			ext: true,
		}
		fn := func() *TransportUPLayerAddressInfoToRemoveItem { return new(TransportUPLayerAddressInfoToRemoveItem) }
		if err = tmp_TransportUPLayerAddressInfoToRemoveList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read TransportUPLayerAddressInfoToRemoveList", err)
			return
		}
		ie.TransportUPLayerAddressInfoToRemoveList = []TransportUPLayerAddressInfoToRemoveItem{}
		for _, i := range tmp_TransportUPLayerAddressInfoToRemoveList.Value {
			ie.TransportUPLayerAddressInfoToRemoveList = append(ie.TransportUPLayerAddressInfoToRemoveList, *i)
		}
	}
	return
}
