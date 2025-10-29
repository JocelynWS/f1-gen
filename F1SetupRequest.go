package f1ap

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type F1SetupRequest struct {
	TransactionID             int64                      `lb:0,ub:255,mandatory,reject`
	GNBDUID                   int64                      `lb:0,ub:68719476735,mandatory,reject`
	GNBDUName                 []byte                     `lb:1,ub:150,optional,ignore,valueExt`
	GNBDUServedCellsList      []GNBDUServedCellsItem     `lb:1,ub:maxCellingNBDU,optional,reject,valueExt`
	GNBDURRCVersion           RRCVersion                 `mandatory,reject`
	TransportLayerAddressInfo *TransportLayerAddressInfo `optional,ignore`
	BAPAddress                *aper.BitString            `lb:10,ub:10,optional,ignore`
	ExtendedGNBCUName         *ExtendedGNBCUName         `optional,ignore`
}

func (msg *F1SetupRequest) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("F1SetupRequest"), err)
		return
	}
	return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_F1Setup, Criticality_PresentReject, ies)
}
func (msg *F1SetupRequest) toIes() (ies []F1apMessageIE, err error) {
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
		Id:          ProtocolIEID{Value: ProtocolIEID_GNBDUID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 68719476735},
			ext:   false,
			Value: aper.Integer(msg.GNBDUID),
		}})
	if msg.GNBDUName != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GNBDUName},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 1, Ub: 150},
				ext:   true,
				Value: msg.GNBDUName,
			}})
	}
	if len(msg.GNBDUServedCellsList) > 0 {
		tmp_GNBDUServedCellsList := Sequence[*GNBDUServedCellsItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		for _, i := range msg.GNBDUServedCellsList {
			tmp_GNBDUServedCellsList.Value = append(tmp_GNBDUServedCellsList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GNBDUServedCellsList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_GNBDUServedCellsList,
		})
	}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_GNBDURRCVersion},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.GNBDURRCVersion,
	})
	if msg.TransportLayerAddressInfo != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TransportLayerAddressInfo},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.TransportLayerAddressInfo,
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
	if msg.ExtendedGNBCUName != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ExtendedGNBCUName},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.ExtendedGNBCUName,
		})
	}
	return
}
func (msg *F1SetupRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("F1SetupRequest"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	//r.ReadBool()
	decoder := F1SetupRequestDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_GNBDUID]; !ok {
		err = fmt.Errorf("Mandatory field GNBDUID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_GNBDUID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_GNBDURRCVersion]; !ok {
		err = fmt.Errorf("Mandatory field GNBDURRCVersion is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_GNBDURRCVersion},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type F1SetupRequestDecoder struct {
	msg      *F1SetupRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *F1SetupRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_GNBDUID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 68719476735},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read GNBDUID", err)
			return
		}
		msg.GNBDUID = int64(tmp.Value)
	case ProtocolIEID_GNBDUName:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 150},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read GNBDUName", err)
			return
		}
		msg.GNBDUName = tmp.Value
	case ProtocolIEID_GNBDUServedCellsList:
		tmp := Sequence[*GNBDUServedCellsItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		fn := func() *GNBDUServedCellsItem { return new(GNBDUServedCellsItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read GNBDUServedCellsList", err)
			return
		}
		msg.GNBDUServedCellsList = []GNBDUServedCellsItem{}
		for _, i := range tmp.Value {
			msg.GNBDUServedCellsList = append(msg.GNBDUServedCellsList, *i)
		}
	case ProtocolIEID_GNBDURRCVersion:
		var tmp RRCVersion
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read GNBDURRCVersion", err)
			return
		}
		msg.GNBDURRCVersion = tmp
	case ProtocolIEID_TransportLayerAddressInfo:
		var tmp TransportLayerAddressInfo
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read TransportLayerAddressInfo", err)
			return
		}
		msg.TransportLayerAddressInfo = &tmp
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
	case ProtocolIEID_ExtendedGNBCUName:
		var tmp ExtendedGNBCUName
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ExtendedGNBCUName", err)
			return
		}
		msg.ExtendedGNBCUName = &tmp
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
