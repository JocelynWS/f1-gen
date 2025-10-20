package ies

import "github.com/lvdund/ngap/aper"

const (
	ResourceTypePeriodicPeriodicitySlot1    aper.Enumerated = 0
	ResourceTypePeriodicPeriodicitySlot2    aper.Enumerated = 1
	ResourceTypePeriodicPeriodicitySlot4    aper.Enumerated = 2
	ResourceTypePeriodicPeriodicitySlot5    aper.Enumerated = 3
	ResourceTypePeriodicPeriodicitySlot8    aper.Enumerated = 4
	ResourceTypePeriodicPeriodicitySlot10   aper.Enumerated = 5
	ResourceTypePeriodicPeriodicitySlot16   aper.Enumerated = 6
	ResourceTypePeriodicPeriodicitySlot20   aper.Enumerated = 7
	ResourceTypePeriodicPeriodicitySlot32   aper.Enumerated = 8
	ResourceTypePeriodicPeriodicitySlot40   aper.Enumerated = 9
	ResourceTypePeriodicPeriodicitySlot64   aper.Enumerated = 10
	ResourceTypePeriodicPeriodicitySlot80   aper.Enumerated = 11
	ResourceTypePeriodicPeriodicitySlot160  aper.Enumerated = 12
	ResourceTypePeriodicPeriodicitySlot320  aper.Enumerated = 13
	ResourceTypePeriodicPeriodicitySlot640  aper.Enumerated = 14
	ResourceTypePeriodicPeriodicitySlot1280 aper.Enumerated = 15
	ResourceTypePeriodicPeriodicitySlot2560 aper.Enumerated = 16
)

type ResourceTypePeriodicPeriodicity struct {
	Value aper.Enumerated
}

func (ie *ResourceTypePeriodicPeriodicity) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 16}, true)
	return
}

func (ie *ResourceTypePeriodicPeriodicity) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 16}, true)
	ie.Value = aper.Enumerated(v)
	return
}
