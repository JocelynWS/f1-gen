package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SRSResourceSet struct {
	SRSResourceSetID  int64             `lb:0,ub:15,mandatory`
	SRSResourceIDList SRSResourceIDList `mandatory`
	ResourceSetType   ResourceSetType   `mandatory`
	// IEExtensions * `optional`
}

func (ie *SRSResourceSet) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_SRSResourceSetID := NewINTEGER(ie.SRSResourceSetID, aper.Constraint{Lb: 0, Ub: 15}, false)
	if err = tmp_SRSResourceSetID.Encode(w); err != nil {
		err = utils.WrapError("Encode SRSResourceSetID", err)
		return
	}
	if err = ie.SRSResourceIDList.Encode(w); err != nil {
		err = utils.WrapError("Encode SRSResourceIDList", err)
		return
	}
	if err = ie.ResourceSetType.Encode(w); err != nil {
		err = utils.WrapError("Encode ResourceSetType", err)
		return
	}
	return
}
func (ie *SRSResourceSet) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_SRSResourceSetID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 15},
		ext: false,
	}
	if err = tmp_SRSResourceSetID.Decode(r); err != nil {
		err = utils.WrapError("Read SRSResourceSetID", err)
		return
	}
	ie.SRSResourceSetID = int64(tmp_SRSResourceSetID.Value)
	if err = ie.SRSResourceIDList.Decode(r); err != nil {
		err = utils.WrapError("Read SRSResourceIDList", err)
		return
	}
	if err = ie.ResourceSetType.Decode(r); err != nil {
		err = utils.WrapError("Read ResourceSetType", err)
		return
	}
	return
}
