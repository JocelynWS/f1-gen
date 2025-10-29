package ies

import "github.com/lvdund/ngap/aper"

const (
	SSBSubcarrierSpacingKHz15  aper.Enumerated = 0
	SSBSubcarrierSpacingKHz30  aper.Enumerated = 1
	SSBSubcarrierSpacingKHz60  aper.Enumerated = 2
	SSBSubcarrierSpacingKHz120 aper.Enumerated = 3
	SSBSubcarrierSpacingKHz240 aper.Enumerated = 4
)

type SSBSubcarrierSpacing struct {
	Value aper.Enumerated
}

func (ie *SSBSubcarrierSpacing) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, true)
	return
}

func (ie *SSBSubcarrierSpacing) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, true)
	ie.Value = aper.Enumerated(v)
	return
}
