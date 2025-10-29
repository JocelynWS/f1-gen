package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type RACHReportInformationItem struct {
	RACHReportContainer  []byte `lb:0,ub:0,mandatory`
	UEAssitantIdentifier *int64 `lb:0,ub:4294967295,optional`
	// IEExtensions * `optional`
}

func (ie *RACHReportInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.UEAssitantIdentifier != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_RACHReportContainer := NewOCTETSTRING(ie.RACHReportContainer, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_RACHReportContainer.Encode(w); err != nil {
		err = utils.WrapError("Encode RACHReportContainer", err)
		return
	}
	if ie.UEAssitantIdentifier != nil {
		tmp_UEAssitantIdentifier := NewINTEGER(*ie.UEAssitantIdentifier, aper.Constraint{Lb: 0, Ub: 4294967295}, false)
		if err = tmp_UEAssitantIdentifier.Encode(w); err != nil {
			err = utils.WrapError("Encode UEAssitantIdentifier", err)
			return
		}
	}
	return
}
func (ie *RACHReportInformationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_RACHReportContainer := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_RACHReportContainer.Decode(r); err != nil {
		err = utils.WrapError("Read RACHReportContainer", err)
		return
	}
	ie.RACHReportContainer = tmp_RACHReportContainer.Value
	if aper.IsBitSet(optionals, 1) {
		tmp_UEAssitantIdentifier := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4294967295},
			ext: false,
		}
		if err = tmp_UEAssitantIdentifier.Decode(r); err != nil {
			err = utils.WrapError("Read UEAssitantIdentifier", err)
			return
		}
		ie.UEAssitantIdentifier = (*int64)(&tmp_UEAssitantIdentifier.Value)
	}
	return
}
