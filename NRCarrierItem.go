package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type NRCarrierItem struct {
	CarrierSCS       NRSCS `mandatory`
	OffsetToCarrier  int64 `lb:0,ub:2199,mandatory,valExt`
	CarrierBandwidth int64 `lb:0,ub:0,mandatory`
	// IEExtension * `optional`
}

func (ie *NRCarrierItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.CarrierSCS.Encode(w); err != nil {
		err = utils.WrapError("Encode CarrierSCS", err)
		return
	}
	tmp_OffsetToCarrier := NewINTEGER(ie.OffsetToCarrier, aper.Constraint{Lb: 0, Ub: 2199}, true)
	if err = tmp_OffsetToCarrier.Encode(w); err != nil {
		err = utils.WrapError("Encode OffsetToCarrier", err)
		return
	}
	tmp_CarrierBandwidth := NewINTEGER(ie.CarrierBandwidth, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_CarrierBandwidth.Encode(w); err != nil {
		err = utils.WrapError("Encode CarrierBandwidth", err)
		return
	}
	return
}
func (ie *NRCarrierItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.CarrierSCS.Decode(r); err != nil {
		err = utils.WrapError("Read CarrierSCS", err)
		return
	}
	tmp_OffsetToCarrier := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 2199},
		ext: true,
	}
	if err = tmp_OffsetToCarrier.Decode(r); err != nil {
		err = utils.WrapError("Read OffsetToCarrier", err)
		return
	}
	ie.OffsetToCarrier = int64(tmp_OffsetToCarrier.Value)
	tmp_CarrierBandwidth := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_CarrierBandwidth.Decode(r); err != nil {
		err = utils.WrapError("Read CarrierBandwidth", err)
		return
	}
	ie.CarrierBandwidth = int64(tmp_CarrierBandwidth.Value)
	return
}
