package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DRBsRequiredToBeModifiedItem struct {
	DRBID                           int64                             `lb:1,ub:32,mandatory,valueExt`
	DLUPTNLInformationToBeSetupList []DLUPTNLInformationToBeSetupItem `lb:1,ub:maxnoofDLUPTNLInformation,mandatory`
	// IEExtensions * `optional`
}

func (ie *DRBsRequiredToBeModifiedItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_DRBID := NewINTEGER(ie.DRBID, aper.Constraint{Lb: 1, Ub: 32}, true)
	if err = tmp_DRBID.Encode(w); err != nil {
		err = utils.WrapError("Encode DRBID", err)
		return
	}
	if len(ie.DLUPTNLInformationToBeSetupList) > 0 {
		tmp := Sequence[*DLUPTNLInformationToBeSetupItem]{
			Value: []*DLUPTNLInformationToBeSetupItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofDLUPTNLInformation},
			ext:   false,
		}
		for _, i := range ie.DLUPTNLInformationToBeSetupList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode DLUPTNLInformationToBeSetupList", err)
			return
		}
	} else {
		err = utils.WrapError("DLUPTNLInformationToBeSetupList is nil", err)
		return
	}
	return
}
func (ie *DRBsRequiredToBeModifiedItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_DRBID := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 32},
		ext: true,
	}
	if err = tmp_DRBID.Decode(r); err != nil {
		err = utils.WrapError("Read DRBID", err)
		return
	}
	ie.DRBID = int64(tmp_DRBID.Value)
	tmp_DLUPTNLInformationToBeSetupList := Sequence[*DLUPTNLInformationToBeSetupItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofDLUPTNLInformation},
		ext: false,
	}
	fn := func() *DLUPTNLInformationToBeSetupItem { return new(DLUPTNLInformationToBeSetupItem) }
	if err = tmp_DLUPTNLInformationToBeSetupList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read DLUPTNLInformationToBeSetupList", err)
		return
	}
	ie.DLUPTNLInformationToBeSetupList = []DLUPTNLInformationToBeSetupItem{}
	for _, i := range tmp_DLUPTNLInformationToBeSetupList.Value {
		ie.DLUPTNLInformationToBeSetupList = append(ie.DLUPTNLInformationToBeSetupList, *i)
	}
	return
}
