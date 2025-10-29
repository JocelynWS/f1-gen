package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PathlossReferenceInfo struct {
	PathlossReferenceSignal PathlossReferenceSignal `mandatory`
	// IEExtensions * `optional`
}

func (ie *PathlossReferenceInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.PathlossReferenceSignal.Encode(w); err != nil {
		err = utils.WrapError("Encode PathlossReferenceSignal", err)
		return
	}
	return
}
func (ie *PathlossReferenceInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.PathlossReferenceSignal.Decode(r); err != nil {
		err = utils.WrapError("Read PathlossReferenceSignal", err)
		return
	}
	return
}
