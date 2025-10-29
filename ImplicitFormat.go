package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ImplicitFormat struct {
	DUFSlotformatIndex DUFSlotformatIndex `mandatory`
	// IEExtensions * `optional`
}

func (ie *ImplicitFormat) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.DUFSlotformatIndex.Encode(w); err != nil {
		err = utils.WrapError("Encode DUFSlotformatIndex", err)
		return
	}
	return
}
func (ie *ImplicitFormat) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.DUFSlotformatIndex.Decode(r); err != nil {
		err = utils.WrapError("Read DUFSlotformatIndex", err)
		return
	}
	return
}
