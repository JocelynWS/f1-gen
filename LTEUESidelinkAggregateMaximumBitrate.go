package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type LTEUESidelinkAggregateMaximumBitrate struct {
	UELTESidelinkAggregateMaximumBitrate int64 `lb:0,ub:4000000000000,mandatory,valExt`
	// IEExtensions * `optional`
}

func (ie *LTEUESidelinkAggregateMaximumBitrate) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_UELTESidelinkAggregateMaximumBitrate := NewINTEGER(ie.UELTESidelinkAggregateMaximumBitrate, aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
	if err = tmp_UELTESidelinkAggregateMaximumBitrate.Encode(w); err != nil {
		err = utils.WrapError("Encode UELTESidelinkAggregateMaximumBitrate", err)
		return
	}
	return
}
func (ie *LTEUESidelinkAggregateMaximumBitrate) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_UELTESidelinkAggregateMaximumBitrate := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: true,
	}
	if err = tmp_UELTESidelinkAggregateMaximumBitrate.Decode(r); err != nil {
		err = utils.WrapError("Read UELTESidelinkAggregateMaximumBitrate", err)
		return
	}
	ie.UELTESidelinkAggregateMaximumBitrate = int64(tmp_UELTESidelinkAggregateMaximumBitrate.Value)
	return
}
