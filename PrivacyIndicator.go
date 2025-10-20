package ies

import "github.com/lvdund/ngap/aper"

const (
	PrivacyIndicatorImmediateMdt aper.Enumerated = 0
	PrivacyIndicatorLoggedMdt    aper.Enumerated = 1
)

type PrivacyIndicator struct {
	Value aper.Enumerated
}

func (ie *PrivacyIndicator) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *PrivacyIndicator) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
