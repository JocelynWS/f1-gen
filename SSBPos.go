package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SSBPos struct {
	PCINR    NRPCI     `mandatory`
	SsbIndex *SSBIndex `optional`
	// IEExtensions * `optional`
}

func (ie *SSBPos) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.SsbIndex != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.PCINR.Encode(w); err != nil {
		err = utils.WrapError("Encode PCINR", err)
		return
	}
	if ie.SsbIndex != nil {
		if err = ie.SsbIndex.Encode(w); err != nil {
			err = utils.WrapError("Encode SsbIndex", err)
			return
		}
	}
	return
}
func (ie *SSBPos) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.PCINR.Decode(r); err != nil {
		err = utils.WrapError("Read PCINR", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(SSBIndex)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read SsbIndex", err)
			return
		}
		ie.SsbIndex = tmp
	}
	return
}
