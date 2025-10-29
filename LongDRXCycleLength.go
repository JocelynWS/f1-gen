package f1ap

import "github.com/lvdund/ngap/aper"

const (
	LongDRXCycleLengthLowerlayerpresencestatuschange aper.Enumerated = 0
	LongDRXCycleLengthSuspendlowerlayers             aper.Enumerated = 1
	LongDRXCycleLengthResumelowerlayers              aper.Enumerated = 2
)

type LongDRXCycleLength struct {
	Value aper.Enumerated
}

func (ie *LongDRXCycleLength) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
	return
}

func (ie *LongDRXCycleLength) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	ie.Value = aper.Enumerated(v)
	return
}
