package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TRPPositionDirect struct {
	Accuracy TRPPositionDirectAccuracy `mandatory`
	//IEExtensions *ProtocolExtensionContainer optional
}

func (ie *TRPPositionDirect) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 0)
	if err = ie.Accuracy.Encode(w); err != nil {
		err = utils.WrapError("Encode Accuracy", err)
		return
	}
	return
}

func (ie *TRPPositionDirect) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if err = ie.Accuracy.Decode(r); err != nil {
		err = utils.WrapError("Read Accuracy", err)
		return
	}
	return
}
