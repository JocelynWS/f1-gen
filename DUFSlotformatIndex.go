package f1ap

import "github.com/lvdund/ngap/aper"

type DUFSlotformatIndex struct {
	Value int64 `lb:0,ub:254,mandatory`
}

func (m *DUFSlotformatIndex) Encode(w *aper.AperWriter) error {
	c := &aper.Constraint{Lb: 0, Ub: 254}
	return w.WriteInteger(m.Value, c, false)
}

func (m *DUFSlotformatIndex) Decode(r *aper.AperReader) error {
	c := &aper.Constraint{Lb: 0, Ub: 254}
	val, err := r.ReadInteger(c, false)
	if err != nil {
		return err
	}
	m.Value = val
	return nil
}
