package f1ap

import "github.com/lvdund/ngap/aper"

const (
	FR2BandwidthBw50  aper.Enumerated = 0
	FR2BandwidthBw100 aper.Enumerated = 1
	FR2BandwidthBw200 aper.Enumerated = 2
	FR2BandwidthBw400 aper.Enumerated = 3
)

type FR2Bandwidth struct {
	Value aper.Enumerated
}

func (ie *FR2Bandwidth) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}

func (ie *FR2Bandwidth) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
