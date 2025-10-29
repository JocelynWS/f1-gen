package f1ap

import "github.com/lvdund/ngap/aper"

const (
	HSNAUplinkHard         aper.Enumerated = 0
	HSNAUplinkSoft         aper.Enumerated = 1
	HSNAUplinkNotavailable aper.Enumerated = 2
)

type HSNAUplink struct {
	Value aper.Enumerated
}

func (ie *HSNAUplink) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false)
	return
}

func (ie *HSNAUplink) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	ie.Value = aper.Enumerated(v)
	return
}
