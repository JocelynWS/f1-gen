package f1ap

import "github.com/lvdund/ngap/aper"

const (
	PosBroadcastStart aper.Enumerated = 0
	PosBroadcastStop  aper.Enumerated = 1
)

type PosBroadcast struct {
	Value aper.Enumerated
}

func (ie *PosBroadcast) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *PosBroadcast) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
