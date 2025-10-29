package f1ap

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PositioningAssistanceInformationControl struct {
	TransactionID             int64         `lb:0,ub:255,mandatory,reject`
	PosAssistanceInformation  []byte        `lb:0,ub:0,mandatory,reject`
	PosBroadcast              *PosBroadcast `mandatory,reject`
	PositioningBroadcastCells []NRCGI       `lb:1,ub:maxCellingNBDU,mandatory,reject,valueExt`
	RoutingID                 []byte        `lb:0,ub:0,mandatory,reject`
}

func (msg *PositioningAssistanceInformationControl) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("PositioningAssistanceInformationControl"), err)
		return
	}
	return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_PositioningAssistanceInformationControl, Criticality_PresentIgnore, ies)
}
func (msg *PositioningAssistanceInformationControl) toIes() (ies []F1apMessageIE, err error) {
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
		Id:          ProtocolIEID{Value: ProtocolIEID_PosAssistanceInformation},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 0, Ub: 0},
			ext:   false,
			Value: msg.PosAssistanceInformation,
		}})
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_PosBroadcast},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       msg.PosBroadcast,
	})
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
	return
}
func (msg *PositioningAssistanceInformationControl) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("PositioningAssistanceInformationControl"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := PositioningAssistanceInformationControlDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_PosAssistanceInformation]; !ok {
		err = fmt.Errorf("Mandatory field PosAssistanceInformation is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_PosAssistanceInformation},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_PosBroadcast]; !ok {
		err = fmt.Errorf("Mandatory field PosBroadcast is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_PosBroadcast},
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
	return
}

type PositioningAssistanceInformationControlDecoder struct {
	msg      *PositioningAssistanceInformationControl
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *PositioningAssistanceInformationControlDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_PosAssistanceInformation:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read PosAssistanceInformation", err)
			return
		}
		msg.PosAssistanceInformation = tmp.Value
	case ProtocolIEID_PosBroadcast:
		var tmp PosBroadcast
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read PosBroadcast", err)
			return
		}
		msg.PosBroadcast = &tmp
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
