package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type WriteReplaceWarningResponse struct {
	TransactionID                   int64                                 `lb:0,ub:255,mandatory,reject`
	CellsBroadcastCompletedList     []CellsBroadcastCompletedItem `lb:1,ub:maxCellingNBDU,optional,reject,valueExt`
	CriticalityDiagnostics          *CriticalityDiagnostics               `optional,ignore`
	DedicatedSIDeliveryNeededUEList *DedicatedSIDeliveryNeededUEItem      `optional,ignore`
}

func (msg *WriteReplaceWarningResponse) Encode(w io.Writer) (err error) {
    var ies []F1apMessageIE
    if ies, err = msg.toIes(); err != nil {
        err = msgErrors(fmt.Errorf("WriteReplaceWarningResponse"), err)
        return
    }
    return encodeMessage(w, F1apPduSuccessfulOutcome, ProcedureCode_WriteReplaceWarning, Criticality_PresentReject, ies)
}
func (msg *WriteReplaceWarningResponse) toIes() (ies []F1apMessageIE, err error) {
	ies = []F1apMessageIE{}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TransactionID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 255},
			ext:   false,
			Value: aper.Integer(msg.TransactionID),
		}})
	if len(msg.CellsBroadcastCompletedList) > 0 {
		tmp_CellsBroadcastCompletedList := Sequence[*CellsBroadcastCompletedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		for _, i := range msg.CellsBroadcastCompletedList {
			tmp_CellsBroadcastCompletedList.Value = append(tmp_CellsBroadcastCompletedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CellsBroadcastCompletedList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_CellsBroadcastCompletedList,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if msg.DedicatedSIDeliveryNeededUEList != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DedicatedSIDeliveryNeededUEList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.DedicatedSIDeliveryNeededUEList,
		})
	}
	return
}
func (msg *WriteReplaceWarningResponse) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("WriteReplaceWarningResponse"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := WriteReplaceWarningResponseDecoder{
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
	return
}

type WriteReplaceWarningResponseDecoder struct {
	msg      *WriteReplaceWarningResponse
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *WriteReplaceWarningResponseDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read TransactionID", err)
			return
		}
		msg.TransactionID = int64(tmp.Value)
	case ProtocolIEID_CellsBroadcastCompletedList:
		tmp := Sequence[*CellsBroadcastCompletedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		fn := func() *CellsBroadcastCompletedItem { return new(CellsBroadcastCompletedItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read CellsBroadcastCompletedList", err)
			return
		}
		msg.CellsBroadcastCompletedList = []CellsBroadcastCompletedItem{}
		for _, i := range tmp.Value {
			msg.CellsBroadcastCompletedList = append(msg.CellsBroadcastCompletedList, *i)
		}
	case ProtocolIEID_CriticalityDiagnostics:
		var tmp CriticalityDiagnostics
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read CriticalityDiagnostics", err)
			return
		}
		msg.CriticalityDiagnostics = &tmp
	case ProtocolIEID_DedicatedSIDeliveryNeededUEList:
		var tmp DedicatedSIDeliveryNeededUEItem
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read DedicatedSIDeliveryNeededUEList", err)
			return
		}
		msg.DedicatedSIDeliveryNeededUEList = &tmp
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
