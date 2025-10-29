package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type RANUEPagingIdentity struct {
	IRNTI aper.BitString `lb:40,ub:40,mandatory`
	// IEExtensions * `optional`
}

func (ie *RANUEPagingIdentity) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)

	tmp_IRNTI := NewBITSTRING(ie.IRNTI, aper.Constraint{Lb: 40, Ub: 40}, false)
	if err = tmp_IRNTI.Encode(w); err != nil {
		err = utils.WrapError("Encode IRNTI", err)
		return
	}

	return
}

func (ie *RANUEPagingIdentity) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	if _, err = r.ReadBits(1); err != nil {
		return
	}

	tmp_IRNTI := BITSTRING{
		c:   aper.Constraint{Lb: 40, Ub: 40},
		ext: false,
	}

	if err = tmp_IRNTI.Decode(r); err != nil {
		err = utils.WrapError("Read IRNTI", err)
		return
	}

	ie.IRNTI = aper.BitString{
		Bytes:   tmp_IRNTI.Value.Bytes,
		NumBits: tmp_IRNTI.Value.NumBits,
	}

	return
}
