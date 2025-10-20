package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

// PosResourceSetTypeAP struct
type PosResourceSetTypeAP struct {
	SRSResourceTriggerList int64 `lb:1,ub:3,mandatory`
	// IEExtensions * `optional`
}

// Encode APER
func (ie *PosResourceSetTypeAP) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0} // currently no optional
	w.WriteBits(optionals, 1)

	tmp := NewINTEGER(ie.SRSResourceTriggerList, aper.Constraint{Lb: 1, Ub: 3}, false)
	if err = tmp.Encode(w); err != nil {
		err = utils.WrapError("Encode SRSResourceTriggerList", err)
		return
	}
	return
}

// Decode APER
func (ie *PosResourceSetTypeAP) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil { // no optional
		return
	}

	tmp := INTEGER{c: aper.Constraint{Lb: 1, Ub: 3}, ext: false}
	if err = tmp.Decode(r); err != nil {
		err = utils.WrapError("Read SRSResourceTriggerList", err)
		return
	}
	ie.SRSResourceTriggerList = int64(tmp.Value)
	return
}
