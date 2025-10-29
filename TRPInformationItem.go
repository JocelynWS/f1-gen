package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TRPInformationItem struct {
	TRPInformation TRPInformation `mandatory`
	// IEExtensions * `optional`
}

func (ie *TRPInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.TRPInformation.Encode(w); err != nil {
		err = utils.WrapError("Encode TRPInformation", err)
		return
	}
	return
}
func (ie *TRPInformationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.TRPInformation.Decode(r); err != nil {
		err = utils.WrapError("Read TRPInformation", err)
		return
	}
	return
}
