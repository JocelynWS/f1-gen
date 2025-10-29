package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ServedCellInformation struct {
	NRCGI                          NRCGI             `mandatory`
	NRPCI                          NRPCI             `mandatory`
	FiveGSTAC                      []byte            `lb:3,ub:3,optional`
	ConfiguredEPSTAC               []byte            `lb:2,ub:2,optional`
	ServedPLMNs                    []ServedPLMNsItem `lb:1,ub:maxnoofBPLMNs,mandatory`
	NRModeInfo                     NRModeInfo        `mandatory`
	MeasurementTimingConfiguration []byte            `lb:0,ub:0,mandatory`
	// IEExtensions * `optional`
}

func (ie *ServedCellInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.FiveGSTAC != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.ConfiguredEPSTAC != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if err = ie.NRCGI.Encode(w); err != nil {
		err = utils.WrapError("Encode NRCGI", err)
		return
	}
	if err = ie.NRPCI.Encode(w); err != nil {
		err = utils.WrapError("Encode NRPCI", err)
		return
	}
	if ie.FiveGSTAC != nil {
		tmp_FiveGSTAC := NewOCTETSTRING(ie.FiveGSTAC, aper.Constraint{Lb: 3, Ub: 3}, false)
		if err = tmp_FiveGSTAC.Encode(w); err != nil {
			err = utils.WrapError("Encode FiveGSTAC", err)
			return
		}
	}
	if ie.ConfiguredEPSTAC != nil {
		tmp_ConfiguredEPSTAC := NewOCTETSTRING(ie.ConfiguredEPSTAC, aper.Constraint{Lb: 2, Ub: 2}, false)
		if err = tmp_ConfiguredEPSTAC.Encode(w); err != nil {
			err = utils.WrapError("Encode ConfiguredEPSTAC", err)
			return
		}
	}
	if len(ie.ServedPLMNs) > 0 {
		tmp := Sequence[*ServedPLMNsItem]{
			Value: []*ServedPLMNsItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofBPLMNs},
			ext:   false,
		}
		for _, i := range ie.ServedPLMNs {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode ServedPLMNs", err)
			return
		}
	} else {
		err = utils.WrapError("ServedPLMNs is nil", err)
		return
	}
	if err = ie.NRModeInfo.Encode(w); err != nil {
		err = utils.WrapError("Encode NRModeInfo", err)
		return
	}
	tmp_MeasurementTimingConfiguration := NewOCTETSTRING(ie.MeasurementTimingConfiguration, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_MeasurementTimingConfiguration.Encode(w); err != nil {
		err = utils.WrapError("Encode MeasurementTimingConfiguration", err)
		return
	}
	return
}
func (ie *ServedCellInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if err = ie.NRCGI.Decode(r); err != nil {
		err = utils.WrapError("Read NRCGI", err)
		return
	}
	if err = ie.NRPCI.Decode(r); err != nil {
		err = utils.WrapError("Read NRPCI", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_FiveGSTAC := OCTETSTRING{
			c:   aper.Constraint{Lb: 3, Ub: 3},
			ext: false,
		}
		if err = tmp_FiveGSTAC.Decode(r); err != nil {
			err = utils.WrapError("Read FiveGSTAC", err)
			return
		}
		ie.FiveGSTAC = tmp_FiveGSTAC.Value
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_ConfiguredEPSTAC := OCTETSTRING{
			c:   aper.Constraint{Lb: 2, Ub: 2},
			ext: false,
		}
		if err = tmp_ConfiguredEPSTAC.Decode(r); err != nil {
			err = utils.WrapError("Read ConfiguredEPSTAC", err)
			return
		}
		ie.ConfiguredEPSTAC = tmp_ConfiguredEPSTAC.Value
	}
	tmp_ServedPLMNs := Sequence[*ServedPLMNsItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofBPLMNs},
		ext: false,
	}
	fn := func() *ServedPLMNsItem { return new(ServedPLMNsItem) }
	if err = tmp_ServedPLMNs.Decode(r, fn); err != nil {
		err = utils.WrapError("Read ServedPLMNs", err)
		return
	}
	ie.ServedPLMNs = []ServedPLMNsItem{}
	for _, i := range tmp_ServedPLMNs.Value {
		ie.ServedPLMNs = append(ie.ServedPLMNs, *i)
	}
	if err = ie.NRModeInfo.Decode(r); err != nil {
		err = utils.WrapError("Read NRModeInfo", err)
		return
	}
	tmp_MeasurementTimingConfiguration := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_MeasurementTimingConfiguration.Decode(r); err != nil {
		err = utils.WrapError("Read MeasurementTimingConfiguration", err)
		return
	}
	ie.MeasurementTimingConfiguration = tmp_MeasurementTimingConfiguration.Value
	return
}
