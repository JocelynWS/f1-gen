package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DUtoCURRCInformation struct {
	CellGroupConfig  []byte `lb:0,ub:0,mandatory`
	MeasGapConfig    []byte `lb:0,ub:0,optional`
	RequestedPMaxFR1 []byte `lb:0,ub:0,optional`
	// IEExtensions * `optional`
}

func (ie *DUtoCURRCInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.MeasGapConfig != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.RequestedPMaxFR1 != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	tmp_CellGroupConfig := NewOCTETSTRING(ie.CellGroupConfig, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_CellGroupConfig.Encode(w); err != nil {
		err = utils.WrapError("Encode CellGroupConfig", err)
		return
	}
	if ie.MeasGapConfig != nil {
		tmp_MeasGapConfig := NewOCTETSTRING(ie.MeasGapConfig, aper.Constraint{Lb: 0, Ub: 0}, false)
		if err = tmp_MeasGapConfig.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasGapConfig", err)
			return
		}
	}
	if ie.RequestedPMaxFR1 != nil {
		tmp_RequestedPMaxFR1 := NewOCTETSTRING(ie.RequestedPMaxFR1, aper.Constraint{Lb: 0, Ub: 0}, false)
		if err = tmp_RequestedPMaxFR1.Encode(w); err != nil {
			err = utils.WrapError("Encode RequestedPMaxFR1", err)
			return
		}
	}
	return
}
func (ie *DUtoCURRCInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	tmp_CellGroupConfig := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_CellGroupConfig.Decode(r); err != nil {
		err = utils.WrapError("Read CellGroupConfig", err)
		return
	}
	ie.CellGroupConfig = tmp_CellGroupConfig.Value
	if aper.IsBitSet(optionals, 1) {
		tmp_MeasGapConfig := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp_MeasGapConfig.Decode(r); err != nil {
			err = utils.WrapError("Read MeasGapConfig", err)
			return
		}
		ie.MeasGapConfig = tmp_MeasGapConfig.Value
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_RequestedPMaxFR1 := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp_RequestedPMaxFR1.Decode(r); err != nil {
			err = utils.WrapError("Read RequestedPMaxFR1", err)
			return
		}
		ie.RequestedPMaxFR1 = tmp_RequestedPMaxFR1.Value
	}
	return
}
