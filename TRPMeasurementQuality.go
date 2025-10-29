package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TRPMeasurementQuality struct {
	TRPmeasurementQualityItem TRPMeasurementQualityItem `mandatory`
	// IEExtensions * `optional`
}

func (ie *TRPMeasurementQuality) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.TRPmeasurementQualityItem.Encode(w); err != nil {
		err = utils.WrapError("Encode TRPmeasurementQualityItem", err)
		return
	}
	return
}
func (ie *TRPMeasurementQuality) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.TRPmeasurementQualityItem.Decode(r); err != nil {
		err = utils.WrapError("Read TRPmeasurementQualityItem", err)
		return
	}
	return
}
