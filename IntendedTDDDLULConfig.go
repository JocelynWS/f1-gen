package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type IntendedTDDDLULConfig struct {
	NRSCS                 NRSCS                   `madatory,valExt`
	NRCP                  NRCP                    `madatory,valExt`
	NRDLULTxPeriodicity   NRDLULTxPeriodicity     `madatory,valExt`
	SlotConfigurationList []SlotConfigurationItem `lb:1,ub:maxnoofSlots,mandatory,valExt`
	// IEExtensions *IntendedTDDDLULConfigExtIEs `optional`
}

func (ie *IntendedTDDDLULConfig) Encode(w *aper.AperWriter) (err error) {
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)

	if err = NewENUMERATED(int64(ie.NRSCS), aper.Constraint{Lb: 0, Ub: 3}, false).Encode(w); err != nil {
		return utils.WrapError("Encode NRSCS", err)
	}
	if err = NewENUMERATED(int64(ie.NRCP), aper.Constraint{Lb: 0, Ub: 1}, false).Encode(w); err != nil {
		return utils.WrapError("Encode NRCP", err)
	}
	if err = NewENUMERATED(int64(ie.NRDLULTxPeriodicity), aper.Constraint{Lb: 0, Ub: 17}, false).Encode(w); err != nil {
		return utils.WrapError("Encode NRDLULTxPeriodicity", err)
	}

	tmp := Sequence[*SlotConfigurationItem]{
		Value: []*SlotConfigurationItem{},
		c:     aper.Constraint{Lb: 1, Ub: maxnoofSlots},
		ext:   true,
	}
	for _, i := range ie.SlotConfigurationList {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SlotConfigurationList", err)
	}
	return
}

func (ie *IntendedTDDDLULConfig) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBits(1); err != nil {
		return
	}

	tmpNRSCS := ENUMERATED{c: aper.Constraint{Lb: 0, Ub: 3}, ext: false}
	if err = tmpNRSCS.Decode(r); err != nil {
		return utils.WrapError("Read NRSCS", err)
	}
	ie.NRSCS = NRSCS(tmpNRSCS.Value)

	tmpNRCP := ENUMERATED{c: aper.Constraint{Lb: 0, Ub: 1}, ext: false}
	if err = tmpNRCP.Decode(r); err != nil {
		return utils.WrapError("Read NRCP", err)
	}
	ie.NRCP = NRCP(tmpNRCP.Value)

	tmpNRDLULTx := ENUMERATED{c: aper.Constraint{Lb: 0, Ub: 17}, ext: false}
	if err = tmpNRDLULTx.Decode(r); err != nil {
		return utils.WrapError("Read NRDLULTxPeriodicity", err)
	}
	ie.NRDLULTxPeriodicity = NRDLULTxPeriodicity(tmpNRDLULTx.Value)

	tmpSlotList := Sequence[*SlotConfigurationItem]{
		c: aper.Constraint{Lb: 1, Ub: maxnoofSlots}, ext: true,
	}
	fn := func() *SlotConfigurationItem { return new(SlotConfigurationItem) }
	if err = tmpSlotList.Decode(r, fn); err != nil {
		return utils.WrapError("Read SlotConfigurationList", err)
	}

	ie.SlotConfigurationList = []SlotConfigurationItem{}
	for _, i := range tmpSlotList.Value {
		ie.SlotConfigurationList = append(ie.SlotConfigurationList, *i)
	}
	return
}
