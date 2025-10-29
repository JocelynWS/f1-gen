package f1ap

import "github.com/lvdund/ngap/aper"

const (
	FR1BandwidthBw5   aper.Enumerated = 0
	FR1BandwidthBw10  aper.Enumerated = 1
	FR1BandwidthBw20  aper.Enumerated = 2
	FR1BandwidthBw40  aper.Enumerated = 3
	FR1BandwidthBw50  aper.Enumerated = 4
	FR1BandwidthBw80  aper.Enumerated = 5
	FR1BandwidthBw100 aper.Enumerated = 6
)

type FR1Bandwidth struct {
	Value aper.Enumerated
}

func (ie *FR1Bandwidth) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, true)
	return
}

func (ie *FR1Bandwidth) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, true)
	ie.Value = aper.Enumerated(v)
	return
}
