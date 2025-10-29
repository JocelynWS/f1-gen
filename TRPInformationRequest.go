package f1ap

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TRPInformationRequest struct {
	TransactionID                int64                `lb:0,ub:255,mandatory,reject`
	TRPList                      []TRPListItem        `lb:1,ub:maxnoofTRPs,optional,ignore,valueExt`
	TRPInformationTypeListTRPReq []TRPInformationItem `lb:1,ub:maxnoofTRPInfoTypes,mandatory,reject,valueExt`
}

func (msg *TRPInformationRequest) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("TRPInformationRequest"), err)
		return
	}
	return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_TRPInformationExchange, Criticality_PresentReject, ies)
}
func (msg *TRPInformationRequest) toIes() (ies []F1apMessageIE, err error) {
	ies = []F1apMessageIE{}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TransactionID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 255},
			ext:   false,
			Value: aper.Integer(msg.TransactionID),
		}})
	if len(msg.TRPList) > 0 {
		tmp_TRPList := Sequence[*TRPListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTRPs},
			ext: true,
		}
		for _, i := range msg.TRPList {
			tmp_TRPList.Value = append(tmp_TRPList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TRPList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_TRPList,
		})
	}
	if len(msg.TRPInformationTypeListTRPReq) > 0 {
		tmp_TRPInformationTypeListTRPReq := Sequence[*TRPInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTRPInfoTypes},
			ext: true,
		}
		for _, i := range msg.TRPInformationTypeListTRPReq {
			tmp_TRPInformationTypeListTRPReq.Value = append(tmp_TRPInformationTypeListTRPReq.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TRPInformationTypeListTRPReq},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_TRPInformationTypeListTRPReq,
		})
	} else {
		err = utils.WrapError("TRPInformationTypeListTRPReq is nil", err)
		return
	}
	return
}
func (msg *TRPInformationRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("TRPInformationRequest"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := TRPInformationRequestDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_TRPInformationTypeListTRPReq]; !ok {
		err = fmt.Errorf("Mandatory field TRPInformationTypeListTRPReq is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_TRPInformationTypeListTRPReq},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type TRPInformationRequestDecoder struct {
	msg      *TRPInformationRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *TRPInformationRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_TRPList:
		tmp := Sequence[*TRPListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTRPs},
			ext: true,
		}
		fn := func() *TRPListItem { return new(TRPListItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read TRPList", err)
			return
		}
		msg.TRPList = []TRPListItem{}
		for _, i := range tmp.Value {
			msg.TRPList = append(msg.TRPList, *i)
		}
	case ProtocolIEID_TRPInformationTypeListTRPReq:
		tmp := Sequence[*TRPInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTRPInfoTypes},
			ext: true,
		}
		fn := func() *TRPInformationItem { return new(TRPInformationItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read TRPInformationTypeListTRPReq", err)
			return
		}
		msg.TRPInformationTypeListTRPReq = []TRPInformationItem{}
		for _, i := range tmp.Value {
			msg.TRPInformationTypeListTRPReq = append(msg.TRPInformationTypeListTRPReq, *i)
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
