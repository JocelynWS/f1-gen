package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type M6Configuration struct {
	M6reportInterval M6reportInterval `mandatory`
	M6LinksToLog     M6LinksToLog     `mandatory`
	// IEExtensions * `optional`
}

func (ie *M6Configuration) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.M6reportInterval.Encode(w); err != nil {
		err = utils.WrapError("Encode M6reportInterval", err)
		return
	}
	if err = ie.M6LinksToLog.Encode(w); err != nil {
		err = utils.WrapError("Encode M6LinksToLog", err)
		return
	}
	return
}
func (ie *M6Configuration) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.M6reportInterval.Decode(r); err != nil {
		err = utils.WrapError("Read M6reportInterval", err)
		return
	}
	if err = ie.M6LinksToLog.Decode(r); err != nil {
		err = utils.WrapError("Read M6LinksToLog", err)
		return
	}
	return
}
