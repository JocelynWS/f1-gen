package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SLDRBsToBeSetupItem struct {
	SLDRBID          int64            `lb:1,ub:512,mandatory`
	SLDRBInformation SLDRBInformation `mandatory`
	RLCMode          RLCMode          `mandatory`
	// IEExtensions * `optional`
}

func (ie *SLDRBsToBeSetupItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_SLDRBID := NewINTEGER(ie.SLDRBID, aper.Constraint{Lb: 1, Ub: 512}, false)
	if err = tmp_SLDRBID.Encode(w); err != nil {
		err = utils.WrapError("Encode SLDRBID", err)
		return
	}
	if err = ie.SLDRBInformation.Encode(w); err != nil {
		err = utils.WrapError("Encode SLDRBInformation", err)
		return
	}
	if err = ie.RLCMode.Encode(w); err != nil {
		err = utils.WrapError("Encode RLCMode", err)
		return
	}
	return
}
func (ie *SLDRBsToBeSetupItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_SLDRBID := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 512},
		ext: false,
	}
	if err = tmp_SLDRBID.Decode(r); err != nil {
		err = utils.WrapError("Read SLDRBID", err)
		return
	}
	ie.SLDRBID = int64(tmp_SLDRBID.Value)
	if err = ie.SLDRBInformation.Decode(r); err != nil {
		err = utils.WrapError("Read SLDRBInformation", err)
		return
	}
	if err = ie.RLCMode.Decode(r); err != nil {
		err = utils.WrapError("Read RLCMode", err)
		return
	}
	return
}
