package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ECIDMeasurementQuantitiesItem struct {
	ECIDmeasurementQuantitiesValue ECIDMeasurementQuantitiesValue `mandatory`
	// IEExtensions * `optional`
}

func (ie *ECIDMeasurementQuantitiesItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.ECIDmeasurementQuantitiesValue.Encode(w); err != nil {
		err = utils.WrapError("Encode ECIDmeasurementQuantitiesValue", err)
		return
	}
	return
}
func (ie *ECIDMeasurementQuantitiesItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.ECIDmeasurementQuantitiesValue.Decode(r); err != nil {
		err = utils.WrapError("Read ECIDmeasurementQuantitiesValue", err)
		return
	}
	return
}
