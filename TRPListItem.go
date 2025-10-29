package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TRPListItem struct {
	TRPID int64 `lb:0,ub:4095,mandatory,valExt`
	// IEExtensions * `optional`
}

func (ie *TRPListItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_TRPID := NewINTEGER(ie.TRPID, aper.Constraint{Lb: 0, Ub: 4095}, true)
	if err = tmp_TRPID.Encode(w); err != nil {
		err = utils.WrapError("Encode TRPID", err)
		return
	}
	return
}
func (ie *TRPListItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_TRPID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4095},
		ext: true,
	}
	if err = tmp_TRPID.Decode(r); err != nil {
		err = utils.WrapError("Read TRPID", err)
		return
	}
	ie.TRPID = int64(tmp_TRPID.Value)
	return
}
