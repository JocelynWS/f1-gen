package ies

import "github.com/lvdund/ngap/aper"

const (
	NPNSupportInfoPresentNothing uint64 = iota
	NPNSupportInfoPresentSNPNInformation
)

type NPNSupportInfo struct {
	Choice          uint64
	SNPNInformation *NID
	// ChoiceExtension
}

func (ie *NPNSupportInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case NPNSupportInfoPresentSNPNInformation:
		err = ie.SNPNInformation.Encode(w)
	}
	return
}

func (ie *NPNSupportInfo) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return
	}
	switch ie.Choice {
	case NPNSupportInfoPresentSNPNInformation:
		var tmp NID
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SNPNInformation = &tmp
	}
	return
}
