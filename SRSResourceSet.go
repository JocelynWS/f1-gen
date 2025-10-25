package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SRSResourceSet struct {
	SRSResourceSetID  int64           `lb:0,ub:15,mandatory`
	SRSResourceIDList []int64         `lb:1,ub:maxnoSRSResourcePerSet,mandatory`
	ResourceSetType   ResourceSetType `mandatory`
	// IEExtensions *ProtocolExtensionContainerSRSResourceSetExtIEs `optional`
}

func (ie *SRSResourceSet) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	if err = w.WriteBits([]byte{0x0}, 1); err != nil {
		return
	}

	tmp_SRSResourceSetID := NewINTEGER(ie.SRSResourceSetID, aper.Constraint{Lb: 0, Ub: 15}, false)
	if err = tmp_SRSResourceSetID.Encode(w); err != nil {
		err = utils.WrapError("Encode SRSResourceSetID", err)
		return
	}

	tmp_List := Sequence[*INTEGER]{
		Value: []*INTEGER{},
		c:     aper.Constraint{Lb: 1, Ub: maxnoSRSResourcePerSet},
		ext:   false,
	}
	for _, id := range ie.SRSResourceIDList {
		tmp_List.Value = append(tmp_List.Value, &INTEGER{
			Value: aper.Integer(id),
			c:     aper.Constraint{Lb: 0, Ub: 63},
			ext:   false,
		})
	}

	if err = tmp_List.Encode(w); err != nil {
		err = utils.WrapError("Encode SRSResourceIDList", err)
		return
	}

	if err = ie.ResourceSetType.Encode(w); err != nil {
		err = utils.WrapError("Encode ResourceSetType", err)
		return
	}

	return
}

func (ie *SRSResourceSet) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	if _, err = r.ReadBits(1); err != nil {
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

	tmp_List := Sequence[*INTEGER]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoSRSResourcePerSet},
		ext: false,
	}
	fn := func() *INTEGER { return &INTEGER{c: aper.Constraint{Lb: 0, Ub: 63}, ext: false} }
	if err = tmp_List.Decode(r, fn); err != nil {
		err = utils.WrapError("Read SRSResourceIDList", err)
		return
	}
	ie.SRSResourceIDList = []int64{}
	for _, item := range tmp_List.Value {
		ie.SRSResourceIDList = append(ie.SRSResourceIDList, int64(item.Value))
	}

	if err = ie.ResourceSetType.Decode(r); err != nil {
		err = utils.WrapError("Read ResourceSetType", err)
		return
	}

	return
}
