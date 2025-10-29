package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type MeasurementBeamInfo struct {
	PRSResourceID    *int64            `lb:0,ub:63,optional`
	PRSResourceSetID *PRSResourceSetID `optional`
	SSBIndex         *SSBIndex         `optional`
	// IEExtensions * `optional`
}

func (ie *MeasurementBeamInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.PRSResourceID != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.PRSResourceSetID != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.SSBIndex != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)
	if ie.PRSResourceID != nil {
		tmp_PRSResourceID := NewINTEGER(*ie.PRSResourceID, aper.Constraint{Lb: 0, Ub: 63}, false)
		if err = tmp_PRSResourceID.Encode(w); err != nil {
			err = utils.WrapError("Encode PRSResourceID", err)
			return
		}
	}
	if ie.PRSResourceSetID != nil {
		if err = ie.PRSResourceSetID.Encode(w); err != nil {
			err = utils.WrapError("Encode PRSResourceSetID", err)
			return
		}
	}
	if ie.SSBIndex != nil {
		if err = ie.SSBIndex.Encode(w); err != nil {
			err = utils.WrapError("Encode SSBIndex", err)
			return
		}
	}
	return
}
func (ie *MeasurementBeamInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_PRSResourceID := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 63},
			ext: false,
		}
		if err = tmp_PRSResourceID.Decode(r); err != nil {
			err = utils.WrapError("Read PRSResourceID", err)
			return
		}
		ie.PRSResourceID = (*int64)(&tmp_PRSResourceID.Value)
	}
	if aper.IsBitSet(optionals, 2) {
		tmp := new(PRSResourceSetID)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read PRSResourceSetID", err)
			return
		}
		ie.PRSResourceSetID = tmp
	}
	if aper.IsBitSet(optionals, 3) {
		tmp := new(SSBIndex)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read SSBIndex", err)
			return
		}
		ie.SSBIndex = tmp
	}
	return
}
