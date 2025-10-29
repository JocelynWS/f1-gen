package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PeriodicityListItem struct {
	PeriodicitySRS PeriodicitySRS `mandatory`
	// IEExtensions * `optional`
}

func (ie *PeriodicityListItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.PeriodicitySRS.Encode(w); err != nil {
		err = utils.WrapError("Encode PeriodicitySRS", err)
		return
	}
	return
}
func (ie *PeriodicityListItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.PeriodicitySRS.Decode(r); err != nil {
		err = utils.WrapError("Read PeriodicitySRS", err)
		return
	}
	return
}
