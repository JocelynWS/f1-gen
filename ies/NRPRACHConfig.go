package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type NRPRACHConfig struct {
	UlPRACHConfigList  []NRPRACHConfigItem `lb:0,ub:maxnoofPRACHconfigs,optional`
	SulPRACHConfigList []NRPRACHConfigItem `lb:0,ub:maxnoofPRACHconfigs,optional`
	// IEExtension * `optional`
}

func (ie *NRPRACHConfig) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.UlPRACHConfigList != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.SulPRACHConfigList != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if len(ie.UlPRACHConfigList) > 0 {
		tmp := Sequence[*NRPRACHConfigItem]{
			Value: []*NRPRACHConfigItem{},
			c:     aper.Constraint{Lb: 0, Ub: maxnoofPRACHconfigs},
			ext:   false,
		}
		for _, i := range ie.UlPRACHConfigList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode UlPRACHConfigList", err)
			return
		}
	}
	if len(ie.SulPRACHConfigList) > 0 {
		tmp := Sequence[*NRPRACHConfigItem]{
			Value: []*NRPRACHConfigItem{},
			c:     aper.Constraint{Lb: 0, Ub: maxnoofPRACHconfigs},
			ext:   false,
		}
		for _, i := range ie.SulPRACHConfigList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode SulPRACHConfigList", err)
			return
		}
	}
	return
}
func (ie *NRPRACHConfig) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_UlPRACHConfigList := Sequence[*NRPRACHConfigItem]{
			c:   aper.Constraint{Lb: 0, Ub: maxnoofPRACHconfigs},
			ext: false,
		}
		fn := func() *NRPRACHConfigItem { return new(NRPRACHConfigItem) }
		if err = tmp_UlPRACHConfigList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read UlPRACHConfigList", err)
			return
		}
		ie.UlPRACHConfigList = []NRPRACHConfigItem{}
		for _, i := range tmp_UlPRACHConfigList.Value {
			ie.UlPRACHConfigList = append(ie.UlPRACHConfigList, *i)
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_SulPRACHConfigList := Sequence[*NRPRACHConfigItem]{
			c:   aper.Constraint{Lb: 0, Ub: maxnoofPRACHconfigs},
			ext: false,
		}
		fn := func() *NRPRACHConfigItem { return new(NRPRACHConfigItem) }
		if err = tmp_SulPRACHConfigList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read SulPRACHConfigList", err)
			return
		}
		ie.SulPRACHConfigList = []NRPRACHConfigItem{}
		for _, i := range tmp_SulPRACHConfigList.Value {
			ie.SulPRACHConfigList = append(ie.SulPRACHConfigList, *i)
		}
	}
	return
}
