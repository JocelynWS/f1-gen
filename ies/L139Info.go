package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type L139Info struct {
	Msg1SCS           Msg1SCS `madatory,valueExt`
	RootSequenceIndex *int64  `optional,lb:0,ub:137`
	// IEExtension *L139InfoExtIEs `optional`
}

func (ie *L139Info) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.RootSequenceIndex != nil {
		aper.SetBit(optionals, 0)
	}
	if err = w.WriteBits(optionals, 2); err != nil {
		return
	}
	if err = ie.Msg1SCS.Encode(w); err != nil {
		err = utils.WrapError("Encode Msg1SCS", err)
		return
	}
	if ie.RootSequenceIndex != nil {
		tmp_RootSequenceIndex := NewINTEGER(*ie.RootSequenceIndex, aper.Constraint{Lb: 0, Ub: 137}, false)
		if err = tmp_RootSequenceIndex.Encode(w); err != nil {
			err = utils.WrapError("Encode RootSequenceIndex", err)
			return
		}
	}
	return
}

func (ie *L139Info) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.Msg1SCS.Decode(r); err != nil {
		err = utils.WrapError("Read Msg1SCS", err)
		return
	}
	if aper.IsBitSet(optionals, 0) {
		tmp_RootSequenceIndex := INTEGER{c: aper.Constraint{Lb: 0, Ub: 137}}
		if err = tmp_RootSequenceIndex.Decode(r); err != nil {
			err = utils.WrapError("Read RootSequenceIndex", err)
			return
		}
		val := int64(tmp_RootSequenceIndex.Value)
		ie.RootSequenceIndex = &val
	}
	return
}
