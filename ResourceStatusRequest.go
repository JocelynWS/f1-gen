package f1ap

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ResourceStatusRequest struct {
	TransactionID         int64                 `lb:0,ub:255,mandatory,reject`
	GNBCUMeasurementID    int64                 `lb:0,ub:4095,mandatory,reject,valueExt`
	GNBDUMeasurementID    *int64                `lb:0,ub:4095,conditional,ignore,valueExt`
	RegistrationRequest   RegistrationRequest   `mandatory,ignore`
	ReportCharacteristics *aper.BitString       `lb:32,ub:32,conditional,ignore`
	CellToReportList      []CellToReportItem    `lb:1,ub:maxCellingNBDU,optional,ignore,valueExt`
	ReportingPeriodicity  *ReportingPeriodicity `optional,ignore`
}

func (msg *ResourceStatusRequest) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("ResourceStatusRequest"), err)
		return
	}
	return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_ResourceStatusReportingInitiation, Criticality_PresentReject, ies)
}

func (msg *ResourceStatusRequest) toIes() (ies []F1apMessageIE, err error) {
	ies = []F1apMessageIE{}

	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TransactionID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 255},
			ext:   false,
			Value: aper.Integer(msg.TransactionID),
		},
	})

	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_GNBCUMeasurementID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4095},
			ext:   true,
			Value: aper.Integer(msg.GNBCUMeasurementID),
		},
	})

	if msg.GNBDUMeasurementID != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GNBDUMeasurementID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 4095},
				ext:   true,
				Value: aper.Integer(*msg.GNBDUMeasurementID),
			},
		})
	}

	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RegistrationRequest},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.RegistrationRequest,
	})

	if msg.ReportCharacteristics != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ReportCharacteristics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &BITSTRING{
				c:   aper.Constraint{Lb: 32, Ub: 32},
				ext: false,
				Value: aper.BitString{
					Bytes:   msg.ReportCharacteristics.Bytes,
					NumBits: msg.ReportCharacteristics.NumBits,
				},
			},
		})
	}

	if len(msg.CellToReportList) > 0 {
		tmp_CellToReportList := Sequence[*CellToReportItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		for _, i := range msg.CellToReportList {
			tmp_CellToReportList.Value = append(tmp_CellToReportList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CellToReportList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_CellToReportList,
		})
	}

	if msg.ReportingPeriodicity != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ReportingPeriodicity},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.ReportingPeriodicity,
		})
	}

	return
}

func (msg *ResourceStatusRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("ResourceStatusRequest"), err)
		}
	}()

	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()

	decoder := ResourceStatusRequestDecoder{
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

	if _, ok := decoder.list[ProtocolIEID_GNBCUMeasurementID]; !ok {
		err = fmt.Errorf("Mandatory field GNBCUMeasurementID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_GNBCUMeasurementID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}

	if _, ok := decoder.list[ProtocolIEID_RegistrationRequest]; !ok {
		err = fmt.Errorf("Mandatory field RegistrationRequest is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_RegistrationRequest},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}

	return
}

type ResourceStatusRequestDecoder struct {
	msg      *ResourceStatusRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *ResourceStatusRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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

	case ProtocolIEID_GNBCUMeasurementID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4095},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read GNBCUMeasurementID", err)
			return
		}
		msg.GNBCUMeasurementID = int64(tmp.Value)

	case ProtocolIEID_GNBDUMeasurementID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4095},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read GNBDUMeasurementID", err)
			return
		}
		v := int64(tmp.Value)
		msg.GNBDUMeasurementID = &v

	case ProtocolIEID_RegistrationRequest:
		var tmp RegistrationRequest
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RegistrationRequest", err)
			return
		}
		msg.RegistrationRequest = tmp

	case ProtocolIEID_ReportCharacteristics:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 32, Ub: 32},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ReportCharacteristics", err)
			return
		}
		msg.ReportCharacteristics = &aper.BitString{
			Bytes:   tmp.Value.Bytes,
			NumBits: tmp.Value.NumBits,
		}

	case ProtocolIEID_CellToReportList:
		tmp := Sequence[*CellToReportItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		fn := func() *CellToReportItem { return new(CellToReportItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read CellToReportList", err)
			return
		}
		msg.CellToReportList = []CellToReportItem{}
		for _, i := range tmp.Value {
			msg.CellToReportList = append(msg.CellToReportList, *i)
		}

	case ProtocolIEID_ReportingPeriodicity:
		var tmp ReportingPeriodicity
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ReportingPeriodicity", err)
			return
		}
		msg.ReportingPeriodicity = &tmp

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
