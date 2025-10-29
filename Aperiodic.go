package f1ap

import "github.com/lvdund/ngap/aper"

const (
	AperiodicTrue aper.Enumerated = 0
)

type Aperiodic struct {
	Value aper.Enumerated
}

func (ie *Aperiodic) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, true)
	return
}

func (ie *Aperiodic) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true)
	ie.Value = aper.Enumerated(v)
	return
}
