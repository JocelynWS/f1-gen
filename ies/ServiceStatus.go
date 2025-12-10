package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ServiceStatus struct {
	ServiceState        ServiceState         `madatory`
	SwitchingOffOngoing *SwitchingOffOngoing `optional,valueExt`
	// IEExtensions *ServiceStatusExtIEs `optional`
}

func (ie *ServiceStatus) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.SwitchingOffOngoing != nil {
		aper.SetBit(optionals, 0)
	}
	if err = w.WriteBits(optionals, 2); err != nil {
		return
	}

	if err = ie.ServiceState.Encode(w); err != nil {
		err = utils.WrapError("Encode ServiceState", err)
		return
	}
	if ie.SwitchingOffOngoing != nil {
		if err = ie.SwitchingOffOngoing.Encode(w); err != nil {
			err = utils.WrapError("Encode SwitchingOffOngoing", err)
			return
		}
	}

	return
}

func (ie *ServiceStatus) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}

	if err = ie.ServiceState.Decode(r); err != nil {
		err = utils.WrapError("Read ServiceState", err)
		return
	}

	if aper.IsBitSet(optionals, 0) {
		tmp := new(SwitchingOffOngoing)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read SwitchingOffOngoing", err)
			return
		}
		ie.SwitchingOffOngoing = tmp
	}

	return
}
