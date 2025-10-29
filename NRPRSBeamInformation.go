package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type NRPRSBeamInformation struct {
	NRPRSBeamInformationList []NRPRSBeamInformationItem `lb:1,ub:maxnoofPRSResourcesPerSet,mandatory,valExt`
	LCStoGCSTranslationList  []LCStoGCSTranslation      `lb:1,ub:maxnoofLCSGCSTranslation,optional,valExt`
	// IEExtensions * `optional`
}

func (ie *NRPRSBeamInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.LCStoGCSTranslationList != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if len(ie.NRPRSBeamInformationList) > 0 {
		tmp := Sequence[*NRPRSBeamInformationItem]{
			Value: []*NRPRSBeamInformationItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofPRSResourcesPerSet},
			ext:   true,
		}
		for _, i := range ie.NRPRSBeamInformationList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode NRPRSBeamInformationList", err)
			return
		}
	} else {
		err = utils.WrapError("NRPRSBeamInformationList is nil", err)
		return
	}
	if len(ie.LCStoGCSTranslationList) > 0 {
		tmp := Sequence[*LCStoGCSTranslation]{
			Value: []*LCStoGCSTranslation{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofLCSGCSTranslation},
			ext:   true,
		}
		for _, i := range ie.LCStoGCSTranslationList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode LCStoGCSTranslationList", err)
			return
		}
	}
	return
}
func (ie *NRPRSBeamInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_NRPRSBeamInformationList := Sequence[*NRPRSBeamInformationItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofPRSResourcesPerSet},
		ext: true,
	}
	fn := func() *NRPRSBeamInformationItem { return new(NRPRSBeamInformationItem) }
	if err = tmp_NRPRSBeamInformationList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read NRPRSBeamInformationList", err)
		return
	}
	ie.NRPRSBeamInformationList = []NRPRSBeamInformationItem{}
	for _, i := range tmp_NRPRSBeamInformationList.Value {
		ie.NRPRSBeamInformationList = append(ie.NRPRSBeamInformationList, *i)
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_LCStoGCSTranslationList := Sequence[*LCStoGCSTranslation]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofLCSGCSTranslation},
			ext: true,
		}
		fn := func() *LCStoGCSTranslation { return new(LCStoGCSTranslation) }
		if err = tmp_LCStoGCSTranslationList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read LCStoGCSTranslationList", err)
			return
		}
		ie.LCStoGCSTranslationList = []LCStoGCSTranslation{}
		for _, i := range tmp_LCStoGCSTranslationList.Value {
			ie.LCStoGCSTranslationList = append(ie.LCStoGCSTranslationList, *i)
		}
	}
	return
}
