package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UEContextSetupRequest struct {
	GNBCUUEF1APID                           int64                                    `lb:0,ub:4294967295,mandatory,reject`
	GNBDUUEF1APID                           *int64                                   `lb:0,ub:4294967295,optional,ignore`
	SpCellID                                NRCGI                                    `mandatory,reject`
	ServCellIndex                           int64                                    `lb:0,ub:31,mandatory,reject`
	SpCellULConfigured                      *CellULConfigured                        `optional,ignore`
	CUtoDURRCInformation                    *CUtoDURRCInformation                    `mandatory,reject`
	CandidateSpCellList                     []CandidateSpCellItem                    `lb:1,ub:maxnoofCandidateSpCells,optional,ignore`
	DRXCycle                                *DRXCycle                                `optional,ignore`
	ResourceCoordinationTransferContainer   []byte                                   `lb:0,ub:0,optional,ignore`
	SCellToBeSetupList                      []SCellToBeSetupItem                     `lb:1,ub:maxnoofSCells,optional,ignore`
	SRBsToBeSetupList                       []SRBsToBeSetupItem                      `lb:1,ub:maxnoofSRBs,optional,reject`
	DRBsToBeSetupList                       []DRBsToBeSetupItem                      `lb:1,ub:maxnoofDRBs,optional,reject`
	InactivityMonitoringRequest             *InactivityMonitoringRequest             `optional,reject`
	RATFrequencyPriorityInformation         *RATFrequencyPriorityInformation         `optional,reject`
	RRCContainer                            []byte                                   `lb:0,ub:0,optional,ignore`
	MaskedIMEISV                            *aper.BitString                          `lb:64,ub:64,optional,ignore`
	ServingPLMN                             []byte                                   `lb:3,ub:3,optional,ignore`
	GNBDUUEAMBRUL                           int64                                    `lb:0,ub:4000000000000,conditional,ignore,valueExt`
	RRCDeliveryStatusRequest                *RRCDeliveryStatusRequest                `optional,ignore`
	ResourceCoordinationTransferInformation *ResourceCoordinationTransferInformation `optional,ignore`
	ServingCellMO                           *int64                                   `lb:1,ub:64,optional,ignore`
	NewgNBCUUEF1APID                        *int64                                   `lb:0,ub:4294967295,optional,reject`
	RANUEID                                 []byte                                   `lb:8,ub:8,optional,ignore`
	TraceActivation                         *TraceActivation                         `optional,ignore`
	AdditionalRRMPriorityIndex              *aper.BitString                          `lb:32,ub:32,optional,ignore`
	BHChannelsToBeSetupList                 []BHChannelsToBeSetupItem                `lb:1,ub:maxnoofBHRLCChannels,optional,reject`
	ConfiguredBAPAddress                    *aper.BitString                          `lb:10,ub:10,optional,reject`
	NRV2XServicesAuthorized                 *NRV2XServicesAuthorized                 `optional,ignore`
	LTEV2XServicesAuthorized                *LTEV2XServicesAuthorized                `optional,ignore`
	NRUESidelinkAggregateMaximumBitrate     *NRUESidelinkAggregateMaximumBitrate     `optional,ignore`
	LTEUESidelinkAggregateMaximumBitrate    *LTEUESidelinkAggregateMaximumBitrate    `optional,ignore`
	PC5LinkAMBR                             int64                                    `lb:0,ub:4000000000000,mandatory,ignore,valueExt`
	SLDRBsToBeSetupList                     []SLDRBsToBeSetupItem                    `lb:1,ub:maxnoofSLDRBs,optional,reject`
	ConditionalInterDUMobilityInformation   *ConditionalInterDUMobilityInformation   `mandatory,reject`
	ManagementBasedMDTPLMNList              []PLMNIdentity                           `lb:1,ub:maxnoofMDTPLMNs,optional,ignore`
	ServingNID                              *aper.BitString                          `lb:44,ub:44,optional,reject`
	F1CTransferPath                         *F1CTransferPath                         `optional,reject`
}

func (msg *UEContextSetupRequest) Encode(w io.Writer) (err error) {
    var ies []F1apMessageIE
    if ies, err = msg.toIes(); err != nil {
        err = msgErrors(fmt.Errorf("UEContextSetupRequest"), err)
        return
    }
    return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_UEContextSetup, Criticality_PresentReject, ies)
}
func (msg *UEContextSetupRequest) toIes() (ies []F1apMessageIE, err error) {
	ies = []F1apMessageIE{}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_GNBCUUEF1APID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: aper.Integer(msg.GNBCUUEF1APID),
		}})
	if msg.GNBDUUEF1APID != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GNBDUUEF1APID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 4294967295},
				ext:   false,
				Value: aper.Integer(*msg.GNBDUUEF1APID),
			}})
	}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_SpCellID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.SpCellID,
	})
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_ServCellIndex},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 31},
			ext:   false,
			Value: aper.Integer(msg.ServCellIndex),
		}})
	if msg.SpCellULConfigured != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SpCellULConfigured},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.SpCellULConfigured,
		})
	}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_CUtoDURRCInformation},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.CUtoDURRCInformation,
	})
	if len(msg.CandidateSpCellList) > 0 {
		tmp_CandidateSpCellList := Sequence[*CandidateSpCellItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCandidateSpCells},
			ext: false,
		}
		for _, i := range msg.CandidateSpCellList {
			tmp_CandidateSpCellList.Value = append(tmp_CandidateSpCellList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CandidateSpCellList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_CandidateSpCellList,
		})
	}
	if msg.DRXCycle != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DRXCycle},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.DRXCycle,
		})
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
	if len(msg.SCellToBeSetupList) > 0 {
		tmp_SCellToBeSetupList := Sequence[*SCellToBeSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSCells},
			ext: false,
		}
		for _, i := range msg.SCellToBeSetupList {
			tmp_SCellToBeSetupList.Value = append(tmp_SCellToBeSetupList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SCellToBeSetupList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_SCellToBeSetupList,
		})
	}
	if len(msg.SRBsToBeSetupList) > 0 {
		tmp_SRBsToBeSetupList := Sequence[*SRBsToBeSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSRBs},
			ext: false,
		}
		for _, i := range msg.SRBsToBeSetupList {
			tmp_SRBsToBeSetupList.Value = append(tmp_SRBsToBeSetupList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SRBsToBeSetupList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_SRBsToBeSetupList,
		})
	}
	if len(msg.DRBsToBeSetupList) > 0 {
		tmp_DRBsToBeSetupList := Sequence[*DRBsToBeSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		for _, i := range msg.DRBsToBeSetupList {
			tmp_DRBsToBeSetupList.Value = append(tmp_DRBsToBeSetupList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DRBsToBeSetupList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_DRBsToBeSetupList,
		})
	}
	if msg.InactivityMonitoringRequest != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_InactivityMonitoringRequest},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.InactivityMonitoringRequest,
		})
	}
	if msg.RATFrequencyPriorityInformation != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RATFrequencyPriorityInformation},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.RATFrequencyPriorityInformation,
		})
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
	if msg.MaskedIMEISV != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_MaskedIMEISV},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &BITSTRING{
				c:   aper.Constraint{Lb: 64, Ub: 64},
				ext: false,
				Value: aper.BitString{
					Bytes: msg.MaskedIMEISV.Bytes, NumBits: msg.MaskedIMEISV.NumBits},
			}})
	}
	if msg.ServingPLMN != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ServingPLMN},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 3, Ub: 3},
				ext:   false,
				Value: msg.ServingPLMN,
			}})
	}
	if msg.GNBDUUEAMBRUL != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GNBDUUEAMBRUL},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 4000000000000},
				ext:   true,
				Value: aper.Integer(*msg.GNBDUUEAMBRUL),
			}})
	}
	if msg.RRCDeliveryStatusRequest != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RRCDeliveryStatusRequest},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RRCDeliveryStatusRequest,
		})
	}
	if msg.ResourceCoordinationTransferInformation != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ResourceCoordinationTransferInformation},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.ResourceCoordinationTransferInformation,
		})
	}
	if msg.ServingCellMO != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ServingCellMO},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 1, Ub: 64},
				ext:   false,
				Value: aper.Integer(*msg.ServingCellMO),
			}})
	}
	if msg.NewgNBCUUEF1APID != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NewgNBCUUEF1APID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 4294967295},
				ext:   false,
				Value: aper.Integer(*msg.NewgNBCUUEF1APID),
			}})
	}
	if msg.RANUEID != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANUEID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 8, Ub: 8},
				ext:   false,
				Value: msg.RANUEID,
			}})
	}
	if msg.TraceActivation != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TraceActivation},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.TraceActivation,
		})
	}
	if msg.AdditionalRRMPriorityIndex != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AdditionalRRMPriorityIndex},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &BITSTRING{
				c:   aper.Constraint{Lb: 32, Ub: 32},
				ext: false,
				Value: aper.BitString{
					Bytes: msg.AdditionalRRMPriorityIndex.Bytes, NumBits: msg.AdditionalRRMPriorityIndex.NumBits},
			}})
	}
	if len(msg.BHChannelsToBeSetupList) > 0 {
		tmp_BHChannelsToBeSetupList := Sequence[*BHChannelsToBeSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofBHRLCChannels},
			ext: false,
		}
		for _, i := range msg.BHChannelsToBeSetupList {
			tmp_BHChannelsToBeSetupList.Value = append(tmp_BHChannelsToBeSetupList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_BHChannelsToBeSetupList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_BHChannelsToBeSetupList,
		})
	}
	if msg.ConfiguredBAPAddress != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ConfiguredBAPAddress},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value: &BITSTRING{
				c:   aper.Constraint{Lb: 10, Ub: 10},
				ext: false,
				Value: aper.BitString{
					Bytes: msg.ConfiguredBAPAddress.Bytes, NumBits: msg.ConfiguredBAPAddress.NumBits},
			}})
	}
	if msg.NRV2XServicesAuthorized != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NRV2XServicesAuthorized},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.NRV2XServicesAuthorized,
		})
	}
	if msg.LTEV2XServicesAuthorized != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_LTEV2XServicesAuthorized},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.LTEV2XServicesAuthorized,
		})
	}
	if msg.NRUESidelinkAggregateMaximumBitrate != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NRUESidelinkAggregateMaximumBitrate},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.NRUESidelinkAggregateMaximumBitrate,
		})
	}
	if msg.LTEUESidelinkAggregateMaximumBitrate != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_LTEUESidelinkAggregateMaximumBitrate},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.LTEUESidelinkAggregateMaximumBitrate,
		})
	}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_PC5LinkAMBR},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4000000000000},
			ext:   true,
			Value: aper.Integer(msg.PC5LinkAMBR),
		}})
	if len(msg.SLDRBsToBeSetupList) > 0 {
		tmp_SLDRBsToBeSetupList := Sequence[*SLDRBsToBeSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		for _, i := range msg.SLDRBsToBeSetupList {
			tmp_SLDRBsToBeSetupList.Value = append(tmp_SLDRBsToBeSetupList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SLDRBsToBeSetupList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_SLDRBsToBeSetupList,
		})
	}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_ConditionalInterDUMobilityInformation},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.ConditionalInterDUMobilityInformation,
	})
	if len(msg.ManagementBasedMDTPLMNList) > 0 {
		tmp_ManagementBasedMDTPLMNList := Sequence[*PLMNIdentity]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMDTPLMNs},
			ext: false,
		}
		for _, i := range msg.ManagementBasedMDTPLMNList {
			tmp_ManagementBasedMDTPLMNList.Value = append(tmp_ManagementBasedMDTPLMNList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ManagementBasedMDTPLMNList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_ManagementBasedMDTPLMNList,
		})
	}
	if msg.ServingNID != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ServingNID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value: &BITSTRING{
				c:   aper.Constraint{Lb: 44, Ub: 44},
				ext: false,
				Value: aper.BitString{
					Bytes: msg.ServingNID.Bytes, NumBits: msg.ServingNID.NumBits},
			}})
	}
	if msg.F1CTransferPath != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_F1CTransferPath},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.F1CTransferPath,
		})
	}
	return
}
func (msg *UEContextSetupRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("UEContextSetupRequest"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := UEContextSetupRequestDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_SpCellID]; !ok {
		err = fmt.Errorf("Mandatory field SpCellID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_SpCellID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_ServCellIndex]; !ok {
		err = fmt.Errorf("Mandatory field ServCellIndex is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_ServCellIndex},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_CUtoDURRCInformation]; !ok {
		err = fmt.Errorf("Mandatory field CUtoDURRCInformation is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_CUtoDURRCInformation},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_PC5LinkAMBR]; !ok {
		err = fmt.Errorf("Mandatory field PC5LinkAMBR is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_PC5LinkAMBR},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_ConditionalInterDUMobilityInformation]; !ok {
		err = fmt.Errorf("Mandatory field ConditionalInterDUMobilityInformation is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_ConditionalInterDUMobilityInformation},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type UEContextSetupRequestDecoder struct {
	msg      *UEContextSetupRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *UEContextSetupRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
		msg.GNBDUUEF1APID = (*int64)(&tmp.Value)
	case ProtocolIEID_SpCellID:
		var tmp NRCGI
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SpCellID", err)
			return
		}
		msg.SpCellID = tmp
	case ProtocolIEID_ServCellIndex:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 31},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ServCellIndex", err)
			return
		}
		msg.ServCellIndex = int64(tmp.Value)
	case ProtocolIEID_SpCellULConfigured:
		var tmp CellULConfigured
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SpCellULConfigured", err)
			return
		}
		msg.SpCellULConfigured = &tmp
	case ProtocolIEID_CUtoDURRCInformation:
		var tmp CUtoDURRCInformation
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read CUtoDURRCInformation", err)
			return
		}
		msg.CUtoDURRCInformation = tmp
	case ProtocolIEID_CandidateSpCellList:
		tmp := Sequence[*CandidateSpCellItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCandidateSpCells},
			ext: false,
		}
		fn := func() *CandidateSpCellItem { return new(CandidateSpCellItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read CandidateSpCellList", err)
			return
		}
		msg.CandidateSpCellList = []CandidateSpCellItem{}
		for _, i := range tmp.Value {
			msg.CandidateSpCellList = append(msg.CandidateSpCellList, *i)
		}
	case ProtocolIEID_DRXCycle:
		var tmp DRXCycle
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read DRXCycle", err)
			return
		}
		msg.DRXCycle = &tmp
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
	case ProtocolIEID_SCellToBeSetupList:
		tmp := Sequence[*SCellToBeSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSCells},
			ext: false,
		}
		fn := func() *SCellToBeSetupItem { return new(SCellToBeSetupItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SCellToBeSetupList", err)
			return
		}
		msg.SCellToBeSetupList = []SCellToBeSetupItem{}
		for _, i := range tmp.Value {
			msg.SCellToBeSetupList = append(msg.SCellToBeSetupList, *i)
		}
	case ProtocolIEID_SRBsToBeSetupList:
		tmp := Sequence[*SRBsToBeSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSRBs},
			ext: false,
		}
		fn := func() *SRBsToBeSetupItem { return new(SRBsToBeSetupItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SRBsToBeSetupList", err)
			return
		}
		msg.SRBsToBeSetupList = []SRBsToBeSetupItem{}
		for _, i := range tmp.Value {
			msg.SRBsToBeSetupList = append(msg.SRBsToBeSetupList, *i)
		}
	case ProtocolIEID_DRBsToBeSetupList:
		tmp := Sequence[*DRBsToBeSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		fn := func() *DRBsToBeSetupItem { return new(DRBsToBeSetupItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read DRBsToBeSetupList", err)
			return
		}
		msg.DRBsToBeSetupList = []DRBsToBeSetupItem{}
		for _, i := range tmp.Value {
			msg.DRBsToBeSetupList = append(msg.DRBsToBeSetupList, *i)
		}
	case ProtocolIEID_InactivityMonitoringRequest:
		var tmp InactivityMonitoringRequest
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read InactivityMonitoringRequest", err)
			return
		}
		msg.InactivityMonitoringRequest = &tmp
	case ProtocolIEID_RATFrequencyPriorityInformation:
		var tmp RATFrequencyPriorityInformation
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RATFrequencyPriorityInformation", err)
			return
		}
		msg.RATFrequencyPriorityInformation = &tmp
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
	case ProtocolIEID_MaskedIMEISV:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 64, Ub: 64},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read MaskedIMEISV", err)
			return
		}
		msg.MaskedIMEISV = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	case ProtocolIEID_ServingPLMN:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 3, Ub: 3},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ServingPLMN", err)
			return
		}
		msg.ServingPLMN = tmp.Value
	case ProtocolIEID_GNBDUUEAMBRUL:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read GNBDUUEAMBRUL", err)
			return
		}
		msg.GNBDUUEAMBRUL = (*int64)(&tmp.Value)
	case ProtocolIEID_RRCDeliveryStatusRequest:
		var tmp RRCDeliveryStatusRequest
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RRCDeliveryStatusRequest", err)
			return
		}
		msg.RRCDeliveryStatusRequest = &tmp
	case ProtocolIEID_ResourceCoordinationTransferInformation:
		var tmp ResourceCoordinationTransferInformation
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ResourceCoordinationTransferInformation", err)
			return
		}
		msg.ResourceCoordinationTransferInformation = &tmp
	case ProtocolIEID_ServingCellMO:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 64},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ServingCellMO", err)
			return
		}
		msg.ServingCellMO = (*int64)(&tmp.Value)
	case ProtocolIEID_NewgNBCUUEF1APID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4294967295},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read NewgNBCUUEF1APID", err)
			return
		}
		msg.NewgNBCUUEF1APID = (*int64)(&tmp.Value)
	case ProtocolIEID_RANUEID:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 8, Ub: 8},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RANUEID", err)
			return
		}
		msg.RANUEID = tmp.Value
	case ProtocolIEID_TraceActivation:
		var tmp TraceActivation
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read TraceActivation", err)
			return
		}
		msg.TraceActivation = &tmp
	case ProtocolIEID_AdditionalRRMPriorityIndex:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 32, Ub: 32},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read AdditionalRRMPriorityIndex", err)
			return
		}
		msg.AdditionalRRMPriorityIndex = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	case ProtocolIEID_BHChannelsToBeSetupList:
		tmp := Sequence[*BHChannelsToBeSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofBHRLCChannels},
			ext: false,
		}
		fn := func() *BHChannelsToBeSetupItem { return new(BHChannelsToBeSetupItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read BHChannelsToBeSetupList", err)
			return
		}
		msg.BHChannelsToBeSetupList = []BHChannelsToBeSetupItem{}
		for _, i := range tmp.Value {
			msg.BHChannelsToBeSetupList = append(msg.BHChannelsToBeSetupList, *i)
		}
	case ProtocolIEID_ConfiguredBAPAddress:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 10, Ub: 10},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ConfiguredBAPAddress", err)
			return
		}
		msg.ConfiguredBAPAddress = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	case ProtocolIEID_NRV2XServicesAuthorized:
		var tmp NRV2XServicesAuthorized
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read NRV2XServicesAuthorized", err)
			return
		}
		msg.NRV2XServicesAuthorized = &tmp
	case ProtocolIEID_LTEV2XServicesAuthorized:
		var tmp LTEV2XServicesAuthorized
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read LTEV2XServicesAuthorized", err)
			return
		}
		msg.LTEV2XServicesAuthorized = &tmp
	case ProtocolIEID_NRUESidelinkAggregateMaximumBitrate:
		var tmp NRUESidelinkAggregateMaximumBitrate
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read NRUESidelinkAggregateMaximumBitrate", err)
			return
		}
		msg.NRUESidelinkAggregateMaximumBitrate = &tmp
	case ProtocolIEID_LTEUESidelinkAggregateMaximumBitrate:
		var tmp LTEUESidelinkAggregateMaximumBitrate
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read LTEUESidelinkAggregateMaximumBitrate", err)
			return
		}
		msg.LTEUESidelinkAggregateMaximumBitrate = &tmp
	case ProtocolIEID_PC5LinkAMBR:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read PC5LinkAMBR", err)
			return
		}
		msg.PC5LinkAMBR = int64(tmp.Value)
	case ProtocolIEID_SLDRBsToBeSetupList:
		tmp := Sequence[*SLDRBsToBeSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		fn := func() *SLDRBsToBeSetupItem { return new(SLDRBsToBeSetupItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SLDRBsToBeSetupList", err)
			return
		}
		msg.SLDRBsToBeSetupList = []SLDRBsToBeSetupItem{}
		for _, i := range tmp.Value {
			msg.SLDRBsToBeSetupList = append(msg.SLDRBsToBeSetupList, *i)
		}
	case ProtocolIEID_ConditionalInterDUMobilityInformation:
		var tmp ConditionalInterDUMobilityInformation
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ConditionalInterDUMobilityInformation", err)
			return
		}
		msg.ConditionalInterDUMobilityInformation = tmp
	case ProtocolIEID_ManagementBasedMDTPLMNList:
		tmp := Sequence[*PLMNIdentity]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMDTPLMNs},
			ext: false,
		}
		fn := func() *PLMNIdentity { return new(PLMNIdentity) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read ManagementBasedMDTPLMNList", err)
			return
		}
		msg.ManagementBasedMDTPLMNList = []PLMNIdentity{}
		for _, i := range tmp.Value {
			msg.ManagementBasedMDTPLMNList = append(msg.ManagementBasedMDTPLMNList, *i)
		}
	case ProtocolIEID_ServingNID:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 44, Ub: 44},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ServingNID", err)
			return
		}
		msg.ServingNID = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	case ProtocolIEID_F1CTransferPath:
		var tmp F1CTransferPath
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read F1CTransferPath", err)
			return
		}
		msg.F1CTransferPath = &tmp
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
