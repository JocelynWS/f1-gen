package ies

import "github.com/lvdund/ngap/aper"

const (
	M5periodMs1024  aper.Enumerated = 0
	M5periodMs2048  aper.Enumerated = 1
	M5periodMs5120  aper.Enumerated = 2
	M5periodMs10240 aper.Enumerated = 3
	M5periodMin1    aper.Enumerated = 4
)

type M5period struct {
	Value aper.Enumerated
}

func (ie *M5period) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, true)
	return
}

func (ie *M5period) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, true)
	ie.Value = aper.Enumerated(v)
	return
}
