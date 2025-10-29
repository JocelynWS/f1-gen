package f1ap

import "github.com/lvdund/ngap/aper"

const (
	RequestTypeOffer     aper.Enumerated = 0
	RequestTypeExecution aper.Enumerated = 1
)

type RequestType struct {
	Value aper.Enumerated
}

func (ie *RequestType) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *RequestType) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
