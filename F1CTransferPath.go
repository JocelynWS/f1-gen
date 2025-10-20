package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type F1CTransferPath struct {
	F1CPathNSA F1CPathNSA `mandatory`
	// IEExtensions * `optional`
}

func (ie *F1CTransferPath) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.F1CPathNSA.Encode(w); err != nil {
		err = utils.WrapError("Encode F1CPathNSA", err)
		return
	}
	return
}
func (ie *F1CTransferPath) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.F1CPathNSA.Decode(r); err != nil {
		err = utils.WrapError("Read F1CPathNSA", err)
		return
	}
	return
}
