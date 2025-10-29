package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TRPPositionReferenced struct {
	ReferencePoint     ReferencePoint        `mandatory`
	ReferencePointType TRPReferencePointType `mandatory`
	//IEExtensions *ProtocolExtensionContainer optional
}

func (ie *TRPPositionReferenced) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 0)
	if err = ie.ReferencePoint.Encode(w); err != nil {
		err = utils.WrapError("Encode ReferencePoint", err)
		return
	}
	if err = ie.ReferencePointType.Encode(w); err != nil {
		err = utils.WrapError("Encode ReferencePointType", err)
		return
	}
	return
}

func (ie *TRPPositionReferenced) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if err = ie.ReferencePoint.Decode(r); err != nil {
		err = utils.WrapError("Read ReferencePoint", err)
		return
	}
	if err = ie.ReferencePointType.Decode(r); err != nil {
		err = utils.WrapError("Read ReferencePointType", err)
		return
	}
	return
}
