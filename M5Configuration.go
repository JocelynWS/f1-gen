package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type M5Configuration struct {
	M5period     M5period     `mandatory`
	M5LinksToLog M5LinksToLog `mandatory`
	// IEExtensions * `optional`
}

func (ie *M5Configuration) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.M5period.Encode(w); err != nil {
		err = utils.WrapError("Encode M5period", err)
		return
	}
	if err = ie.M5LinksToLog.Encode(w); err != nil {
		err = utils.WrapError("Encode M5LinksToLog", err)
		return
	}
	return
}
func (ie *M5Configuration) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.M5period.Decode(r); err != nil {
		err = utils.WrapError("Read M5period", err)
		return
	}
	if err = ie.M5LinksToLog.Decode(r); err != nil {
		err = utils.WrapError("Read M5LinksToLog", err)
		return
	}
	return
}
