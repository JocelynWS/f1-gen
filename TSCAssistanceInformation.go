package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TSCAssistanceInformation struct {
	Periodicity      int64  `lb:0,ub:640000,mandatory`
	BurstArrivalTime []byte `lb:0,ub:0,optional`
	// IEExtensions * `optional`
}

func (ie *TSCAssistanceInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.BurstArrivalTime != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_Periodicity := NewINTEGER(ie.Periodicity, aper.Constraint{Lb: 0, Ub: 640000}, false)
	if err = tmp_Periodicity.Encode(w); err != nil {
		err = utils.WrapError("Encode Periodicity", err)
		return
	}
	if ie.BurstArrivalTime != nil {
		tmp_BurstArrivalTime := NewOCTETSTRING(ie.BurstArrivalTime, aper.Constraint{Lb: 0, Ub: 0}, false)
		if err = tmp_BurstArrivalTime.Encode(w); err != nil {
			err = utils.WrapError("Encode BurstArrivalTime", err)
			return
		}
	}
	return
}
func (ie *TSCAssistanceInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_Periodicity := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 640000},
		ext: false,
	}
	if err = tmp_Periodicity.Decode(r); err != nil {
		err = utils.WrapError("Read Periodicity", err)
		return
	}
	ie.Periodicity = int64(tmp_Periodicity.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp_BurstArrivalTime := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp_BurstArrivalTime.Decode(r); err != nil {
			err = utils.WrapError("Read BurstArrivalTime", err)
			return
		}
		ie.BurstArrivalTime = tmp_BurstArrivalTime.Value
	}
	return
}
