package f1ap

import "github.com/lvdund/ngap/aper"

const (
	EUTRASubframeAssignmentSa0 aper.Enumerated = 0
	EUTRASubframeAssignmentSa1 aper.Enumerated = 1
	EUTRASubframeAssignmentSa2 aper.Enumerated = 2
	EUTRASubframeAssignmentSa3 aper.Enumerated = 3
	EUTRASubframeAssignmentSa4 aper.Enumerated = 4
	EUTRASubframeAssignmentSa5 aper.Enumerated = 5
	EUTRASubframeAssignmentSa6 aper.Enumerated = 6
)

type EUTRASubframeAssignment struct {
	Value aper.Enumerated
}

func (ie *EUTRASubframeAssignment) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, true)
	return
}

func (ie *EUTRASubframeAssignment) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, true)
	ie.Value = aper.Enumerated(v)
	return
}
