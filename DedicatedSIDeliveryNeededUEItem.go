package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DedicatedSIDeliveryNeededUEItem struct {
	GNBCUUEF1APID int64 `lb:0,ub:4294967295,mandatory`
	NRCGI         NRCGI `mandatory`
	// IEExtensions * `optional`
}

func (ie *DedicatedSIDeliveryNeededUEItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_GNBCUUEF1APID := NewINTEGER(ie.GNBCUUEF1APID, aper.Constraint{Lb: 0, Ub: 4294967295}, false)
	if err = tmp_GNBCUUEF1APID.Encode(w); err != nil {
		err = utils.WrapError("Encode GNBCUUEF1APID", err)
		return
	}
	if err = ie.NRCGI.Encode(w); err != nil {
		err = utils.WrapError("Encode NRCGI", err)
		return
	}
	return
}
func (ie *DedicatedSIDeliveryNeededUEItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_GNBCUUEF1APID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4294967295},
		ext: false,
	}
	if err = tmp_GNBCUUEF1APID.Decode(r); err != nil {
		err = utils.WrapError("Read GNBCUUEF1APID", err)
		return
	}
	ie.GNBCUUEF1APID = int64(tmp_GNBCUUEF1APID.Value)
	if err = ie.NRCGI.Decode(r); err != nil {
		err = utils.WrapError("Read NRCGI", err)
		return
	}
	return
}
