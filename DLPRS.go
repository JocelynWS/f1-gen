package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DLPRS struct {
	Prsid              int64            `lb:0,ub:255,mandatory`
	DlPRSResourceSetID PRSResourceSetID `mandatory`
	DlPRSResourceID    *int64           `lb:0,ub:63,optional`
	// IEExtensions * `optional`
}

func (ie *DLPRS) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.DlPRSResourceID != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_Prsid := NewINTEGER(ie.Prsid, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_Prsid.Encode(w); err != nil {
		err = utils.WrapError("Encode Prsid", err)
		return
	}
	if err = ie.DlPRSResourceSetID.Encode(w); err != nil {
		err = utils.WrapError("Encode DlPRSResourceSetID", err)
		return
	}
	if ie.DlPRSResourceID != nil {
		tmp_DlPRSResourceID := NewINTEGER(*ie.DlPRSResourceID, aper.Constraint{Lb: 0, Ub: 63}, false)
		if err = tmp_DlPRSResourceID.Encode(w); err != nil {
			err = utils.WrapError("Encode DlPRSResourceID", err)
			return
		}
	}
	return
}
func (ie *DLPRS) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_Prsid := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 255},
		ext: false,
	}
	if err = tmp_Prsid.Decode(r); err != nil {
		err = utils.WrapError("Read Prsid", err)
		return
	}
	ie.Prsid = int64(tmp_Prsid.Value)
	if err = ie.DlPRSResourceSetID.Decode(r); err != nil {
		err = utils.WrapError("Read DlPRSResourceSetID", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_DlPRSResourceID := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 63},
			ext: false,
		}
		if err = tmp_DlPRSResourceID.Decode(r); err != nil {
			err = utils.WrapError("Read DlPRSResourceID", err)
			return
		}
		ie.DlPRSResourceID = (*int64)(&tmp_DlPRSResourceID.Value)
	}
	return
}
