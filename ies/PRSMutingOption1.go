package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PRSMutingOption1 struct {
	MutingPattern             DLPRSMutingPattern        `madatory`
	MutingBitRepetitionFactor MutingBitRepetitionFactor `madatory,valExt`
	// IEExtensions *PRSMutingOption1ExtIEs `optional`
}

func (ie *PRSMutingOption1) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if err = w.WriteBits(optionals, 1); err != nil {
		return
	}
	if err = ie.MutingPattern.Encode(w); err != nil {
		err = utils.WrapError("Encode MutingPattern", err)
		return
	}
	if err = ie.MutingBitRepetitionFactor.Encode(w); err != nil {
		err = utils.WrapError("Encode MutingBitRepetitionFactor", err)
		return
	}
	return
}

func (ie *PRSMutingOption1) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.MutingPattern.Decode(r); err != nil {
		err = utils.WrapError("Read MutingPattern", err)
		return
	}
	if err = ie.MutingBitRepetitionFactor.Decode(r); err != nil {
		err = utils.WrapError("Read MutingBitRepetitionFactor", err)
		return
	}
	return
}
