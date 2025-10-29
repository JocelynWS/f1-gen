package f1ap

import "github.com/lvdund/ngap/aper"

const (
	NRSCSscs15  aper.Enumerated = 0
	NRSCSscs30  aper.Enumerated = 1
	NRSCSscs60  aper.Enumerated = 2
	NRSCSscs120 aper.Enumerated = 3
)

type NRSCS struct {
	Value aper.Enumerated
}

func (ie *NRSCS) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
	return
}

func (ie *NRSCS) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	ie.Value = aper.Enumerated(v)
	return
}
