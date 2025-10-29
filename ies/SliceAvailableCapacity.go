package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SliceAvailableCapacity struct {
	SliceAvailableCapacityList []SliceAvailableCapacityItem `lb:1,ub:maxnoofSliceItems,mandatory,valExt`
	// IEExtensions * `optional`
}

func (ie *SliceAvailableCapacity) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if len(ie.SliceAvailableCapacityList) > 0 {
		tmp := Sequence[*SliceAvailableCapacityItem]{
			Value: []*SliceAvailableCapacityItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofSliceItems},
			ext:   true,
		}
		for _, i := range ie.SliceAvailableCapacityList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode SliceAvailableCapacityList", err)
			return
		}
	} else {
		err = utils.WrapError("SliceAvailableCapacityList is nil", err)
		return
	}
	return
}
func (ie *SliceAvailableCapacity) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_SliceAvailableCapacityList := Sequence[*SliceAvailableCapacityItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofSliceItems},
		ext: true,
	}
	fn := func() *SliceAvailableCapacityItem { return new(SliceAvailableCapacityItem) }
	if err = tmp_SliceAvailableCapacityList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read SliceAvailableCapacityList", err)
		return
	}
	ie.SliceAvailableCapacityList = []SliceAvailableCapacityItem{}
	for _, i := range tmp_SliceAvailableCapacityList.Value {
		ie.SliceAvailableCapacityList = append(ie.SliceAvailableCapacityList, *i)
	}
	return
}
