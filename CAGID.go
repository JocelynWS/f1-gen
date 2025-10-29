package f1ap

import "github.com/lvdund/ngap/aper"

type CAGID struct {
	Value aper.BitString `lb:32,ub:32,mandatory`
}

func (m *CAGID) Encode(w *aper.AperWriter) error {
	c := &aper.Constraint{Lb: 32, Ub: 32}
	return w.WriteBitString(m.Value.Bytes, uint(m.Value.NumBits), c, false)
}

func (m *CAGID) Decode(r *aper.AperReader) error {
	c := &aper.Constraint{Lb: 32, Ub: 32}
	bytes, numBits, err := r.ReadBitString(c, false)
	if err != nil {
		return err
	}
	m.Value = aper.BitString{Bytes: bytes, NumBits: uint64(numBits)}
	return nil
}
