package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SSBAreaRadioResourceStatusItem struct {
	SSBIndex                  int64  `lb:0,ub:63,mandatory`
	SSBAreaDLGBRPRBusage      int64  `lb:0,ub:100,mandatory`
	SSBAreaULGBRPRBusage      int64  `lb:0,ub:100,mandatory`
	SSBAreaDLnonGBRPRBusage   int64  `lb:0,ub:100,mandatory`
	SSBAreaULnonGBRPRBusage   int64  `lb:0,ub:100,mandatory`
	SSBAreaDLTotalPRBusage    int64  `lb:0,ub:100,mandatory`
	SSBAreaULTotalPRBusage    int64  `lb:0,ub:100,mandatory`
	DLschedulingPDCCHCCEusage *int64 `lb:0,ub:100,optional`
	ULschedulingPDCCHCCEusage *int64 `lb:0,ub:100,optional`
}

func (ie *SSBAreaRadioResourceStatusItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.DLschedulingPDCCHCCEusage != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.ULschedulingPDCCHCCEusage != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)

	tmp_SSBIndex := NewINTEGER(ie.SSBIndex, aper.Constraint{Lb: 0, Ub: 63}, false)
	if err = tmp_SSBIndex.Encode(w); err != nil {
		err = utils.WrapError("Encode SSBIndex", err)
		return
	}

	tmp := []struct {
		val  int64
		lb   int64
		ub   int64
		name string
	}{
		{ie.SSBAreaDLGBRPRBusage, 0, 100, "SSBAreaDLGBRPRBusage"},
		{ie.SSBAreaULGBRPRBusage, 0, 100, "SSBAreaULGBRPRBusage"},
		{ie.SSBAreaDLnonGBRPRBusage, 0, 100, "SSBAreaDLnonGBRPRBusage"},
		{ie.SSBAreaULnonGBRPRBusage, 0, 100, "SSBAreaULnonGBRPRBusage"},
		{ie.SSBAreaDLTotalPRBusage, 0, 100, "SSBAreaDLTotalPRBusage"},
		{ie.SSBAreaULTotalPRBusage, 0, 100, "SSBAreaULTotalPRBusage"},
	}
	for _, t := range tmp {
		tmpInt := NewINTEGER(t.val, aper.Constraint{Lb: t.lb, Ub: t.ub}, false)
		if err = (&tmpInt).Encode(w); err != nil {
			err = utils.WrapError("Encode "+t.name, err)
			return
		}
	}

	if ie.DLschedulingPDCCHCCEusage != nil {
		tmp := NewINTEGER(*ie.DLschedulingPDCCHCCEusage, aper.Constraint{Lb: 0, Ub: 100}, false)
		if err = (&tmp).Encode(w); err != nil {
			err = utils.WrapError("Encode DLschedulingPDCCHCCEusage", err)
			return
		}
	}

	if ie.ULschedulingPDCCHCCEusage != nil {
		tmpInt := NewINTEGER(*ie.ULschedulingPDCCHCCEusage, aper.Constraint{Lb: 0, Ub: 100}, false)
		if err = (&tmpInt).Encode(w); err != nil {
			err = utils.WrapError("Encode ULschedulingPDCCHCCEusage", err)
			return
		}
	}

	return
}

func (ie *SSBAreaRadioResourceStatusItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}

	tmp_SSBIndex := INTEGER{c: aper.Constraint{Lb: 0, Ub: 63}, ext: false}
	if err = tmp_SSBIndex.Decode(r); err != nil {
		err = utils.WrapError("Read SSBIndex", err)
		return
	}
	ie.SSBIndex = int64(tmp_SSBIndex.Value)

	tmp := []struct {
		field *int64
		lb    int64
		ub    int64
		name  string
	}{
		{&ie.SSBAreaDLGBRPRBusage, 0, 100, "SSBAreaDLGBRPRBusage"},
		{&ie.SSBAreaULGBRPRBusage, 0, 100, "SSBAreaULGBRPRBusage"},
		{&ie.SSBAreaDLnonGBRPRBusage, 0, 100, "SSBAreaDLnonGBRPRBusage"},
		{&ie.SSBAreaULnonGBRPRBusage, 0, 100, "SSBAreaULnonGBRPRBusage"},
		{&ie.SSBAreaDLTotalPRBusage, 0, 100, "SSBAreaDLTotalPRBusage"},
		{&ie.SSBAreaULTotalPRBusage, 0, 100, "SSBAreaULTotalPRBusage"},
	}

	for _, t := range tmp {
		tmpInt := INTEGER{c: aper.Constraint{Lb: t.lb, Ub: t.ub}, ext: false}
		if err = tmpInt.Decode(r); err != nil {
			err = utils.WrapError("Read "+t.name, err)
			return
		}
		*t.field = int64(tmpInt.Value)
	}

	if aper.IsBitSet(optionals, 1) {
		tmp := INTEGER{c: aper.Constraint{Lb: 0, Ub: 100}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read DLschedulingPDCCHCCEusage", err)
			return
		}
		ie.DLschedulingPDCCHCCEusage = (*int64)(&tmp.Value)
	}

	if aper.IsBitSet(optionals, 2) {
		tmp := INTEGER{c: aper.Constraint{Lb: 0, Ub: 100}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read ULschedulingPDCCHCCEusage", err)
			return
		}
		ie.ULschedulingPDCCHCCEusage = (*int64)(&tmp.Value)
	}

	return
}

// IEExtensions * `optional`
