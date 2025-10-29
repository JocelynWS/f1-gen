package ies

import "github.com/lvdund/ngap/aper"

const (
	ResourceSetPeriodicityN4     aper.Enumerated = 0
	ResourceSetPeriodicityN5     aper.Enumerated = 1
	ResourceSetPeriodicityN8     aper.Enumerated = 2
	ResourceSetPeriodicityN10    aper.Enumerated = 3
	ResourceSetPeriodicityN16    aper.Enumerated = 4
	ResourceSetPeriodicityN20    aper.Enumerated = 5
	ResourceSetPeriodicityN32    aper.Enumerated = 6
	ResourceSetPeriodicityN40    aper.Enumerated = 7
	ResourceSetPeriodicityN64    aper.Enumerated = 8
	ResourceSetPeriodicityN80    aper.Enumerated = 9
	ResourceSetPeriodicityN160   aper.Enumerated = 10
	ResourceSetPeriodicityN320   aper.Enumerated = 11
	ResourceSetPeriodicityN640   aper.Enumerated = 12
	ResourceSetPeriodicityN1280  aper.Enumerated = 13
	ResourceSetPeriodicityN2560  aper.Enumerated = 14
	ResourceSetPeriodicityN5120  aper.Enumerated = 15
	ResourceSetPeriodicityN10240 aper.Enumerated = 16
	ResourceSetPeriodicityN20480 aper.Enumerated = 17
	ResourceSetPeriodicityN40960 aper.Enumerated = 18
	ResourceSetPeriodicityN81920 aper.Enumerated = 19
)

type ResourceSetPeriodicity struct {
	Value aper.Enumerated
}

func (ie *ResourceSetPeriodicity) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 19}, true)
	return
}

func (ie *ResourceSetPeriodicity) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 19}, true)
	ie.Value = aper.Enumerated(v)
	return
}
