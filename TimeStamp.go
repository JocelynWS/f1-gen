package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TimeStamp struct {
	SystemFrameNumber int64              `lb:0,ub:1023,mandatory`
	SlotIndex         TimeStampSlotIndex `mandatory`
	MeasurementTime   *aper.BitString    `lb:64,ub:64,optional`
	// IEExtension * `optional`
}

func (ie *TimeStamp) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.MeasurementTime != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_SystemFrameNumber := NewINTEGER(ie.SystemFrameNumber, aper.Constraint{Lb: 0, Ub: 1023}, false)
	if err = tmp_SystemFrameNumber.Encode(w); err != nil {
		err = utils.WrapError("Encode SystemFrameNumber", err)
		return
	}
	if err = ie.SlotIndex.Encode(w); err != nil {
		err = utils.WrapError("Encode SlotIndex", err)
		return
	}
	if ie.MeasurementTime != nil {
		tmp_MeasurementTime := NewBITSTRING(ie.MeasurementTime, aper.Constraint{Lb: 64, Ub: 64}, false)
		if err = tmp_MeasurementTime.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasurementTime", err)
			return
		}
	}
	return
}
func (ie *TimeStamp) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_SystemFrameNumber := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 1023},
		ext: false,
	}
	if err = tmp_SystemFrameNumber.Decode(r); err != nil {
		err = utils.WrapError("Read SystemFrameNumber", err)
		return
	}
	ie.SystemFrameNumber = int64(tmp_SystemFrameNumber.Value)
	if err = ie.SlotIndex.Decode(r); err != nil {
		err = utils.WrapError("Read SlotIndex", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_MeasurementTime := BITSTRING{
			c:   aper.Constraint{Lb: 64, Ub: 64},
			ext: false,
		}
		if err = tmp_MeasurementTime.Decode(r); err != nil {
			err = utils.WrapError("Read MeasurementTime", err)
			return
		}
		ie.MeasurementTime = &aper.BitString{Bytes: tmp_MeasurementTime.Value.Bytes, NumBits: tmp_MeasurementTime.Value.NumBits}
	}
	return
}
