package f1ap

import "github.com/lvdund/ngap/aper"

const (
	SubcarrierSpacingKHz15  aper.Enumerated = 0
	SubcarrierSpacingKHz30  aper.Enumerated = 1
	SubcarrierSpacingKHz60  aper.Enumerated = 2
	SubcarrierSpacingKHz120 aper.Enumerated = 3
)

type SubcarrierSpacing struct {
	Value aper.Enumerated
}

func (ie *SubcarrierSpacing) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}

func (ie *SubcarrierSpacing) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
