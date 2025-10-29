package f1ap

import "github.com/lvdund/ngap/aper"

const (
	EUTRATransmissionBandwidthBw6   aper.Enumerated = 0
	EUTRATransmissionBandwidthBw15  aper.Enumerated = 1
	EUTRATransmissionBandwidthBw25  aper.Enumerated = 2
	EUTRATransmissionBandwidthBw50  aper.Enumerated = 3
	EUTRATransmissionBandwidthBw75  aper.Enumerated = 4
	EUTRATransmissionBandwidthBw100 aper.Enumerated = 5
)

type EUTRATransmissionBandwidth struct {
	Value aper.Enumerated
}

func (ie *EUTRATransmissionBandwidth) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, true)
	return
}

func (ie *EUTRATransmissionBandwidth) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, true)
	ie.Value = aper.Enumerated(v)
	return
}
