package ies

import "github.com/lvdund/ngap/aper"

const (
	CauseProtocolTransferSyntaxError                          aper.Enumerated = 0
	CauseProtocolAbstractSyntaxErrorReject                    aper.Enumerated = 1
	CauseProtocolAbstractSyntaxErrorIgnoreAndNotify           aper.Enumerated = 2
	CauseProtocolMessageNotCompatibleWithReceiverState        aper.Enumerated = 3
	CauseProtocolSemanticerror                                aper.Enumerated = 4
	CauseProtocolAbstractSyntaxErrorFalselyConstructedMessage aper.Enumerated = 5
	CauseProtocolUnspecified                                  aper.Enumerated = 6
)

type CauseProtocol struct {
	Value aper.Enumerated
}

func (ie *CauseProtocol) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, true)
	return
}

func (ie *CauseProtocol) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, true)
	ie.Value = aper.Enumerated(v)
	return
}
