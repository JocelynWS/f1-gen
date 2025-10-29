package f1ap

import "github.com/lvdund/ngap/aper"

const (
	SubcarrierSpacingPRSKHz15  aper.Enumerated = 0
	SubcarrierSpacingPRSKHz30  aper.Enumerated = 1
	SubcarrierSpacingPRSKHz60  aper.Enumerated = 2
	SubcarrierSpacingPRSKHz120 aper.Enumerated = 3
)

type SubcarrierSpacingPRS struct {
	Value aper.Enumerated
}

func (ie *SubcarrierSpacingPRS) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}

func (ie *SubcarrierSpacingPRS) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
