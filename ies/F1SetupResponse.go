package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type F1SetupResponse struct {
	TransactionID             int64                        `lb:0,ub:255,mandatory,reject,valueExt`
	GNBCUName                 []byte                       `lb:1,ub:150,optional,ignore,valueExt`
	CellstobeActivatedList    []CellstobeActivatedListItem `lb:1,ub:maxCellingNBDU,optional,reject,valueExt`
	GNBCURRCVersion           RRCVersion                   `mandatory,reject`
	TransportLayerAddressInfo *TransportLayerAddressInfo   `optional,ignore`
	ULBHNonUPTrafficMapping   *ULBHNonUPTrafficMapping     `optional,reject`
	BAPAddress                *aper.BitString              `lb:10,ub:10,optional,ignore`
	ExtendedGNBDUName         *ExtendedGNBDUName           `optional,ignore`
}

func (msg *F1SetupResponse) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("F1SetupResponse"), err)
		return
	}
	return encodeMessage(w, F1apPduSuccessfulOutcome, ProcedureCode_F1Setup, Criticality_PresentReject, ies)
}
func (msg *F1SetupResponse) toIes() (ies []F1apMessageIE, err error) {
	ies = []F1apMessageIE{}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TransactionID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 255},
			ext:   true,
			Value: aper.Integer(msg.TransactionID),
		}})
	if msg.GNBCUName != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GNBCUName},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 1, Ub: 150},
				ext:   true,
				Value: msg.GNBCUName,
			}})
	}
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
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_GNBCURRCVersion},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.GNBCURRCVersion,
	})
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
	if msg.ExtendedGNBDUName != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ExtendedGNBDUName},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.ExtendedGNBDUName,
		})
	}
	return
}
func (msg *F1SetupResponse) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("F1SetupResponse"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := F1SetupResponseDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_GNBCURRCVersion]; !ok {
		err = fmt.Errorf("Mandatory field GNBCURRCVersion is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_GNBCURRCVersion},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type F1SetupResponseDecoder struct {
	msg      *F1SetupResponse
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *F1SetupResponseDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_GNBCUName:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 150},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read GNBCUName", err)
			return
		}
		msg.GNBCUName = tmp.Value
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
	case ProtocolIEID_GNBCURRCVersion:
		var tmp RRCVersion
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read GNBCURRCVersion", err)
			return
		}
		msg.GNBCURRCVersion = tmp
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
	case ProtocolIEID_ExtendedGNBDUName:
		var tmp ExtendedGNBDUName
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ExtendedGNBDUName", err)
			return
		}
		msg.ExtendedGNBDUName = &tmp
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
