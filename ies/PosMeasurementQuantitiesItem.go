package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PosMeasurementQuantitiesItem struct {
	PosMeasurementType               PosMeasurementType `mandatory`
	TimingReportingGranularityFactor *int64             `lb:0,ub:5,optional`
	// IEExtensions * `optional`
}

func (ie *PosMeasurementQuantitiesItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.TimingReportingGranularityFactor != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.PosMeasurementType.Encode(w); err != nil {
		err = utils.WrapError("Encode PosMeasurementType", err)
		return
	}
	if ie.TimingReportingGranularityFactor != nil {
		tmp_TimingReportingGranularityFactor := NewINTEGER(*ie.TimingReportingGranularityFactor, aper.Constraint{Lb: 0, Ub: 5}, false)
		if err = tmp_TimingReportingGranularityFactor.Encode(w); err != nil {
			err = utils.WrapError("Encode TimingReportingGranularityFactor", err)
			return
		}
	}
	return
}
func (ie *PosMeasurementQuantitiesItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.PosMeasurementType.Decode(r); err != nil {
		err = utils.WrapError("Read PosMeasurementType", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_TimingReportingGranularityFactor := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 5},
			ext: false,
		}
		if err = tmp_TimingReportingGranularityFactor.Decode(r); err != nil {
			err = utils.WrapError("Read TimingReportingGranularityFactor", err)
			return
		}
		ie.TimingReportingGranularityFactor = (*int64)(&tmp_TimingReportingGranularityFactor.Value)
	}
	return
}
