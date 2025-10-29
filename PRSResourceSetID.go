package f1ap

import "github.com/lvdund/ngap/aper"

type PRSResourceSetID struct {
	Value int64 `lb:0,ub:7,mandatory`
}

func (m *PRSResourceSetID) Encode(w *aper.AperWriter) error {
	c := &aper.Constraint{Lb: 0, Ub: 7}
	return w.WriteInteger(m.Value, c, false)
}

func (m *PRSResourceSetID) Decode(r *aper.AperReader) error {
	c := &aper.Constraint{Lb: 0, Ub: 7}
	val, err := r.ReadInteger(c, false)
	if err != nil {
		return err
	}
	m.Value = val
	return nil
}
