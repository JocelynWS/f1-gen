package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type IABDUCellResourceConfigurationTDDInfo struct {
	GNBDUCellResourcConfigurationTDD GNBDUCellResourceConfiguration `mandatory`
	// IEExtensions * `optional`
}

func (ie *IABDUCellResourceConfigurationTDDInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.GNBDUCellResourcConfigurationTDD.Encode(w); err != nil {
		err = utils.WrapError("Encode GNBDUCellResourcConfigurationTDD", err)
		return
	}
	return
}
func (ie *IABDUCellResourceConfigurationTDDInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.GNBDUCellResourcConfigurationTDD.Decode(r); err != nil {
		err = utils.WrapError("Read GNBDUCellResourcConfigurationTDD", err)
		return
	}
	return
}
