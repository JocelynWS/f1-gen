package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type BroadcastSNPNIDListItem struct {
	PLMNIdentity     []byte `lb:3,ub:3,mandatory`
	BroadcastNIDList []NID  `lb:1,ub:maxnoofNIDsupported,mandatory`
	// IEExtensions * `optional`
}

func (ie *BroadcastSNPNIDListItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PLMNIdentity := NewOCTETSTRING(ie.PLMNIdentity, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_PLMNIdentity.Encode(w); err != nil {
		err = utils.WrapError("Encode PLMNIdentity", err)
		return
	}
	if len(ie.BroadcastNIDList) > 0 {
		tmp := Sequence[*NID]{
			Value: []*NID{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofNIDsupported},
			ext:   false,
		}
		for _, i := range ie.BroadcastNIDList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode BroadcastNIDList", err)
			return
		}
	} else {
		err = utils.WrapError("BroadcastNIDList is nil", err)
		return
	}
	return
}
func (ie *BroadcastSNPNIDListItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_PLMNIdentity := OCTETSTRING{
		c:   aper.Constraint{Lb: 3, Ub: 3},
		ext: false,
	}
	if err = tmp_PLMNIdentity.Decode(r); err != nil {
		err = utils.WrapError("Read PLMNIdentity", err)
		return
	}
	ie.PLMNIdentity = tmp_PLMNIdentity.Value
	tmp_BroadcastNIDList := Sequence[*NID]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofNIDsupported},
		ext: false,
	}
	fn := func() *NID { return new(NID) }
	if err = tmp_BroadcastNIDList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read BroadcastNIDList", err)
		return
	}
	ie.BroadcastNIDList = []NID{}
	for _, i := range tmp_BroadcastNIDList.Value {
		ie.BroadcastNIDList = append(ie.BroadcastNIDList, *i)
	}
	return
}
