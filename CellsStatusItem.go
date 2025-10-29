package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type CellsStatusItem struct {
	NRCGI         NRCGI         `mandatory`
	ServiceStatus ServiceStatus `mandatory`
	// IEExtensions * `optional`
}

func (ie *CellsStatusItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.NRCGI.Encode(w); err != nil {
		err = utils.WrapError("Encode NRCGI", err)
		return
	}
	if err = ie.ServiceStatus.Encode(w); err != nil {
		err = utils.WrapError("Encode ServiceStatus", err)
		return
	}
	return
}
func (ie *CellsStatusItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.NRCGI.Decode(r); err != nil {
		err = utils.WrapError("Read NRCGI", err)
		return
	}
	if err = ie.ServiceStatus.Decode(r); err != nil {
		err = utils.WrapError("Read ServiceStatus", err)
		return
	}
	return
}
