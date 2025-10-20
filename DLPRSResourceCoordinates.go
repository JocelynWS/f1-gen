package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DLPRSResourceCoordinates struct {
	ListofDLPRSResourceSetARP SEQUENCE `mandatory`
	// IEExtensions * `optional`
}

func (ie *DLPRSResourceCoordinates) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.ListofDLPRSResourceSetARP.Encode(w); err != nil {
		err = utils.WrapError("Encode ListofDLPRSResourceSetARP", err)
		return
	}
	return
}
func (ie *DLPRSResourceCoordinates) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.ListofDLPRSResourceSetARP.Decode(r); err != nil {
		err = utils.WrapError("Read ListofDLPRSResourceSetARP", err)
		return
	}
	return
}
