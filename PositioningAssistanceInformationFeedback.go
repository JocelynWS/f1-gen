package f1ap

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PositioningAssistanceInformationFeedback struct {
	TransactionID                       int64                   `lb:0,ub:255,mandatory,reject`
	PosAssistanceInformationFailureList []byte                  `lb:0,ub:0,mandatory,reject`
	PositioningBroadcastCells           []NRCGI                 `lb:1,ub:maxCellingNBDU,mandatory,reject,valueExt`
	RoutingID                           []byte                  `lb:0,ub:0,mandatory,reject`
	CriticalityDiagnostics              *CriticalityDiagnostics `mandatory,ignore`
}

func (msg *PositioningAssistanceInformationFeedback) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("PositioningAssistanceInformationFeedback"), err)
		return
	}
	return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_PositioningAssistanceInformationFeedback, Criticality_PresentIgnore, ies)
}
func (msg *PositioningAssistanceInformationFeedback) toIes() (ies []F1apMessageIE, err error) {
	ies = []F1apMessageIE{}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TransactionID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 255},
			ext:   false,
			Value: aper.Integer(msg.TransactionID),
		}})
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_PosAssistanceInformationFailureList},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 0, Ub: 0},
			ext:   false,
			Value: msg.PosAssistanceInformationFailureList,
		}})
	if len(msg.PositioningBroadcastCells) > 0 {
		tmp_PositioningBroadcastCells := Sequence[*NRCGI]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		for _, i := range msg.PositioningBroadcastCells {
			tmp_PositioningBroadcastCells.Value = append(tmp_PositioningBroadcastCells.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PositioningBroadcastCells},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_PositioningBroadcastCells,
		})
	} else {
		err = utils.WrapError("PositioningBroadcastCells is nil", err)
		return
	}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RoutingID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 0, Ub: 0},
			ext:   false,
			Value: msg.RoutingID,
		}})
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       msg.CriticalityDiagnostics,
	})
	return
}
func (msg *PositioningAssistanceInformationFeedback) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("PositioningAssistanceInformationFeedback"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := PositioningAssistanceInformationFeedbackDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_PosAssistanceInformationFailureList]; !ok {
		err = fmt.Errorf("Mandatory field PosAssistanceInformationFailureList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_PosAssistanceInformationFailureList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_PositioningBroadcastCells]; !ok {
		err = fmt.Errorf("Mandatory field PositioningBroadcastCells is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_PositioningBroadcastCells},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_RoutingID]; !ok {
		err = fmt.Errorf("Mandatory field RoutingID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_RoutingID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_CriticalityDiagnostics]; !ok {
		err = fmt.Errorf("Mandatory field CriticalityDiagnostics is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type PositioningAssistanceInformationFeedbackDecoder struct {
	msg      *PositioningAssistanceInformationFeedback
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *PositioningAssistanceInformationFeedbackDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_PosAssistanceInformationFailureList:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read PosAssistanceInformationFailureList", err)
			return
		}
		msg.PosAssistanceInformationFailureList = tmp.Value
	case ProtocolIEID_PositioningBroadcastCells:
		tmp := Sequence[*NRCGI]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		fn := func() *NRCGI { return new(NRCGI) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read PositioningBroadcastCells", err)
			return
		}
		msg.PositioningBroadcastCells = []NRCGI{}
		for _, i := range tmp.Value {
			msg.PositioningBroadcastCells = append(msg.PositioningBroadcastCells, *i)
		}
	case ProtocolIEID_RoutingID:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RoutingID", err)
			return
		}
		msg.RoutingID = tmp.Value
	case ProtocolIEID_CriticalityDiagnostics:
		var tmp CriticalityDiagnostics
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read CriticalityDiagnostics", err)
			return
		}
		msg.CriticalityDiagnostics = &tmp
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
