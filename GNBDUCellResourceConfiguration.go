package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GNBDUCellResourceConfiguration struct {
	SubcarrierSpacing           SubcarrierSpacing           `mandatory`
	DUFTransmissionPeriodicity  *DUFTransmissionPeriodicity `optional`
	DUFSlotConfigList           []DUFSlotConfigItem         `lb:1,ub:maxnoofDUFSlots,optional`
	HSNATransmissionPeriodicity HSNATransmissionPeriodicity `mandatory`
	HNSASlotConfigList          []HSNASlotConfigItem        `lb:1,ub:maxnoofHSNASlots,optional`
	// IEExtensions * `optional`
}

func (ie *GNBDUCellResourceConfiguration) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.DUFTransmissionPeriodicity != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.DUFSlotConfigList != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.HNSASlotConfigList != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)
	if err = ie.SubcarrierSpacing.Encode(w); err != nil {
		err = utils.WrapError("Encode SubcarrierSpacing", err)
		return
	}
	if ie.DUFTransmissionPeriodicity != nil {
		if err = ie.DUFTransmissionPeriodicity.Encode(w); err != nil {
			err = utils.WrapError("Encode DUFTransmissionPeriodicity", err)
			return
		}
	}
	if len(ie.DUFSlotConfigList) > 0 {
		tmp := Sequence[*DUFSlotConfigItem]{
			Value: []*DUFSlotConfigItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofDUFSlots},
			ext:   false,
		}
		for _, i := range ie.DUFSlotConfigList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode DUFSlotConfigList", err)
			return
		}
	}
	if err = ie.HSNATransmissionPeriodicity.Encode(w); err != nil {
		err = utils.WrapError("Encode HSNATransmissionPeriodicity", err)
		return
	}
	if len(ie.HNSASlotConfigList) > 0 {
		tmp := Sequence[*HSNASlotConfigItem]{
			Value: []*HSNASlotConfigItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofHSNASlots},
			ext:   false,
		}
		for _, i := range ie.HNSASlotConfigList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode HNSASlotConfigList", err)
			return
		}
	}
	return
}
func (ie *GNBDUCellResourceConfiguration) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	if err = ie.SubcarrierSpacing.Decode(r); err != nil {
		err = utils.WrapError("Read SubcarrierSpacing", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(DUFTransmissionPeriodicity)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read DUFTransmissionPeriodicity", err)
			return
		}
		ie.DUFTransmissionPeriodicity = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_DUFSlotConfigList := Sequence[*DUFSlotConfigItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDUFSlots},
			ext: false,
		}
		fn := func() *DUFSlotConfigItem { return new(DUFSlotConfigItem) }
		if err = tmp_DUFSlotConfigList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read DUFSlotConfigList", err)
			return
		}
		ie.DUFSlotConfigList = []DUFSlotConfigItem{}
		for _, i := range tmp_DUFSlotConfigList.Value {
			ie.DUFSlotConfigList = append(ie.DUFSlotConfigList, *i)
		}
	}
	if err = ie.HSNATransmissionPeriodicity.Decode(r); err != nil {
		err = utils.WrapError("Read HSNATransmissionPeriodicity", err)
		return
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_HNSASlotConfigList := Sequence[*HSNASlotConfigItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofHSNASlots},
			ext: false,
		}
		fn := func() *HSNASlotConfigItem { return new(HSNASlotConfigItem) }
		if err = tmp_HNSASlotConfigList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read HNSASlotConfigList", err)
			return
		}
		ie.HNSASlotConfigList = []HSNASlotConfigItem{}
		for _, i := range tmp_HNSASlotConfigList.Value {
			ie.HNSASlotConfigList = append(ie.HNSASlotConfigList, *i)
		}
	}
	return
}
