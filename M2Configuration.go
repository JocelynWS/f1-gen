package f1ap

import "github.com/lvdund/ngap/aper"

const (
	M2ConfigurationTrue aper.Enumerated = 0
)

type M2Configuration struct {
	Value aper.Enumerated
}

func (ie *M2Configuration) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, true)
	return
}

func (ie *M2Configuration) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true)
	ie.Value = aper.Enumerated(v)
	return
}
