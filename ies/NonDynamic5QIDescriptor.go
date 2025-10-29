package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type NonDynamic5QIDescriptor struct {
	FiveQI             int64  `lb:0,ub:255,mandatory,valExt`
	QoSPriorityLevel   *int64 `lb:1,ub:127,optional`
	AveragingWindow    *int64 `lb:0,ub:4095,optional`
	MaxDataBurstVolume *int64 `lb:0,ub:4095,optional`
	// IEExtensions * `optional`
}

func (ie *NonDynamic5QIDescriptor) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.QoSPriorityLevel != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.AveragingWindow != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.MaxDataBurstVolume != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)
	tmp_FiveQI := NewINTEGER(ie.FiveQI, aper.Constraint{Lb: 0, Ub: 255}, true)
	if err = tmp_FiveQI.Encode(w); err != nil {
		err = utils.WrapError("Encode FiveQI", err)
		return
	}
	if ie.QoSPriorityLevel != nil {
		tmp_QoSPriorityLevel := NewINTEGER(*ie.QoSPriorityLevel, aper.Constraint{Lb: 1, Ub: 127}, false)
		if err = tmp_QoSPriorityLevel.Encode(w); err != nil {
			err = utils.WrapError("Encode QoSPriorityLevel", err)
			return
		}
	}
	if ie.AveragingWindow != nil {
		tmp_AveragingWindow := NewINTEGER(*ie.AveragingWindow, aper.Constraint{Lb: 0, Ub: 4095}, false)
		if err = tmp_AveragingWindow.Encode(w); err != nil {
			err = utils.WrapError("Encode AveragingWindow", err)
			return
		}
	}
	if ie.MaxDataBurstVolume != nil {
		tmp_MaxDataBurstVolume := NewINTEGER(*ie.MaxDataBurstVolume, aper.Constraint{Lb: 0, Ub: 4095}, false)
		if err = tmp_MaxDataBurstVolume.Encode(w); err != nil {
			err = utils.WrapError("Encode MaxDataBurstVolume", err)
			return
		}
	}
	return
}
func (ie *NonDynamic5QIDescriptor) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	tmp_FiveQI := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 255},
		ext: true,
	}
	if err = tmp_FiveQI.Decode(r); err != nil {
		err = utils.WrapError("Read FiveQI", err)
		return
	}
	ie.FiveQI = int64(tmp_FiveQI.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp_QoSPriorityLevel := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 127},
			ext: false,
		}
		if err = tmp_QoSPriorityLevel.Decode(r); err != nil {
			err = utils.WrapError("Read QoSPriorityLevel", err)
			return
		}
		ie.QoSPriorityLevel = (*int64)(&tmp_QoSPriorityLevel.Value)
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_AveragingWindow := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4095},
			ext: false,
		}
		if err = tmp_AveragingWindow.Decode(r); err != nil {
			err = utils.WrapError("Read AveragingWindow", err)
			return
		}
		ie.AveragingWindow = (*int64)(&tmp_AveragingWindow.Value)
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_MaxDataBurstVolume := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4095},
			ext: false,
		}
		if err = tmp_MaxDataBurstVolume.Decode(r); err != nil {
			err = utils.WrapError("Read MaxDataBurstVolume", err)
			return
		}
		ie.MaxDataBurstVolume = (*int64)(&tmp_MaxDataBurstVolume.Value)
	}
	return
}
