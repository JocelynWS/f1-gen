package ies

import "github.com/lvdund/ngap/aper"

// NZP-CSI-RS-ResourceID ::= INTEGER (0..191, ...)
type NZPCSIRSResourceID struct {
	Value int64 `lb:0,ub:191,mandatory,valueExt`
}

func (m *NZPCSIRSResourceID) Encode(w *aper.AperWriter) error {
	c := &aper.Constraint{Lb: 0, Ub: 191}
	return w.WriteInteger(m.Value, c, true)
}

func (m *NZPCSIRSResourceID) Decode(r *aper.AperReader) error {
	c := &aper.Constraint{Lb: 0, Ub: 191}
	val, err := r.ReadInteger(c, true)
	if err != nil {
		return err
	}
	m.Value = val
	return nil
}
