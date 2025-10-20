package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SSBToReportItem struct {
	SSBIndex int64 `lb:0,ub:63,mandatory`
	// IEExtensions * `optional`
}

func (ie *SSBToReportItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)

	tmp_SSBIndex := NewINTEGER(ie.SSBIndex, aper.Constraint{Lb: 0, Ub: 63}, false)
	if err = tmp_SSBIndex.Encode(w); err != nil {
		err = utils.WrapError("Encode SSBIndex", err)
		return
	}

	return
}

func (ie *SSBToReportItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}

	tmp_SSBIndex := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 63},
		ext: false,
	}
	if err = tmp_SSBIndex.Decode(r); err != nil {
		err = utils.WrapError("Read SSBIndex", err)
		return
	}
	ie.SSBIndex = int64(tmp_SSBIndex.Value)

	return
}
