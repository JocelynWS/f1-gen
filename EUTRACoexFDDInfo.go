package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type EUTRACoexFDDInfo struct {
	ULEARFCN                *int64                      `lb:0,ub:262143,optional`
	DLEARFCN                int64                       `lb:0,ub:262143,mandatory`
	ULTransmissionBandwidth *EUTRATransmissionBandwidth `optional`
	DLTransmissionBandwidth EUTRATransmissionBandwidth  `mandatory`
	// IEExtensions * `optional`
}

func (ie *EUTRACoexFDDInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.ULEARFCN != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.ULTransmissionBandwidth != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if ie.ULEARFCN != nil {
		tmp_ULEARFCN := NewINTEGER(*ie.ULEARFCN, aper.Constraint{Lb: 0, Ub: 262143}, false)
		if err = tmp_ULEARFCN.Encode(w); err != nil {
			err = utils.WrapError("Encode ULEARFCN", err)
			return
		}
	}
	tmp_DLEARFCN := NewINTEGER(ie.DLEARFCN, aper.Constraint{Lb: 0, Ub: 262143}, false)
	if err = tmp_DLEARFCN.Encode(w); err != nil {
		err = utils.WrapError("Encode DLEARFCN", err)
		return
	}
	if ie.ULTransmissionBandwidth != nil {
		if err = ie.ULTransmissionBandwidth.Encode(w); err != nil {
			err = utils.WrapError("Encode ULTransmissionBandwidth", err)
			return
		}
	}
	if err = ie.DLTransmissionBandwidth.Encode(w); err != nil {
		err = utils.WrapError("Encode DLTransmissionBandwidth", err)
		return
	}
	return
}
func (ie *EUTRACoexFDDInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_ULEARFCN := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 262143},
			ext: false,
		}
		if err = tmp_ULEARFCN.Decode(r); err != nil {
			err = utils.WrapError("Read ULEARFCN", err)
			return
		}
		ie.ULEARFCN = (*int64)(&tmp_ULEARFCN.Value)
	}
	tmp_DLEARFCN := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 262143},
		ext: false,
	}
	if err = tmp_DLEARFCN.Decode(r); err != nil {
		err = utils.WrapError("Read DLEARFCN", err)
		return
	}
	ie.DLEARFCN = int64(tmp_DLEARFCN.Value)
	if aper.IsBitSet(optionals, 2) {
		tmp := new(EUTRATransmissionBandwidth)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read ULTransmissionBandwidth", err)
			return
		}
		ie.ULTransmissionBandwidth = tmp
	}
	if err = ie.DLTransmissionBandwidth.Decode(r); err != nil {
		err = utils.WrapError("Read DLTransmissionBandwidth", err)
		return
	}
	return
}
