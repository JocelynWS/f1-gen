package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DRBsModifiedItem struct {
	DRBID                           int64                             `lb:1,ub:32,mandatory`
	LCID                            *int64                            `lb:1,ub:32,optional`
	DLUPTNLInformationToBeSetupList []DLUPTNLInformationToBeSetupItem `lb:1,ub:maxnoofDLUPTNLInformation,mandatory`
	// IEExtensions * `optional`
}

func (ie *DRBsModifiedItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.LCID != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_DRBID := NewINTEGER(ie.DRBID, aper.Constraint{Lb: 1, Ub: 32}, false)
	if err = tmp_DRBID.Encode(w); err != nil {
		err = utils.WrapError("Encode DRBID", err)
		return
	}
	if ie.LCID != nil {
		tmp_LCID := NewINTEGER(*ie.LCID, aper.Constraint{Lb: 1, Ub: 32}, false)
		if err = tmp_LCID.Encode(w); err != nil {
			err = utils.WrapError("Encode LCID", err)
			return
		}
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
func (ie *DRBsModifiedItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
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
	if aper.IsBitSet(optionals, 1) {
		tmp_LCID := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 32},
			ext: false,
		}
		if err = tmp_LCID.Decode(r); err != nil {
			err = utils.WrapError("Read LCID", err)
			return
		}
		ie.LCID = (*int64)(&tmp_LCID.Value)
	}
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
