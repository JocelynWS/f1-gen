package ies

import "github.com/lvdund/ngap/aper"

const (
	ShortDRXCycleLengthShortdrxcycletimer aper.Enumerated = 0
	ShortDRXCycleLengthSib1Message        aper.Enumerated = 1
	ShortDRXCycleLengthSib10Message       aper.Enumerated = 2
	ShortDRXCycleLengthSib12Message       aper.Enumerated = 3
	ShortDRXCycleLengthSib13Message       aper.Enumerated = 4
	ShortDRXCycleLengthSib14Message       aper.Enumerated = 5
	ShortDRXCycleLengthSitypelist         aper.Enumerated = 6
	ShortDRXCycleLengthSitypeitem         aper.Enumerated = 7
	ShortDRXCycleLengthSitype             aper.Enumerated = 8
	ShortDRXCycleLengthIeextensions       aper.Enumerated = 9
)

type ShortDRXCycleLength struct {
	Value aper.Enumerated
}

func (ie *ShortDRXCycleLength) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 9}, true)
	return
}

func (ie *ShortDRXCycleLength) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 9}, true)
	ie.Value = aper.Enumerated(v)
	return
}
