package f1ap

import "github.com/lvdund/ngap/aper"

const (
	MDTActivationImmediatemdtonly     aper.Enumerated = 0
	MDTActivationImmediatemdtandtrace aper.Enumerated = 1
)

type MDTActivation struct {
	Value aper.Enumerated
}

func (ie *MDTActivation) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *MDTActivation) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
