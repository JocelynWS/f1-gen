package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	CNUEPagingIdentityPresentNothing uint64 = iota
	CNUEPagingIdentityPresentFivegSTmsi
	CNUEPagingIdentityPresentChoiceExtension
)

type CNUEPagingIdentity struct {
	Choice     uint64
	FiveGSTMSI *aper.BitString
	// ChoiceExtension
}

func (ie *CNUEPagingIdentity) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case CNUEPagingIdentityPresentFivegSTmsi:
		tmp := NewBITSTRING(*ie.FiveGSTMSI, aper.Constraint{Lb: 48, Ub: 48}, false)
		err = tmp.Encode(w)
	}
	return
}

func (ie *CNUEPagingIdentity) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return
	}
	switch ie.Choice {
	case CNUEPagingIdentityPresentFivegSTmsi:
		tmp := BITSTRING{c: aper.Constraint{Lb: 48, Ub: 48}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read FiveGSTMSI", err)
			return
		}
		ie.FiveGSTMSI = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	}
	return
}
