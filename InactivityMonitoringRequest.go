package f1ap

import "github.com/lvdund/ngap/aper"

const (
	InactivityMonitoringRequestTrue aper.Enumerated = 0
)

type InactivityMonitoringRequest struct {
	Value aper.Enumerated
}

func (ie *InactivityMonitoringRequest) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, true)
	return
}

func (ie *InactivityMonitoringRequest) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true)
	ie.Value = aper.Enumerated(v)
	return
}
