package f1ap

import "github.com/lvdund/ngap/aper"

const (
	ResourceTypeSemiPersistentPeriodicitySlot1    aper.Enumerated = 0
	ResourceTypeSemiPersistentPeriodicitySlot2    aper.Enumerated = 1
	ResourceTypeSemiPersistentPeriodicitySlot4    aper.Enumerated = 2
	ResourceTypeSemiPersistentPeriodicitySlot5    aper.Enumerated = 3
	ResourceTypeSemiPersistentPeriodicitySlot8    aper.Enumerated = 4
	ResourceTypeSemiPersistentPeriodicitySlot10   aper.Enumerated = 5
	ResourceTypeSemiPersistentPeriodicitySlot16   aper.Enumerated = 6
	ResourceTypeSemiPersistentPeriodicitySlot20   aper.Enumerated = 7
	ResourceTypeSemiPersistentPeriodicitySlot32   aper.Enumerated = 8
	ResourceTypeSemiPersistentPeriodicitySlot40   aper.Enumerated = 9
	ResourceTypeSemiPersistentPeriodicitySlot64   aper.Enumerated = 10
	ResourceTypeSemiPersistentPeriodicitySlot80   aper.Enumerated = 11
	ResourceTypeSemiPersistentPeriodicitySlot160  aper.Enumerated = 12
	ResourceTypeSemiPersistentPeriodicitySlot320  aper.Enumerated = 13
	ResourceTypeSemiPersistentPeriodicitySlot640  aper.Enumerated = 14
	ResourceTypeSemiPersistentPeriodicitySlot1280 aper.Enumerated = 15
	ResourceTypeSemiPersistentPeriodicitySlot2560 aper.Enumerated = 16
)

type ResourceTypeSemiPersistentPeriodicity struct {
	Value aper.Enumerated
}

func (ie *ResourceTypeSemiPersistentPeriodicity) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 16}, true)
	return
}

func (ie *ResourceTypeSemiPersistentPeriodicity) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 16}, true)
	ie.Value = aper.Enumerated(v)
	return
}
