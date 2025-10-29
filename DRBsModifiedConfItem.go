package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DRBsModifiedConfItem struct {
	DRBID                           int64                             `lb:1,ub:32,mandatory`
	ULUPTNLInformationToBeSetupList []ULUPTNLInformationToBeSetupItem `lb:1,ub:maxnoofULUPTNLInformation,mandatory`
	// IEExtensions * `optional`
}

func (ie *DRBsModifiedConfItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_DRBID := NewINTEGER(ie.DRBID, aper.Constraint{Lb: 1, Ub: 32}, false)
	if err = tmp_DRBID.Encode(w); err != nil {
		err = utils.WrapError("Encode DRBID", err)
		return
	}
	if len(ie.ULUPTNLInformationToBeSetupList) > 0 {
		tmp := Sequence[*ULUPTNLInformationToBeSetupItem]{
			Value: []*ULUPTNLInformationToBeSetupItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofULUPTNLInformation},
			ext:   false,
		}
		for _, i := range ie.ULUPTNLInformationToBeSetupList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode ULUPTNLInformationToBeSetupList", err)
			return
		}
	} else {
		err = utils.WrapError("ULUPTNLInformationToBeSetupList is nil", err)
		return
	}
	return
}
func (ie *DRBsModifiedConfItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_DRBID := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 32},
		ext: false,
	}
	if err = tmp_DRBID.Decode(r); err != nil {
		err = utils.WrapError("Read DRBID", err)
		return
	}
	ie.DRBID = int64(tmp_DRBID.Value)
	tmp_ULUPTNLInformationToBeSetupList := Sequence[*ULUPTNLInformationToBeSetupItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofULUPTNLInformation},
		ext: false,
	}
	fn := func() *ULUPTNLInformationToBeSetupItem { return new(ULUPTNLInformationToBeSetupItem) }
	if err = tmp_ULUPTNLInformationToBeSetupList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read ULUPTNLInformationToBeSetupList", err)
		return
	}
	ie.ULUPTNLInformationToBeSetupList = []ULUPTNLInformationToBeSetupItem{}
	for _, i := range tmp_ULUPTNLInformationToBeSetupList.Value {
		ie.ULUPTNLInformationToBeSetupList = append(ie.ULUPTNLInformationToBeSetupList, *i)
	}
	return
}
