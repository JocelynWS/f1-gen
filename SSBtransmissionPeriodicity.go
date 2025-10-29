package f1ap

import "github.com/lvdund/ngap/aper"

const (
	SSBtransmissionPeriodicitySf10  aper.Enumerated = 0
	SSBtransmissionPeriodicitySf20  aper.Enumerated = 1
	SSBtransmissionPeriodicitySf40  aper.Enumerated = 2
	SSBtransmissionPeriodicitySf80  aper.Enumerated = 3
	SSBtransmissionPeriodicitySf160 aper.Enumerated = 4
	SSBtransmissionPeriodicitySf320 aper.Enumerated = 5
	SSBtransmissionPeriodicitySf640 aper.Enumerated = 6
)

type SSBTransmissionPeriodicity struct {
	Value aper.Enumerated
}

func (ie *SSBTransmissionPeriodicity) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, true)
	return
}

func (ie *SSBTransmissionPeriodicity) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, true)
	ie.Value = aper.Enumerated(v)
	return
}
