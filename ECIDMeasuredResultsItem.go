package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ECIDMeasuredResultsItem struct {
	ECIDMeasuredResultsValue ECIDMeasuredResultsValue `mandatory`
	// IEExtensions * `optional`
}

func (ie *ECIDMeasuredResultsItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.ECIDMeasuredResultsValue.Encode(w); err != nil {
		err = utils.WrapError("Encode ECIDMeasuredResultsValue", err)
		return
	}
	return
}
func (ie *ECIDMeasuredResultsItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.ECIDMeasuredResultsValue.Decode(r); err != nil {
		err = utils.WrapError("Read ECIDMeasuredResultsValue", err)
		return
	}
	return
}
