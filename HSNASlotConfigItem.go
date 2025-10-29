package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type HSNASlotConfigItem struct {
	HSNADownlink *HSNADownlink `optional`
	HSNAUplink   *HSNAUplink   `optional`
	HSNAFlexible *HSNAFlexible `optional`
	// IEExtensions * `optional`
}

func (ie *HSNASlotConfigItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.HSNADownlink != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.HSNAUplink != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.HSNAFlexible != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)
	if ie.HSNADownlink != nil {
		if err = ie.HSNADownlink.Encode(w); err != nil {
			err = utils.WrapError("Encode HSNADownlink", err)
			return
		}
	}
	if ie.HSNAUplink != nil {
		if err = ie.HSNAUplink.Encode(w); err != nil {
			err = utils.WrapError("Encode HSNAUplink", err)
			return
		}
	}
	if ie.HSNAFlexible != nil {
		if err = ie.HSNAFlexible.Encode(w); err != nil {
			err = utils.WrapError("Encode HSNAFlexible", err)
			return
		}
	}
	return
}
func (ie *HSNASlotConfigItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(HSNADownlink)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read HSNADownlink", err)
			return
		}
		ie.HSNADownlink = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp := new(HSNAUplink)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read HSNAUplink", err)
			return
		}
		ie.HSNAUplink = tmp
	}
	if aper.IsBitSet(optionals, 3) {
		tmp := new(HSNAFlexible)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read HSNAFlexible", err)
			return
		}
		ie.HSNAFlexible = tmp
	}
	return
}
