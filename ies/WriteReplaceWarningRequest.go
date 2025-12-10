package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type WriteReplaceWarningRequest struct {
	TransactionID            int64                        `lb:0,ub:255,mandatory,reject,valueExt`
	PWSSystemInformation     PWSSystemInformation         `mandatory,reject`
	RepetitionPeriod         int64                        `lb:0,ub:131071,mandatory,reject,valueExt`
	NumberofBroadcastRequest int64                        `lb:0,ub:65535,mandatory,reject`
	CellsToBeBroadcastList   []CellsToBeBroadcastListItem `lb:1,ub:maxCellingNBDU,optional,reject,valueExt`
}

func (msg *WriteReplaceWarningRequest) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("WriteReplaceWarningRequest"), err)
		return
	}
	return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_WriteReplaceWarning, Criticality_PresentReject, ies)
}
func (msg *WriteReplaceWarningRequest) toIes() (ies []F1apMessageIE, err error) {
	ies = []F1apMessageIE{}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TransactionID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 255},
			ext:   true,
			Value: aper.Integer(msg.TransactionID),
		}})
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_PWSSystemInformation},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.PWSSystemInformation,
	})
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RepetitionPeriod},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 131071},
			ext:   true,
			Value: aper.Integer(msg.RepetitionPeriod),
		}})
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_NumberofBroadcastRequest},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 65535},
			ext:   false,
			Value: aper.Integer(msg.NumberofBroadcastRequest),
		}})
	if len(msg.CellsToBeBroadcastList) > 0 {
		tmp_CellsToBeBroadcastList := Sequence[*CellsToBeBroadcastListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		for _, i := range msg.CellsToBeBroadcastList {
			tmp_CellsToBeBroadcastList.Value = append(tmp_CellsToBeBroadcastList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CellsToBeBroadcastList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_CellsToBeBroadcastList,
		})
	}
	return
}
func (msg *WriteReplaceWarningRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("WriteReplaceWarningRequest"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := WriteReplaceWarningRequestDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*F1apMessageIE),
	}
	if _, err = aper.ReadSequenceOf[F1apMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	if _, ok := decoder.list[ProtocolIEID_TransactionID]; !ok {
		err = fmt.Errorf("Mandatory field TransactionID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_TransactionID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_PWSSystemInformation]; !ok {
		err = fmt.Errorf("Mandatory field PWSSystemInformation is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_PWSSystemInformation},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_RepetitionPeriod]; !ok {
		err = fmt.Errorf("Mandatory field RepetitionPeriod is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_RepetitionPeriod},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_NumberofBroadcastRequest]; !ok {
		err = fmt.Errorf("Mandatory field NumberofBroadcastRequest is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_NumberofBroadcastRequest},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type WriteReplaceWarningRequestDecoder struct {
	msg      *WriteReplaceWarningRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *WriteReplaceWarningRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
	var id int64
	var c uint64
	var buf []byte
	if id, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: int64(aper.POW_16) - 1}, false); err != nil {
		return
	}
	msgIe = new(F1apMessageIE)
	msgIe.Id.Value = aper.Integer(id)
	if c, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return
	}
	msgIe.Criticality.Value = aper.Enumerated(c)
	if buf, err = r.ReadOpenType(); err != nil {
		return
	}
	ieId := msgIe.Id.Value
	if _, ok := decoder.list[ieId]; ok {
		err = fmt.Errorf("Duplicated protocol IEID[%d] found", ieId)
		return
	}
	decoder.list[ieId] = msgIe
	ieR := aper.NewReader(bytes.NewReader(buf))
	msg := decoder.msg
	switch msgIe.Id.Value {
	case ProtocolIEID_TransactionID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 255},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read TransactionID", err)
			return
		}
		msg.TransactionID = int64(tmp.Value)
	case ProtocolIEID_PWSSystemInformation:
		var tmp PWSSystemInformation
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read PWSSystemInformation", err)
			return
		}
		msg.PWSSystemInformation = tmp
	case ProtocolIEID_RepetitionPeriod:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 131071},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RepetitionPeriod", err)
			return
		}
		msg.RepetitionPeriod = int64(tmp.Value)
	case ProtocolIEID_NumberofBroadcastRequest:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 65535},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read NumberofBroadcastRequest", err)
			return
		}
		msg.NumberofBroadcastRequest = int64(tmp.Value)
	case ProtocolIEID_CellsToBeBroadcastList:
		tmp := Sequence[*CellsToBeBroadcastListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		fn := func() *CellsToBeBroadcastListItem { return new(CellsToBeBroadcastListItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read CellsToBeBroadcastList", err)
			return
		}
		msg.CellsToBeBroadcastList = []CellsToBeBroadcastListItem{}
		for _, i := range tmp.Value {
			msg.CellsToBeBroadcastList = append(msg.CellsToBeBroadcastList, *i)
		}
	default:
		switch msgIe.Criticality.Value {
		case Criticality_PresentReject:
			fmt.Errorf("Not comprehended IE ID 0x%04x (criticality: reject)", msgIe.Id.Value)
		case Criticality_PresentIgnore:
			fmt.Errorf("Not comprehended IE ID 0x%04x (criticality: ignore)", msgIe.Id.Value)
		case Criticality_PresentNotify:
			fmt.Errorf("Not comprehended IE ID 0x%04x (criticality: notify)", msgIe.Id.Value)
		}
		if msgIe.Criticality.Value != Criticality_PresentIgnore {
			decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
				IECriticality: msgIe.Criticality,
				IEID:          msgIe.Id,
				TypeOfError:   TypeOfError{Value: TypeOfErrorNotunderstood},
			})
		}
	}
	return
}
