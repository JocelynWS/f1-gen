package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TRPInformation struct {
	TRPID                          int64                            `lb:0,ub:4095,mandatory,valExt`
	TRPInformationTypeResponseList []TRPInformationTypeResponseItem `lb:1,ub:maxnoofTRPInfoTypes,mandatory,valExt`
	// IEExtensions * `optional`
}

func (ie *TRPInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_TRPID := NewINTEGER(ie.TRPID, aper.Constraint{Lb: 0, Ub: 4095}, true)
	if err = tmp_TRPID.Encode(w); err != nil {
		err = utils.WrapError("Encode TRPID", err)
		return
	}
	if len(ie.TRPInformationTypeResponseList) > 0 {
		tmp := Sequence[*TRPInformationTypeResponseItem]{
			Value: []*TRPInformationTypeResponseItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofTRPInfoTypes},
			ext:   true,
		}
		for _, i := range ie.TRPInformationTypeResponseList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode TRPInformationTypeResponseList", err)
			return
		}
	} else {
		err = utils.WrapError("TRPInformationTypeResponseList is nil", err)
		return
	}
	return
}
func (ie *TRPInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_TRPID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4095},
		ext: true,
	}
	if err = tmp_TRPID.Decode(r); err != nil {
		err = utils.WrapError("Read TRPID", err)
		return
	}
	ie.TRPID = int64(tmp_TRPID.Value)
	tmp_TRPInformationTypeResponseList := Sequence[*TRPInformationTypeResponseItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofTRPInfoTypes},
		ext: true,
	}
	fn := func() *TRPInformationTypeResponseItem { return new(TRPInformationTypeResponseItem) }
	if err = tmp_TRPInformationTypeResponseList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read TRPInformationTypeResponseList", err)
		return
	}
	ie.TRPInformationTypeResponseList = []TRPInformationTypeResponseItem{}
	for _, i := range tmp_TRPInformationTypeResponseList.Value {
		ie.TRPInformationTypeResponseList = append(ie.TRPInformationTypeResponseList, *i)
	}
	return
}
