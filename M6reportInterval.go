package ies

import "github.com/lvdund/ngap/aper"

const (
	M6reportIntervalMs120   aper.Enumerated = 0
	M6reportIntervalMs240   aper.Enumerated = 1
	M6reportIntervalMs640   aper.Enumerated = 2
	M6reportIntervalMs1024  aper.Enumerated = 3
	M6reportIntervalMs2048  aper.Enumerated = 4
	M6reportIntervalMs5120  aper.Enumerated = 5
	M6reportIntervalMs10240 aper.Enumerated = 6
	M6reportIntervalMs20480 aper.Enumerated = 7
	M6reportIntervalMs40960 aper.Enumerated = 8
	M6reportIntervalMin1    aper.Enumerated = 9
	M6reportIntervalMin6    aper.Enumerated = 10
	M6reportIntervalMin12   aper.Enumerated = 11
	M6reportIntervalMin30   aper.Enumerated = 12
)

type M6reportInterval struct {
	Value aper.Enumerated
}

func (ie *M6reportInterval) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 12}, true)
	return
}

func (ie *M6reportInterval) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 12}, true)
	ie.Value = aper.Enumerated(v)
	return
}
