package f1ap

import "github.com/lvdund/ngap/aper"

const (
	DURXMTTXSupported    aper.Enumerated = 0
	DURXMTTXNotSupported aper.Enumerated = 1
)

type DURXMTTX struct {
	Value aper.Enumerated
}

func (ie *DURXMTTX) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false)
	return
}

func (ie *DURXMTTX) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false)
	ie.Value = aper.Enumerated(v)
	return
}
