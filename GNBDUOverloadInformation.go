package f1ap

import "github.com/lvdund/ngap/aper"

const (
	GNBDUOverloadInformationOverloaded    aper.Enumerated = 0
	GNBDUOverloadInformationNotOverloaded aper.Enumerated = 1
)

type GNBDUOverloadInformation struct {
	Value aper.Enumerated
}

func (ie *GNBDUOverloadInformation) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false)
	return
}

func (ie *GNBDUOverloadInformation) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false)
	ie.Value = aper.Enumerated(v)
	return
}
