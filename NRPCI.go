package ies

import "github.com/lvdund/ngap/aper"

type NRPCI struct {
	Value int64 `lb:0,ub:1007,mandatory`
}

func (m *NRPCI) Encode(w *aper.AperWriter) error {
	c := &aper.Constraint{Lb: 0, Ub: 1007}
	return w.WriteInteger(m.Value, c, false)
}

func (m *NRPCI) Decode(r *aper.AperReader) error {
	c := &aper.Constraint{Lb: 0, Ub: 1007}
	val, err := r.ReadInteger(c, false)
	if err != nil {
		return err
	}
	m.Value = val
	return nil
}
