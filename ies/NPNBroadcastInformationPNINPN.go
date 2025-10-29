package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type NPNBroadcastInformationPNINPN struct {
	BroadcastPNINPNIDInformation []BroadcastPNINPNIDListItem `lb:1,ub:maxnoofCAGsupported,mandatory`
	// IEExtension * `optional`
}

func (ie *NPNBroadcastInformationPNINPN) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if len(ie.BroadcastPNINPNIDInformation) > 0 {
		tmp := Sequence[*BroadcastPNINPNIDListItem]{
			Value: []*BroadcastPNINPNIDListItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofCAGsupported},
			ext:   false,
		}
		for _, i := range ie.BroadcastPNINPNIDInformation {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode BroadcastPNINPNIDInformation", err)
			return
		}
	} else {
		err = utils.WrapError("BroadcastPNINPNIDInformation is nil", err)
		return
	}
	return
}
func (ie *NPNBroadcastInformationPNINPN) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_BroadcastPNINPNIDInformation := Sequence[*BroadcastPNINPNIDListItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofCAGsupported},
		ext: false,
	}
	fn := func() *BroadcastPNINPNIDListItem { return new(BroadcastPNINPNIDListItem) }
	if err = tmp_BroadcastPNINPNIDInformation.Decode(r, fn); err != nil {
		err = utils.WrapError("Read BroadcastPNINPNIDInformation", err)
		return
	}
	ie.BroadcastPNINPNIDInformation = []BroadcastPNINPNIDListItem{}
	for _, i := range tmp_BroadcastPNINPNIDInformation.Value {
		ie.BroadcastPNINPNIDInformation = append(ie.BroadcastPNINPNIDInformation, *i)
	}
	return
}
