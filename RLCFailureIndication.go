package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type RLCFailureIndication struct {
	AssocatedLCID int64 `lb:1,ub:32,mandatory`
	// IEExtensions * `optional`
}

func (ie *RLCFailureIndication) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_AssocatedLCID := NewINTEGER(ie.AssocatedLCID, aper.Constraint{Lb: 1, Ub: 32}, false)
	if err = tmp_AssocatedLCID.Encode(w); err != nil {
		err = utils.WrapError("Encode AssocatedLCID", err)
		return
	}
	return
}
func (ie *RLCFailureIndication) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_AssocatedLCID := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 32},
		ext: false,
	}
	if err = tmp_AssocatedLCID.Decode(r); err != nil {
		err = utils.WrapError("Read AssocatedLCID", err)
		return
	}
	ie.AssocatedLCID = int64(tmp_AssocatedLCID.Value)
	return
}
