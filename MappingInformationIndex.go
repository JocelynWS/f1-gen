package ies

import "github.com/lvdund/ngap/aper"

type MappingInformationIndex struct {
	Value aper.BitString `lb:26,ub:26,mandatory`
}

func (m *MappingInformationIndex) Encode(w *aper.AperWriter) error {
	c := &aper.Constraint{Lb: 26, Ub: 26}
	return w.WriteBitString(m.Value.Bytes, uint(m.Value.NumBits), c, false)
}

func (m *MappingInformationIndex) Decode(r *aper.AperReader) error {
	c := &aper.Constraint{Lb: 26, Ub: 26}
	bytes, numBits, err := r.ReadBitString(c, false)
	if err != nil {
		return err
	}
	m.Value = aper.BitString{Bytes: bytes, NumBits: uint64(numBits)}
	return nil
}
