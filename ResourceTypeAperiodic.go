package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ResourceTypeAperiodic struct {
	AperiodicResourceType AperiodicResourceType `madatory,valExt`
	// IEExtensions *ResourceTypeAperiodicExtIEs `optional`
}

func (ie *ResourceTypeAperiodic) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if err = w.WriteBits(optionals, 1); err != nil {
		return
	}
	if err = ie.AperiodicResourceType.Encode(w); err != nil {
		err = utils.WrapError("Encode AperiodicResourceType", err)
		return
	}
	return
}

func (ie *ResourceTypeAperiodic) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.AperiodicResourceType.Decode(r); err != nil {
		err = utils.WrapError("Read AperiodicResourceType", err)
		return
	}
	return
}
