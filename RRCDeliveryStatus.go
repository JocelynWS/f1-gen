package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type RRCDeliveryStatus struct {
	DeliveryStatus    int64 `lb:0,ub:4095,mandatory`
	TriggeringMessage int64 `lb:0,ub:4095,mandatory`
	// IEExtensions * `optional`
}

func (ie *RRCDeliveryStatus) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)

	tmp_DeliveryStatus := NewINTEGER(ie.DeliveryStatus, aper.Constraint{Lb: 0, Ub: 4095}, false)
	if err = tmp_DeliveryStatus.Encode(w); err != nil {
		err = utils.WrapError("Encode DeliveryStatus", err)
		return
	}

	tmp_TriggeringMessage := NewINTEGER(ie.TriggeringMessage, aper.Constraint{Lb: 0, Ub: 4095}, false)
	if err = tmp_TriggeringMessage.Encode(w); err != nil {
		err = utils.WrapError("Encode TriggeringMessage", err)
		return
	}

	return
}

func (ie *RRCDeliveryStatus) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	if _, err = r.ReadBits(1); err != nil {
		return
	}

	tmp_DeliveryStatus := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4095},
		ext: false,
	}
	if err = tmp_DeliveryStatus.Decode(r); err != nil {
		err = utils.WrapError("Read DeliveryStatus", err)
		return
	}
	ie.DeliveryStatus = int64(tmp_DeliveryStatus.Value)

	tmp_TriggeringMessage := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4095},
		ext: false,
	}
	if err = tmp_TriggeringMessage.Decode(r); err != nil {
		err = utils.WrapError("Read TriggeringMessage", err)
		return
	}
	ie.TriggeringMessage = int64(tmp_TriggeringMessage.Value)

	return
}
