package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type CompositeAvailableCapacityGroup struct {
	CompositeAvailableCapacityDownlink CompositeAvailableCapacity `mandatory`
	CompositeAvailableCapacityUplink   CompositeAvailableCapacity `mandatory`
	// IEExtensions * `optional`
}

func (ie *CompositeAvailableCapacityGroup) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.CompositeAvailableCapacityDownlink.Encode(w); err != nil {
		err = utils.WrapError("Encode CompositeAvailableCapacityDownlink", err)
		return
	}
	if err = ie.CompositeAvailableCapacityUplink.Encode(w); err != nil {
		err = utils.WrapError("Encode CompositeAvailableCapacityUplink", err)
		return
	}
	return
}
func (ie *CompositeAvailableCapacityGroup) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.CompositeAvailableCapacityDownlink.Decode(r); err != nil {
		err = utils.WrapError("Read CompositeAvailableCapacityDownlink", err)
		return
	}
	if err = ie.CompositeAvailableCapacityUplink.Decode(r); err != nil {
		err = utils.WrapError("Read CompositeAvailableCapacityUplink", err)
		return
	}
	return
}
