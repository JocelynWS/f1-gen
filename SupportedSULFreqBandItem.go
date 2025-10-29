package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SupportedSULFreqBandItem struct {
	FreqBandIndicatorNr int64 `lb:1,ub:1024,mandatory,valExt`
	// IEExtensions * `optional`
}

func (ie *SupportedSULFreqBandItem) Encode(w *aper.AperWriter) (err error) {
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
	return
}
func (ie *SupportedSULFreqBandItem) Decode(r *aper.AperReader) (err error) {
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
	return
}
