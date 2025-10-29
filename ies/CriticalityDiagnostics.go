package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type CriticalityDiagnostics struct {
	ProcedureCode             *int64                         `lb:0,ub:255,optional`
	TriggeringMessage         *TriggeringMessage             `optional`
	ProcedureCriticality      *Criticality                   `optional`
	TransactionID             *int64                         `lb:0,ub:255,optional`
	IEsCriticalityDiagnostics []CriticalityDiagnosticsIEItem `lb:1,ub:maxnoofErrors,optional,valExt`
	// IEExtensions *CriticalityDiagnosticsExtIEs `optional`
}

func (ie *CriticalityDiagnostics) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.ProcedureCode != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.TriggeringMessage != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.ProcedureCriticality != nil {
		aper.SetBit(optionals, 3)
	}
	if ie.TransactionID != nil {
		aper.SetBit(optionals, 4)
	}
	if ie.IEsCriticalityDiagnostics != nil {
		aper.SetBit(optionals, 5)
	}
	w.WriteBits(optionals, 6)
	if ie.ProcedureCode != nil {
		tmp_ProcedureCode := NewINTEGER(*ie.ProcedureCode, aper.Constraint{Lb: 0, Ub: 255}, false)
		if err = tmp_ProcedureCode.Encode(w); err != nil {
			err = utils.WrapError("Encode ProcedureCode", err)
			return
		}
	}
	if ie.TriggeringMessage != nil {
		if err = ie.TriggeringMessage.Encode(w); err != nil {
			err = utils.WrapError("Encode TriggeringMessage", err)
			return
		}
	}
	if ie.ProcedureCriticality != nil {
		if err = ie.ProcedureCriticality.Encode(w); err != nil {
			err = utils.WrapError("Encode ProcedureCriticality", err)
			return
		}
	}
	if ie.TransactionID != nil {
		tmp_TransactionID := NewINTEGER(*ie.TransactionID, aper.Constraint{Lb: 0, Ub: 255}, false)
		if err = tmp_TransactionID.Encode(w); err != nil {
			err = utils.WrapError("Encode TransactionID", err)
			return
		}
	}
	if len(ie.IEsCriticalityDiagnostics) > 0 {
		tmp := Sequence[*CriticalityDiagnosticsIEItem]{
			Value: []*CriticalityDiagnosticsIEItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofErrors},
			ext:   true,
		}
		for _, i := range ie.IEsCriticalityDiagnostics {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode IEsCriticalityDiagnostics", err)
			return
		}
	}
	return
}
func (ie *CriticalityDiagnostics) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(6); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_ProcedureCode := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 255},
			ext: false,
		}
		if err = tmp_ProcedureCode.Decode(r); err != nil {
			err = utils.WrapError("Read ProcedureCode", err)
			return
		}
		ie.ProcedureCode = (*int64)(&tmp_ProcedureCode.Value)
	}
	if aper.IsBitSet(optionals, 2) {
		tmp := new(TriggeringMessage)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read TriggeringMessage", err)
			return
		}
		ie.TriggeringMessage = tmp
	}
	if aper.IsBitSet(optionals, 3) {
		tmp := new(Criticality)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read ProcedureCriticality", err)
			return
		}
		ie.ProcedureCriticality = tmp
	}
	if aper.IsBitSet(optionals, 4) {
		tmp_TransactionID := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 255},
			ext: false,
		}
		if err = tmp_TransactionID.Decode(r); err != nil {
			err = utils.WrapError("Read TransactionID", err)
			return
		}
		ie.TransactionID = (*int64)(&tmp_TransactionID.Value)
	}
	if aper.IsBitSet(optionals, 5) {
		tmp_IEsCriticalityDiagnostics := Sequence[*CriticalityDiagnosticsIEItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofErrors},
			ext: true,
		}
		fn := func() *CriticalityDiagnosticsIEItem { return new(CriticalityDiagnosticsIEItem) }
		if err = tmp_IEsCriticalityDiagnostics.Decode(r, fn); err != nil {
			err = utils.WrapError("Read IEsCriticalityDiagnostics", err)
			return
		}
		ie.IEsCriticalityDiagnostics = []CriticalityDiagnosticsIEItem{}
		for _, i := range tmp_IEsCriticalityDiagnostics.Value {
			ie.IEsCriticalityDiagnostics = append(ie.IEsCriticalityDiagnostics, *i)
		}
	}
	return
}
