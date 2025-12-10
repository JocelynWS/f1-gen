package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PosSRSResourceSetItem struct {
	PossrsResourceSetID  int64              `lb:0,ub:15,mandatory`
	PossRSResourceIDList []SRSPosResourceID `lb:1,ub:maxnoSRSPosResourcePerSet,mandatory,valueExt`
	PosresourceSetType   PosResourceSetType `mandatory`
	// IEExtensions * `optional`
}

func (ie *PosSRSResourceSetItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)

	tmp_PossrsResourceSetID := NewINTEGER(ie.PossrsResourceSetID, aper.Constraint{Lb: 0, Ub: 15}, false)
	if err = tmp_PossrsResourceSetID.Encode(w); err != nil {
		err = utils.WrapError("Encode PossrsResourceSetID", err)
		return
	}
	if len(ie.PossRSResourceIDList) > 0 {
		tmp := Sequence[*SRSPosResourceID]{
			Value: []*SRSPosResourceID{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoSRSPosResourcePerSet},
			ext:   true,
		}
		for i := range ie.PossRSResourceIDList {
			tmp.Value = append(tmp.Value, &ie.PossRSResourceIDList[i])
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode PossRSResourceIDList", err)
			return
		}
	} else {
		err = utils.WrapError("PossRSResourceIDList is nil", err)
		return
	}
	if err = ie.PosresourceSetType.Encode(w); err != nil {
		err = utils.WrapError("Encode PosresourceSetType", err)
		return
	}

	return
}

func (ie *PosSRSResourceSetItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_PossrsResourceSetID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 15},
		ext: false,
	}
	if err = tmp_PossrsResourceSetID.Decode(r); err != nil {
		err = utils.WrapError("Read PossrsResourceSetID", err)
		return
	}
	ie.PossrsResourceSetID = int64(tmp_PossrsResourceSetID.Value)
	tmp_PossRSResourceIDList := Sequence[*SRSPosResourceID]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoSRSPosResourcePerSet},
		ext: true,
	}
	fn := func() *SRSPosResourceID { return new(SRSPosResourceID) }
	if err = tmp_PossRSResourceIDList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read PossRSResourceIDList", err)
		return
	}
	ie.PossRSResourceIDList = []SRSPosResourceID{}
	for _, i := range tmp_PossRSResourceIDList.Value {
		ie.PossRSResourceIDList = append(ie.PossRSResourceIDList, *i)
	}
	if err = ie.PosresourceSetType.Decode(r); err != nil {
		err = utils.WrapError("Read PosresourceSetType", err)
		return
	}
	return
}
