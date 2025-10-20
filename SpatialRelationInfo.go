package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SpatialRelationInfo struct {
	SpatialRelationforResourceID []SpatialRelationforResourceIDItem `lb:1,ub:maxnoofSpatialRelations,mandatory`
	// IEExtensions * `optional`
}

func (ie *SpatialRelationInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if len(ie.SpatialRelationforResourceID) > 0 {
		tmp := Sequence[*SpatialRelationforResourceIDItem]{
			Value: []*SpatialRelationforResourceIDItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofSpatialRelations},
			ext:   false,
		}
		for _, i := range ie.SpatialRelationforResourceID {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode SpatialRelationforResourceID", err)
			return
		}
	} else {
		err = utils.WrapError("SpatialRelationforResourceID is nil", err)
		return
	}
	return
}
func (ie *SpatialRelationInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_SpatialRelationforResourceID := Sequence[*SpatialRelationforResourceIDItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofSpatialRelations},
		ext: false,
	}
	fn := func() *SpatialRelationforResourceIDItem { return new(SpatialRelationforResourceIDItem) }
	if err = tmp_SpatialRelationforResourceID.Decode(r, fn); err != nil {
		err = utils.WrapError("Read SpatialRelationforResourceID", err)
		return
	}
	ie.SpatialRelationforResourceID = []SpatialRelationforResourceIDItem{}
	for _, i := range tmp_SpatialRelationforResourceID.Value {
		ie.SpatialRelationforResourceID = append(ie.SpatialRelationforResourceID, *i)
	}
	return
}
