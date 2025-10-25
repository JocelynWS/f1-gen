package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type L839Info struct {
	RootSequenceIndex   int64               `lb:0,ub:837,mandatory`
	RestrictedSetConfig RestrictedSetConfig `madatory,valExt`
	// IEExtension *L839InfoExtIEs `optional`
}

func (ie *L839Info) Encode(w *aper.AperWriter) (err error) {
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)

	tmp_RootSequenceIndex := NewINTEGER(ie.RootSequenceIndex, aper.Constraint{Lb: 0, Ub: 837}, false)
	if err = tmp_RootSequenceIndex.Encode(w); err != nil {
		return utils.WrapError("Encode RootSequenceIndex", err)
	}

	tmp_RestrictedSetConfig := NewENUMERATED(int64(ie.RestrictedSetConfig.Value), aper.Constraint{Lb: 0, Ub: 2}, false)
	if err = tmp_RestrictedSetConfig.Encode(w); err != nil {
		return utils.WrapError("Encode RestrictedSetConfig", err)
	}

	return
}

func (ie *L839Info) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBits(1); err != nil {
		return
	}

	tmp_RootSequenceIndex := INTEGER{c: aper.Constraint{Lb: 0, Ub: 837}, ext: false}
	if err = tmp_RootSequenceIndex.Decode(r); err != nil {
		return utils.WrapError("Read RootSequenceIndex", err)
	}
	ie.RootSequenceIndex = int64(tmp_RootSequenceIndex.Value)

	tmp_RestrictedSetConfig := ENUMERATED{c: aper.Constraint{Lb: 0, Ub: 2}, ext: false}
	if err = tmp_RestrictedSetConfig.Decode(r); err != nil {
		return utils.WrapError("Read RestrictedSetConfig", err)
	}
	ie.RestrictedSetConfig = RestrictedSetConfig{Value: tmp_RestrictedSetConfig.Value}
	return
}
