package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UEContextSetupResponse struct {
	GNBCUUEF1APID                         int64                           `lb:0,ub:4294967295,mandatory,reject`
	GNBDUUEF1APID                         int64                           `lb:0,ub:4294967295,mandatory,reject`
	DUtoCURRCInformation                  DUtoCURRCInformation            `mandatory,reject`
	CRNTI                                 *int64                          `lb:0,ub:65535,optional,ignore,valueExt`
	ResourceCoordinationTransferContainer []byte                          `lb:0,ub:0,optional,ignore`
	FullConfiguration                     *FullConfiguration              `optional,reject`
	DRBsSetupList                         []DRBsSetupItem                 `lb:1,ub:maxnoofDRBs,optional,ignore`
	SRBsFailedToBeSetupList               []SRBsFailedToBeSetupItem       `lb:1,ub:maxnoofSRBs,optional,ignore`
	DRBsFailedToBeSetupList               []DRBsFailedToBeSetupItem       `lb:1,ub:maxnoofDRBs,optional,ignore`
	SCellFailedtoSetupList                []SCellFailedtoSetupItem        `lb:1,ub:maxnoofSCells,optional,ignore`
	InactivityMonitoringResponse          *InactivityMonitoringResponse   `optional,reject`
	CriticalityDiagnostics                *CriticalityDiagnostics         `optional,ignore`
	SRBsSetupList                         []SRBsSetupItem                 `lb:1,ub:maxnoofSRBs,optional,ignore`
	BHChannelsSetupList                   []BHChannelsSetupItem           `lb:1,ub:maxnoofBHRLCChannels,optional,ignore`
	BHChannelsFailedToBeSetupList         []BHChannelsFailedToBeSetupItem `lb:1,ub:maxnoofBHRLCChannels,optional,ignore`
	SLDRBsSetupList                       []SLDRBsSetupItem               `lb:1,ub:maxnoofSLDRBs,optional,ignore`
	SLDRBsFailedToBeSetupList             []SLDRBsFailedToBeSetupItem     `lb:1,ub:maxnoofSLDRBs,optional,ignore`
	RequestedTargetCellGlobalID           *NRCGI                          `optional,reject`
}

func (msg *UEContextSetupResponse) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("UEContextSetupResponse"), err)
		return
	}
	return encodeMessage(w, F1apPduSuccessfulOutcome, ProcedureCode_UEContextSetup, Criticality_PresentReject, ies)
}
func (msg *UEContextSetupResponse) toIes() (ies []F1apMessageIE, err error) {
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
		Id:          ProtocolIEID{Value: ProtocolIEID_DUtoCURRCInformation},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.DUtoCURRCInformation,
	})
	if msg.CRNTI != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CRNTI},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 65535},
				ext:   true,
				Value: aper.Integer(*msg.CRNTI),
			}})
	}
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
	if msg.FullConfiguration != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_FullConfiguration},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.FullConfiguration,
		})
	}
	if len(msg.DRBsSetupList) > 0 {
		tmp_DRBsSetupList := Sequence[*DRBsSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		for _, i := range msg.DRBsSetupList {
			tmp_DRBsSetupList.Value = append(tmp_DRBsSetupList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DRBsSetupList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_DRBsSetupList,
		})
	}
	if len(msg.SRBsFailedToBeSetupList) > 0 {
		tmp_SRBsFailedToBeSetupList := Sequence[*SRBsFailedToBeSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSRBs},
			ext: false,
		}
		for _, i := range msg.SRBsFailedToBeSetupList {
			tmp_SRBsFailedToBeSetupList.Value = append(tmp_SRBsFailedToBeSetupList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SRBsFailedToBeSetupList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_SRBsFailedToBeSetupList,
		})
	}
	if len(msg.DRBsFailedToBeSetupList) > 0 {
		tmp_DRBsFailedToBeSetupList := Sequence[*DRBsFailedToBeSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		for _, i := range msg.DRBsFailedToBeSetupList {
			tmp_DRBsFailedToBeSetupList.Value = append(tmp_DRBsFailedToBeSetupList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DRBsFailedToBeSetupList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_DRBsFailedToBeSetupList,
		})
	}
	if len(msg.SCellFailedtoSetupList) > 0 {
		tmp_SCellFailedtoSetupList := Sequence[*SCellFailedtoSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSCells},
			ext: false,
		}
		for _, i := range msg.SCellFailedtoSetupList {
			tmp_SCellFailedtoSetupList.Value = append(tmp_SCellFailedtoSetupList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SCellFailedtoSetupList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_SCellFailedtoSetupList,
		})
	}
	if msg.InactivityMonitoringResponse != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_InactivityMonitoringResponse},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.InactivityMonitoringResponse,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if len(msg.SRBsSetupList) > 0 {
		tmp_SRBsSetupList := Sequence[*SRBsSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSRBs},
			ext: false,
		}
		for _, i := range msg.SRBsSetupList {
			tmp_SRBsSetupList.Value = append(tmp_SRBsSetupList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SRBsSetupList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_SRBsSetupList,
		})
	}
	if len(msg.BHChannelsSetupList) > 0 {
		tmp_BHChannelsSetupList := Sequence[*BHChannelsSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofBHRLCChannels},
			ext: false,
		}
		for _, i := range msg.BHChannelsSetupList {
			tmp_BHChannelsSetupList.Value = append(tmp_BHChannelsSetupList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_BHChannelsSetupList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_BHChannelsSetupList,
		})
	}
	if len(msg.BHChannelsFailedToBeSetupList) > 0 {
		tmp_BHChannelsFailedToBeSetupList := Sequence[*BHChannelsFailedToBeSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofBHRLCChannels},
			ext: false,
		}
		for _, i := range msg.BHChannelsFailedToBeSetupList {
			tmp_BHChannelsFailedToBeSetupList.Value = append(tmp_BHChannelsFailedToBeSetupList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_BHChannelsFailedToBeSetupList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_BHChannelsFailedToBeSetupList,
		})
	}
	if len(msg.SLDRBsSetupList) > 0 {
		tmp_SLDRBsSetupList := Sequence[*SLDRBsSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		for _, i := range msg.SLDRBsSetupList {
			tmp_SLDRBsSetupList.Value = append(tmp_SLDRBsSetupList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SLDRBsSetupList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_SLDRBsSetupList,
		})
	}
	if len(msg.SLDRBsFailedToBeSetupList) > 0 {
		tmp_SLDRBsFailedToBeSetupList := Sequence[*SLDRBsFailedToBeSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		for _, i := range msg.SLDRBsFailedToBeSetupList {
			tmp_SLDRBsFailedToBeSetupList.Value = append(tmp_SLDRBsFailedToBeSetupList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SLDRBsFailedToBeSetupList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_SLDRBsFailedToBeSetupList,
		})
	}
	// FIX: RequestedTargetCellGlobalID is optional, not mandatory
	if msg.RequestedTargetCellGlobalID != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RequestedTargetCellGlobalID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.RequestedTargetCellGlobalID,
		})
	}
	return
}
func (msg *UEContextSetupResponse) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("UEContextSetupResponse"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := UEContextSetupResponseDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_DUtoCURRCInformation]; !ok {
		err = fmt.Errorf("Mandatory field DUtoCURRCInformation is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_DUtoCURRCInformation},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	// FIX: RequestedTargetCellGlobalID is optional, removed mandatory check
	return
}

type UEContextSetupResponseDecoder struct {
	msg      *UEContextSetupResponse
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *UEContextSetupResponseDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_DUtoCURRCInformation:
		var tmp DUtoCURRCInformation
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read DUtoCURRCInformation", err)
			return
		}
		msg.DUtoCURRCInformation = tmp
	case ProtocolIEID_CRNTI:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 65535},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read CRNTI", err)
			return
		}
		msg.CRNTI = (*int64)(&tmp.Value)
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
	case ProtocolIEID_FullConfiguration:
		var tmp FullConfiguration
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read FullConfiguration", err)
			return
		}
		msg.FullConfiguration = &tmp
	case ProtocolIEID_DRBsSetupList:
		tmp := Sequence[*DRBsSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		fn := func() *DRBsSetupItem { return new(DRBsSetupItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read DRBsSetupList", err)
			return
		}
		msg.DRBsSetupList = []DRBsSetupItem{}
		for _, i := range tmp.Value {
			msg.DRBsSetupList = append(msg.DRBsSetupList, *i)
		}
	case ProtocolIEID_SRBsFailedToBeSetupList:
		tmp := Sequence[*SRBsFailedToBeSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSRBs},
			ext: false,
		}
		fn := func() *SRBsFailedToBeSetupItem { return new(SRBsFailedToBeSetupItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SRBsFailedToBeSetupList", err)
			return
		}
		msg.SRBsFailedToBeSetupList = []SRBsFailedToBeSetupItem{}
		for _, i := range tmp.Value {
			msg.SRBsFailedToBeSetupList = append(msg.SRBsFailedToBeSetupList, *i)
		}
	case ProtocolIEID_DRBsFailedToBeSetupList:
		tmp := Sequence[*DRBsFailedToBeSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		fn := func() *DRBsFailedToBeSetupItem { return new(DRBsFailedToBeSetupItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read DRBsFailedToBeSetupList", err)
			return
		}
		msg.DRBsFailedToBeSetupList = []DRBsFailedToBeSetupItem{}
		for _, i := range tmp.Value {
			msg.DRBsFailedToBeSetupList = append(msg.DRBsFailedToBeSetupList, *i)
		}
	case ProtocolIEID_SCellFailedtoSetupList:
		tmp := Sequence[*SCellFailedtoSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSCells},
			ext: false,
		}
		fn := func() *SCellFailedtoSetupItem { return new(SCellFailedtoSetupItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SCellFailedtoSetupList", err)
			return
		}
		msg.SCellFailedtoSetupList = []SCellFailedtoSetupItem{}
		for _, i := range tmp.Value {
			msg.SCellFailedtoSetupList = append(msg.SCellFailedtoSetupList, *i)
		}
	case ProtocolIEID_InactivityMonitoringResponse:
		var tmp InactivityMonitoringResponse
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read InactivityMonitoringResponse", err)
			return
		}
		msg.InactivityMonitoringResponse = &tmp
	case ProtocolIEID_CriticalityDiagnostics:
		var tmp CriticalityDiagnostics
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read CriticalityDiagnostics", err)
			return
		}
		msg.CriticalityDiagnostics = &tmp
	case ProtocolIEID_SRBsSetupList:
		tmp := Sequence[*SRBsSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSRBs},
			ext: false,
		}
		fn := func() *SRBsSetupItem { return new(SRBsSetupItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SRBsSetupList", err)
			return
		}
		msg.SRBsSetupList = []SRBsSetupItem{}
		for _, i := range tmp.Value {
			msg.SRBsSetupList = append(msg.SRBsSetupList, *i)
		}
	case ProtocolIEID_BHChannelsSetupList:
		tmp := Sequence[*BHChannelsSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofBHRLCChannels},
			ext: false,
		}
		fn := func() *BHChannelsSetupItem { return new(BHChannelsSetupItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read BHChannelsSetupList", err)
			return
		}
		msg.BHChannelsSetupList = []BHChannelsSetupItem{}
		for _, i := range tmp.Value {
			msg.BHChannelsSetupList = append(msg.BHChannelsSetupList, *i)
		}
	case ProtocolIEID_BHChannelsFailedToBeSetupList:
		tmp := Sequence[*BHChannelsFailedToBeSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofBHRLCChannels},
			ext: false,
		}
		fn := func() *BHChannelsFailedToBeSetupItem { return new(BHChannelsFailedToBeSetupItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read BHChannelsFailedToBeSetupList", err)
			return
		}
		msg.BHChannelsFailedToBeSetupList = []BHChannelsFailedToBeSetupItem{}
		for _, i := range tmp.Value {
			msg.BHChannelsFailedToBeSetupList = append(msg.BHChannelsFailedToBeSetupList, *i)
		}
	case ProtocolIEID_SLDRBsSetupList:
		tmp := Sequence[*SLDRBsSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		fn := func() *SLDRBsSetupItem { return new(SLDRBsSetupItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SLDRBsSetupList", err)
			return
		}
		msg.SLDRBsSetupList = []SLDRBsSetupItem{}
		for _, i := range tmp.Value {
			msg.SLDRBsSetupList = append(msg.SLDRBsSetupList, *i)
		}
	case ProtocolIEID_SLDRBsFailedToBeSetupList:
		tmp := Sequence[*SLDRBsFailedToBeSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		fn := func() *SLDRBsFailedToBeSetupItem { return new(SLDRBsFailedToBeSetupItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SLDRBsFailedToBeSetupList", err)
			return
		}
		msg.SLDRBsFailedToBeSetupList = []SLDRBsFailedToBeSetupItem{}
		for _, i := range tmp.Value {
			msg.SLDRBsFailedToBeSetupList = append(msg.SLDRBsFailedToBeSetupList, *i)
		}
	case ProtocolIEID_RequestedTargetCellGlobalID:
		var tmp NRCGI
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RequestedTargetCellGlobalID", err)
			return
		}
		msg.RequestedTargetCellGlobalID = &tmp
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