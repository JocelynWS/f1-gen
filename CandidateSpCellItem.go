package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type CandidateSpCellItem struct {
	CandidateSpCellID NRCGI `mandatory`
	// IEExtensions * `optional`
}

func (ie *CandidateSpCellItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.CandidateSpCellID.Encode(w); err != nil {
		err = utils.WrapError("Encode CandidateSpCellID", err)
		return
	}
	return
}
func (ie *CandidateSpCellItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.CandidateSpCellID.Decode(r); err != nil {
		err = utils.WrapError("Read CandidateSpCellID", err)
		return
	}
	return
}
