package f1ap

import "github.com/lvdund/ngap/aper"

const (
	PresenceOptional    aper.Enumerated = 0
	PresenceConditional aper.Enumerated = 1
	PresenceMandatory   aper.Enumerated = 2
)

type Presence struct {
	Value aper.Enumerated
}

func (ie *Presence) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false)
	return
}

func (ie *Presence) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	ie.Value = aper.Enumerated(v)
	return
}
