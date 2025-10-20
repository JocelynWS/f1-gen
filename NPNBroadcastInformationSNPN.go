package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type NPNBroadcastInformationSNPN struct {
	BroadcastSNPNIDList []BroadcastSNPNIDListItem `lb:1,ub:maxnoofNIDsupported,mandatory`
	// IEExtension * `optional`
}

func (ie *NPNBroadcastInformationSNPN) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if len(ie.BroadcastSNPNIDList) > 0 {
		tmp := Sequence[*BroadcastSNPNIDListItem]{
			Value: []*BroadcastSNPNIDListItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofNIDsupported},
			ext:   false,
		}
		for _, i := range ie.BroadcastSNPNIDList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode BroadcastSNPNIDList", err)
			return
		}
	} else {
		err = utils.WrapError("BroadcastSNPNIDList is nil", err)
		return
	}
	return
}
func (ie *NPNBroadcastInformationSNPN) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_BroadcastSNPNIDList := Sequence[*BroadcastSNPNIDListItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofNIDsupported},
		ext: false,
	}
	fn := func() *BroadcastSNPNIDListItem { return new(BroadcastSNPNIDListItem) }
	if err = tmp_BroadcastSNPNIDList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read BroadcastSNPNIDList", err)
		return
	}
	ie.BroadcastSNPNIDList = []BroadcastSNPNIDListItem{}
	for _, i := range tmp_BroadcastSNPNIDList.Value {
		ie.BroadcastSNPNIDList = append(ie.BroadcastSNPNIDList, *i)
	}
	return
}
