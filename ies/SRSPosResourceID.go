package ies

import "github.com/lvdund/ngap/aper"

// SRSPosResourceID ::= INTEGER (0..63)
type SRSPosResourceID struct {
	Value int64 `lb:0,ub:63,mandatory`
}

func (m *SRSPosResourceID) Encode(w *aper.AperWriter) error {
	c := &aper.Constraint{Lb: 0, Ub: 63}
	return w.WriteInteger(m.Value, c, false)
}

func (m *SRSPosResourceID) Decode(r *aper.AperReader) error {
	c := &aper.Constraint{Lb: 0, Ub: 63}
	val, err := r.ReadInteger(c, false)
	if err != nil {
		return err
	}
	m.Value = val
	return nil
}
