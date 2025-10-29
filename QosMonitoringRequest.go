package f1ap

import "github.com/lvdund/ngap/aper"

const (
	QosMonitoringRequestUl   aper.Enumerated = 0
	QosMonitoringRequestDl   aper.Enumerated = 1
	QosMonitoringRequestBoth aper.Enumerated = 2
)

type QosMonitoringRequest struct {
	Value aper.Enumerated
}

func (ie *QosMonitoringRequest) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
	return
}

func (ie *QosMonitoringRequest) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	ie.Value = aper.Enumerated(v)
	return
}
