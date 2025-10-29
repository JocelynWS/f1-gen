package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type LTEV2XServicesAuthorized struct {
	VehicleUE    *VehicleUE    `optional`
	PedestrianUE *PedestrianUE `optional`
	// IEExtensions * `optional`
}

func (ie *LTEV2XServicesAuthorized) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.VehicleUE != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.PedestrianUE != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if ie.VehicleUE != nil {
		if err = ie.VehicleUE.Encode(w); err != nil {
			err = utils.WrapError("Encode VehicleUE", err)
			return
		}
	}
	if ie.PedestrianUE != nil {
		if err = ie.PedestrianUE.Encode(w); err != nil {
			err = utils.WrapError("Encode PedestrianUE", err)
			return
		}
	}
	return
}
func (ie *LTEV2XServicesAuthorized) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(VehicleUE)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read VehicleUE", err)
			return
		}
		ie.VehicleUE = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp := new(PedestrianUE)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read PedestrianUE", err)
			return
		}
		ie.PedestrianUE = tmp
	}
	return
}
