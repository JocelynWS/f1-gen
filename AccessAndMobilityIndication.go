package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AccessAndMobilityIndication struct {
	TransactionID             int64                       `lb:0,ub:255,mandatory,reject`
	RACHReportInformationList []RACHReportInformationItem `lb:1,ub:maxnoofRACHReports,optional,ignore,valueExt`
	RLFReportInformationList  []RLFReportInformationItem  `lb:1,ub:maxnoofRLFReports,optional,ignore,valueExt`
}

func (msg *AccessAndMobilityIndication) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("AccessAndMobilityIndication"), err)
		return
	}
	return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_F1Setup, Criticality_PresentReject, ies)
}
func (msg *AccessAndMobilityIndication) toIes() (ies []F1apMessageIE, err error) {
	ies = []F1apMessageIE{}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TransactionID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 255},
			ext:   false,
			Value: aper.Integer(msg.TransactionID),
		}})
	if len(msg.RACHReportInformationList) > 0 {
		tmp_RACHReportInformationList := Sequence[*RACHReportInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofRACHReports},
			ext: true,
		}
		for _, i := range msg.RACHReportInformationList {
			tmp_RACHReportInformationList.Value = append(tmp_RACHReportInformationList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RACHReportInformationList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_RACHReportInformationList,
		})
	}
	if len(msg.RLFReportInformationList) > 0 {
		tmp_RLFReportInformationList := Sequence[*RLFReportInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofRLFReports},
			ext: true,
		}
		for _, i := range msg.RLFReportInformationList {
			tmp_RLFReportInformationList.Value = append(tmp_RLFReportInformationList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RLFReportInformationList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_RLFReportInformationList,
		})
	}
	return
}
func (msg *AccessAndMobilityIndication) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("AccessAndMobilityIndication"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := AccessAndMobilityIndicationDecoder{
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

type AccessAndMobilityIndicationDecoder struct {
	msg      *AccessAndMobilityIndication
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *AccessAndMobilityIndicationDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_RACHReportInformationList:
		tmp := Sequence[*RACHReportInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofRACHReports},
			ext: true,
		}
		fn := func() *RACHReportInformationItem { return new(RACHReportInformationItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read RACHReportInformationList", err)
			return
		}
		msg.RACHReportInformationList = []RACHReportInformationItem{}
		for _, i := range tmp.Value {
			msg.RACHReportInformationList = append(msg.RACHReportInformationList, *i)
		}
	case ProtocolIEID_RLFReportInformationList:
		tmp := Sequence[*RLFReportInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofRLFReports},
			ext: true,
		}
		fn := func() *RLFReportInformationItem { return new(RLFReportInformationItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read RLFReportInformationList", err)
			return
		}
		msg.RLFReportInformationList = []RLFReportInformationItem{}
		for _, i := range tmp.Value {
			msg.RLFReportInformationList = append(msg.RLFReportInformationList, *i)
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
