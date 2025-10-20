package ies

import "github.com/lvdund/ngap/aper"

const (
	PC5QoSCharacteristicsPresentNothing uint64 = iota
	PC5QoSCharacteristicsPresentNonDynamicPQI
	PC5QoSCharacteristicsPresentDynamicPQI
)

type PC5QoSCharacteristics struct {
	Choice        uint64
	NonDynamicPQI *NonDynamicPQIDescriptor
	DynamicPQI    *DynamicPQIDescriptor
	// ChoiceExtension // ChoiceExtensions
}

func (ie *PC5QoSCharacteristics) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case PC5QoSCharacteristicsPresentNonDynamicPQI:
		err = ie.NonDynamicPQI.Encode(w)
	case PC5QoSCharacteristicsPresentDynamicPQI:
		err = ie.DynamicPQI.Encode(w)
	}
	return
}

func (ie *PC5QoSCharacteristics) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case PC5QoSCharacteristicsPresentNonDynamicPQI:
		var tmp NonDynamicPQIDescriptor
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NonDynamicPQI = &tmp
	case PC5QoSCharacteristicsPresentDynamicPQI:
		var tmp DynamicPQIDescriptor
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.DynamicPQI = &tmp
	}
	return
}
