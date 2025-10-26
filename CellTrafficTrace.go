package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type CellTrafficTrace struct {
	GNBCUUEF1APID                  int64                 `lb:0,ub:4294967295,mandatory,reject`
	GNBDUUEF1APID                  int64                 `lb:0,ub:4294967295,mandatory,reject`
	TraceID                        TraceID               `mandatory,ignore`
	TraceCollectionEntityIPAddress TransportLayerAddress `mandatory,ignore`
	PrivacyIndicator               *PrivacyIndicator     `optional,ignore`
	TraceCollectionEntityURI       []byte                `optional,ignore`
}

func (msg *CellTrafficTrace) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("CellTrafficTrace"), err)
		return
	}
	return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_CellTrafficTrace, Criticality_PresentIgnore, ies)
}

func (msg *CellTrafficTrace) toIes() (ies []F1apMessageIE, err error) {
	ies = []F1apMessageIE{}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_GNBCUUEF1APID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: aper.Integer(msg.GNBCUUEF1APID),
		},
	})
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_GNBDUUEF1APID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: aper.Integer(msg.GNBDUUEF1APID),
		},
	})
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TraceID},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.TraceID,
	})
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TraceCollectionEntityIPAddress},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.TraceCollectionEntityIPAddress,
	})
	if msg.PrivacyIndicator != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PrivacyIndicator},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.PrivacyIndicator,
		})
	}
	if msg.TraceCollectionEntityURI != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TraceCollectionEntityURI},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				Value: msg.TraceCollectionEntityURI,
				c:     aper.Constraint{Lb: 0, Ub: 255},
				ext:   false,
			},
		})
	}

	return
}

func (msg *CellTrafficTrace) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("CellTrafficTrace"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := CellTrafficTraceDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*F1apMessageIE),
	}
	if _, err = aper.ReadSequenceOf[F1apMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}

	mandatoryFields := []struct {
		id       aper.Integer
		name     string
		critical Criticality
	}{
		{ProtocolIEID_GNBCUUEF1APID, "GNBCUUEF1APID", Criticality{Value: Criticality_PresentReject}},
		{ProtocolIEID_GNBDUUEF1APID, "GNBDUUEF1APID", Criticality{Value: Criticality_PresentReject}},
		{ProtocolIEID_TraceID, "TraceID", Criticality{Value: Criticality_PresentIgnore}},
		{ProtocolIEID_TraceCollectionEntityIPAddress, "TraceCollectionEntityIPAddress", Criticality{Value: Criticality_PresentIgnore}},
	}

	for _, f := range mandatoryFields {
		if _, ok := decoder.list[f.id]; !ok {
			err = fmt.Errorf("Mandatory field %s is missing", f.name)
			decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
				IECriticality: f.critical,
				IEID:          ProtocolIEID{Value: f.id},
				TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
			})
			return
		}
	}

	return
}

type CellTrafficTraceDecoder struct {
	msg      *CellTrafficTrace
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *CellTrafficTraceDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
	var id int64
	var c uint64
	var buf []byte

	if id, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
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
		tmp := INTEGER{c: aper.Constraint{Lb: 0, Ub: 4294967295}, ext: false}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read GNBCUUEF1APID", err)
			return
		}
		msg.GNBCUUEF1APID = int64(tmp.Value)

	case ProtocolIEID_GNBDUUEF1APID:
		tmp := INTEGER{c: aper.Constraint{Lb: 0, Ub: 4294967295}, ext: false}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read GNBDUUEF1APID", err)
			return
		}
		msg.GNBDUUEF1APID = int64(tmp.Value)

	case ProtocolIEID_TraceID:
		var tmp TraceID
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read TraceID", err)
			return
		}
		msg.TraceID = tmp

	case ProtocolIEID_TraceCollectionEntityIPAddress:
		var tmp TransportLayerAddress
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read TraceCollectionEntityIPAddress", err)
			return
		}
		msg.TraceCollectionEntityIPAddress = tmp

	case ProtocolIEID_PrivacyIndicator:
		var tmp PrivacyIndicator
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read PrivacyIndicator", err)
			return
		}
		msg.PrivacyIndicator = &tmp

	case ProtocolIEID_TraceCollectionEntityURI:
		var tmp OCTETSTRING
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read TraceCollectionEntityURI", err)
			return
		}
		msg.TraceCollectionEntityURI = tmp.Value

	default:
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
