package ies

import "github.com/lvdund/ngap/aper"

const (
	QoSCharacteristicsPresentNothing uint64 = iota
	QoSCharacteristicsPresentNonDynamic5QI
	QoSCharacteristicsPresentDynamic5QI
)

type QoSCharacteristics struct {
	Choice        uint64
	NonDynamic5QI *NonDynamic5QIDescriptor
	Dynamic5QI    *Dynamic5QIDescriptor
	// ChoiceExtension // ChoiceExtensions
}

func (ie *QoSCharacteristics) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case QoSCharacteristicsPresentNonDynamic5QI:
		err = ie.NonDynamic5QI.Encode(w)
	case QoSCharacteristicsPresentDynamic5QI:
		err = ie.Dynamic5QI.Encode(w)
	}
	return
}

func (ie *QoSCharacteristics) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case QoSCharacteristicsPresentNonDynamic5QI:
		var tmp NonDynamic5QIDescriptor
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NonDynamic5QI = &tmp
	case QoSCharacteristicsPresentDynamic5QI:
		var tmp Dynamic5QIDescriptor
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.Dynamic5QI = &tmp
	}
	return
}
