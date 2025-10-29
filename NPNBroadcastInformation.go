package f1ap

import "github.com/lvdund/ngap/aper"

const (
	NPNBroadcastInformationPresentNothing uint64 = iota
	NPNBroadcastInformationPresentSNPNBroadcastInformation
	NPNBroadcastInformationPresentPNINPNBroadcastInformation
)

type NPNBroadcastInformation struct {
	Choice                     uint64
	SNPNBroadcastInformation   *NPNBroadcastInformationSNPN
	PNINPNBroadcastInformation *NPNBroadcastInformationPNINPN
	// ChoiceExtension
}

func (ie *NPNBroadcastInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case NPNBroadcastInformationPresentSNPNBroadcastInformation:
		err = ie.SNPNBroadcastInformation.Encode(w)
	case NPNBroadcastInformationPresentPNINPNBroadcastInformation:
		err = ie.PNINPNBroadcastInformation.Encode(w)
	}
	return
}

func (ie *NPNBroadcastInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case NPNBroadcastInformationPresentSNPNBroadcastInformation:
		var tmp NPNBroadcastInformationSNPN
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SNPNBroadcastInformation = &tmp
	case NPNBroadcastInformationPresentPNINPNBroadcastInformation:
		var tmp NPNBroadcastInformationPNINPN
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.PNINPNBroadcastInformation = &tmp
	}
	return
}
