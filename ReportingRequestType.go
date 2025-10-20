package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ReportingRequestType struct {
	EventType                 EventType `mandatory`
	ReportingPeriodicityValue *int64    `lb:0,ub:512,optional`
	// IEExtensions *ProtocolExtensionContainer `optional`
}

func (ie *ReportingRequestType) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.ReportingPeriodicityValue != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)

	if err = ie.EventType.Encode(w); err != nil {
		err = utils.WrapError("Encode EventType", err)
		return
	}

	if ie.ReportingPeriodicityValue != nil {
		tmp := NewINTEGER(*ie.ReportingPeriodicityValue, aper.Constraint{Lb: 0, Ub: 512}, false)
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode ReportingPeriodicityValue", err)
			return
		}
	}

	return
}

func (ie *ReportingRequestType) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	optionals, err := r.ReadBits(2)
	if err != nil {
		return
	}

	if err = ie.EventType.Decode(r); err != nil {
		err = utils.WrapError("Read EventType", err)
		return
	}

	if aper.IsBitSet(optionals, 1) {
		tmp := INTEGER{c: aper.Constraint{Lb: 0, Ub: 512}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read ReportingPeriodicityValue", err)
			return
		}
		ie.ReportingPeriodicityValue = (*int64)(&tmp.Value)
	}

	return
}
