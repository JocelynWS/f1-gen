package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type EUTRASpecialSubframeInfo struct {
	SpecialSubframePatterns EUTRASpecialSubframePatterns `mandatory`
	CyclicPrefixDL          EUTRACyclicPrefixDL          `mandatory`
	CyclicPrefixUL          EUTRACyclicPrefixUL          `mandatory`
	// IEExtensions * `optional`
}

func (ie *EUTRASpecialSubframeInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.SpecialSubframePatterns.Encode(w); err != nil {
		err = utils.WrapError("Encode SpecialSubframePatterns", err)
		return
	}
	if err = ie.CyclicPrefixDL.Encode(w); err != nil {
		err = utils.WrapError("Encode CyclicPrefixDL", err)
		return
	}
	if err = ie.CyclicPrefixUL.Encode(w); err != nil {
		err = utils.WrapError("Encode CyclicPrefixUL", err)
		return
	}
	return
}
func (ie *EUTRASpecialSubframeInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.SpecialSubframePatterns.Decode(r); err != nil {
		err = utils.WrapError("Read SpecialSubframePatterns", err)
		return
	}
	if err = ie.CyclicPrefixDL.Decode(r); err != nil {
		err = utils.WrapError("Read CyclicPrefixDL", err)
		return
	}
	if err = ie.CyclicPrefixUL.Decode(r); err != nil {
		err = utils.WrapError("Read CyclicPrefixUL", err)
		return
	}
	return
}
