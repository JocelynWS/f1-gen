package f1ap

import "github.com/lvdund/ngap/aper"

const (
	DirectionOfAltitudeHeight aper.Enumerated = 0
	DirectionOfAltitudeDepth  aper.Enumerated = 1
)

type DirectionOfAltitude struct {
	Value aper.Enumerated
}

func (ie *DirectionOfAltitude) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false)
	return
}

func (ie *DirectionOfAltitude) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false)
	ie.Value = aper.Enumerated(v)
	return
}
