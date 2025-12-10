package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ResourceStatusUpdate struct {
	TransactionID             int64                       `lb:0,ub:255,mandatory,reject,valueExt`
	GNBCUMeasurementID        int64                       `lb:0,ub:4095,mandatory,reject,valueExt`
	GNBDUMeasurementID        int64                       `lb:0,ub:4095,mandatory,ignore,valueExt`
	HardwareLoadIndicator     *HardwareLoadIndicator      `optional,ignore`
	TNLCapacityIndicator      *TNLCapacityIndicator       `optional,ignore`
	CellMeasurementResultList []CellMeasurementResultItem `lb:1,ub:maxCellingNBDU,optional,ignore,valueExt`
}

func (msg *ResourceStatusUpdate) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("ResourceStatusUpdate"), err)
		return
	}
	return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_ResourceStatusReporting, Criticality_PresentIgnore, ies)
}
func (msg *ResourceStatusUpdate) toIes() (ies []F1apMessageIE, err error) {
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
		Id:          ProtocolIEID{Value: ProtocolIEID_GNBCUMeasurementID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4095},
			ext:   true,
			Value: aper.Integer(msg.GNBCUMeasurementID),
		}})
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_GNBDUMeasurementID},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4095},
			ext:   true,
			Value: aper.Integer(msg.GNBDUMeasurementID),
		}})
	if msg.HardwareLoadIndicator != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_HardwareLoadIndicator},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.HardwareLoadIndicator,
		})
	}
	if msg.TNLCapacityIndicator != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TNLCapacityIndicator},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.TNLCapacityIndicator,
		})
	}
	if len(msg.CellMeasurementResultList) > 0 {
		tmp_CellMeasurementResultList := Sequence[*CellMeasurementResultItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		for _, i := range msg.CellMeasurementResultList {
			tmp_CellMeasurementResultList.Value = append(tmp_CellMeasurementResultList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CellMeasurementResultList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_CellMeasurementResultList,
		})
	}
	return
}
func (msg *ResourceStatusUpdate) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("ResourceStatusUpdate"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := ResourceStatusUpdateDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_GNBDUMeasurementID]; !ok {
		err = fmt.Errorf("Mandatory field GNBDUMeasurementID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_GNBDUMeasurementID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type ResourceStatusUpdateDecoder struct {
	msg      *ResourceStatusUpdate
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *ResourceStatusUpdateDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
		msg.GNBDUMeasurementID = int64(tmp.Value)
	case ProtocolIEID_HardwareLoadIndicator:
		var tmp HardwareLoadIndicator
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read HardwareLoadIndicator", err)
			return
		}
		msg.HardwareLoadIndicator = &tmp
	case ProtocolIEID_TNLCapacityIndicator:
		var tmp TNLCapacityIndicator
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read TNLCapacityIndicator", err)
			return
		}
		msg.TNLCapacityIndicator = &tmp
	case ProtocolIEID_CellMeasurementResultList:
		tmp := Sequence[*CellMeasurementResultItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext: true,
		}
		fn := func() *CellMeasurementResultItem { return new(CellMeasurementResultItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read CellMeasurementResultList", err)
			return
		}
		msg.CellMeasurementResultList = []CellMeasurementResultItem{}
		for _, i := range tmp.Value {
			msg.CellMeasurementResultList = append(msg.CellMeasurementResultList, *i)
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
