package ies

import "github.com/lvdund/ngap/aper"

const (
	CyclicPrefixNormal   aper.Enumerated = 0
	CyclicPrefixExtended aper.Enumerated = 1
)

type CyclicPrefix struct {
	Value aper.Enumerated
}

func (ie *CyclicPrefix) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false)
	return
}

func (ie *CyclicPrefix) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false)
	ie.Value = aper.Enumerated(v)
	return
}
