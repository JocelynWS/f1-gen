package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type FreqBandNrItem struct {
	FreqBandIndicatorNr  int64                      `lb:1,ub:1024,mandatory,valueExt`
	SupportedSULBandList []SupportedSULFreqBandItem `mandatory,lb:0,ub:maxnoofNrCellBands`
	// IEExtensions * `optional,ignore`
}

func (ie *FreqBandNrItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_FreqBandIndicatorNr := NewINTEGER(ie.FreqBandIndicatorNr, aper.Constraint{Lb: 1, Ub: 1024}, true)
	if err = tmp_FreqBandIndicatorNr.Encode(w); err != nil {
		err = utils.WrapError("Encode FreqBandIndicatorNr", err)
		return
	}
	tmp := Sequence[*SupportedSULFreqBandItem]{Value: []*SupportedSULFreqBandItem{}, c: aper.Constraint{Lb: 0, Ub: maxnoofNrCellBands}, ext: false}
	for i := range ie.SupportedSULBandList {
		tmp.Value = append(tmp.Value, &ie.SupportedSULBandList[i])
	}
	if err = tmp.Encode(w); err != nil {
		err = utils.WrapError("Encode SupportedSULBandList", err)
		return
	}
	return
}

func (ie *FreqBandNrItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_FreqBandIndicatorNr := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 1024},
		ext: true,
	}
	if err = tmp_FreqBandIndicatorNr.Decode(r); err != nil {
		err = utils.WrapError("Read FreqBandIndicatorNr", err)
		return
	}
	ie.FreqBandIndicatorNr = int64(tmp_FreqBandIndicatorNr.Value)
	tmp := Sequence[*SupportedSULFreqBandItem]{c: aper.Constraint{Lb: 0, Ub: maxnoofNrCellBands}, ext: false}
	fn := func() *SupportedSULFreqBandItem { return new(SupportedSULFreqBandItem) }
	if err = tmp.Decode(r, fn); err != nil {
		err = utils.WrapError("Read SupportedSULBandList", err)
		return
	}
	ie.SupportedSULBandList = []SupportedSULFreqBandItem{}
	for _, v := range tmp.Value {
		ie.SupportedSULBandList = append(ie.SupportedSULBandList, *v)
	}
	return
}
