package f1ap

import "github.com/lvdund/ngap/aper"

const (
	GroupOrSequenceHoppingNeither         aper.Enumerated = 0
	GroupOrSequenceHoppingGroupHopping    aper.Enumerated = 1
	GroupOrSequenceHoppingSequenceHopping aper.Enumerated = 2
)

type GroupOrSequenceHopping struct {
	Value aper.Enumerated
}

func (ie *GroupOrSequenceHopping) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false)
	return
}

func (ie *GroupOrSequenceHopping) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	ie.Value = aper.Enumerated(v)
	return
}
