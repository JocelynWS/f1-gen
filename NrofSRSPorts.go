package f1ap

import "github.com/lvdund/ngap/aper"

const (
	NrofSRSPortsPort1  aper.Enumerated = 0
	NrofSRSPortsPorts2 aper.Enumerated = 1
	NrofSRSPortsPorts4 aper.Enumerated = 2
)

type NrofSRSPorts struct {
	Value aper.Enumerated
}

func (ie *NrofSRSPorts) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false)
	return
}

func (ie *NrofSRSPorts) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	ie.Value = aper.Enumerated(v)
	return
}
