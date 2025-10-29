package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TNLCapacityIndicator struct {
	DLTNLOfferedCapacity   int64 `lb:1,ub:16777216,mandatory,valExt`
	DLTNLAvailableCapacity int64 `lb:0,ub:100,mandatory,valExt`
	ULTNLOfferedCapacity   int64 `lb:1,ub:16777216,mandatory,valExt`
	ULTNLAvailableCapacity int64 `lb:0,ub:100,mandatory,valExt`
	// IEExtensions * `optional`
}

func (ie *TNLCapacityIndicator) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_DLTNLOfferedCapacity := NewINTEGER(ie.DLTNLOfferedCapacity, aper.Constraint{Lb: 1, Ub: 16777216}, true)
	if err = tmp_DLTNLOfferedCapacity.Encode(w); err != nil {
		err = utils.WrapError("Encode DLTNLOfferedCapacity", err)
		return
	}
	tmp_DLTNLAvailableCapacity := NewINTEGER(ie.DLTNLAvailableCapacity, aper.Constraint{Lb: 0, Ub: 100}, true)
	if err = tmp_DLTNLAvailableCapacity.Encode(w); err != nil {
		err = utils.WrapError("Encode DLTNLAvailableCapacity", err)
		return
	}
	tmp_ULTNLOfferedCapacity := NewINTEGER(ie.ULTNLOfferedCapacity, aper.Constraint{Lb: 1, Ub: 16777216}, true)
	if err = tmp_ULTNLOfferedCapacity.Encode(w); err != nil {
		err = utils.WrapError("Encode ULTNLOfferedCapacity", err)
		return
	}
	tmp_ULTNLAvailableCapacity := NewINTEGER(ie.ULTNLAvailableCapacity, aper.Constraint{Lb: 0, Ub: 100}, true)
	if err = tmp_ULTNLAvailableCapacity.Encode(w); err != nil {
		err = utils.WrapError("Encode ULTNLAvailableCapacity", err)
		return
	}
	return
}
func (ie *TNLCapacityIndicator) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_DLTNLOfferedCapacity := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 16777216},
		ext: true,
	}
	if err = tmp_DLTNLOfferedCapacity.Decode(r); err != nil {
		err = utils.WrapError("Read DLTNLOfferedCapacity", err)
		return
	}
	ie.DLTNLOfferedCapacity = int64(tmp_DLTNLOfferedCapacity.Value)
	tmp_DLTNLAvailableCapacity := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 100},
		ext: true,
	}
	if err = tmp_DLTNLAvailableCapacity.Decode(r); err != nil {
		err = utils.WrapError("Read DLTNLAvailableCapacity", err)
		return
	}
	ie.DLTNLAvailableCapacity = int64(tmp_DLTNLAvailableCapacity.Value)
	tmp_ULTNLOfferedCapacity := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 16777216},
		ext: true,
	}
	if err = tmp_ULTNLOfferedCapacity.Decode(r); err != nil {
		err = utils.WrapError("Read ULTNLOfferedCapacity", err)
		return
	}
	ie.ULTNLOfferedCapacity = int64(tmp_ULTNLOfferedCapacity.Value)
	tmp_ULTNLAvailableCapacity := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 100},
		ext: true,
	}
	if err = tmp_ULTNLAvailableCapacity.Decode(r); err != nil {
		err = utils.WrapError("Read ULTNLAvailableCapacity", err)
		return
	}
	ie.ULTNLAvailableCapacity = int64(tmp_ULTNLAvailableCapacity.Value)
	return
}
