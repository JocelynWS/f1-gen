package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GNBDUConfigurationUpdate struct {
	TransactionID                   int64                             `lb:0,ub:255,mandatory,reject`
	ServedCellsToAddList            []ServedCellsToAddItem            `lb:1,ub:maxnoofCellsingNB,optional,reject,valueExt`
	ServedCellsToModifyList         []ServedCellsToModifyItem         `lb:1,ub:maxCellingNBDU,optional,reject,valueExt`
	ServedCellsToDeleteList         []ServedCellsToDeleteItem         `lb:1,ub:maxCellingNBDU,optional,reject,valueExt`
	CellsStatusList                 []CellsStatusItem                 `lb:0,ub:maxCellingNBDU,optional,reject,valueExt`
	DedicatedSIDeliveryNeededUEList *DedicatedSIDeliveryNeededUEList  `optional,ignore`
	GNBDUID                         *int64                            `lb:0,ub:68719476735,optional,reject`
	GNBDUTNLAssociationToRemoveList []GNBDUTNLAssociationToRemoveItem `lb:1,ub:maxnoofTNLAssociations,optional,reject,valueExt`
	TransportLayerAddressInfo       *TransportLayerAddressInfo        `optional,ignore`
}

func (msg *GNBDUConfigurationUpdate) Encode(w io.Writer) (err error) {
    var ies []F1apMessageIE
    if ies, err = msg.toIes(); err != nil {
        err = msgErrors(fmt.Errorf("GNBDUConfigurationUpdate"), err)
        return
    }
    return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_GNBDUConfigurationUpdate, Criticality_PresentReject, ies)
}
func (msg *GNBDUConfigurationUpdate) toIes() (ies []F1apMessageIE, err error) {
	ies = []F1apMessageIE{}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TransactionID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 255},
			ext:   false,
			Value: aper.Integer(msg.TransactionID),
		}})
	if len(msg.ServedCellsToAddList) > 0 {
		tmp_ServedCellsToAddList := Sequence[*ServedCellsToAddItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellsingNB},
			ext: true,
		}
		for _, i := range msg.ServedCellsToAddList {
			tmp_ServedCellsToAddList.Value = append(tmp_ServedCellsToAddList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ServedCellsToAddList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_ServedCellsToAddList,
		})
	}
	if len(msg.ServedCellsToModifyList) > 0 {
		tmp_ServedCellsToModifyList := Sequence[*ServedCellsToModifyItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		for _, i := range msg.ServedCellsToModifyList {
			tmp_ServedCellsToModifyList.Value = append(tmp_ServedCellsToModifyList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ServedCellsToModifyList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_ServedCellsToModifyList,
		})
	}
	if len(msg.ServedCellsToDeleteList) > 0 {
		tmp_ServedCellsToDeleteList := Sequence[*ServedCellsToDeleteItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		for _, i := range msg.ServedCellsToDeleteList {
			tmp_ServedCellsToDeleteList.Value = append(tmp_ServedCellsToDeleteList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ServedCellsToDeleteList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_ServedCellsToDeleteList,
		})
	}
	if len(msg.CellsStatusList) > 0 {
		tmp_CellsStatusList := Sequence[*CellsStatusItem]{
			c:   aper.Constraint{Lb: 0, Ub: maxCellingNBDU},
			ext: true,
		}
		for _, i := range msg.CellsStatusList {
			tmp_CellsStatusList.Value = append(tmp_CellsStatusList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CellsStatusList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_CellsStatusList,
		})
	}
	if msg.DedicatedSIDeliveryNeededUEList != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DedicatedSIDeliveryNeededUEList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.DedicatedSIDeliveryNeededUEList,
		})
	}
	if msg.GNBDUID != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GNBDUID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 68719476735},
				ext:   false,
				Value: aper.Integer(*msg.GNBDUID),
			}})
	}
	if len(msg.GNBDUTNLAssociationToRemoveList) > 0 {
		tmp_GNBDUTNLAssociationToRemoveList := Sequence[*GNBDUTNLAssociationToRemoveItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: true,
		}
		for _, i := range msg.GNBDUTNLAssociationToRemoveList {
			tmp_GNBDUTNLAssociationToRemoveList.Value = append(tmp_GNBDUTNLAssociationToRemoveList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GNBDUTNLAssociationToRemoveList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_GNBDUTNLAssociationToRemoveList,
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
func (msg *GNBDUConfigurationUpdate) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("GNBDUConfigurationUpdate"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := GNBDUConfigurationUpdateDecoder{
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

type GNBDUConfigurationUpdateDecoder struct {
	msg      *GNBDUConfigurationUpdate
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *GNBDUConfigurationUpdateDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_ServedCellsToAddList:
		tmp := Sequence[*ServedCellsToAddItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellsingNB},
			ext: true,
		}
		fn := func() *ServedCellsToAddItem { return new(ServedCellsToAddItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read ServedCellsToAddList", err)
			return
		}
		msg.ServedCellsToAddList = []ServedCellsToAddItem{}
		for _, i := range tmp.Value {
			msg.ServedCellsToAddList = append(msg.ServedCellsToAddList, *i)
		}
	case ProtocolIEID_ServedCellsToModifyList:
		tmp := Sequence[*ServedCellsToModifyItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		fn := func() *ServedCellsToModifyItem { return new(ServedCellsToModifyItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read ServedCellsToModifyList", err)
			return
		}
		msg.ServedCellsToModifyList = []ServedCellsToModifyItem{}
		for _, i := range tmp.Value {
			msg.ServedCellsToModifyList = append(msg.ServedCellsToModifyList, *i)
		}
	case ProtocolIEID_ServedCellsToDeleteList:
		tmp := Sequence[*ServedCellsToDeleteItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		fn := func() *ServedCellsToDeleteItem { return new(ServedCellsToDeleteItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read ServedCellsToDeleteList", err)
			return
		}
		msg.ServedCellsToDeleteList = []ServedCellsToDeleteItem{}
		for _, i := range tmp.Value {
			msg.ServedCellsToDeleteList = append(msg.ServedCellsToDeleteList, *i)
		}
	case ProtocolIEID_CellsStatusList:
		tmp := Sequence[*CellsStatusItem]{
			c:   aper.Constraint{Lb: 0, Ub: maxCellingNBDU},
			ext: true,
		}
		fn := func() *CellsStatusItem { return new(CellsStatusItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read CellsStatusList", err)
			return
		}
		msg.CellsStatusList = []CellsStatusItem{}
		for _, i := range tmp.Value {
			msg.CellsStatusList = append(msg.CellsStatusList, *i)
		}
	case ProtocolIEID_DedicatedSIDeliveryNeededUEList:
		var tmp DedicatedSIDeliveryNeededUEList
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read DedicatedSIDeliveryNeededUEList", err)
			return
		}
		msg.DedicatedSIDeliveryNeededUEList = &tmp
	case ProtocolIEID_GNBDUID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 68719476735},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read GNBDUID", err)
			return
		}
		msg.GNBDUID = (*int64)(&tmp.Value)
	case ProtocolIEID_GNBDUTNLAssociationToRemoveList:
		tmp := Sequence[*GNBDUTNLAssociationToRemoveItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: true,
		}
		fn := func() *GNBDUTNLAssociationToRemoveItem { return new(GNBDUTNLAssociationToRemoveItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read GNBDUTNLAssociationToRemoveList", err)
			return
		}
		msg.GNBDUTNLAssociationToRemoveList = []GNBDUTNLAssociationToRemoveItem{}
		for _, i := range tmp.Value {
			msg.GNBDUTNLAssociationToRemoveList = append(msg.GNBDUTNLAssociationToRemoveList, *i)
		}
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
