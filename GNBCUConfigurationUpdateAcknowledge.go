package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GNBCUConfigurationUpdateAcknowledge struct {
	TransactionID                        int64                                  `lb:0,ub:255,mandatory,reject`
	CellsFailedtobeActivatedList         []CellsFailedToBeActivatedListItem `lb:1,ub:maxCellingNBDU,mandatory,reject,valueExt`
	CriticalityDiagnostics               *CriticalityDiagnostics                `optional,ignore`
	GNBCUTNLAssociationSetupList         []GNBCUTNLAssociationSetupItem     `lb:1,ub:maxnoofTNLAssociations,optional,ignore,valueExt`
	GNBCUTNLAssociationFailedToSetupList []GNBCUTNLAssociationFailedToSetupItem `lb:1,ub:maxnoofTNLAssociations,optional,ignore,valueExt`
	DedicatedSIDeliveryNeededUEList      *DedicatedSIDeliveryNeededUEItem       `optional,ignore`
	TransportLayerAddressInfo            *TransportLayerAddressInfo             `optional,ignore`
}

func (msg *GNBCUConfigurationUpdateAcknowledge) Encode(w io.Writer) (err error) {
    var ies []F1apMessageIE
    if ies, err = msg.toIes(); err != nil {
        err = msgErrors(fmt.Errorf("GNBCUConfigurationUpdateAcknowledge"), err)
        return
    }
    return encodeMessage(w, F1apPduSuccessfulOutcome, ProcedureCode_GNBCUConfigurationUpdate, Criticality_PresentReject, ies)
}
func (msg *GNBCUConfigurationUpdateAcknowledge) toIes() (ies []F1apMessageIE, err error) {
	ies = []F1apMessageIE{}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TransactionID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 255},
			ext:   false,
			Value: aper.Integer(msg.TransactionID),
		}})
	if len(msg.CellsFailedtobeActivatedList) > 0 {
		tmp_CellsFailedtobeActivatedList := Sequence[*CellsFailedToBeActivatedListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		for _, i := range msg.CellsFailedtobeActivatedList {
			tmp_CellsFailedtobeActivatedList.Value = append(tmp_CellsFailedtobeActivatedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CellsFailedToBeActivatedList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_CellsFailedtobeActivatedList,
		})
	} else {
		err = utils.WrapError("CellsFailedtobeActivatedList is nil", err)
		return
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if len(msg.GNBCUTNLAssociationSetupList) > 0 {
		tmp_GNBCUTNLAssociationSetupList := Sequence[*GNBCUTNLAssociationSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: true,
		}
		for _, i := range msg.GNBCUTNLAssociationSetupList {
			tmp_GNBCUTNLAssociationSetupList.Value = append(tmp_GNBCUTNLAssociationSetupList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GNBCUTNLAssociationSetupList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_GNBCUTNLAssociationSetupList,
		})
	}
	if len(msg.GNBCUTNLAssociationFailedToSetupList) > 0 {
		tmp_GNBCUTNLAssociationFailedToSetupList := Sequence[*GNBCUTNLAssociationFailedToSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: true,
		}
		for _, i := range msg.GNBCUTNLAssociationFailedToSetupList {
			tmp_GNBCUTNLAssociationFailedToSetupList.Value = append(tmp_GNBCUTNLAssociationFailedToSetupList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GNBCUTNLAssociationFailedToSetupList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_GNBCUTNLAssociationFailedToSetupList,
		})
	}
	if msg.DedicatedSIDeliveryNeededUEList != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DedicatedSIDeliveryNeededUEList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.DedicatedSIDeliveryNeededUEList,
		})
	}
	if msg.TransportLayerAddressInfo != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TransportLayerAddressInfo},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.TransportLayerAddressInfo,
		})
	}
	return
}
func (msg *GNBCUConfigurationUpdateAcknowledge) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("GNBCUConfigurationUpdateAcknowledge"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := GNBCUConfigurationUpdateAcknowledgeDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_CellsFailedToBeActivatedList]; !ok {
		err = fmt.Errorf("Mandatory field CellsFailedtobeActivatedList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_CellsFailedToBeActivatedList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type GNBCUConfigurationUpdateAcknowledgeDecoder struct {
	msg      *GNBCUConfigurationUpdateAcknowledge
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *GNBCUConfigurationUpdateAcknowledgeDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_CellsFailedToBeActivatedList:
		tmp := Sequence[*CellsFailedToBeActivatedListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		fn := func() *CellsFailedToBeActivatedListItem { return new(CellsFailedToBeActivatedListItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read CellsFailedtobeActivatedList", err)
			return
		}
		msg.CellsFailedtobeActivatedList = []CellsFailedToBeActivatedListItem{}
		for _, i := range tmp.Value {
			msg.CellsFailedtobeActivatedList = append(msg.CellsFailedtobeActivatedList, *i)
		}
	case ProtocolIEID_CriticalityDiagnostics:
		var tmp CriticalityDiagnostics
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read CriticalityDiagnostics", err)
			return
		}
		msg.CriticalityDiagnostics = &tmp
	case ProtocolIEID_GNBCUTNLAssociationSetupList:
		tmp := Sequence[*GNBCUTNLAssociationSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: true,
		}
		fn := func() *GNBCUTNLAssociationSetupItem { return new(GNBCUTNLAssociationSetupItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read GNBCUTNLAssociationSetupList", err)
			return
		}
		msg.GNBCUTNLAssociationSetupList = []GNBCUTNLAssociationSetupItem{}
		for _, i := range tmp.Value {
			msg.GNBCUTNLAssociationSetupList = append(msg.GNBCUTNLAssociationSetupList, *i)
		}
	case ProtocolIEID_GNBCUTNLAssociationFailedToSetupList:
		tmp := Sequence[*GNBCUTNLAssociationFailedToSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: true,
		}
		fn := func() *GNBCUTNLAssociationFailedToSetupItem { return new(GNBCUTNLAssociationFailedToSetupItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read GNBCUTNLAssociationFailedToSetupList", err)
			return
		}
		msg.GNBCUTNLAssociationFailedToSetupList = []GNBCUTNLAssociationFailedToSetupItem{}
		for _, i := range tmp.Value {
			msg.GNBCUTNLAssociationFailedToSetupList = append(msg.GNBCUTNLAssociationFailedToSetupList, *i)
		}
	case ProtocolIEID_DedicatedSIDeliveryNeededUEList:
		var tmp DedicatedSIDeliveryNeededUEItem
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read DedicatedSIDeliveryNeededUEList", err)
			return
		}
		msg.DedicatedSIDeliveryNeededUEList = &tmp
	case ProtocolIEID_TransportLayerAddressInfo:
		var tmp TransportLayerAddressInfo
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read TransportLayerAddressInfo", err)
			return
		}
		msg.TransportLayerAddressInfo = &tmp
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
