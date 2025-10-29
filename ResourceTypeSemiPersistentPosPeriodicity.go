package f1ap

import "github.com/lvdund/ngap/aper"

const (
	ResourceTypeSemiPersistentPosPeriodicitySlot1     aper.Enumerated = 0
	ResourceTypeSemiPersistentPosPeriodicitySlot2     aper.Enumerated = 1
	ResourceTypeSemiPersistentPosPeriodicitySlot4     aper.Enumerated = 2
	ResourceTypeSemiPersistentPosPeriodicitySlot5     aper.Enumerated = 3
	ResourceTypeSemiPersistentPosPeriodicitySlot8     aper.Enumerated = 4
	ResourceTypeSemiPersistentPosPeriodicitySlot10    aper.Enumerated = 5
	ResourceTypeSemiPersistentPosPeriodicitySlot16    aper.Enumerated = 6
	ResourceTypeSemiPersistentPosPeriodicitySlot20    aper.Enumerated = 7
	ResourceTypeSemiPersistentPosPeriodicitySlot32    aper.Enumerated = 8
	ResourceTypeSemiPersistentPosPeriodicitySlot40    aper.Enumerated = 9
	ResourceTypeSemiPersistentPosPeriodicitySlot64    aper.Enumerated = 10
	ResourceTypeSemiPersistentPosPeriodicitySlot80    aper.Enumerated = 11
	ResourceTypeSemiPersistentPosPeriodicitySlot160   aper.Enumerated = 12
	ResourceTypeSemiPersistentPosPeriodicitySlot320   aper.Enumerated = 13
	ResourceTypeSemiPersistentPosPeriodicitySlot640   aper.Enumerated = 14
	ResourceTypeSemiPersistentPosPeriodicitySlot1280  aper.Enumerated = 15
	ResourceTypeSemiPersistentPosPeriodicitySlot2560  aper.Enumerated = 16
	ResourceTypeSemiPersistentPosPeriodicitySlot5120  aper.Enumerated = 17
	ResourceTypeSemiPersistentPosPeriodicitySlot10240 aper.Enumerated = 18
	ResourceTypeSemiPersistentPosPeriodicitySlot20480 aper.Enumerated = 19
	ResourceTypeSemiPersistentPosPeriodicitySlot40960 aper.Enumerated = 20
	ResourceTypeSemiPersistentPosPeriodicitySlot81920 aper.Enumerated = 21
)

type ResourceTypeSemiPersistentPosPeriodicity struct {
	Value aper.Enumerated
}

func (ie *ResourceTypeSemiPersistentPosPeriodicity) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 20}, true)
	return
}

func (ie *ResourceTypeSemiPersistentPosPeriodicity) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 20}, true)
	ie.Value = aper.Enumerated(v)
	return
}
