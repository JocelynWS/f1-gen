package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SRSResourceSetItem struct {
	NumSRSresourcesperset *int64                 `lb:1,ub:16,optional,valExt`
	PeriodicityList       []PeriodicityListItem  `lb:1,ub:maxnoSRSResourcePerSet,optional,valExt`
	SpatialRelationInfo   *SpatialRelationInfo   `optional`
	PathlossReferenceInfo *PathlossReferenceInfo `optional`
	// IEExtensions * `optional`
}

func (ie *SRSResourceSetItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.NumSRSresourcesperset != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.PeriodicityList != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.SpatialRelationInfo != nil {
		aper.SetBit(optionals, 3)
	}
	if ie.PathlossReferenceInfo != nil {
		aper.SetBit(optionals, 4)
	}
	w.WriteBits(optionals, 5)
	if ie.NumSRSresourcesperset != nil {
		tmp_NumSRSresourcesperset := NewINTEGER(*ie.NumSRSresourcesperset, aper.Constraint{Lb: 1, Ub: 16}, true)
		if err = tmp_NumSRSresourcesperset.Encode(w); err != nil {
			err = utils.WrapError("Encode NumSRSresourcesperset", err)
			return
		}
	}
	if len(ie.PeriodicityList) > 0 {
		tmp := Sequence[*PeriodicityListItem]{
			Value: []*PeriodicityListItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoSRSResourcePerSet},
			ext:   true,
		}
		for _, i := range ie.PeriodicityList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode PeriodicityList", err)
			return
		}
	}
	if ie.SpatialRelationInfo != nil {
		if err = ie.SpatialRelationInfo.Encode(w); err != nil {
			err = utils.WrapError("Encode SpatialRelationInfo", err)
			return
		}
	}
	if ie.PathlossReferenceInfo != nil {
		if err = ie.PathlossReferenceInfo.Encode(w); err != nil {
			err = utils.WrapError("Encode PathlossReferenceInfo", err)
			return
		}
	}
	return
}
func (ie *SRSResourceSetItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(5); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_NumSRSresourcesperset := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 16},
			ext: true,
		}
		if err = tmp_NumSRSresourcesperset.Decode(r); err != nil {
			err = utils.WrapError("Read NumSRSresourcesperset", err)
			return
		}
		ie.NumSRSresourcesperset = (*int64)(&tmp_NumSRSresourcesperset.Value)
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_PeriodicityList := Sequence[*PeriodicityListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoSRSResourcePerSet},
			ext: true,
		}
		fn := func() *PeriodicityListItem { return new(PeriodicityListItem) }
		if err = tmp_PeriodicityList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read PeriodicityList", err)
			return
		}
		ie.PeriodicityList = []PeriodicityListItem{}
		for _, i := range tmp_PeriodicityList.Value {
			ie.PeriodicityList = append(ie.PeriodicityList, *i)
		}
	}
	if aper.IsBitSet(optionals, 3) {
		tmp := new(SpatialRelationInfo)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read SpatialRelationInfo", err)
			return
		}
		ie.SpatialRelationInfo = tmp
	}
	if aper.IsBitSet(optionals, 4) {
		tmp := new(PathlossReferenceInfo)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read PathlossReferenceInfo", err)
			return
		}
		ie.PathlossReferenceInfo = tmp
	}
	return
}
