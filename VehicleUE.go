package f1ap

import "github.com/lvdund/ngap/aper"

const (
	VehicleUEAuthorized    aper.Enumerated = 0
	VehicleUENotauthorized aper.Enumerated = 1
)

type VehicleUE struct {
	Value aper.Enumerated
}

func (ie *VehicleUE) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *VehicleUE) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
