package f1ap

import "github.com/lvdund/ngap/aper"

const (
	EUTRACyclicPrefixULNormal   aper.Enumerated = 0
	EUTRACyclicPrefixULExtended aper.Enumerated = 1
)

type EUTRACyclicPrefixUL struct {
	Value aper.Enumerated
}

func (ie *EUTRACyclicPrefixUL) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *EUTRACyclicPrefixUL) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
