package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SCellFailedtoSetupModItem struct {
	SCellID NRCGI  `mandatory`
	Cause   *Cause `optional`
	// IEExtensions * `optional`
}

func (ie *SCellFailedtoSetupModItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.Cause != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.SCellID.Encode(w); err != nil {
		err = utils.WrapError("Encode SCellID", err)
		return
	}
	if ie.Cause != nil {
		if err = ie.Cause.Encode(w); err != nil {
			err = utils.WrapError("Encode Cause", err)
			return
		}
	}
	return
}
func (ie *SCellFailedtoSetupModItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.SCellID.Decode(r); err != nil {
		err = utils.WrapError("Read SCellID", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(Cause)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read Cause", err)
			return
		}
		ie.Cause = tmp
	}
	return
}
