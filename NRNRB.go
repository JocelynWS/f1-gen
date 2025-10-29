package f1ap

import "github.com/lvdund/ngap/aper"

const (
	NRNRBNrpci                aper.Enumerated = 0
	NRNRBNrprachconfiglist    aper.Enumerated = 1
	NRNRBNrprachconfigitem    aper.Enumerated = 2
	NRNRBNrscs                aper.Enumerated = 3
	NRNRBMsg1Fdm              aper.Enumerated = 4
	NRNRBSsbperrachoccasion   aper.Enumerated = 5
	NRNRBFreqdomainlength     aper.Enumerated = 6
	NRNRBZerocorrelzoneconfig aper.Enumerated = 7
	NRNRBIeextension          aper.Enumerated = 8
)

type NRNRB struct {
	Value aper.Enumerated
}

func (ie *NRNRB) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 8}, true)
	return
}

func (ie *NRNRB) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 8}, true)
	ie.Value = aper.Enumerated(v)
	return
}
