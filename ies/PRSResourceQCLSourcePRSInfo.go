package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PRSResourceQCLSourcePRSInfo struct {
	QCLSourcePRSResourceSetID int64  `lb:0,ub:7,mandatory`
	QCLSourcePRSResourceID    *int64 `lb:0,ub:63,optional`
	// IEExtensions	*ProtocolExtensionContainerPRSResourceQCLSourcePRSInfoExtIEs	`optional`
}

func (ie *PRSResourceQCLSourcePRSInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.QCLSourcePRSResourceID != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_QCLSourcePRSResourceSetID := NewINTEGER(ie.QCLSourcePRSResourceSetID, aper.Constraint{Lb: 0, Ub: 7}, false)
	if err = tmp_QCLSourcePRSResourceSetID.Encode(w); err != nil {
		err = utils.WrapError("Encode QCLSourcePRSResourceSetID", err)
		return
	}
	if ie.QCLSourcePRSResourceID != nil {
		tmp_QCLSourcePRSResourceID := NewINTEGER(*ie.QCLSourcePRSResourceID, aper.Constraint{Lb: 0, Ub: 63}, false)
		if err = tmp_QCLSourcePRSResourceID.Encode(w); err != nil {
			err = utils.WrapError("Encode QCLSourcePRSResourceID", err)
			return
		}
	}
	return
}

func (ie *PRSResourceQCLSourcePRSInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_QCLSourcePRSResourceSetID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 7},
		ext: false,
	}
	if err = tmp_QCLSourcePRSResourceSetID.Decode(r); err != nil {
		err = utils.WrapError("Read QCLSourcePRSResourceSetID", err)
		return
	}
	ie.QCLSourcePRSResourceSetID = int64(tmp_QCLSourcePRSResourceSetID.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp_QCLSourcePRSResourceID := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 63},
			ext: false,
		}
		if err = tmp_QCLSourcePRSResourceID.Decode(r); err != nil {
			err = utils.WrapError("Read QCLSourcePRSResourceID", err)
			return
		}
		ie.QCLSourcePRSResourceID = (*int64)(&tmp_QCLSourcePRSResourceID.Value)
	}
	return
}
