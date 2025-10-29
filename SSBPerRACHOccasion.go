package f1ap

import "github.com/lvdund/ngap/aper"

const (
	SSBPerRACHOccasionOneEighth aper.Enumerated = 0
	SSBPerRACHOccasionOneFourth aper.Enumerated = 1
	SSBPerRACHOccasionOneHalf   aper.Enumerated = 2
	SSBPerRACHOccasionOne       aper.Enumerated = 3
	SSBPerRACHOccasionTwo       aper.Enumerated = 4
	SSBPerRACHOccasionFour      aper.Enumerated = 5
	SSBPerRACHOccasionEight     aper.Enumerated = 6
	SSBPerRACHOccasionSixteen   aper.Enumerated = 7
)

type SSBPerRACHOccasion struct {
	Value aper.Enumerated
}

func (ie *SSBPerRACHOccasion) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, true)
	return
}

func (ie *SSBPerRACHOccasion) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, true)
	ie.Value = aper.Enumerated(v)
	return
}
