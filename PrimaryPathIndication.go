package f1ap

import "github.com/lvdund/ngap/aper"

const (
	PrimaryPathIndicationTrue  aper.Enumerated = 0
	PrimaryPathIndicationFalse aper.Enumerated = 1
)

type PrimaryPathIndication struct {
	Value aper.Enumerated
}

func (ie *PrimaryPathIndication) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *PrimaryPathIndication) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
