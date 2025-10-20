package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ECIDMeasurementInitiationRequest struct {
	GNBCUUEF1APID              int64                           `lb:0,ub:4294967295,mandatory,reject`
	GNBDUUEF1APID              int64                           `lb:0,ub:4294967295,mandatory,reject`
	LMFUEMeasurementID         int64                           `lb:1,ub:256,mandatory,reject,valueExt`
	RANUEMeasurementID         int64                           `lb:1,ub:256,mandatory,reject,valueExt`
	ECIDReportCharacteristics  ECIDReportCharacteristics       `mandatory,reject`
	ECIDMeasurementPeriodicity *MeasurementPeriodicity         `conditional,reject`
	ECIDMeasurementQuantities  []ECIDMeasurementQuantitiesItem `lb:1,ub:maxnoofMeasECID,mandatory,reject,valueExt`
}

func (msg *ECIDMeasurementInitiationRequest) Encode(w io.Writer) (err error) {
    var ies []F1apMessageIE
    if ies, err = msg.toIes(); err != nil {
        err = msgErrors(fmt.Errorf("ECIDMeasurementInitiationRequest"), err)
        return
    }
    return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_ECIDMeasurementInitiation, Criticality_PresentReject, ies)
}
func (msg *ECIDMeasurementInitiationRequest) toIes() (ies []F1apMessageIE, err error) {
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
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_LMFUEMeasurementID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 1, Ub: 256},
			ext:   true,
			Value: aper.Integer(msg.LMFUEMeasurementID),
		}})
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RANUEMeasurementID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 1, Ub: 256},
			ext:   true,
			Value: aper.Integer(msg.RANUEMeasurementID),
		}})
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_ECIDReportCharacteristics},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.ECIDReportCharacteristics,
	})
	if msg.ECIDMeasurementPeriodicity != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ECIDMeasurementPeriodicity},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.ECIDMeasurementPeriodicity,
		})
	}
	if len(msg.ECIDMeasurementQuantities) > 0 {
		tmp_ECIDMeasurementQuantities := Sequence[*ECIDMeasurementQuantitiesItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMeasECID},
			ext: true,
		}
		for _, i := range msg.ECIDMeasurementQuantities {
			tmp_ECIDMeasurementQuantities.Value = append(tmp_ECIDMeasurementQuantities.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ECIDMeasurementQuantities},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_ECIDMeasurementQuantities,
		})
	} else {
		err = utils.WrapError("ECIDMeasurementQuantities is nil", err)
		return
	}
	return
}
func (msg *ECIDMeasurementInitiationRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("ECIDMeasurementInitiationRequest"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := ECIDMeasurementInitiationRequestDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_LMFUEMeasurementID]; !ok {
		err = fmt.Errorf("Mandatory field LMFUEMeasurementID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_LMFUEMeasurementID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_RANUEMeasurementID]; !ok {
		err = fmt.Errorf("Mandatory field RANUEMeasurementID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_RANUEMeasurementID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_ECIDReportCharacteristics]; !ok {
		err = fmt.Errorf("Mandatory field ECIDReportCharacteristics is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_ECIDReportCharacteristics},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_ECIDMeasurementQuantities]; !ok {
		err = fmt.Errorf("Mandatory field ECIDMeasurementQuantities is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_ECIDMeasurementQuantities},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type ECIDMeasurementInitiationRequestDecoder struct {
	msg      *ECIDMeasurementInitiationRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *ECIDMeasurementInitiationRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_LMFUEMeasurementID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 256},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read LMFUEMeasurementID", err)
			return
		}
		msg.LMFUEMeasurementID = int64(tmp.Value)
	case ProtocolIEID_RANUEMeasurementID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 256},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RANUEMeasurementID", err)
			return
		}
		msg.RANUEMeasurementID = int64(tmp.Value)
	case ProtocolIEID_ECIDReportCharacteristics:
		var tmp ECIDReportCharacteristics
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ECIDReportCharacteristics", err)
			return
		}
		msg.ECIDReportCharacteristics = tmp
	case ProtocolIEID_ECIDMeasurementPeriodicity:
		var tmp MeasurementPeriodicity
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ECIDMeasurementPeriodicity", err)
			return
		}
		msg.ECIDMeasurementPeriodicity = &tmp
	case ProtocolIEID_ECIDMeasurementQuantities:
		tmp := Sequence[*ECIDMeasurementQuantitiesItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMeasECID},
			ext: true,
		}
		fn := func() *ECIDMeasurementQuantitiesItem { return new(ECIDMeasurementQuantitiesItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read ECIDMeasurementQuantities", err)
			return
		}
		msg.ECIDMeasurementQuantities = []ECIDMeasurementQuantitiesItem{}
		for _, i := range tmp.Value {
			msg.ECIDMeasurementQuantities = append(msg.ECIDMeasurementQuantities, *i)
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
