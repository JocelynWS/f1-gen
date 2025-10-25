package ies

import "github.com/lvdund/ngap/aper"

// PacketDelayBudget ::= INTEGER (0..1023, ...)
type PacketDelayBudget struct {
	Value int64 `lb:0,ub:1023,mandatory,ext`
}

func (m *PacketDelayBudget) Encode(w *aper.AperWriter) error {
	c := &aper.Constraint{Lb: 0, Ub: 1023}
	return w.WriteInteger(m.Value, c, true)
}

func (m *PacketDelayBudget) Decode(r *aper.AperReader) error {
	c := &aper.Constraint{Lb: 0, Ub: 1023}
	val, err := r.ReadInteger(c, true)
	if err != nil {
		return err
	}
	m.Value = val
	return nil
}
