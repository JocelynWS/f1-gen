package f1ap

import "github.com/lvdund/ngap/aper"

const (
	QoSInformationPresentNothing uint64 = iota
	QoSInformationPresentEUTRANQoS
)

type QoSInformation struct {
	Choice    uint64
	EUTRANQoS *EUTRANQoS
	// ChoiceExtension
}

func (ie *QoSInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case QoSInformationPresentEUTRANQoS:
		err = ie.EUTRANQoS.Encode(w)
	}
	return
}

func (ie *QoSInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return
	}
	switch ie.Choice {
	case QoSInformationPresentEUTRANQoS:
		var tmp EUTRANQoS
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRANQoS = &tmp
	}
	return
}
