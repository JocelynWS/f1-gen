package f1ap

import "github.com/lvdund/ngap/aper"

const (
	IABTNLAddressUsageF1C   aper.Enumerated = 0
	IABTNLAddressUsageF1U   aper.Enumerated = 1
	IABTNLAddressUsageNonf1 aper.Enumerated = 2
)

type IABTNLAddressUsage struct {
	Value aper.Enumerated
}

func (ie *IABTNLAddressUsage) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
	return
}

func (ie *IABTNLAddressUsage) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	ie.Value = aper.Enumerated(v)
	return
}
