package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type IABUPConfigurationUpdateRequest struct {
	TransactionID                  int64                                `lb:0,ub:255,mandatory,reject,valueExt`
	ULUPTNLInformationtoUpdateList []ULUPTNLInformationtoUpdateListItem `lb:1,ub:maxnoofULUPTNLInformationforIAB,optional,ignore,valueExt`
	ULUPTNLAddresstoUpdateList     []ULUPTNLInformationtoUpdateListItem `lb:1,ub:maxnoofUPTNLAddresses,optional,ignore,valueExt`
}

func (msg *IABUPConfigurationUpdateRequest) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("IABUPConfigurationUpdateRequest"), err)
		return
	}
	return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_IABUPConfigurationUpdate, Criticality_PresentReject, ies)
}
func (msg *IABUPConfigurationUpdateRequest) toIes() (ies []F1apMessageIE, err error) {
	ies = []F1apMessageIE{}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TransactionID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 255},
			ext:   true,
			Value: aper.Integer(msg.TransactionID),
		}})
	if len(msg.ULUPTNLInformationtoUpdateList) > 0 {
		tmp_ULUPTNLInformationtoUpdateList := Sequence[*ULUPTNLInformationtoUpdateListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofULUPTNLInformationforIAB},
			ext: true,
		}
		for _, i := range msg.ULUPTNLInformationtoUpdateList {
			tmp_ULUPTNLInformationtoUpdateList.Value = append(tmp_ULUPTNLInformationtoUpdateList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ULUPTNLInformationtoUpdateList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_ULUPTNLInformationtoUpdateList,
		})
	}
	if len(msg.ULUPTNLAddresstoUpdateList) > 0 {
		tmp_ULUPTNLAddresstoUpdateList := Sequence[*ULUPTNLInformationtoUpdateListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofUPTNLAddresses},
			ext: true,
		}
		for _, i := range msg.ULUPTNLAddresstoUpdateList {
			tmp_ULUPTNLAddresstoUpdateList.Value = append(tmp_ULUPTNLAddresstoUpdateList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ULUPTNLAddressToUpdateList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_ULUPTNLAddresstoUpdateList,
		})
	}
	return
}
func (msg *IABUPConfigurationUpdateRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("IABUPConfigurationUpdateRequest"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := IABUPConfigurationUpdateRequestDecoder{
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

type IABUPConfigurationUpdateRequestDecoder struct {
	msg      *IABUPConfigurationUpdateRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *IABUPConfigurationUpdateRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_ULUPTNLInformationtoUpdateList:
		tmp := Sequence[*ULUPTNLInformationtoUpdateListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofULUPTNLInformationforIAB},
			ext: true,
		}
		fn := func() *ULUPTNLInformationtoUpdateListItem { return new(ULUPTNLInformationtoUpdateListItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read ULUPTNLInformationtoUpdateList", err)
			return
		}
		msg.ULUPTNLInformationtoUpdateList = []ULUPTNLInformationtoUpdateListItem{}
		for _, i := range tmp.Value {
			msg.ULUPTNLInformationtoUpdateList = append(msg.ULUPTNLInformationtoUpdateList, *i)
		}
	case ProtocolIEID_ULUPTNLAddressToUpdateList:
		tmp := Sequence[*ULUPTNLInformationtoUpdateListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofUPTNLAddresses},
			ext: true,
		}
		fn := func() *ULUPTNLInformationtoUpdateListItem { return new(ULUPTNLInformationtoUpdateListItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read ULUPTNLAddresstoUpdateList", err)
			return
		}
		msg.ULUPTNLAddresstoUpdateList = []ULUPTNLInformationtoUpdateListItem{}
		for _, i := range tmp.Value {
			msg.ULUPTNLAddresstoUpdateList = append(msg.ULUPTNLAddresstoUpdateList, *i)
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
