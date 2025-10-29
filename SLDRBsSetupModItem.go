package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SLDRBsSetupModItem struct {
	SLDRBID int64 `lb:1,ub:512,mandatory`
	// IEExtensions * `optional`
}

func (ie *SLDRBsSetupModItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_SLDRBID := NewINTEGER(ie.SLDRBID, aper.Constraint{Lb: 1, Ub: 512}, false)
	if err = tmp_SLDRBID.Encode(w); err != nil {
		err = utils.WrapError("Encode SLDRBID", err)
		return
	}
	return
}
func (ie *SLDRBsSetupModItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_SLDRBID := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 512},
		ext: false,
	}
	if err = tmp_SLDRBID.Decode(r); err != nil {
		err = utils.WrapError("Read SLDRBID", err)
		return
	}
	ie.SLDRBID = int64(tmp_SLDRBID.Value)
	return
}
