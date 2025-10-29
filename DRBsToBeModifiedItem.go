package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DRBsToBeModifiedItem struct {
	DRBID                           int64                             `lb:1,ub:32,mandatory`
	QoSInformation                  *QoSInformation                   `optional`
	ULUPTNLInformationToBeSetupList []ULUPTNLInformationToBeSetupItem `lb:1,ub:maxnoofULUPTNLInformation,mandatory`
	ULConfiguration                 *ULConfiguration                  `optional`
	// IEExtensions * `optional`
}

func (ie *DRBsToBeModifiedItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.QoSInformation != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.ULConfiguration != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	tmp_DRBID := NewINTEGER(ie.DRBID, aper.Constraint{Lb: 1, Ub: 32}, false)
	if err = tmp_DRBID.Encode(w); err != nil {
		err = utils.WrapError("Encode DRBID", err)
		return
	}
	if ie.QoSInformation != nil {
		if err = ie.QoSInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode QoSInformation", err)
			return
		}
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
	if ie.ULConfiguration != nil {
		if err = ie.ULConfiguration.Encode(w); err != nil {
			err = utils.WrapError("Encode ULConfiguration", err)
			return
		}
	}
	return
}
func (ie *DRBsToBeModifiedItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
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
		tmp := new(QoSInformation)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read QoSInformation", err)
			return
		}
		ie.QoSInformation = tmp
	}
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
	if aper.IsBitSet(optionals, 2) {
		tmp := new(ULConfiguration)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read ULConfiguration", err)
			return
		}
		ie.ULConfiguration = tmp
	}
	return
}
