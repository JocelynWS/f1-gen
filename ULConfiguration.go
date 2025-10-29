package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ULConfiguration struct {
	ULUEConfiguration ULUEConfiguration `mandatory`
	// IEExtensions * `optional`
}

func (ie *ULConfiguration) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.ULUEConfiguration.Encode(w); err != nil {
		err = utils.WrapError("Encode ULUEConfiguration", err)
		return
	}
	return
}
func (ie *ULConfiguration) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.ULUEConfiguration.Decode(r); err != nil {
		err = utils.WrapError("Read ULUEConfiguration", err)
		return
	}
	return
}
