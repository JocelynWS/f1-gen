package ies

import "github.com/lvdund/ngap/aper"

const (
	DUFTransmissionPeriodicityMs0P5   aper.Enumerated = 0
	DUFTransmissionPeriodicityMs0P625 aper.Enumerated = 1
	DUFTransmissionPeriodicityMs1     aper.Enumerated = 2
	DUFTransmissionPeriodicityMs1P25  aper.Enumerated = 3
	DUFTransmissionPeriodicityMs2     aper.Enumerated = 4
	DUFTransmissionPeriodicityMs2P5   aper.Enumerated = 5
	DUFTransmissionPeriodicityMs5     aper.Enumerated = 6
	DUFTransmissionPeriodicityMs10    aper.Enumerated = 7
)

type DUFTransmissionPeriodicity struct {
	Value aper.Enumerated
}

func (ie *DUFTransmissionPeriodicity) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, true)
	return
}

func (ie *DUFTransmissionPeriodicity) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, true)
	ie.Value = aper.Enumerated(v)
	return
}
