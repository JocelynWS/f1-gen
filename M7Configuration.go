package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type M7Configuration struct {
	M7period     int64        `lb:1,ub:60,mandatory,valExt`
	M7LinksToLog M7LinksToLog `mandatory`
	// IEExtensions * `optional`
}

func (ie *M7Configuration) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_M7period := NewINTEGER(ie.M7period, aper.Constraint{Lb: 1, Ub: 60}, true)
	if err = tmp_M7period.Encode(w); err != nil {
		err = utils.WrapError("Encode M7period", err)
		return
	}
	if err = ie.M7LinksToLog.Encode(w); err != nil {
		err = utils.WrapError("Encode M7LinksToLog", err)
		return
	}
	return
}
func (ie *M7Configuration) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_M7period := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 60},
		ext: true,
	}
	if err = tmp_M7period.Decode(r); err != nil {
		err = utils.WrapError("Read M7period", err)
		return
	}
	ie.M7period = int64(tmp_M7period.Value)
	if err = ie.M7LinksToLog.Decode(r); err != nil {
		err = utils.WrapError("Read M7LinksToLog", err)
		return
	}
	return
}
