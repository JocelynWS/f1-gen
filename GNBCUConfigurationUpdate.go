package f1ap

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GNBCUConfigurationUpdate struct {
	TransactionID                   int64                             `lb:0,ub:255,mandatory,reject`
	CellstobeActivatedList          []CellstobeActivatedListItem      `lb:1,ub:maxCellingNBDU,optional,reject,valueExt`
	CellstobeDeactivatedList        []CellsToBeDeactivatedListItem    `lb:1,ub:maxCellingNBDU,optional,reject,valueExt`
	GNBCUTNLAssociationToAddList    []GNBCUTNLAssociationToAddItem    `lb:1,ub:maxnoofTNLAssociations,optional,ignore,valueExt`
	GNBCUTNLAssociationToRemoveList []GNBCUTNLAssociationToRemoveItem `lb:1,ub:maxnoofTNLAssociations,optional,ignore,valueExt`
	GNBCUTNLAssociationToUpdateList []GNBCUTNLAssociationToUpdateItem `lb:1,ub:maxnoofTNLAssociations,optional,ignore,valueExt`
	CellstobeBarredList             []CellsToBeBarredItem             `lb:1,ub:maxCellingNBDU,optional,ignore,valueExt`
	ProtectedEUTRAResourcesList     []ProtectedEUTRAResourcesItem     `lb:1,ub:maxCellineNB,optional,reject,valueExt`
	NeighbourCellInformationList    []NeighbourCellInformationItem    `lb:1,ub:maxCellingNBDU,optional,ignore,valueExt`
	TransportLayerAddressInfo       *TransportLayerAddressInfo        `optional,ignore`
	ULBHNonUPTrafficMapping         *ULBHNonUPTrafficMapping          `optional,reject`
}

func (msg *GNBCUConfigurationUpdate) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("GNBCUConfigurationUpdate"), err)
		return
	}
	return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_GNBCUConfigurationUpdate, Criticality_PresentReject, ies)
}

func (msg *GNBCUConfigurationUpdate) toIes() (ies []F1apMessageIE, err error) {
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

	if len(msg.GNBCUTNLAssociationToAddList) > 0 {
		tmp_GNBCUTNLAssociationToAddList := Sequence[*GNBCUTNLAssociationToAddItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: true,
		}
		for _, i := range msg.GNBCUTNLAssociationToAddList {
			tmp_GNBCUTNLAssociationToAddList.Value = append(tmp_GNBCUTNLAssociationToAddList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GNBCUTNLAssociationToAddList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_GNBCUTNLAssociationToAddList,
		})
	}

	if len(msg.GNBCUTNLAssociationToRemoveList) > 0 {
		tmp_GNBCUTNLAssociationToRemoveList := Sequence[*GNBCUTNLAssociationToRemoveItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: true,
		}
		for _, i := range msg.GNBCUTNLAssociationToRemoveList {
			tmp_GNBCUTNLAssociationToRemoveList.Value = append(tmp_GNBCUTNLAssociationToRemoveList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GNBCUTNLAssociationToRemoveList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_GNBCUTNLAssociationToRemoveList,
		})
	}

	if len(msg.GNBCUTNLAssociationToUpdateList) > 0 {
		tmp_GNBCUTNLAssociationToUpdateList := Sequence[*GNBCUTNLAssociationToUpdateItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: true,
		}
		for _, i := range msg.GNBCUTNLAssociationToUpdateList {
			tmp_GNBCUTNLAssociationToUpdateList.Value = append(tmp_GNBCUTNLAssociationToUpdateList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GNBCUTNLAssociationToUpdateList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_GNBCUTNLAssociationToUpdateList,
		})
	}

	if len(msg.CellstobeBarredList) > 0 {
		tmp_CellstobeBarredList := Sequence[*CellsToBeBarredItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		for _, i := range msg.CellstobeBarredList {
			tmp_CellstobeBarredList.Value = append(tmp_CellstobeBarredList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CellsToBeBarredList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_CellstobeBarredList,
		})
	}

	if len(msg.ProtectedEUTRAResourcesList) > 0 {
		tmp_ProtectedEUTRAResourcesList := Sequence[*ProtectedEUTRAResourcesItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellineNB},
			ext: true,
		}
		for _, i := range msg.ProtectedEUTRAResourcesList {
			tmp_ProtectedEUTRAResourcesList.Value = append(tmp_ProtectedEUTRAResourcesList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ProtectedEUTRAResourcesList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_ProtectedEUTRAResourcesList,
		})
	}

	if len(msg.NeighbourCellInformationList) > 0 {
		tmp_NeighbourCellInformationList := Sequence[*NeighbourCellInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		for _, i := range msg.NeighbourCellInformationList {
			tmp_NeighbourCellInformationList.Value = append(tmp_NeighbourCellInformationList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NeighbourCellInformationList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_NeighbourCellInformationList,
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

	return
}

func (msg *GNBCUConfigurationUpdate) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("GNBCUConfigurationUpdate"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := GNBCUConfigurationUpdateDecoder{
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

type GNBCUConfigurationUpdateDecoder struct {
	msg      *GNBCUConfigurationUpdate
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *GNBCUConfigurationUpdateDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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

	case ProtocolIEID_GNBCUTNLAssociationToAddList:
		tmp := Sequence[*GNBCUTNLAssociationToAddItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: true,
		}
		fn := func() *GNBCUTNLAssociationToAddItem { return new(GNBCUTNLAssociationToAddItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read GNBCUTNLAssociationToAddList", err)
			return
		}
		msg.GNBCUTNLAssociationToAddList = []GNBCUTNLAssociationToAddItem{}
		for _, i := range tmp.Value {
			msg.GNBCUTNLAssociationToAddList = append(msg.GNBCUTNLAssociationToAddList, *i)
		}

	case ProtocolIEID_GNBCUTNLAssociationToRemoveList:
		tmp := Sequence[*GNBCUTNLAssociationToRemoveItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: true,
		}
		fn := func() *GNBCUTNLAssociationToRemoveItem { return new(GNBCUTNLAssociationToRemoveItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read GNBCUTNLAssociationToRemoveList", err)
			return
		}
		msg.GNBCUTNLAssociationToRemoveList = []GNBCUTNLAssociationToRemoveItem{}
		for _, i := range tmp.Value {
			msg.GNBCUTNLAssociationToRemoveList = append(msg.GNBCUTNLAssociationToRemoveList, *i)
		}

	case ProtocolIEID_GNBCUTNLAssociationToUpdateList:
		tmp := Sequence[*GNBCUTNLAssociationToUpdateItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: true,
		}
		fn := func() *GNBCUTNLAssociationToUpdateItem { return new(GNBCUTNLAssociationToUpdateItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read GNBCUTNLAssociationToUpdateList", err)
			return
		}
		msg.GNBCUTNLAssociationToUpdateList = []GNBCUTNLAssociationToUpdateItem{}
		for _, i := range tmp.Value {
			msg.GNBCUTNLAssociationToUpdateList = append(msg.GNBCUTNLAssociationToUpdateList, *i)
		}

	case ProtocolIEID_CellsToBeBarredList:
		tmp := Sequence[*CellsToBeBarredItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		fn := func() *CellsToBeBarredItem { return new(CellsToBeBarredItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read CellstobeBarredList", err)
			return
		}
		msg.CellstobeBarredList = []CellsToBeBarredItem{}
		for _, i := range tmp.Value {
			msg.CellstobeBarredList = append(msg.CellstobeBarredList, *i)
		}

	case ProtocolIEID_ProtectedEUTRAResourcesList:
		tmp := Sequence[*ProtectedEUTRAResourcesItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellineNB},
			ext: true,
		}
		fn := func() *ProtectedEUTRAResourcesItem { return new(ProtectedEUTRAResourcesItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read ProtectedEUTRAResourcesList", err)
			return
		}
		msg.ProtectedEUTRAResourcesList = []ProtectedEUTRAResourcesItem{}
		for _, i := range tmp.Value {
			msg.ProtectedEUTRAResourcesList = append(msg.ProtectedEUTRAResourcesList, *i)
		}

	case ProtocolIEID_NeighbourCellInformationList:
		tmp := Sequence[*NeighbourCellInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		fn := func() *NeighbourCellInformationItem { return new(NeighbourCellInformationItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read NeighbourCellInformationList", err)
			return
		}
		msg.NeighbourCellInformationList = []NeighbourCellInformationItem{}
		for _, i := range tmp.Value {
			msg.NeighbourCellInformationList = append(msg.NeighbourCellInformationList, *i)
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
