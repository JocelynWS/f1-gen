package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AperiodicSRS struct {
	Aperiodic          Aperiodic           `madatory,valueExt`
	SRSResourceTrigger *SRSResourceTrigger `optional`
	// IEExtensions *AperiodicSRSExtIEs `optional`
}

func (ie *AperiodicSRS) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.SRSResourceTrigger != nil {
		aper.SetBit(optionals, 0)
	}
	if err = w.WriteBits(optionals, 2); err != nil {
		return
	}

	if err = ie.Aperiodic.Encode(w); err != nil {
		err = utils.WrapError("Encode Aperiodic", err)
		return
	}

	if ie.SRSResourceTrigger != nil {
		if err = ie.SRSResourceTrigger.Encode(w); err != nil {
			err = utils.WrapError("Encode SRSResourceTrigger", err)
			return
		}
	}

	return
}

func (ie *AperiodicSRS) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}

	if err = ie.Aperiodic.Decode(r); err != nil {
		err = utils.WrapError("Read Aperiodic", err)
		return
	}

	if aper.IsBitSet(optionals, 0) {
		tmp := new(SRSResourceTrigger)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read SRSResourceTrigger", err)
			return
		}
		ie.SRSResourceTrigger = tmp
	}

	return
}
