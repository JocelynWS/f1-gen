package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SemipersistentSRS struct {
	SRSResourceSetID   int64                `lb:0,ub:15,mandatory`
	SRSSpatialRelation *SpatialRelationInfo `optional`
	// IEExtensions * `optional`
}

func (ie *SemipersistentSRS) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.SRSSpatialRelation != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_SRSResourceSetID := NewINTEGER(ie.SRSResourceSetID, aper.Constraint{Lb: 0, Ub: 15}, false)
	if err = tmp_SRSResourceSetID.Encode(w); err != nil {
		err = utils.WrapError("Encode SRSResourceSetID", err)
		return
	}
	if ie.SRSSpatialRelation != nil {
		if err = ie.SRSSpatialRelation.Encode(w); err != nil {
			err = utils.WrapError("Encode SRSSpatialRelation", err)
			return
		}
	}
	return
}
func (ie *SemipersistentSRS) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_SRSResourceSetID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 15},
		ext: false,
	}
	if err = tmp_SRSResourceSetID.Decode(r); err != nil {
		err = utils.WrapError("Read SRSResourceSetID", err)
		return
	}
	ie.SRSResourceSetID = int64(tmp_SRSResourceSetID.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp := new(SpatialRelationInfo)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read SRSSpatialRelation", err)
			return
		}
		ie.SRSSpatialRelation = tmp
	}
	return
}
