package ies

import "github.com/lvdund/ngap/aper"

const (
	NRDLULTxPeriodicityMs0p5   aper.Enumerated = 0
	NRDLULTxPeriodicityMs0p625 aper.Enumerated = 1
	NRDLULTxPeriodicityMs1     aper.Enumerated = 2
	NRDLULTxPeriodicityMs1p25  aper.Enumerated = 3
	NRDLULTxPeriodicityMs2     aper.Enumerated = 4
	NRDLULTxPeriodicityMs2p5   aper.Enumerated = 5
	NRDLULTxPeriodicityMs3     aper.Enumerated = 6
	NRDLULTxPeriodicityMs4     aper.Enumerated = 7
	NRDLULTxPeriodicityMs5     aper.Enumerated = 8
	NRDLULTxPeriodicityMs10    aper.Enumerated = 9
	NRDLULTxPeriodicityMs20    aper.Enumerated = 10
	NRDLULTxPeriodicityMs40    aper.Enumerated = 11
	NRDLULTxPeriodicityMs60    aper.Enumerated = 12
	NRDLULTxPeriodicityMs80    aper.Enumerated = 13
	NRDLULTxPeriodicityMs100   aper.Enumerated = 14
	NRDLULTxPeriodicityMs120   aper.Enumerated = 15
	NRDLULTxPeriodicityMs140   aper.Enumerated = 16
	NRDLULTxPeriodicityMs160   aper.Enumerated = 17
)

type NRDLULTxPeriodicity struct {
	Value aper.Enumerated
}

func (ie *NRDLULTxPeriodicity) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 17}, true)
	return
}

func (ie *NRDLULTxPeriodicity) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 17}, true)
	ie.Value = aper.Enumerated(v)
	return
}
