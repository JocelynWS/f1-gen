package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type BroadcastPNINPNIDListItem struct {
	PLMNIdentity     []byte  `lb:3,ub:3,mandatory`
	BroadcastCAGList []CAGID `lb:1,ub:maxnoofCAGsupported,mandatory`
	// IEExtensions * `optional`
}

func (ie *BroadcastPNINPNIDListItem) Encode(w *aper.AperWriter) (err error) {
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
	if len(ie.BroadcastCAGList) > 0 {
		tmp := Sequence[*CAGID]{
			Value: []*CAGID{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofCAGsupported},
			ext:   false,
		}
		for _, i := range ie.BroadcastCAGList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode BroadcastCAGList", err)
			return
		}
	} else {
		err = utils.WrapError("BroadcastCAGList is nil", err)
		return
	}
	return
}
func (ie *BroadcastPNINPNIDListItem) Decode(r *aper.AperReader) (err error) {
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
	tmp_BroadcastCAGList := Sequence[*CAGID]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofCAGsupported},
		ext: false,
	}
	fn := func() *CAGID { return new(CAGID) }
	if err = tmp_BroadcastCAGList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read BroadcastCAGList", err)
		return
	}
	ie.BroadcastCAGList = []CAGID{}
	for _, i := range tmp_BroadcastCAGList.Value {
		ie.BroadcastCAGList = append(ie.BroadcastCAGList, *i)
	}
	return
}
