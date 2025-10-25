package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GNBDUConfigurationUpdateAcknowledge struct {
	TransactionID             int64                          `lb:0,ub:255,mandatory,reject`
	CellstobeActivatedList    []CellstobeActivatedListItem   `lb:1,ub:maxCellingNBDU,optional,reject,valueExt`
	CriticalityDiagnostics    *CriticalityDiagnostics        `optional,ignore`
	CellstobeDeactivatedList  []CellsToBeDeactivatedListItem `lb:1,ub:maxCellingNBDU,optional,reject,valueExt`
	TransportLayerAddressInfo *TransportLayerAddressInfo     `optional,ignore`
	ULBHNonUPTrafficMapping   *ULBHNonUPTrafficMapping       `optional,reject`
	BAPAddress                *aper.BitString                `lb:10,ub:10,optional,ignore`
}

func (msg *GNBDUConfigurationUpdateAcknowledge) Encode(w io.Writer) (err error) {
    var ies []F1apMessageIE
    if ies, err = msg.toIes(); err != nil {
        err = msgErrors(fmt.Errorf("GNBDUConfigurationUpdateAcknowledge"), err)
        return
    }
    return encodeMessage(w, F1apPduSuccessfulOutcome, ProcedureCode_GNBDUConfigurationUpdate, Criticality_PresentReject, ies)
}
func (msg *GNBDUConfigurationUpdateAcknowledge) toIes() (ies []F1apMessageIE, err error) {
	ies = []F1apMessageIE{}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TransactionID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 255},
			ext:   false,
			Value: aper.Integer(msg.TransactionID),
		}})
	if len(msg.CellstobeActivatedList) > 0 {
		tmp_CellstobeActivatedList := Sequence[*CellstobeActivatedListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		for _, i := range msg.CellstobeActivatedList {
			tmp_CellstobeActivatedList.Value = append(tmp_CellstobeActivatedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CellstobeActivatedList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_CellstobeActivatedList,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if len(msg.CellstobeDeactivatedList) > 0 {
		tmp_CellstobeDeactivatedList := Sequence[*CellsToBeDeactivatedListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		for _, i := range msg.CellstobeDeactivatedList {
			tmp_CellstobeDeactivatedList.Value = append(tmp_CellstobeDeactivatedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CellsToBeDeactivatedList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_CellstobeDeactivatedList,
		})
	}
	if msg.TransportLayerAddressInfo != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TransportLayerAddressInfo},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.TransportLayerAddressInfo,
		})
	}
	if msg.ULBHNonUPTrafficMapping != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ULBHNonUPTrafficMapping},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.ULBHNonUPTrafficMapping,
		})
	}
	if msg.BAPAddress != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_BAPAddress},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &BITSTRING{
				c:   aper.Constraint{Lb: 10, Ub: 10},
				ext: false,
				Value: aper.BitString{
					Bytes: msg.BAPAddress.Bytes, NumBits: msg.BAPAddress.NumBits},
			}})
	}
	return
}
func (msg *GNBDUConfigurationUpdateAcknowledge) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("GNBDUConfigurationUpdateAcknowledge"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := GNBDUConfigurationUpdateAcknowledgeDecoder{
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

type GNBDUConfigurationUpdateAcknowledgeDecoder struct {
	msg      *GNBDUConfigurationUpdateAcknowledge
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *GNBDUConfigurationUpdateAcknowledgeDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_CellstobeActivatedList:
		tmp := Sequence[*CellstobeActivatedListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		fn := func() *CellstobeActivatedListItem { return new(CellstobeActivatedListItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read CellstobeActivatedList", err)
			return
		}
		msg.CellstobeActivatedList = []CellstobeActivatedListItem{}
		for _, i := range tmp.Value {
			msg.CellstobeActivatedList = append(msg.CellstobeActivatedList, *i)
		}
	case ProtocolIEID_CriticalityDiagnostics:
		var tmp CriticalityDiagnostics
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read CriticalityDiagnostics", err)
			return
		}
		msg.CriticalityDiagnostics = &tmp
	case ProtocolIEID_CellsToBeDeactivatedList:
		tmp := Sequence[*CellsToBeDeactivatedListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		fn := func() *CellsToBeDeactivatedListItem { return new(CellsToBeDeactivatedListItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read CellstobeDeactivatedList", err)
			return
		}
		msg.CellstobeDeactivatedList = []CellsToBeDeactivatedListItem{}
		for _, i := range tmp.Value {
			msg.CellstobeDeactivatedList = append(msg.CellstobeDeactivatedList, *i)
		}
	case ProtocolIEID_TransportLayerAddressInfo:
		var tmp TransportLayerAddressInfo
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read TransportLayerAddressInfo", err)
			return
		}
		msg.TransportLayerAddressInfo = &tmp
	case ProtocolIEID_ULBHNonUPTrafficMapping:
		var tmp ULBHNonUPTrafficMapping
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ULBHNonUPTrafficMapping", err)
			return
		}
		msg.ULBHNonUPTrafficMapping = &tmp
	case ProtocolIEID_BAPAddress:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 10, Ub: 10},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read BAPAddress", err)
			return
		}
		msg.BAPAddress = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
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
