package f1ap

import "github.com/lvdund/ngap/aper"

const (
	PedestrianUEAuthorized    aper.Enumerated = 0
	PedestrianUENotauthorized aper.Enumerated = 1
)

type PedestrianUE struct {
	Value aper.Enumerated
}

func (ie *PedestrianUE) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *PedestrianUE) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
