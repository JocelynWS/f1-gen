package f1ap

import "github.com/lvdund/ngap/aper"

const (
	NonUPTrafficTypeUeAssociated    aper.Enumerated = 0
	NonUPTrafficTypeNonUeAssociated aper.Enumerated = 1
	NonUPTrafficTypeNonF1           aper.Enumerated = 2
	NonUPTrafficTypeBapControlPdu   aper.Enumerated = 3
)

type NonUPTrafficType struct {
	Value aper.Enumerated
}

func (ie *NonUPTrafficType) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}

func (ie *NonUPTrafficType) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
