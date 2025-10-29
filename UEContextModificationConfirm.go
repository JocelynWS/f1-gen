package f1ap

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UEContextModificationConfirm struct {
	GNBCUUEF1APID                           int64                                    `lb:0,ub:4294967295,mandatory,reject`
	GNBDUUEF1APID                           int64                                    `lb:0,ub:4294967295,mandatory,reject`
	ResourceCoordinationTransferContainer   []byte                                   `lb:0,ub:0,optional,ignore`
	DRBsModifiedConfList                    []DRBsModifiedConfItem                   `lb:1,ub:maxnoofDRBs,mandatory,ignore`
	RRCContainer                            []byte                                   `lb:0,ub:0,optional,ignore`
	CriticalityDiagnostics                  *CriticalityDiagnostics                  `optional,ignore`
	ExecuteDuplication                      *ExecuteDuplication                      `mandatory,ignore`
	ResourceCoordinationTransferInformation *ResourceCoordinationTransferInformation `optional,ignore`
	SLDRBsModifiedConfList                  []SLDRBsModifiedConfItem                 `lb:1,ub:maxnoofSLDRBs,mandatory,ignore`
}

func (msg *UEContextModificationConfirm) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("UEContextModificationConfirm"), err)
		return
	}
	return encodeMessage(w, F1apPduSuccessfulOutcome, ProcedureCode_UEContextModification, Criticality_PresentReject, ies)
}
func (msg *UEContextModificationConfirm) toIes() (ies []F1apMessageIE, err error) {
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
	if msg.ResourceCoordinationTransferContainer != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ResourceCoordinationTransferContainer},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   false,
				Value: msg.ResourceCoordinationTransferContainer,
			}})
	}
	if len(msg.DRBsModifiedConfList) > 0 {
		tmp_DRBsModifiedConfList := Sequence[*DRBsModifiedConfItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		for _, i := range msg.DRBsModifiedConfList {
			tmp_DRBsModifiedConfList.Value = append(tmp_DRBsModifiedConfList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DRBsModifiedConfList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_DRBsModifiedConfList,
		})
	} else {
		err = utils.WrapError("DRBsModifiedConfList is nil", err)
		return
	}
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
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_ExecuteDuplication},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       msg.ExecuteDuplication,
	})
	if msg.ResourceCoordinationTransferInformation != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ResourceCoordinationTransferInformation},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.ResourceCoordinationTransferInformation,
		})
	}
	if len(msg.SLDRBsModifiedConfList) > 0 {
		tmp_SLDRBsModifiedConfList := Sequence[*SLDRBsModifiedConfItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		for _, i := range msg.SLDRBsModifiedConfList {
			tmp_SLDRBsModifiedConfList.Value = append(tmp_SLDRBsModifiedConfList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SLDRBsModifiedConfList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_SLDRBsModifiedConfList,
		})
	} else {
		err = utils.WrapError("SLDRBsModifiedConfList is nil", err)
		return
	}
	return
}
func (msg *UEContextModificationConfirm) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("UEContextModificationConfirm"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := UEContextModificationConfirmDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_DRBsModifiedConfList]; !ok {
		err = fmt.Errorf("Mandatory field DRBsModifiedConfList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_DRBsModifiedConfList},
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
	if _, ok := decoder.list[ProtocolIEID_SLDRBsModifiedConfList]; !ok {
		err = fmt.Errorf("Mandatory field SLDRBsModifiedConfList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_SLDRBsModifiedConfList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type UEContextModificationConfirmDecoder struct {
	msg      *UEContextModificationConfirm
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *UEContextModificationConfirmDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_ResourceCoordinationTransferContainer:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ResourceCoordinationTransferContainer", err)
			return
		}
		msg.ResourceCoordinationTransferContainer = tmp.Value
	case ProtocolIEID_DRBsModifiedConfList:
		tmp := Sequence[*DRBsModifiedConfItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		fn := func() *DRBsModifiedConfItem { return new(DRBsModifiedConfItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read DRBsModifiedConfList", err)
			return
		}
		msg.DRBsModifiedConfList = []DRBsModifiedConfItem{}
		for _, i := range tmp.Value {
			msg.DRBsModifiedConfList = append(msg.DRBsModifiedConfList, *i)
		}
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
	case ProtocolIEID_CriticalityDiagnostics:
		var tmp CriticalityDiagnostics
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read CriticalityDiagnostics", err)
			return
		}
		msg.CriticalityDiagnostics = &tmp
	case ProtocolIEID_ExecuteDuplication:
		var tmp ExecuteDuplication
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ExecuteDuplication", err)
			return
		}
		msg.ExecuteDuplication = &tmp
	case ProtocolIEID_ResourceCoordinationTransferInformation:
		var tmp ResourceCoordinationTransferInformation
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ResourceCoordinationTransferInformation", err)
			return
		}
		msg.ResourceCoordinationTransferInformation = &tmp
	case ProtocolIEID_SLDRBsModifiedConfList:
		tmp := Sequence[*SLDRBsModifiedConfItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		fn := func() *SLDRBsModifiedConfItem { return new(SLDRBsModifiedConfItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SLDRBsModifiedConfList", err)
			return
		}
		msg.SLDRBsModifiedConfList = []SLDRBsModifiedConfItem{}
		for _, i := range tmp.Value {
			msg.SLDRBsModifiedConfList = append(msg.SLDRBsModifiedConfList, *i)
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
