package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UEAssociatedLogicalF1ConnectionItem struct {
	GNBCUUEF1APID *int64 `lb:0,ub:4294967295,optional`
	GNBDUUEF1APID *int64 `lb:0,ub:4294967295,optional`
	// IEExtensions * `optional`
}

func (ie *UEAssociatedLogicalF1ConnectionItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.GNBCUUEF1APID != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.GNBDUUEF1APID != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if ie.GNBCUUEF1APID != nil {
		tmp_GNBCUUEF1APID := NewINTEGER(*ie.GNBCUUEF1APID, aper.Constraint{Lb: 0, Ub: 4294967295}, false)
		if err = tmp_GNBCUUEF1APID.Encode(w); err != nil {
			err = utils.WrapError("Encode GNBCUUEF1APID", err)
			return
		}
	}
	if ie.GNBDUUEF1APID != nil {
		tmp_GNBDUUEF1APID := NewINTEGER(*ie.GNBDUUEF1APID, aper.Constraint{Lb: 0, Ub: 4294967295}, false)
		if err = tmp_GNBDUUEF1APID.Encode(w); err != nil {
			err = utils.WrapError("Encode GNBDUUEF1APID", err)
			return
		}
	}
	return
}
func (ie *UEAssociatedLogicalF1ConnectionItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_GNBCUUEF1APID := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4294967295},
			ext: false,
		}
		if err = tmp_GNBCUUEF1APID.Decode(r); err != nil {
			err = utils.WrapError("Read GNBCUUEF1APID", err)
			return
		}
		ie.GNBCUUEF1APID = (*int64)(&tmp_GNBCUUEF1APID.Value)
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_GNBDUUEF1APID := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4294967295},
			ext: false,
		}
		if err = tmp_GNBDUUEF1APID.Decode(r); err != nil {
			err = utils.WrapError("Read GNBDUUEF1APID", err)
			return
		}
		ie.GNBDUUEF1APID = (*int64)(&tmp_GNBDUUEF1APID.Value)
	}
	return
}
