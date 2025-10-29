package ies

import "github.com/lvdund/ngap/aper"

type DSCP struct {
	Value aper.BitString `lb:6,ub:6,mandatory`
}

func (m *DSCP) Encode(w *aper.AperWriter) error {
	c := &aper.Constraint{Lb: 6, Ub: 6}
	return w.WriteBitString(m.Value.Bytes, uint(m.Value.NumBits), c, false)
}

func (m *DSCP) Decode(r *aper.AperReader) error {
	c := &aper.Constraint{Lb: 6, Ub: 6}
	bytes, numBits, err := r.ReadBitString(c, false)
	if err != nil {
		return err
	}
	m.Value = aper.BitString{Bytes: bytes, NumBits: uint64(numBits)}
	return nil
}
