package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type EUTRANQoS struct {
	QCI                            int64                          `lb:0,ub:255,mandatory`
	AllocationAndRetentionPriority AllocationAndRetentionPriority `mandatory`
	GbrQosInformation              *GBRQosInformation             `optional`
	// IEExtensions * `optional`
}

func (ie *EUTRANQoS) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.GbrQosInformation != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_QCI := NewINTEGER(ie.QCI, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_QCI.Encode(w); err != nil {
		err = utils.WrapError("Encode QCI", err)
		return
	}
	if err = ie.AllocationAndRetentionPriority.Encode(w); err != nil {
		err = utils.WrapError("Encode AllocationAndRetentionPriority", err)
		return
	}
	if ie.GbrQosInformation != nil {
		if err = ie.GbrQosInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode GbrQosInformation", err)
			return
		}
	}
	return
}
func (ie *EUTRANQoS) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_QCI := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 255},
		ext: false,
	}
	if err = tmp_QCI.Decode(r); err != nil {
		err = utils.WrapError("Read QCI", err)
		return
	}
	ie.QCI = int64(tmp_QCI.Value)
	if err = ie.AllocationAndRetentionPriority.Decode(r); err != nil {
		err = utils.WrapError("Read AllocationAndRetentionPriority", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(GBRQosInformation)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read GbrQosInformation", err)
			return
		}
		ie.GbrQosInformation = tmp
	}
	return
}
