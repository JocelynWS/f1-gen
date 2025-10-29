package f1ap

import "github.com/lvdund/ngap/aper"

const (
	UPTransportLayerInformationPresentNothing uint64 = iota
	UPTransportLayerInformationPresentGTPTunnel
)

type UPTransportLayerInformation struct {
	Choice    uint64
	GTPTunnel *GTPTunnel
	// ChoiceExtension
}

func (ie *UPTransportLayerInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case UPTransportLayerInformationPresentGTPTunnel:
		err = ie.GTPTunnel.Encode(w)
	}
	return
}

func (ie *UPTransportLayerInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return
	}
	switch ie.Choice {
	case UPTransportLayerInformationPresentGTPTunnel:
		var tmp GTPTunnel
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GTPTunnel = &tmp
	}
	return
}
