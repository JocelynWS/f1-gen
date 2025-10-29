package f1ap

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UEContextModificationResponse struct {
	GNBCUUEF1APID                         int64                              `lb:0,ub:4294967295,mandatory,reject`
	GNBDUUEF1APID                         int64                              `lb:0,ub:4294967295,mandatory,reject`
	ResourceCoordinationTransferContainer []byte                             `lb:0,ub:0,optional,ignore`
	DUtoCURRCInformation                  *DUtoCURRCInformation              `mandatory,reject`
	DRBsSetupModList                      []DRBsSetupModItem                 `lb:1,ub:maxnoofDRBs,mandatory,ignore`
	DRBsModifiedList                      []DRBsModifiedItem                 `lb:1,ub:maxnoofDRBs,mandatory,ignore`
	SRBsFailedToBeSetupModList            []SRBsFailedToBeSetupModItem       `lb:1,ub:maxnoofSRBs,optional,ignore`
	DRBsFailedToBeSetupModList            []DRBsFailedToBeSetupModItem       `lb:1,ub:maxnoofDRBs,optional,ignore`
	SCellFailedtoSetupModList             []SCellFailedtoSetupModItem        `lb:1,ub:maxnoofSCells,optional,ignore`
	DRBsFailedToBeModifiedList            []DRBsFailedToBeModifiedItem       `lb:1,ub:maxnoofDRBs,optional,ignore`
	InactivityMonitoringResponse          *InactivityMonitoringResponse      `optional,reject`
	CriticalityDiagnostics                *CriticalityDiagnostics            `optional,ignore`
	CRNTI                                 *int64                             `lb:0,ub:65535,optional,ignore`
	AssociatedSCellList                   []AssociatedSCellItem              `lb:1,ub:maxnoofSCells,optional,ignore,valueExt`
	SRBsSetupModList                      []SRBsSetupModItem                 `lb:1,ub:maxnoofSRBs,optional,ignore`
	SRBsModifiedList                      []SRBsModifiedItem                 `lb:1,ub:maxnoofSRBs,optional,ignore`
	FullConfiguration                     *FullConfiguration                 `optional,reject`
	BHChannelsSetupModList                []BHChannelsSetupModItem           `lb:1,ub:maxnoofBHRLCChannels,mandatory,ignore`
	BHChannelsModifiedList                []BHChannelsModifiedItem           `lb:1,ub:maxnoofBHRLCChannels,mandatory,ignore`
	BHChannelsFailedToBeSetupModList      []BHChannelsFailedToBeSetupModItem `lb:1,ub:maxnoofBHRLCChannels,optional,ignore`
	BHChannelsFailedToBeModifiedList      []BHChannelsFailedToBeModifiedItem `lb:1,ub:maxnoofBHRLCChannels,optional,ignore`
	SLDRBsSetupModList                    []SLDRBsSetupModItem               `lb:1,ub:maxnoofSLDRBs,optional,ignore`
	SLDRBsModifiedList                    []SLDRBsModifiedItem               `lb:1,ub:maxnoofSLDRBs,optional,ignore`
	SLDRBsFailedToBeSetupModList          []SLDRBsFailedToBeSetupModItem     `lb:1,ub:maxnoofSLDRBs,optional,ignore`
	SLDRBsFailedToBeModifiedList          []SLDRBsFailedToBeModifiedItem     `lb:1,ub:maxnoofSLDRBs,optional,ignore`
	RequestedTargetCellGlobalID           *NRCGI                             `mandatory,reject`
}

func (msg *UEContextModificationResponse) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("UEContextModificationResponse"), err)
		return
	}
	return encodeMessage(w, F1apPduSuccessfulOutcome, ProcedureCode_UEContextModificationRequired, Criticality_PresentReject, ies)
}
func (msg *UEContextModificationResponse) toIes() (ies []F1apMessageIE, err error) {
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
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_DUtoCURRCInformation},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       msg.DUtoCURRCInformation,
	})
	if len(msg.DRBsSetupModList) > 0 {
		tmp_DRBsSetupModList := Sequence[*DRBsSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		for _, i := range msg.DRBsSetupModList {
			tmp_DRBsSetupModList.Value = append(tmp_DRBsSetupModList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DRBsSetupModList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_DRBsSetupModList,
		})
	} else {
		err = utils.WrapError("DRBsSetupModList is nil", err)
		return
	}
	if len(msg.DRBsModifiedList) > 0 {
		tmp_DRBsModifiedList := Sequence[*DRBsModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		for _, i := range msg.DRBsModifiedList {
			tmp_DRBsModifiedList.Value = append(tmp_DRBsModifiedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DRBsModifiedList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_DRBsModifiedList,
		})
	} else {
		err = utils.WrapError("DRBsModifiedList is nil", err)
		return
	}
	if len(msg.SRBsFailedToBeSetupModList) > 0 {
		tmp_SRBsFailedToBeSetupModList := Sequence[*SRBsFailedToBeSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSRBs},
			ext: false,
		}
		for _, i := range msg.SRBsFailedToBeSetupModList {
			tmp_SRBsFailedToBeSetupModList.Value = append(tmp_SRBsFailedToBeSetupModList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SRBsFailedToBeSetupModList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_SRBsFailedToBeSetupModList,
		})
	}
	if len(msg.DRBsFailedToBeSetupModList) > 0 {
		tmp_DRBsFailedToBeSetupModList := Sequence[*DRBsFailedToBeSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		for _, i := range msg.DRBsFailedToBeSetupModList {
			tmp_DRBsFailedToBeSetupModList.Value = append(tmp_DRBsFailedToBeSetupModList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DRBsFailedToBeSetupModList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_DRBsFailedToBeSetupModList,
		})
	}
	if len(msg.SCellFailedtoSetupModList) > 0 {
		tmp_SCellFailedtoSetupModList := Sequence[*SCellFailedtoSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSCells},
			ext: false,
		}
		for _, i := range msg.SCellFailedtoSetupModList {
			tmp_SCellFailedtoSetupModList.Value = append(tmp_SCellFailedtoSetupModList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SCellFailedtoSetupModList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_SCellFailedtoSetupModList,
		})
	}
	if len(msg.DRBsFailedToBeModifiedList) > 0 {
		tmp_DRBsFailedToBeModifiedList := Sequence[*DRBsFailedToBeModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		for _, i := range msg.DRBsFailedToBeModifiedList {
			tmp_DRBsFailedToBeModifiedList.Value = append(tmp_DRBsFailedToBeModifiedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DRBsFailedToBeModifiedList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_DRBsFailedToBeModifiedList,
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
	if msg.CRNTI != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CRNTI},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 65535},
				ext:   false,
				Value: aper.Integer(*msg.CRNTI),
			}})
	}
	if len(msg.AssociatedSCellList) > 0 {
		tmp_AssociatedSCellList := Sequence[*AssociatedSCellItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSCells},
			ext: true,
		}
		for _, i := range msg.AssociatedSCellList {
			tmp_AssociatedSCellList.Value = append(tmp_AssociatedSCellList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AssociatedSCellList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_AssociatedSCellList,
		})
	}
	if len(msg.SRBsSetupModList) > 0 {
		tmp_SRBsSetupModList := Sequence[*SRBsSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSRBs},
			ext: false,
		}
		for _, i := range msg.SRBsSetupModList {
			tmp_SRBsSetupModList.Value = append(tmp_SRBsSetupModList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SRBsSetupModList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_SRBsSetupModList,
		})
	}
	if len(msg.SRBsModifiedList) > 0 {
		tmp_SRBsModifiedList := Sequence[*SRBsModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSRBs},
			ext: false,
		}
		for _, i := range msg.SRBsModifiedList {
			tmp_SRBsModifiedList.Value = append(tmp_SRBsModifiedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SRBsModifiedList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_SRBsModifiedList,
		})
	}
	if msg.FullConfiguration != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_FullConfiguration},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.FullConfiguration,
		})
	}
	if len(msg.BHChannelsSetupModList) > 0 {
		tmp_BHChannelsSetupModList := Sequence[*BHChannelsSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofBHRLCChannels},
			ext: false,
		}
		for _, i := range msg.BHChannelsSetupModList {
			tmp_BHChannelsSetupModList.Value = append(tmp_BHChannelsSetupModList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_BHChannelsSetupModList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_BHChannelsSetupModList,
		})
	} else {
		err = utils.WrapError("BHChannelsSetupModList is nil", err)
		return
	}
	if len(msg.BHChannelsModifiedList) > 0 {
		tmp_BHChannelsModifiedList := Sequence[*BHChannelsModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofBHRLCChannels},
			ext: false,
		}
		for _, i := range msg.BHChannelsModifiedList {
			tmp_BHChannelsModifiedList.Value = append(tmp_BHChannelsModifiedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_BHChannelsModifiedList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_BHChannelsModifiedList,
		})
	} else {
		err = utils.WrapError("BHChannelsModifiedList is nil", err)
		return
	}
	if len(msg.BHChannelsFailedToBeSetupModList) > 0 {
		tmp_BHChannelsFailedToBeSetupModList := Sequence[*BHChannelsFailedToBeSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofBHRLCChannels},
			ext: false,
		}
		for _, i := range msg.BHChannelsFailedToBeSetupModList {
			tmp_BHChannelsFailedToBeSetupModList.Value = append(tmp_BHChannelsFailedToBeSetupModList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_BHChannelsFailedToBeSetupModList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_BHChannelsFailedToBeSetupModList,
		})
	}
	if len(msg.BHChannelsFailedToBeModifiedList) > 0 {
		tmp_BHChannelsFailedToBeModifiedList := Sequence[*BHChannelsFailedToBeModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofBHRLCChannels},
			ext: false,
		}
		for _, i := range msg.BHChannelsFailedToBeModifiedList {
			tmp_BHChannelsFailedToBeModifiedList.Value = append(tmp_BHChannelsFailedToBeModifiedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_BHChannelsFailedToBeModifiedList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_BHChannelsFailedToBeModifiedList,
		})
	}
	if len(msg.SLDRBsSetupModList) > 0 {
		tmp_SLDRBsSetupModList := Sequence[*SLDRBsSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		for _, i := range msg.SLDRBsSetupModList {
			tmp_SLDRBsSetupModList.Value = append(tmp_SLDRBsSetupModList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SLDRBsSetupModList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_SLDRBsSetupModList,
		})
	}
	if len(msg.SLDRBsModifiedList) > 0 {
		tmp_SLDRBsModifiedList := Sequence[*SLDRBsModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		for _, i := range msg.SLDRBsModifiedList {
			tmp_SLDRBsModifiedList.Value = append(tmp_SLDRBsModifiedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SLDRBsModifiedList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_SLDRBsModifiedList,
		})
	}
	if len(msg.SLDRBsFailedToBeSetupModList) > 0 {
		tmp_SLDRBsFailedToBeSetupModList := Sequence[*SLDRBsFailedToBeSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		for _, i := range msg.SLDRBsFailedToBeSetupModList {
			tmp_SLDRBsFailedToBeSetupModList.Value = append(tmp_SLDRBsFailedToBeSetupModList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SLDRBsFailedToBeSetupModList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_SLDRBsFailedToBeSetupModList,
		})
	}
	if len(msg.SLDRBsFailedToBeModifiedList) > 0 {
		tmp_SLDRBsFailedToBeModifiedList := Sequence[*SLDRBsFailedToBeModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		for _, i := range msg.SLDRBsFailedToBeModifiedList {
			tmp_SLDRBsFailedToBeModifiedList.Value = append(tmp_SLDRBsFailedToBeModifiedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SLDRBsFailedToBeModifiedList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_SLDRBsFailedToBeModifiedList,
		})
	}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RequestedTargetCellGlobalID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       msg.RequestedTargetCellGlobalID,
	})
	return
}
func (msg *UEContextModificationResponse) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("UEContextModificationResponse"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := UEContextModificationResponseDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_DRBsSetupModList]; !ok {
		err = fmt.Errorf("Mandatory field DRBsSetupModList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_DRBsSetupModList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_DRBsModifiedList]; !ok {
		err = fmt.Errorf("Mandatory field DRBsModifiedList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_DRBsModifiedList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_BHChannelsSetupModList]; !ok {
		err = fmt.Errorf("Mandatory field BHChannelsSetupModList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_BHChannelsSetupModList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_BHChannelsModifiedList]; !ok {
		err = fmt.Errorf("Mandatory field BHChannelsModifiedList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_BHChannelsModifiedList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_RequestedTargetCellGlobalID]; !ok {
		err = fmt.Errorf("Mandatory field RequestedTargetCellGlobalID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_RequestedTargetCellGlobalID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type UEContextModificationResponseDecoder struct {
	msg      *UEContextModificationResponse
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *UEContextModificationResponseDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_DUtoCURRCInformation:
		var tmp DUtoCURRCInformation
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read DUtoCURRCInformation", err)
			return
		}
		msg.DUtoCURRCInformation = &tmp
	case ProtocolIEID_DRBsSetupModList:
		tmp := Sequence[*DRBsSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		fn := func() *DRBsSetupModItem { return new(DRBsSetupModItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read DRBsSetupModList", err)
			return
		}
		msg.DRBsSetupModList = []DRBsSetupModItem{}
		for _, i := range tmp.Value {
			msg.DRBsSetupModList = append(msg.DRBsSetupModList, *i)
		}
	case ProtocolIEID_DRBsModifiedList:
		tmp := Sequence[*DRBsModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		fn := func() *DRBsModifiedItem { return new(DRBsModifiedItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read DRBsModifiedList", err)
			return
		}
		msg.DRBsModifiedList = []DRBsModifiedItem{}
		for _, i := range tmp.Value {
			msg.DRBsModifiedList = append(msg.DRBsModifiedList, *i)
		}
	case ProtocolIEID_SRBsFailedToBeSetupModList:
		tmp := Sequence[*SRBsFailedToBeSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSRBs},
			ext: false,
		}
		fn := func() *SRBsFailedToBeSetupModItem { return new(SRBsFailedToBeSetupModItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SRBsFailedToBeSetupModList", err)
			return
		}
		msg.SRBsFailedToBeSetupModList = []SRBsFailedToBeSetupModItem{}
		for _, i := range tmp.Value {
			msg.SRBsFailedToBeSetupModList = append(msg.SRBsFailedToBeSetupModList, *i)
		}
	case ProtocolIEID_DRBsFailedToBeSetupModList:
		tmp := Sequence[*DRBsFailedToBeSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		fn := func() *DRBsFailedToBeSetupModItem { return new(DRBsFailedToBeSetupModItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read DRBsFailedToBeSetupModList", err)
			return
		}
		msg.DRBsFailedToBeSetupModList = []DRBsFailedToBeSetupModItem{}
		for _, i := range tmp.Value {
			msg.DRBsFailedToBeSetupModList = append(msg.DRBsFailedToBeSetupModList, *i)
		}
	case ProtocolIEID_SCellFailedtoSetupModList:
		tmp := Sequence[*SCellFailedtoSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSCells},
			ext: false,
		}
		fn := func() *SCellFailedtoSetupModItem { return new(SCellFailedtoSetupModItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SCellFailedtoSetupModList", err)
			return
		}
		msg.SCellFailedtoSetupModList = []SCellFailedtoSetupModItem{}
		for _, i := range tmp.Value {
			msg.SCellFailedtoSetupModList = append(msg.SCellFailedtoSetupModList, *i)
		}
	case ProtocolIEID_DRBsFailedToBeModifiedList:
		tmp := Sequence[*DRBsFailedToBeModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		fn := func() *DRBsFailedToBeModifiedItem { return new(DRBsFailedToBeModifiedItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read DRBsFailedToBeModifiedList", err)
			return
		}
		msg.DRBsFailedToBeModifiedList = []DRBsFailedToBeModifiedItem{}
		for _, i := range tmp.Value {
			msg.DRBsFailedToBeModifiedList = append(msg.DRBsFailedToBeModifiedList, *i)
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
	case ProtocolIEID_CRNTI:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 65535},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read CRNTI", err)
			return
		}
		msg.CRNTI = (*int64)(&tmp.Value)
	case ProtocolIEID_AssociatedSCellList:
		tmp := Sequence[*AssociatedSCellItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSCells},
			ext: true,
		}
		fn := func() *AssociatedSCellItem { return new(AssociatedSCellItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read AssociatedSCellList", err)
			return
		}
		msg.AssociatedSCellList = []AssociatedSCellItem{}
		for _, i := range tmp.Value {
			msg.AssociatedSCellList = append(msg.AssociatedSCellList, *i)
		}
	case ProtocolIEID_SRBsSetupModList:
		tmp := Sequence[*SRBsSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSRBs},
			ext: false,
		}
		fn := func() *SRBsSetupModItem { return new(SRBsSetupModItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SRBsSetupModList", err)
			return
		}
		msg.SRBsSetupModList = []SRBsSetupModItem{}
		for _, i := range tmp.Value {
			msg.SRBsSetupModList = append(msg.SRBsSetupModList, *i)
		}
	case ProtocolIEID_SRBsModifiedList:
		tmp := Sequence[*SRBsModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSRBs},
			ext: false,
		}
		fn := func() *SRBsModifiedItem { return new(SRBsModifiedItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SRBsModifiedList", err)
			return
		}
		msg.SRBsModifiedList = []SRBsModifiedItem{}
		for _, i := range tmp.Value {
			msg.SRBsModifiedList = append(msg.SRBsModifiedList, *i)
		}
	case ProtocolIEID_FullConfiguration:
		var tmp FullConfiguration
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read FullConfiguration", err)
			return
		}
		msg.FullConfiguration = &tmp
	case ProtocolIEID_BHChannelsSetupModList:
		tmp := Sequence[*BHChannelsSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofBHRLCChannels},
			ext: false,
		}
		fn := func() *BHChannelsSetupModItem { return new(BHChannelsSetupModItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read BHChannelsSetupModList", err)
			return
		}
		msg.BHChannelsSetupModList = []BHChannelsSetupModItem{}
		for _, i := range tmp.Value {
			msg.BHChannelsSetupModList = append(msg.BHChannelsSetupModList, *i)
		}
	case ProtocolIEID_BHChannelsModifiedList:
		tmp := Sequence[*BHChannelsModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofBHRLCChannels},
			ext: false,
		}
		fn := func() *BHChannelsModifiedItem { return new(BHChannelsModifiedItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read BHChannelsModifiedList", err)
			return
		}
		msg.BHChannelsModifiedList = []BHChannelsModifiedItem{}
		for _, i := range tmp.Value {
			msg.BHChannelsModifiedList = append(msg.BHChannelsModifiedList, *i)
		}
	case ProtocolIEID_BHChannelsFailedToBeSetupModList:
		tmp := Sequence[*BHChannelsFailedToBeSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofBHRLCChannels},
			ext: false,
		}
		fn := func() *BHChannelsFailedToBeSetupModItem { return new(BHChannelsFailedToBeSetupModItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read BHChannelsFailedToBeSetupModList", err)
			return
		}
		msg.BHChannelsFailedToBeSetupModList = []BHChannelsFailedToBeSetupModItem{}
		for _, i := range tmp.Value {
			msg.BHChannelsFailedToBeSetupModList = append(msg.BHChannelsFailedToBeSetupModList, *i)
		}
	case ProtocolIEID_BHChannelsFailedToBeModifiedList:
		tmp := Sequence[*BHChannelsFailedToBeModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofBHRLCChannels},
			ext: false,
		}
		fn := func() *BHChannelsFailedToBeModifiedItem { return new(BHChannelsFailedToBeModifiedItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read BHChannelsFailedToBeModifiedList", err)
			return
		}
		msg.BHChannelsFailedToBeModifiedList = []BHChannelsFailedToBeModifiedItem{}
		for _, i := range tmp.Value {
			msg.BHChannelsFailedToBeModifiedList = append(msg.BHChannelsFailedToBeModifiedList, *i)
		}
	case ProtocolIEID_SLDRBsSetupModList:
		tmp := Sequence[*SLDRBsSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		fn := func() *SLDRBsSetupModItem { return new(SLDRBsSetupModItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SLDRBsSetupModList", err)
			return
		}
		msg.SLDRBsSetupModList = []SLDRBsSetupModItem{}
		for _, i := range tmp.Value {
			msg.SLDRBsSetupModList = append(msg.SLDRBsSetupModList, *i)
		}
	case ProtocolIEID_SLDRBsModifiedList:
		tmp := Sequence[*SLDRBsModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		fn := func() *SLDRBsModifiedItem { return new(SLDRBsModifiedItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SLDRBsModifiedList", err)
			return
		}
		msg.SLDRBsModifiedList = []SLDRBsModifiedItem{}
		for _, i := range tmp.Value {
			msg.SLDRBsModifiedList = append(msg.SLDRBsModifiedList, *i)
		}
	case ProtocolIEID_SLDRBsFailedToBeSetupModList:
		tmp := Sequence[*SLDRBsFailedToBeSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		fn := func() *SLDRBsFailedToBeSetupModItem { return new(SLDRBsFailedToBeSetupModItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SLDRBsFailedToBeSetupModList", err)
			return
		}
		msg.SLDRBsFailedToBeSetupModList = []SLDRBsFailedToBeSetupModItem{}
		for _, i := range tmp.Value {
			msg.SLDRBsFailedToBeSetupModList = append(msg.SLDRBsFailedToBeSetupModList, *i)
		}
	case ProtocolIEID_SLDRBsFailedToBeModifiedList:
		tmp := Sequence[*SLDRBsFailedToBeModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		fn := func() *SLDRBsFailedToBeModifiedItem { return new(SLDRBsFailedToBeModifiedItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SLDRBsFailedToBeModifiedList", err)
			return
		}
		msg.SLDRBsFailedToBeModifiedList = []SLDRBsFailedToBeModifiedItem{}
		for _, i := range tmp.Value {
			msg.SLDRBsFailedToBeModifiedList = append(msg.SLDRBsFailedToBeModifiedList, *i)
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
