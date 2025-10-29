package f1ap

import "github.com/lvdund/ngap/aper"

const (
	CellDirectionDlOnly aper.Enumerated = 0
	CellDirectionUlOnly aper.Enumerated = 1
)

type CellDirection struct {
	Value aper.Enumerated
}

func (ie *CellDirection) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false)
	return
}

func (ie *CellDirection) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false)
	ie.Value = aper.Enumerated(v)
	return
}
