package ies

import "github.com/lvdund/ngap/aper"

const (
	ResourceTypePeriodicPosPeriodicitySlot1     aper.Enumerated = 0
	ResourceTypePeriodicPosPeriodicitySlot2     aper.Enumerated = 1
	ResourceTypePeriodicPosPeriodicitySlot4     aper.Enumerated = 2
	ResourceTypePeriodicPosPeriodicitySlot5     aper.Enumerated = 3
	ResourceTypePeriodicPosPeriodicitySlot8     aper.Enumerated = 4
	ResourceTypePeriodicPosPeriodicitySlot10    aper.Enumerated = 5
	ResourceTypePeriodicPosPeriodicitySlot16    aper.Enumerated = 6
	ResourceTypePeriodicPosPeriodicitySlot20    aper.Enumerated = 7
	ResourceTypePeriodicPosPeriodicitySlot32    aper.Enumerated = 8
	ResourceTypePeriodicPosPeriodicitySlot40    aper.Enumerated = 9
	ResourceTypePeriodicPosPeriodicitySlot64    aper.Enumerated = 10
	ResourceTypePeriodicPosPeriodicitySlot80    aper.Enumerated = 11
	ResourceTypePeriodicPosPeriodicitySlot160   aper.Enumerated = 12
	ResourceTypePeriodicPosPeriodicitySlot320   aper.Enumerated = 13
	ResourceTypePeriodicPosPeriodicitySlot640   aper.Enumerated = 14
	ResourceTypePeriodicPosPeriodicitySlot1280  aper.Enumerated = 15
	ResourceTypePeriodicPosPeriodicitySlot2560  aper.Enumerated = 16
	ResourceTypePeriodicPosPeriodicitySlot5120  aper.Enumerated = 17
	ResourceTypePeriodicPosPeriodicitySlot10240 aper.Enumerated = 18
	ResourceTypePeriodicPosPeriodicitySlot20480 aper.Enumerated = 19
	ResourceTypePeriodicPosPeriodicitySlot40960 aper.Enumerated = 20
	ResourceTypePeriodicPosPeriodicitySlot81920 aper.Enumerated = 21
)

type ResourceTypePeriodicPosPeriodicity struct {
	Value aper.Enumerated
}

func (ie *ResourceTypePeriodicPosPeriodicity) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 20}, true)
	return
}

func (ie *ResourceTypePeriodicPosPeriodicity) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 20}, true)
	ie.Value = aper.Enumerated(v)
	return
}
