package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type CUtoDURRCInformation struct {
	CGConfigInfo                 []byte `lb:0,ub:0,optional`
	UECapabilityRATContainerList []byte `lb:0,ub:0,optional`
	MeasConfig                   []byte `lb:0,ub:0,optional`
	// IEExtensions * `optional`
}

func (ie *CUtoDURRCInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.CGConfigInfo != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.UECapabilityRATContainerList != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.MeasConfig != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)
	if ie.CGConfigInfo != nil {
		tmp_CGConfigInfo := NewOCTETSTRING(ie.CGConfigInfo, aper.Constraint{Lb: 0, Ub: 0}, false)
		if err = tmp_CGConfigInfo.Encode(w); err != nil {
			err = utils.WrapError("Encode CGConfigInfo", err)
			return
		}
	}
	if ie.UECapabilityRATContainerList != nil {
		tmp_UECapabilityRATContainerList := NewOCTETSTRING(ie.UECapabilityRATContainerList, aper.Constraint{Lb: 0, Ub: 0}, false)
		if err = tmp_UECapabilityRATContainerList.Encode(w); err != nil {
			err = utils.WrapError("Encode UECapabilityRATContainerList", err)
			return
		}
	}
	if ie.MeasConfig != nil {
		tmp_MeasConfig := NewOCTETSTRING(ie.MeasConfig, aper.Constraint{Lb: 0, Ub: 0}, false)
		if err = tmp_MeasConfig.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasConfig", err)
			return
		}
	}
	return
}
func (ie *CUtoDURRCInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_CGConfigInfo := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp_CGConfigInfo.Decode(r); err != nil {
			err = utils.WrapError("Read CGConfigInfo", err)
			return
		}
		ie.CGConfigInfo = tmp_CGConfigInfo.Value
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_UECapabilityRATContainerList := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp_UECapabilityRATContainerList.Decode(r); err != nil {
			err = utils.WrapError("Read UECapabilityRATContainerList", err)
			return
		}
		ie.UECapabilityRATContainerList = tmp_UECapabilityRATContainerList.Value
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_MeasConfig := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp_MeasConfig.Decode(r); err != nil {
			err = utils.WrapError("Read MeasConfig", err)
			return
		}
		ie.MeasConfig = tmp_MeasConfig.Value
	}
	return
}
