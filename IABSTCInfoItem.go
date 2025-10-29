package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type IABSTCInfoItem struct {
	SSBFreqInfo                 int64                      `lb:0,ub:3279165,mandatory`
	SSBSubcarrierSpacing        SSBSubcarrierSpacing       `mandatory`
	SSBTransmissionPeriodicity  SSBTransmissionPeriodicity `mandatory`
	SSBTransmissionTimingOffset int64                      `lb:0,ub:127,mandatory`
	SSBTransmissionBitmap       SSBTransmissionBitmap      `mandatory`
	// IEExtensions * `optional`
}

func (ie *IABSTCInfoItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_SSBFreqInfo := NewINTEGER(ie.SSBFreqInfo, aper.Constraint{Lb: 0, Ub: 3279165}, false)
	if err = tmp_SSBFreqInfo.Encode(w); err != nil {
		err = utils.WrapError("Encode SSBFreqInfo", err)
		return
	}
	if err = ie.SSBSubcarrierSpacing.Encode(w); err != nil {
		err = utils.WrapError("Encode SSBSubcarrierSpacing", err)
		return
	}
	if err = ie.SSBTransmissionPeriodicity.Encode(w); err != nil {
		err = utils.WrapError("Encode SSBTransmissionPeriodicity", err)
		return
	}
	tmp_SSBTransmissionTimingOffset := NewINTEGER(ie.SSBTransmissionTimingOffset, aper.Constraint{Lb: 0, Ub: 127}, false)
	if err = tmp_SSBTransmissionTimingOffset.Encode(w); err != nil {
		err = utils.WrapError("Encode SSBTransmissionTimingOffset", err)
		return
	}
	if err = ie.SSBTransmissionBitmap.Encode(w); err != nil {
		err = utils.WrapError("Encode SSBTransmissionBitmap", err)
		return
	}
	return
}
func (ie *IABSTCInfoItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_SSBFreqInfo := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 3279165},
		ext: false,
	}
	if err = tmp_SSBFreqInfo.Decode(r); err != nil {
		err = utils.WrapError("Read SSBFreqInfo", err)
		return
	}
	ie.SSBFreqInfo = int64(tmp_SSBFreqInfo.Value)
	if err = ie.SSBSubcarrierSpacing.Decode(r); err != nil {
		err = utils.WrapError("Read SSBSubcarrierSpacing", err)
		return
	}
	if err = ie.SSBTransmissionPeriodicity.Decode(r); err != nil {
		err = utils.WrapError("Read SSBTransmissionPeriodicity", err)
		return
	}
	tmp_SSBTransmissionTimingOffset := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 127},
		ext: false,
	}
	if err = tmp_SSBTransmissionTimingOffset.Decode(r); err != nil {
		err = utils.WrapError("Read SSBTransmissionTimingOffset", err)
		return
	}
	ie.SSBTransmissionTimingOffset = int64(tmp_SSBTransmissionTimingOffset.Value)
	if err = ie.SSBTransmissionBitmap.Decode(r); err != nil {
		err = utils.WrapError("Read SSBTransmissionBitmap", err)
		return
	}
	return
}
