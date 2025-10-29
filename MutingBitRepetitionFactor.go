package f1ap

import "github.com/lvdund/ngap/aper"

const (
	MutingBitRepetitionFactorRf1 aper.Enumerated = 0
	MutingBitRepetitionFactorRf2 aper.Enumerated = 1
	MutingBitRepetitionFactorRf4 aper.Enumerated = 2
	MutingBitRepetitionFactorRf8 aper.Enumerated = 3
)

type MutingBitRepetitionFactor struct {
	Value aper.Enumerated
}

func (ie *MutingBitRepetitionFactor) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}

func (ie *MutingBitRepetitionFactor) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
