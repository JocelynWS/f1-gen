package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SSBInformation struct {
	SSBInformationList []SSBInformationItem `lb:1,ub:maxnoofSSBs,mandatory,valueExt`
	// IEExtensions * `optional`
}

func (ie *SSBInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if len(ie.SSBInformationList) > 0 {
		tmp := Sequence[*SSBInformationItem]{
			Value: []*SSBInformationItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofSSBs},
			ext:   true,
		}
		for _, i := range ie.SSBInformationList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode SSBInformationList", err)
			return
		}
	} else {
		err = utils.WrapError("SSBInformationList is nil", err)
		return
	}
	return
}
func (ie *SSBInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_SSBInformationList := Sequence[*SSBInformationItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofSSBs},
		ext: true,
	}
	fn := func() *SSBInformationItem { return new(SSBInformationItem) }
	if err = tmp_SSBInformationList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read SSBInformationList", err)
		return
	}
	ie.SSBInformationList = []SSBInformationItem{}
	for _, i := range tmp_SSBInformationList.Value {
		ie.SSBInformationList = append(ie.SSBInformationList, *i)
	}
	return
}
