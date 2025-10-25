package ies

import "github.com/lvdund/ngap/aper"

// AperiodicSRSResourceTrigger ::= INTEGER (1..3)
type AperiodicSRSResourceTrigger struct {
	Value int64 `lb:1,ub:3,mandatory`
}

func (m *AperiodicSRSResourceTrigger) Encode(w *aper.AperWriter) error {
	c := &aper.Constraint{Lb: 1, Ub: 3}
	return w.WriteInteger(m.Value, c, false)
}

func (m *AperiodicSRSResourceTrigger) Decode(r *aper.AperReader) error {
	c := &aper.Constraint{Lb: 1, Ub: 3}
	val, err := r.ReadInteger(c, false)
	if err != nil {
		return err
	}
	m.Value = val
	return nil
}
