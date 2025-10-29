package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SRSConfiguration struct {
	SRSCarrierList []SRSCarrierListItem `lb:1,ub:maxnoSRSCarriers,mandatory,valExt`
	// IEExtensions * `optional`
}

func (ie *SRSConfiguration) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if len(ie.SRSCarrierList) > 0 {
		tmp := Sequence[*SRSCarrierListItem]{
			Value: []*SRSCarrierListItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoSRSCarriers},
			ext:   true,
		}
		for _, i := range ie.SRSCarrierList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode SRSCarrierList", err)
			return
		}
	} else {
		err = utils.WrapError("SRSCarrierList is nil", err)
		return
	}
	return
}
func (ie *SRSConfiguration) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_SRSCarrierList := Sequence[*SRSCarrierListItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoSRSCarriers},
		ext: true,
	}
	fn := func() *SRSCarrierListItem { return new(SRSCarrierListItem) }
	if err = tmp_SRSCarrierList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read SRSCarrierList", err)
		return
	}
	ie.SRSCarrierList = []SRSCarrierListItem{}
	for _, i := range tmp_SRSCarrierList.Value {
		ie.SRSCarrierList = append(ie.SRSCarrierList, *i)
	}
	return
}
