package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DLPRSResourceARP struct {
	DlPRSResourceID          int64                    `lb:0,ub:63,mandatory`
	DLPRSResourceARPLocation DLPRSResourceARPLocation `mandatory`
	// IEExtensions * `optional`
}

func (ie *DLPRSResourceARP) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_DlPRSResourceID := NewINTEGER(ie.DlPRSResourceID, aper.Constraint{Lb: 0, Ub: 63}, false)
	if err = tmp_DlPRSResourceID.Encode(w); err != nil {
		err = utils.WrapError("Encode DlPRSResourceID", err)
		return
	}
	if err = ie.DLPRSResourceARPLocation.Encode(w); err != nil {
		err = utils.WrapError("Encode DLPRSResourceARPLocation", err)
		return
	}
	return
}
func (ie *DLPRSResourceARP) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_DlPRSResourceID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 63},
		ext: false,
	}
	if err = tmp_DlPRSResourceID.Decode(r); err != nil {
		err = utils.WrapError("Read DlPRSResourceID", err)
		return
	}
	ie.DlPRSResourceID = int64(tmp_DlPRSResourceID.Value)
	if err = ie.DLPRSResourceARPLocation.Decode(r); err != nil {
		err = utils.WrapError("Read DLPRSResourceARPLocation", err)
		return
	}
	return
}
