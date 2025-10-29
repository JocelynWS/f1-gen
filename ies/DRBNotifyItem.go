package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DRBNotifyItem struct {
	DRBID             int64             `lb:1,ub:32,mandatory`
	NotificationCause NotificationCause `mandatory`
	// IEExtensions * `optional`
}

func (ie *DRBNotifyItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_DRBID := NewINTEGER(ie.DRBID, aper.Constraint{Lb: 1, Ub: 32}, false)
	if err = tmp_DRBID.Encode(w); err != nil {
		err = utils.WrapError("Encode DRBID", err)
		return
	}
	if err = ie.NotificationCause.Encode(w); err != nil {
		err = utils.WrapError("Encode NotificationCause", err)
		return
	}
	return
}
func (ie *DRBNotifyItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_DRBID := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 32},
		ext: false,
	}
	if err = tmp_DRBID.Decode(r); err != nil {
		err = utils.WrapError("Read DRBID", err)
		return
	}
	ie.DRBID = int64(tmp_DRBID.Value)
	if err = ie.NotificationCause.Decode(r); err != nil {
		err = utils.WrapError("Read NotificationCause", err)
		return
	}
	return
}
