package ies

import "github.com/lvdund/ngap/aper"

const (
	UACActionRejectnonemergencymodt                                    aper.Enumerated = 0
	UACActionRejectrrccrsignalling                                     aper.Enumerated = 1
	UACActionPermitemergencysessionsandmobileterminatedservicesonly    aper.Enumerated = 2
	UACActionPermithighprioritysessionsandmobileterminatedservicesonly aper.Enumerated = 3
)

type UACAction struct {
	Value aper.Enumerated
}

func (ie *UACAction) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}

func (ie *UACAction) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
