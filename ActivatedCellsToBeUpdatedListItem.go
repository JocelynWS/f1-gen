package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ActivatedCellsToBeUpdatedListItem struct {
	NRCGI                                  NRCGI
	IABDUCellResourceConfigurationModeInfo IABDUCellResourceConfigurationModeInfo
	// iE-Extensions ProtocolExtensionContainer { { Activated-Cells-to-be-Updated-List-Item-ExtIEs} } OPTIONAL
}

func (ie *ActivatedCellsToBeUpdatedListItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	if err = ie.NRCGI.Encode(w); err != nil {
		return utils.WrapError("Encode NRCGI", err)
	}
	if err = ie.IABDUCellResourceConfigurationModeInfo.Encode(w); err != nil {
		return utils.WrapError("Encode IABDUCellResourceConfigurationModeInfo", err)
	}

	return
}

func (ie *ActivatedCellsToBeUpdatedListItem) Decode(r *aper.AperReader) (err error) {
	var ext bool
	if ext, err = r.ReadBool(); err != nil {
		return
	}

	if err = ie.NRCGI.Decode(r); err != nil {
		return utils.WrapError("Decode NRCGI", err)
	}
	if err = ie.IABDUCellResourceConfigurationModeInfo.Decode(r); err != nil {
		return utils.WrapError("Decode IABDUCellResourceConfigurationModeInfo", err)
	}

	if ext {
	}

	return
}
