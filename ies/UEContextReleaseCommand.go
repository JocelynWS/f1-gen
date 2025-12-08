package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UEContextReleaseCommand struct {
	GNBCUUEF1APID            int64                     `lb:0,ub:4294967295,mandatory,reject`
	GNBDUUEF1APID            int64                     `lb:0,ub:4294967295,mandatory,reject`
	Cause                    Cause                     `mandatory,ignore`
	RRCContainer             []byte                    `lb:0,ub:0,optional,ignore`
	SRBID                    *int64                    `lb:0,ub:3,conditional,ignore`
	OldgNBDUUEF1APID         *int64                    `lb:0,ub:4294967295,optional,ignore`
	ExecuteDuplication       *ExecuteDuplication       `optional,ignore`
	RRCDeliveryStatusRequest *RRCDeliveryStatusRequest `optional,ignore`
	TargetCellsToCancel      []TargetCellListItem      `lb:1,ub:maxnoofCHOcells,optional,reject`
}

func (msg *UEContextReleaseCommand) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("UEContextReleaseCommand"), err)
		return
	}
	return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_UEContextRelease, Criticality_PresentReject, ies)
}

func (msg *UEContextReleaseCommand) toIes() (ies []F1apMessageIE, err error) {
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
		Id:          ProtocolIEID{Value: ProtocolIEID_Cause},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.Cause,
	})

	if msg.RRCContainer != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RRCContainer},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   false,
				Value: msg.RRCContainer,
			}})
	}

	if msg.SRBID != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SRBID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 3},
				ext:   false,
				Value: aper.Integer(*msg.SRBID),
			}})
	}

	if msg.OldgNBDUUEF1APID != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_OldgNBDUUEF1APID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 4294967295},
				ext:   false,
				Value: aper.Integer(*msg.OldgNBDUUEF1APID),
			}})
	}

	if msg.ExecuteDuplication != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ExecuteDuplication},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.ExecuteDuplication,
		})
	}

	if msg.RRCDeliveryStatusRequest != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RRCDeliveryStatusRequest},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RRCDeliveryStatusRequest,
		})
	}

	if len(msg.TargetCellsToCancel) > 0 {
		tmp_TargetCellsToCancel := Sequence[*TargetCellListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCHOcells},
			ext: false,
		}
		for _, i := range msg.TargetCellsToCancel {
			tmp_TargetCellsToCancel.Value = append(tmp_TargetCellsToCancel.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TargetCellsToCancel},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_TargetCellsToCancel,
		})
	}
	return
}

func (msg *UEContextReleaseCommand) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("UEContextReleaseCommand"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := UEContextReleaseCommandDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*F1apMessageIE),
	}
	if _, err = aper.ReadSequenceOf[F1apMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	// Check only mandatory fields
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
	if _, ok := decoder.list[ProtocolIEID_Cause]; !ok {
		err = fmt.Errorf("Mandatory field Cause is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_Cause},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	// All other fields are optional or conditional
	diagList = decoder.diagList
	return
}

type UEContextReleaseCommandDecoder struct {
	msg      *UEContextReleaseCommand
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *UEContextReleaseCommandDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_Cause:
		var tmp Cause
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read Cause", err)
			return
		}
		msg.Cause = tmp
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
	case ProtocolIEID_SRBID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 3},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SRBID", err)
			return
		}
		val := int64(tmp.Value)
		msg.SRBID = &val
	case ProtocolIEID_OldgNBDUUEF1APID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4294967295},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read OldgNBDUUEF1APID", err)
			return
		}
		val := int64(tmp.Value)
		msg.OldgNBDUUEF1APID = &val
	case ProtocolIEID_ExecuteDuplication:
		var tmp ExecuteDuplication
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ExecuteDuplication", err)
			return
		}
		msg.ExecuteDuplication = &tmp
	case ProtocolIEID_RRCDeliveryStatusRequest:
		var tmp RRCDeliveryStatusRequest
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RRCDeliveryStatusRequest", err)
			return
		}
		msg.RRCDeliveryStatusRequest = &tmp
	case ProtocolIEID_TargetCellsToCancel:
		tmp := Sequence[*TargetCellListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCHOcells},
			ext: false,
		}
		fn := func() *TargetCellListItem { return new(TargetCellListItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read TargetCellsToCancel", err)
			return
		}
		msg.TargetCellsToCancel = []TargetCellListItem{}
		for _, i := range tmp.Value {
			msg.TargetCellsToCancel = append(msg.TargetCellsToCancel, *i)
		}
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