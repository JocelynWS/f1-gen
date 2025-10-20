package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PRSConfiguration struct {
	PRSResourceSetList []PRSResourceSetItem `lb:1,ub:maxnoofPRSResourceSets,mandatory,valExt`
	// IEExtensions * `optional`
}

func (ie *PRSConfiguration) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if len(ie.PRSResourceSetList) > 0 {
		tmp := Sequence[*PRSResourceSetItem]{
			Value: []*PRSResourceSetItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofPRSResourceSets},
			ext:   true,
		}
		for _, i := range ie.PRSResourceSetList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode PRSResourceSetList", err)
			return
		}
	} else {
		err = utils.WrapError("PRSResourceSetList is nil", err)
		return
	}
	return
}
func (ie *PRSConfiguration) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_PRSResourceSetList := Sequence[*PRSResourceSetItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofPRSResourceSets},
		ext: true,
	}
	fn := func() *PRSResourceSetItem { return new(PRSResourceSetItem) }
	if err = tmp_PRSResourceSetList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read PRSResourceSetList", err)
		return
	}
	ie.PRSResourceSetList = []PRSResourceSetItem{}
	for _, i := range tmp_PRSResourceSetList.Value {
		ie.PRSResourceSetList = append(ie.PRSResourceSetList, *i)
	}
	return
}
