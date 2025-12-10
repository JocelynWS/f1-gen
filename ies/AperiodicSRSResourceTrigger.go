package ies

import "github.com/lvdund/ngap/aper"

type AperiodicSRSResourceTrigger struct {
	Value int64 `lb:0,ub:3,mandatory,valueExt`
}

func (m *AperiodicSRSResourceTrigger) Encode(w *aper.AperWriter) error {
	c := &aper.Constraint{Lb: 0, Ub: 3}
	return w.WriteInteger(m.Value, c, true)
}

func (m *AperiodicSRSResourceTrigger) Decode(r *aper.AperReader) error {
	c := &aper.Constraint{Lb: 0, Ub: 3}
	val, err := r.ReadInteger(c, true)
	if err != nil {
		return err
	}
	m.Value = val
	return nil
}