package f1ap

import "github.com/lvdund/ngap/aper"

const (
	NotificationControlActive    aper.Enumerated = 0
	NotificationControlNotActive aper.Enumerated = 1
)

type NotificationControl struct {
	Value aper.Enumerated
}

func (ie *NotificationControl) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
	return
}

func (ie *NotificationControl) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	ie.Value = aper.Enumerated(v)
	return
}
