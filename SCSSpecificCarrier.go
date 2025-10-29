package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SCSSpecificCarrier struct {
	OffsetToCarrier   int64                `lb:0,ub:2199,madatory,valExt`
	SubcarrierSpacing SubcarrierSpacingSCS `madatory,valExt`
	CarrierBandwidth  int64                `lb:1,ub:275,madatory,valExt`
	// IEExtensions *SCSSpecificCarrierExtIEs `optional`
}

func (ie *SCSSpecificCarrier) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if err = w.WriteBits(optionals, 1); err != nil {
		return
	}

	tmp_OffsetToCarrier := NewINTEGER(ie.OffsetToCarrier, aper.Constraint{Lb: 0, Ub: 2199}, true)
	if err = tmp_OffsetToCarrier.Encode(w); err != nil {
		err = utils.WrapError("Encode OffsetToCarrier", err)
		return
	}
	if err = ie.SubcarrierSpacing.Encode(w); err != nil {
		err = utils.WrapError("Encode SubcarrierSpacing", err)
		return
	}
	tmp_CarrierBandwidth := NewINTEGER(ie.CarrierBandwidth, aper.Constraint{Lb: 1, Ub: 275}, true)
	if err = tmp_CarrierBandwidth.Encode(w); err != nil {
		err = utils.WrapError("Encode CarrierBandwidth", err)
		return
	}
	return
}

func (ie *SCSSpecificCarrier) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}

	tmp_OffsetToCarrier := INTEGER{c: aper.Constraint{Lb: 0, Ub: 2199}, ext: true}
	if err = tmp_OffsetToCarrier.Decode(r); err != nil {
		err = utils.WrapError("Read OffsetToCarrier", err)
		return
	}
	ie.OffsetToCarrier = int64(tmp_OffsetToCarrier.Value)

	if err = ie.SubcarrierSpacing.Decode(r); err != nil {
		err = utils.WrapError("Read SubcarrierSpacing", err)
		return
	}
	tmp_CarrierBandwidth := INTEGER{c: aper.Constraint{Lb: 1, Ub: 275}, ext: true}
	if err = tmp_CarrierBandwidth.Decode(r); err != nil {
		err = utils.WrapError("Read CarrierBandwidth", err)
		return
	}
	ie.CarrierBandwidth = int64(tmp_CarrierBandwidth.Value)

	return
}
