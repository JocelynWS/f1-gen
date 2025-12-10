package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type NRUESidelinkAggregateMaximumBitrate struct {
	UENRSidelinkAggregateMaximumBitrate int64 `lb:0,ub:4000000000000,mandatory,valueExt`
	// IEExtensions * `optional`
}

func (ie *NRUESidelinkAggregateMaximumBitrate) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_UENRSidelinkAggregateMaximumBitrate := NewINTEGER(ie.UENRSidelinkAggregateMaximumBitrate, aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
	if err = tmp_UENRSidelinkAggregateMaximumBitrate.Encode(w); err != nil {
		err = utils.WrapError("Encode UENRSidelinkAggregateMaximumBitrate", err)
		return
	}
	return
}
func (ie *NRUESidelinkAggregateMaximumBitrate) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_UENRSidelinkAggregateMaximumBitrate := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: true,
	}
	if err = tmp_UENRSidelinkAggregateMaximumBitrate.Decode(r); err != nil {
		err = utils.WrapError("Read UENRSidelinkAggregateMaximumBitrate", err)
		return
	}
	ie.UENRSidelinkAggregateMaximumBitrate = int64(tmp_UENRSidelinkAggregateMaximumBitrate.Value)
	return
}
