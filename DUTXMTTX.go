package f1ap

import "github.com/lvdund/ngap/aper"

const (
	DUTXMTTXSupported    aper.Enumerated = 0
	DUTXMTTXNotSupported aper.Enumerated = 1
)

type DUTXMTTX struct {
	Value aper.Enumerated
}

func (ie *DUTXMTTX) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false)
	return
}

func (ie *DUTXMTTX) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false)
	ie.Value = aper.Enumerated(v)
	return
}
