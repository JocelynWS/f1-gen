package ies

import "github.com/lvdund/ngap/aper"

const (
	NRCPNormal   aper.Enumerated = 0
	NRCPExtended aper.Enumerated = 1
)

type NRCP struct {
	Value aper.Enumerated
}

func (ie *NRCP) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *NRCP) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
