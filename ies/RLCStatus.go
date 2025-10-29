package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type RLCStatus struct {
	ReestablishmentIndication ReestablishmentIndication `mandatory`
	// IEExtensions * `optional`
}

func (ie *RLCStatus) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.ReestablishmentIndication.Encode(w); err != nil {
		err = utils.WrapError("Encode ReestablishmentIndication", err)
		return
	}
	return
}
func (ie *RLCStatus) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.ReestablishmentIndication.Decode(r); err != nil {
		err = utils.WrapError("Read ReestablishmentIndication", err)
		return
	}
	return
}
