package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SystemInformationDeliveryCommand struct {
	TransactionID int64        `lb:0,ub:255,mandatory,reject`
	NRCGI         NRCGI        `mandatory,reject`
	SItypeList    []SItypeItem `lb:1,ub:maxnoofSITypes,mandatory,reject,valueExt`
	ConfirmedUEID int64        `lb:0,ub:4294967295,mandatory,reject`
}

func (msg *SystemInformationDeliveryCommand) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("SystemInformationDeliveryCommand"), err)
		return
	}
	return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_SystemInformationDeliveryCommand, Criticality_PresentIgnore, ies)
}
func (msg *SystemInformationDeliveryCommand) toIes() (ies []F1apMessageIE, err error) {
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
		Id:          ProtocolIEID{Value: ProtocolIEID_NRCGI},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.NRCGI,
	})
	if len(msg.SItypeList) > 0 {
		tmp_SItypeList := Sequence[*SItypeItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSITypes},
			ext: true,
		}
		for _, i := range msg.SItypeList {
			tmp_SItypeList.Value = append(tmp_SItypeList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SItypeList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_SItypeList,
		})
	} else {
		err = utils.WrapError("SItypeList is nil", err)
		return
	}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_ConfirmedUEID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: aper.Integer(msg.ConfirmedUEID),
		}})
	return
}
func (msg *SystemInformationDeliveryCommand) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("SystemInformationDeliveryCommand"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := SystemInformationDeliveryCommandDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_NRCGI]; !ok {
		err = fmt.Errorf("Mandatory field NRCGI is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_NRCGI},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_SItypeList]; !ok {
		err = fmt.Errorf("Mandatory field SItypeList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_SItypeList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_ConfirmedUEID]; !ok {
		err = fmt.Errorf("Mandatory field ConfirmedUEID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_ConfirmedUEID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type SystemInformationDeliveryCommandDecoder struct {
	msg      *SystemInformationDeliveryCommand
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *SystemInformationDeliveryCommandDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_NRCGI:
		var tmp NRCGI
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read NRCGI", err)
			return
		}
		msg.NRCGI = tmp
	case ProtocolIEID_SItypeList:
		tmp := Sequence[*SItypeItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSITypes},
			ext: true,
		}
		fn := func() *SItypeItem { return new(SItypeItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SItypeList", err)
			return
		}
		msg.SItypeList = []SItypeItem{}
		for _, i := range tmp.Value {
			msg.SItypeList = append(msg.SItypeList, *i)
		}
	case ProtocolIEID_ConfirmedUEID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4294967295},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ConfirmedUEID", err)
			return
		}
		msg.ConfirmedUEID = int64(tmp.Value)
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
