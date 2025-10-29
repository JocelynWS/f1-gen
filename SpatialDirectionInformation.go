package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SpatialDirectionInformation struct {
	NRPRSBeamInformation NRPRSBeamInformation `mandatory`
	// IEExtensions * `optional`
}

func (ie *SpatialDirectionInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.NRPRSBeamInformation.Encode(w); err != nil {
		err = utils.WrapError("Encode NRPRSBeamInformation", err)
		return
	}
	return
}
func (ie *SpatialDirectionInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.NRPRSBeamInformation.Decode(r); err != nil {
		err = utils.WrapError("Read NRPRSBeamInformation", err)
		return
	}
	return
}
