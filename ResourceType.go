package f1ap

import "github.com/lvdund/ngap/aper"

const (
	SRBResourceTypePeriodic       aper.Enumerated = 0
	SRBResourceTypeSemiPersistent aper.Enumerated = 1
	SRBResourceTypeAperiodic      aper.Enumerated = 2
)

type SRBResourceType struct {
	Value aper.Enumerated
}

func (ie *SRBResourceType) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
	return
}

func (ie *SRBResourceType) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	ie.Value = aper.Enumerated(v)
	return
}
