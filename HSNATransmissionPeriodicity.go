package f1ap

import "github.com/lvdund/ngap/aper"

const (
	HSNATransmissionPeriodicityMs0P5   aper.Enumerated = 0
	HSNATransmissionPeriodicityMs0P625 aper.Enumerated = 1
	HSNATransmissionPeriodicityMs1     aper.Enumerated = 2
	HSNATransmissionPeriodicityMs1P25  aper.Enumerated = 3
	HSNATransmissionPeriodicityMs2     aper.Enumerated = 4
	HSNATransmissionPeriodicityMs2P5   aper.Enumerated = 5
	HSNATransmissionPeriodicityMs5     aper.Enumerated = 6
	HSNATransmissionPeriodicityMs10    aper.Enumerated = 7
	HSNATransmissionPeriodicityMs20    aper.Enumerated = 8
	HSNATransmissionPeriodicityMs40    aper.Enumerated = 9
	HSNATransmissionPeriodicityMs80    aper.Enumerated = 10
	HSNATransmissionPeriodicityMs160   aper.Enumerated = 11
)

type HSNATransmissionPeriodicity struct {
	Value aper.Enumerated
}

func (ie *HSNATransmissionPeriodicity) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 11}, true)
	return
}

func (ie *HSNATransmissionPeriodicity) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 11}, true)
	ie.Value = aper.Enumerated(v)
	return
}
