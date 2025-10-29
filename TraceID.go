package f1ap

import "github.com/lvdund/ngap/aper"

// TraceID ::= OCTETSTRING STRING (SIZE(8))
type TraceID struct {
	Value []byte `lb:8,ub:8,mandatory`
}

func (m *TraceID) Encode(w *aper.AperWriter) error {
	c := &aper.Constraint{Lb: 8, Ub: 8}
	return w.WriteOctetString(m.Value, c, false)
}

func (m *TraceID) Decode(r *aper.AperReader) error {
	c := &aper.Constraint{Lb: 8, Ub: 8}
	val, err := r.ReadOctetString(c, false)
	if err != nil {
		return err
	}
	m.Value = val
	return nil
}
