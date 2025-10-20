package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type IABTNLAddressResponse struct {
	TransactionID              int64                        `lb:0,ub:255,mandatory,reject`
	IABAllocatedTNLAddressList []IABAllocatedTNLAddressItem `lb:1,ub:maxnoofTLAsIAB,mandatory,reject,valueExt`
}

func (msg *IABTNLAddressResponse) Encode(w io.Writer) (err error) {
    var ies []F1apMessageIE
    if ies, err = msg.toIes(); err != nil {
        err = msgErrors(fmt.Errorf("IABTNLAddressResponse"), err)
        return
    }
    return encodeMessage(w, F1apPduSuccessfulOutcome, ProcedureCode_IABTNLAddressAllocation, Criticality_PresentReject, ies)
}
func (msg *IABTNLAddressResponse) toIes() (ies []F1apMessageIE, err error) {
	ies = []F1apMessageIE{}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TransactionID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 255},
			ext:   false,
			Value: aper.Integer(msg.TransactionID),
		}})
	if len(msg.IABAllocatedTNLAddressList) > 0 {
		tmp_IABAllocatedTNLAddressList := Sequence[*IABAllocatedTNLAddressItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTLAsIAB},
			ext: true,
		}
		for _, i := range msg.IABAllocatedTNLAddressList {
			tmp_IABAllocatedTNLAddressList.Value = append(tmp_IABAllocatedTNLAddressList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_IABAllocatedTNLAddressList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_IABAllocatedTNLAddressList,
		})
	} else {
		err = utils.WrapError("IABAllocatedTNLAddressList is nil", err)
		return
	}
	return
}
func (msg *IABTNLAddressResponse) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("IABTNLAddressResponse"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := IABTNLAddressResponseDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_IABAllocatedTNLAddressList]; !ok {
		err = fmt.Errorf("Mandatory field IABAllocatedTNLAddressList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_IABAllocatedTNLAddressList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type IABTNLAddressResponseDecoder struct {
	msg      *IABTNLAddressResponse
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *IABTNLAddressResponseDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_IABAllocatedTNLAddressList:
		tmp := Sequence[*IABAllocatedTNLAddressItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTLAsIAB},
			ext: true,
		}
		fn := func() *IABAllocatedTNLAddressItem { return new(IABAllocatedTNLAddressItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read IABAllocatedTNLAddressList", err)
			return
		}
		msg.IABAllocatedTNLAddressList = []IABAllocatedTNLAddressItem{}
		for _, i := range tmp.Value {
			msg.IABAllocatedTNLAddressList = append(msg.IABAllocatedTNLAddressList, *i)
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
