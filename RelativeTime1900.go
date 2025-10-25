package ies

import "github.com/lvdund/ngap/aper"

type RelativeTime1900 struct {
	Value aper.BitString `lb:64,ub:64,mandatory`
}

func (m *RelativeTime1900) Encode(w *aper.AperWriter) error {
	c := &aper.Constraint{Lb: 64, Ub: 64}
	return w.WriteBitString(m.Value.Bytes, uint(m.Value.NumBits), c, false)
}

func (m *RelativeTime1900) Decode(r *aper.AperReader) error {
	c := &aper.Constraint{Lb: 64, Ub: 64}
	bytes, numBits, err := r.ReadBitString(c, false)
	if err != nil {
		return err
	}
	m.Value = aper.BitString{Bytes: bytes, NumBits: uint64(numBits)}
	return nil
}
