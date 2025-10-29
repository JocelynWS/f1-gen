package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type IABDUCellResourceConfigurationFDDInfo struct {
	GNBDUCellResourceConfigurationFDDUL GNBDUCellResourceConfiguration `mandatory`
	GNBDUCellResourceConfigurationFDDDL GNBDUCellResourceConfiguration `mandatory`
	// IEExtensions * `optional`
}

func (ie *IABDUCellResourceConfigurationFDDInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.GNBDUCellResourceConfigurationFDDUL.Encode(w); err != nil {
		err = utils.WrapError("Encode GNBDUCellResourceConfigurationFDDUL", err)
		return
	}
	if err = ie.GNBDUCellResourceConfigurationFDDDL.Encode(w); err != nil {
		err = utils.WrapError("Encode GNBDUCellResourceConfigurationFDDDL", err)
		return
	}
	return
}
func (ie *IABDUCellResourceConfigurationFDDInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.GNBDUCellResourceConfigurationFDDUL.Decode(r); err != nil {
		err = utils.WrapError("Read GNBDUCellResourceConfigurationFDDUL", err)
		return
	}
	if err = ie.GNBDUCellResourceConfigurationFDDDL.Decode(r); err != nil {
		err = utils.WrapError("Read GNBDUCellResourceConfigurationFDDDL", err)
		return
	}
	return
}
