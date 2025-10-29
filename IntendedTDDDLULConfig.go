package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type IntendedTDDDLULConfig struct {
	NRSCS                 NRSCS                   `mandatory`
	NRCP                  NRCP                    `mandatory`
	NRDLULTxPeriodicity   NRDLULTxPeriodicity     `mandatory`
	SlotConfigurationList []SlotConfigurationItem `lb:1,ub:maxnoofslots,mandatory`
	// IEExtensions *ProtocolExtensionContainerIntendedTDDDLULConfigExtIEs `optional`
}

func (ie *IntendedTDDDLULConfig) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)

	if err = ie.NRSCS.Encode(w); err != nil {
		err = utils.WrapError("Encode NRSCS", err)
		return
	}

	if err = ie.NRCP.Encode(w); err != nil {
		err = utils.WrapError("Encode NRCP", err)
		return
	}

	if err = ie.NRDLULTxPeriodicity.Encode(w); err != nil {
		err = utils.WrapError("Encode NRDLULTxPeriodicity", err)
		return
	}

	tmpList := Sequence[*SlotConfigurationItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofslots},
		ext: false,
	}
	for i := range ie.SlotConfigurationList {
		tmpList.Value = append(tmpList.Value, &ie.SlotConfigurationList[i])
	}
	if err = tmpList.Encode(w); err != nil {
		err = utils.WrapError("Encode SlotConfigurationList", err)
		return
	}

	return
}

func (ie *IntendedTDDDLULConfig) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	if _, err = r.ReadBits(1); err != nil {
		return
	}

	if err = ie.NRSCS.Decode(r); err != nil {
		err = utils.WrapError("Read NRSCS", err)
		return
	}

	if err = ie.NRCP.Decode(r); err != nil {
		err = utils.WrapError("Read NRCP", err)
		return
	}

	if err = ie.NRDLULTxPeriodicity.Decode(r); err != nil {
		err = utils.WrapError("Read NRDLULTxPeriodicity", err)
		return
	}

	var tmpList Sequence[*SlotConfigurationItem]
	tmpList.c = aper.Constraint{Lb: 1, Ub: maxnoofslots}
	tmpList.ext = false
	if err = tmpList.Decode(r, func() *SlotConfigurationItem {
		return new(SlotConfigurationItem)
	}); err != nil {
		err = utils.WrapError("Read SlotConfigurationList", err)
		return
	}
	ie.SlotConfigurationList = make([]SlotConfigurationItem, len(tmpList.Value))
	for i, item := range tmpList.Value {
		ie.SlotConfigurationList[i] = *item
	}

	return
}
