package f1ap

import "github.com/lvdund/ngap/aper"

const (
	SubcarrierSpacingSCSKHz15  aper.Enumerated = 0
	SubcarrierSpacingSCSKHz30  aper.Enumerated = 1
	SubcarrierSpacingSCSKHz60  aper.Enumerated = 2
	SubcarrierSpacingSCSKHz120 aper.Enumerated = 3
)

type SubcarrierSpacingSCS struct {
	Value aper.Enumerated
}

func (ie *SubcarrierSpacingSCS) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}

func (ie *SubcarrierSpacingSCS) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
