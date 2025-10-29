package f1ap

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PositioningMeasurementResponse struct {
	TransactionID            int64                          `lb:0,ub:255,mandatory,reject`
	LMFMeasurementID         int64                          `lb:1,ub:65536,mandatory,reject,valueExt`
	RANMeasurementID         int64                          `lb:1,ub:65536,mandatory,reject,valueExt`
	PosMeasurementResultList []PosMeasurementResultListItem `lb:1,ub:maxNoOfMeasTRPs,optional,reject,valueExt`
	CriticalityDiagnostics   *CriticalityDiagnostics        `optional,ignore`
}

func (msg *PositioningMeasurementResponse) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("PositioningMeasurementResponse"), err)
		return
	}
	return encodeMessage(w, F1apPduSuccessfulOutcome, ProcedureCode_PositioningMeasurementExchange, Criticality_PresentReject, ies)
}
func (msg *PositioningMeasurementResponse) toIes() (ies []F1apMessageIE, err error) {
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
	if len(msg.PosMeasurementResultList) > 0 {
		tmp_PosMeasurementResultList := Sequence[*PosMeasurementResultListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxNoOfMeasTRPs},
			ext: true,
		}
		for _, i := range msg.PosMeasurementResultList {
			tmp_PosMeasurementResultList.Value = append(tmp_PosMeasurementResultList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PosMeasurementResultList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_PosMeasurementResultList,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	return
}
func (msg *PositioningMeasurementResponse) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("PositioningMeasurementResponse"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := PositioningMeasurementResponseDecoder{
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
	return
}

type PositioningMeasurementResponseDecoder struct {
	msg      *PositioningMeasurementResponse
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *PositioningMeasurementResponseDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_PosMeasurementResultList:
		tmp := Sequence[*PosMeasurementResultListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxNoOfMeasTRPs},
			ext: true,
		}
		fn := func() *PosMeasurementResultListItem { return new(PosMeasurementResultListItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read PosMeasurementResultList", err)
			return
		}
		msg.PosMeasurementResultList = []PosMeasurementResultListItem{}
		for _, i := range tmp.Value {
			msg.PosMeasurementResultList = append(msg.PosMeasurementResultList, *i)
		}
	case ProtocolIEID_CriticalityDiagnostics:
		var tmp CriticalityDiagnostics
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read CriticalityDiagnostics", err)
			return
		}
		msg.CriticalityDiagnostics = &tmp
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
