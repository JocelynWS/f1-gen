package f1ap

import "github.com/lvdund/ngap/aper"

// NZP-CSI-RS-ResourceID ::= INTEGER (0..191)
type NZPCSIRSResourceID struct {
	Value int64 `lb:0,ub:191,mandatory`
}

func (m *NZPCSIRSResourceID) Encode(w *aper.AperWriter) error {
	c := &aper.Constraint{Lb: 0, Ub: 191}
	return w.WriteInteger(m.Value, c, false)
}

func (m *NZPCSIRSResourceID) Decode(r *aper.AperReader) error {
	c := &aper.Constraint{Lb: 0, Ub: 191}
	val, err := r.ReadInteger(c, false)
	if err != nil {
		return err
	}
	m.Value = val
	return nil
}
