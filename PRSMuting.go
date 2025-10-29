package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PRSMuting struct {
	PRSMutingOption1 PRSMutingOption1 `mandatory`
	PRSMutingOption2 PRSMutingOption2 `mandatory`
	// IEExtensions * `optional`
}

func (ie *PRSMuting) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.PRSMutingOption1.Encode(w); err != nil {
		err = utils.WrapError("Encode PRSMutingOption1", err)
		return
	}
	if err = ie.PRSMutingOption2.Encode(w); err != nil {
		err = utils.WrapError("Encode PRSMutingOption2", err)
		return
	}
	return
}
func (ie *PRSMuting) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.PRSMutingOption1.Decode(r); err != nil {
		err = utils.WrapError("Read PRSMutingOption1", err)
		return
	}
	if err = ie.PRSMutingOption2.Decode(r); err != nil {
		err = utils.WrapError("Read PRSMutingOption2", err)
		return
	}
	return
}
