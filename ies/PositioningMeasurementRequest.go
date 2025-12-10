package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PositioningMeasurementRequest struct {
	TransactionID              int64                          `lb:0,ub:255,mandatory,reject,valueExt`
	LMFMeasurementID           int64                          `lb:1,ub:65536,mandatory,reject,valueExt`
	RANMeasurementID           int64                          `lb:1,ub:65536,mandatory,reject,valueExt`
	TRPMeasurementRequestList  []TRPMeasurementRequestItem    `lb:1,ub:maxnoofTRPs,mandatory,reject,valueExt`
	PosReportCharacteristics   *PosReportCharacteristics      `mandatory,reject`
	PosMeasurementPeriodicity  *PosMeasurementPeriodicity     `conditional,reject`
	PosMeasurementQuantities   []PosMeasurementQuantitiesItem `lb:1,ub:maxnoofPosMeas,mandatory,reject,valueExt`
	SFNInitialisationTime      *aper.BitString                `lb:64,ub:64,optional,ignore`
	SRSConfiguration           *SRSConfiguration              `optional,ignore`
	MeasurementBeamInfoRequest *MeasurementBeamInfoRequest    `optional,ignore`
	SystemFrameNumber          *int64                         `lb:0,ub:1023,optional,ignore`
	SlotNumber                 *int64                         `lb:0,ub:79,optional,ignore`
}

func (msg *PositioningMeasurementRequest) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("PositioningMeasurementRequest"), err)
		return
	}
	return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_PositioningMeasurementExchange, Criticality_PresentReject, ies)
}

func (msg *PositioningMeasurementRequest) toIes() (ies []F1apMessageIE, err error) {
	ies = []F1apMessageIE{}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TransactionID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 255},
			ext:   true,
			Value: aper.Integer(msg.TransactionID),
		}})
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_LMFMeasurementID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 1, Ub: 65536},
			ext:   true,
			Value: aper.Integer(msg.LMFMeasurementID),
		}})
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RANMeasurementID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 1, Ub: 65536},
			ext:   true,
			Value: aper.Integer(msg.RANMeasurementID),
		}})
	if len(msg.TRPMeasurementRequestList) > 0 {
		tmp_TRPMeasurementRequestList := Sequence[*TRPMeasurementRequestItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTRPs},
			ext: true,
		}
		for _, i := range msg.TRPMeasurementRequestList {
			tmp_TRPMeasurementRequestList.Value = append(tmp_TRPMeasurementRequestList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TRPMeasurementRequestList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_TRPMeasurementRequestList,
		})
	} else {
		err = utils.WrapError("TRPMeasurementRequestList is nil", err)
		return
	}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_PosReportCharacteristics},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       msg.PosReportCharacteristics,
	})
	if msg.PosMeasurementPeriodicity != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PosMeasurementPeriodicity},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.PosMeasurementPeriodicity,
		})
	}
	if len(msg.PosMeasurementQuantities) > 0 {
		tmp_PosMeasurementQuantities := Sequence[*PosMeasurementQuantitiesItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPosMeas},
			ext: true,
		}
		for _, i := range msg.PosMeasurementQuantities {
			tmp_PosMeasurementQuantities.Value = append(tmp_PosMeasurementQuantities.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PosMeasurementQuantities},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_PosMeasurementQuantities,
		})
	} else {
		err = utils.WrapError("PosMeasurementQuantities is nil", err)
		return
	}
	// Optional fields - only encode if present
	if msg.SFNInitialisationTime != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SFNInitialisationTime},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &BITSTRING{
				c:   aper.Constraint{Lb: 64, Ub: 64},
				ext: false,
				Value: aper.BitString{
					Bytes: msg.SFNInitialisationTime.Bytes, NumBits: msg.SFNInitialisationTime.NumBits},
			}})
	}
	if msg.SRSConfiguration != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SRSConfiguration},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.SRSConfiguration,
		})
	}
	if msg.MeasurementBeamInfoRequest != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_MeasurementBeamInfoRequest},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.MeasurementBeamInfoRequest,
		})
	}
	if msg.SystemFrameNumber != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SystemFrameNumber},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 1023},
				ext:   false,
				Value: aper.Integer(*msg.SystemFrameNumber),
			}})
	}
	if msg.SlotNumber != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SlotNumber},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 79},
				ext:   false,
				Value: aper.Integer(*msg.SlotNumber),
			}})
	}
	return
}

func (msg *PositioningMeasurementRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("PositioningMeasurementRequest"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := PositioningMeasurementRequestDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*F1apMessageIE),
	}
	if _, err = aper.ReadSequenceOf[F1apMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	// Check only mandatory fields
	if _, ok := decoder.list[ProtocolIEID_TransactionID]; !ok {
		err = fmt.Errorf("Mandatory field TransactionID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_TransactionID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_LMFMeasurementID]; !ok {
		err = fmt.Errorf("Mandatory field LMFMeasurementID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_LMFMeasurementID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_RANMeasurementID]; !ok {
		err = fmt.Errorf("Mandatory field RANMeasurementID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_RANMeasurementID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_TRPMeasurementRequestList]; !ok {
		err = fmt.Errorf("Mandatory field TRPMeasurementRequestList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_TRPMeasurementRequestList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_PosReportCharacteristics]; !ok {
		err = fmt.Errorf("Mandatory field PosReportCharacteristics is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_PosReportCharacteristics},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_PosMeasurementQuantities]; !ok {
		err = fmt.Errorf("Mandatory field PosMeasurementQuantities is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_PosMeasurementQuantities},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	// SFNInitialisationTime, SRSConfiguration, MeasurementBeamInfoRequest,
	// SystemFrameNumber, SlotNumber are optional - no check needed
	// PosMeasurementPeriodicity is conditional - no check here
	diagList = decoder.diagList
	return
}

type PositioningMeasurementRequestDecoder struct {
	msg      *PositioningMeasurementRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *PositioningMeasurementRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_LMFMeasurementID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 65536},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read LMFMeasurementID", err)
			return
		}
		msg.LMFMeasurementID = int64(tmp.Value)
	case ProtocolIEID_RANMeasurementID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 65536},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RANMeasurementID", err)
			return
		}
		msg.RANMeasurementID = int64(tmp.Value)
	case ProtocolIEID_TRPMeasurementRequestList:
		tmp := Sequence[*TRPMeasurementRequestItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTRPs},
			ext: true,
		}
		fn := func() *TRPMeasurementRequestItem { return new(TRPMeasurementRequestItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read TRPMeasurementRequestList", err)
			return
		}
		msg.TRPMeasurementRequestList = []TRPMeasurementRequestItem{}
		for _, i := range tmp.Value {
			msg.TRPMeasurementRequestList = append(msg.TRPMeasurementRequestList, *i)
		}
	case ProtocolIEID_PosReportCharacteristics:
		var tmp PosReportCharacteristics
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read PosReportCharacteristics", err)
			return
		}
		msg.PosReportCharacteristics = &tmp
	case ProtocolIEID_PosMeasurementPeriodicity:
		var tmp PosMeasurementPeriodicity
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read PosMeasurementPeriodicity", err)
			return
		}
		msg.PosMeasurementPeriodicity = &tmp
	case ProtocolIEID_PosMeasurementQuantities:
		tmp := Sequence[*PosMeasurementQuantitiesItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPosMeas},
			ext: true,
		}
		fn := func() *PosMeasurementQuantitiesItem { return new(PosMeasurementQuantitiesItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read PosMeasurementQuantities", err)
			return
		}
		msg.PosMeasurementQuantities = []PosMeasurementQuantitiesItem{}
		for _, i := range tmp.Value {
			msg.PosMeasurementQuantities = append(msg.PosMeasurementQuantities, *i)
		}
	case ProtocolIEID_SFNInitialisationTime:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 64, Ub: 64},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SFNInitialisationTime", err)
			return
		}
		msg.SFNInitialisationTime = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	case ProtocolIEID_SRSConfiguration:
		var tmp SRSConfiguration
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SRSConfiguration", err)
			return
		}
		msg.SRSConfiguration = &tmp
	case ProtocolIEID_MeasurementBeamInfoRequest:
		var tmp MeasurementBeamInfoRequest
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read MeasurementBeamInfoRequest", err)
			return
		}
		msg.MeasurementBeamInfoRequest = &tmp
	case ProtocolIEID_SystemFrameNumber:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 1023},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SystemFrameNumber", err)
			return
		}
		val := int64(tmp.Value)
		msg.SystemFrameNumber = &val
	case ProtocolIEID_SlotNumber:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 79},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SlotNumber", err)
			return
		}
		val := int64(tmp.Value)
		msg.SlotNumber = &val
	default:
		switch msgIe.Criticality.Value {
		case Criticality_PresentReject:
			err = fmt.Errorf("Not comprehended IE ID 0x%04x (criticality: reject)", msgIe.Id.Value)
		case Criticality_PresentIgnore:
			// Just log, don't return error for ignore criticality
		case Criticality_PresentNotify:
			err = fmt.Errorf("Not comprehended IE ID 0x%04x (criticality: notify)", msgIe.Id.Value)
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
