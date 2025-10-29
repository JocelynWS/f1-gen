package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PRSResourceQCLInfo struct {
	QCLSourceSSBIndex *int64                       `lb:0,ub:63,optional`
	QCLSourcePRSInfo  *PRSResourceQCLSourcePRSInfo `optional`
	// IEExtensions *ProtocolExtensionContainerPRSResourceQCLInfoExtIEs `optional`
}

func (ie *PRSResourceQCLInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.QCLSourceSSBIndex != nil {
		aper.SetBit(optionals, 0)
	}
	if ie.QCLSourcePRSInfo != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)

	if ie.QCLSourceSSBIndex != nil {
		tmpSSB := INTEGER{
			Value: aper.Integer(*ie.QCLSourceSSBIndex),
			c:     aper.Constraint{Lb: 0, Ub: 63},
			ext:   false,
		}
		if err = tmpSSB.Encode(w); err != nil {
			err = utils.WrapError("Encode QCLSourceSSBIndex", err)
			return
		}
	}

	if ie.QCLSourcePRSInfo != nil {
		if err = ie.QCLSourcePRSInfo.Encode(w); err != nil {
			err = utils.WrapError("Encode QCLSourcePRSInfo", err)
			return
		}
	}

	return
}

func (ie *PRSResourceQCLInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}

	if aper.IsBitSet(optionals, 0) {
		tmpSSB := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 63},
			ext: false,
		}
		if err = tmpSSB.Decode(r); err != nil {
			err = utils.WrapError("Read QCLSourceSSBIndex", err)
			return
		}
		ie.QCLSourceSSBIndex = (*int64)(&tmpSSB.Value)
	}

	if aper.IsBitSet(optionals, 1) {
		ie.QCLSourcePRSInfo = new(PRSResourceQCLSourcePRSInfo)
		if err = ie.QCLSourcePRSInfo.Decode(r); err != nil {
			err = utils.WrapError("Read QCLSourcePRSInfo", err)
			return
		}
	}

	return
}
