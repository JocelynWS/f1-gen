package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ULBHNonUPTrafficMapping struct {
	ULBHNonUPTrafficMappingList []ULBHNonUPTrafficMappingItem `lb:1,ub:maxnoofNonUPTrafficMappings,mandatory`
	// IEExtensions * `optional`
}

func (ie *ULBHNonUPTrafficMapping) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if len(ie.ULBHNonUPTrafficMappingList) > 0 {
		tmp := Sequence[*ULBHNonUPTrafficMappingItem]{
			Value: []*ULBHNonUPTrafficMappingItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofNonUPTrafficMappings},
			ext:   false,
		}
		for _, i := range ie.ULBHNonUPTrafficMappingList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode ULBHNonUPTrafficMappingList", err)
			return
		}
	} else {
		err = utils.WrapError("ULBHNonUPTrafficMappingList is nil", err)
		return
	}
	return
}
func (ie *ULBHNonUPTrafficMapping) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_ULBHNonUPTrafficMappingList := Sequence[*ULBHNonUPTrafficMappingItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofNonUPTrafficMappings},
		ext: false,
	}
	fn := func() *ULBHNonUPTrafficMappingItem { return new(ULBHNonUPTrafficMappingItem) }
	if err = tmp_ULBHNonUPTrafficMappingList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read ULBHNonUPTrafficMappingList", err)
		return
	}
	ie.ULBHNonUPTrafficMappingList = []ULBHNonUPTrafficMappingItem{}
	for _, i := range tmp_ULBHNonUPTrafficMappingList.Value {
		ie.ULBHNonUPTrafficMappingList = append(ie.ULBHNonUPTrafficMappingList, *i)
	}
	return
}
