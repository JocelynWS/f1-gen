package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DLRRCMessageTransfer struct {
	GNBCUUEF1APID                   int64                            `lb:0,ub:4294967295,mandatory,reject`
	GNBDUUEF1APID                   int64                            `lb:0,ub:4294967295,mandatory,reject`
	OldgNBDUUEF1APID                *int64                           `lb:0,ub:4294967295,optional,reject`
	SRBID                           int64                            `lb:0,ub:3,mandatory,reject`
	ExecuteDuplication              *ExecuteDuplication              `mandatory,ignore`
	RRCContainer                    []byte                           `lb:0,ub:0,mandatory,reject`
	RATFrequencyPriorityInformation *RATFrequencyPriorityInformation `optional,reject`
	RRCDeliveryStatusRequest        *RRCDeliveryStatusRequest        `optional,ignore`
	UEContextNotRetrievable         *UEContextNotRetrievable         `optional,reject`
	RedirectedRRCmessage            []byte                           `lb:0,ub:0,mandatory,reject`
	PLMNAssistanceInfoForNetShar    []byte                           `lb:3,ub:3,optional,ignore`
	NewgNBCUUEF1APID                *int64                           `lb:0,ub:4294967295,optional,reject`
	AdditionalRRMPriorityIndex      *aper.BitString                  `lb:32,ub:32,optional,ignore`
}

func (msg *DLRRCMessageTransfer) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("DLRRCMessageTransfer"), err)
		return
	}
	return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_DLRRCMessageTransfer, Criticality_PresentIgnore, ies)
}
func (msg *DLRRCMessageTransfer) toIes() (ies []F1apMessageIE, err error) {
	ies = []F1apMessageIE{}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_GNBCUUEF1APID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: aper.Integer(msg.GNBCUUEF1APID),
		}})
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_GNBDUUEF1APID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: aper.Integer(msg.GNBDUUEF1APID),
		}})
	if msg.OldgNBDUUEF1APID != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_OldgNBDUUEF1APID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 4294967295},
				ext:   false,
				Value: aper.Integer(*msg.OldgNBDUUEF1APID),
			}})
	}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_SRBID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 3},
			ext:   false,
			Value: aper.Integer(msg.SRBID),
		}})
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_ExecuteDuplication},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       msg.ExecuteDuplication,
	})
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RRCContainer},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 0, Ub: 0},
			ext:   false,
			Value: msg.RRCContainer,
		}})
	if msg.RATFrequencyPriorityInformation != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RATFrequencyPriorityInformation},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.RATFrequencyPriorityInformation,
		})
	}
	if msg.RRCDeliveryStatusRequest != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RRCDeliveryStatusRequest},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RRCDeliveryStatusRequest,
		})
	}
	if msg.UEContextNotRetrievable != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UEContextNotRetrievable},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.UEContextNotRetrievable,
		})
	}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RedirectedRRCmessage},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 0, Ub: 0},
			ext:   false,
			Value: msg.RedirectedRRCmessage,
		}})
	if msg.PLMNAssistanceInfoForNetShar != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PLMNAssistanceInfoForNetShar},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 3, Ub: 3},
				ext:   false,
				Value: msg.PLMNAssistanceInfoForNetShar,
			}})
	}
	if msg.NewgNBCUUEF1APID != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NewgNBCUUEF1APID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 4294967295},
				ext:   false,
				Value: aper.Integer(*msg.NewgNBCUUEF1APID),
			}})
	}
	if msg.AdditionalRRMPriorityIndex != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AdditionalRRMPriorityIndex},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &BITSTRING{
				c:   aper.Constraint{Lb: 32, Ub: 32},
				ext: false,
				Value: aper.BitString{
					Bytes: msg.AdditionalRRMPriorityIndex.Bytes, NumBits: msg.AdditionalRRMPriorityIndex.NumBits},
			}})
	}
	return
}
func (msg *DLRRCMessageTransfer) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("DLRRCMessageTransfer"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := DLRRCMessageTransferDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*F1apMessageIE),
	}
	if _, err = aper.ReadSequenceOf[F1apMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	if _, ok := decoder.list[ProtocolIEID_GNBCUUEF1APID]; !ok {
		err = fmt.Errorf("Mandatory field GNBCUUEF1APID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_GNBCUUEF1APID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_GNBDUUEF1APID]; !ok {
		err = fmt.Errorf("Mandatory field GNBDUUEF1APID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_GNBDUUEF1APID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_SRBID]; !ok {
		err = fmt.Errorf("Mandatory field SRBID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_SRBID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_ExecuteDuplication]; !ok {
		err = fmt.Errorf("Mandatory field ExecuteDuplication is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_ExecuteDuplication},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_RRCContainer]; !ok {
		err = fmt.Errorf("Mandatory field RRCContainer is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_RRCContainer},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_RedirectedRRCmessage]; !ok {
		err = fmt.Errorf("Mandatory field RedirectedRRCmessage is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_RedirectedRRCmessage},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type DLRRCMessageTransferDecoder struct {
	msg      *DLRRCMessageTransfer
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *DLRRCMessageTransferDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_GNBCUUEF1APID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4294967295},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read GNBCUUEF1APID", err)
			return
		}
		msg.GNBCUUEF1APID = int64(tmp.Value)
	case ProtocolIEID_GNBDUUEF1APID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4294967295},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read GNBDUUEF1APID", err)
			return
		}
		msg.GNBDUUEF1APID = int64(tmp.Value)
	case ProtocolIEID_OldgNBDUUEF1APID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4294967295},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read OldgNBDUUEF1APID", err)
			return
		}
		msg.OldgNBDUUEF1APID = (*int64)(&tmp.Value)
	case ProtocolIEID_SRBID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 3},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SRBID", err)
			return
		}
		msg.SRBID = int64(tmp.Value)
	case ProtocolIEID_ExecuteDuplication:
		var tmp ExecuteDuplication
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ExecuteDuplication", err)
			return
		}
		msg.ExecuteDuplication = &tmp
	case ProtocolIEID_RRCContainer:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RRCContainer", err)
			return
		}
		msg.RRCContainer = tmp.Value
	case ProtocolIEID_RATFrequencyPriorityInformation:
		var tmp RATFrequencyPriorityInformation
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RATFrequencyPriorityInformation", err)
			return
		}
		msg.RATFrequencyPriorityInformation = &tmp
	case ProtocolIEID_RRCDeliveryStatusRequest:
		var tmp RRCDeliveryStatusRequest
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RRCDeliveryStatusRequest", err)
			return
		}
		msg.RRCDeliveryStatusRequest = &tmp
	case ProtocolIEID_UEContextNotRetrievable:
		var tmp UEContextNotRetrievable
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read UEContextNotRetrievable", err)
			return
		}
		msg.UEContextNotRetrievable = &tmp
	case ProtocolIEID_RedirectedRRCmessage:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RedirectedRRCmessage", err)
			return
		}
		msg.RedirectedRRCmessage = tmp.Value
	case ProtocolIEID_PLMNAssistanceInfoForNetShar:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 3, Ub: 3},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read PLMNAssistanceInfoForNetShar", err)
			return
		}
		msg.PLMNAssistanceInfoForNetShar = tmp.Value
	case ProtocolIEID_NewgNBCUUEF1APID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4294967295},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read NewgNBCUUEF1APID", err)
			return
		}
		msg.NewgNBCUUEF1APID = (*int64)(&tmp.Value)
	case ProtocolIEID_AdditionalRRMPriorityIndex:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 32, Ub: 32},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read AdditionalRRMPriorityIndex", err)
			return
		}
		msg.AdditionalRRMPriorityIndex = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
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
