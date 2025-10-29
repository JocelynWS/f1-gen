package f1ap

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UEContextModificationRequest struct {
	GNBCUUEF1APID                           int64                                    `lb:0,ub:4294967295,mandatory,reject`
	GNBDUUEF1APID                           int64                                    `lb:0,ub:4294967295,mandatory,reject`
	SpCellID                                *NRCGI                                   `optional,ignore`
	ServCellIndex                           *int64                                   `lb:0,ub:31,optional,reject`
	SpCellULConfigured                      *CellULConfigured                        `optional,ignore`
	DRXCycle                                *DRXCycle                                `optional,ignore`
	CUtoDURRCInformation                    *CUtoDURRCInformation                    `optional,reject`
	TransmissionActionIndicator             *TransmissionActionIndicator             `optional,ignore`
	ResourceCoordinationTransferContainer   []byte                                   `lb:0,ub:0,optional,ignore`
	RRCReconfigurationCompleteIndicator     *RRCReconfigurationCompleteIndicator     `optional,ignore`
	RRCContainer                            []byte                                   `lb:0,ub:0,optional,reject`
	SCellToBeSetupModList                   []SCellToBeSetupModItem                  `lb:1,ub:maxnoofSCells,optional,ignore`
	SCellToBeRemovedList                    []SCellToBeRemovedItem                   `lb:1,ub:maxnoofSCells,optional,ignore`
	SRBsToBeSetupModList                    []SRBsToBeSetupModItem                   `lb:1,ub:maxnoofSRBs,optional,reject`
	DRBsToBeSetupModList                    []DRBsToBeSetupModItem                   `lb:1,ub:maxnoofDRBs,optional,reject`
	DRBsToBeModifiedList                    []DRBsToBeModifiedItem                   `lb:1,ub:maxnoofDRBs,optional,reject`
	SRBsToBeReleasedList                    []SRBsToBeReleasedItem                   `lb:1,ub:maxnoofSRBs,optional,reject`
	DRBsToBeReleasedList                    []DRBsToBeReleasedItem                   `lb:1,ub:maxnoofDRBs,optional,reject`
	InactivityMonitoringRequest             *InactivityMonitoringRequest             `optional,reject`
	RATFrequencyPriorityInformation         *RATFrequencyPriorityInformation         `optional,reject`
	DRXConfigurationIndicator               *DRXConfigurationIndicator               `optional,ignore`
	RLCFailureIndication                    *RLCFailureIndication                    `optional,ignore`
	UplinkTxDirectCurrentListInformation    []byte                                   `lb:0,ub:0,optional,ignore`
	GNBDUConfigurationQuery                 *GNBDUConfigurationQuery                 `optional,reject`
	GNBDUUEAMBRUL                           *int64                                   `lb:0,ub:4000000000000,optional,ignore,valueExt`
	ExecuteDuplication                      *ExecuteDuplication                      `mandatory,ignore`
	RRCDeliveryStatusRequest                *RRCDeliveryStatusRequest                `optional,ignore`
	ResourceCoordinationTransferInformation *ResourceCoordinationTransferInformation `optional,ignore`
	ServingCellMO                           *int64                                   `lb:1,ub:64,optional,ignore`
	NeedforGap                              *NeedforGap                              `optional,ignore`
	FullConfiguration                       *FullConfiguration                       `optional,reject`
	AdditionalRRMPriorityIndex              *aper.BitString                          `lb:32,ub:32,optional,ignore`
	LowerLayerPresenceStatusChange          *LowerLayerPresenceStatusChange          `optional,ignore`
	BHChannelsToBeSetupModList              []BHChannelsToBeSetupModItem             `lb:1,ub:maxnoofBHRLCChannels,optional,reject`
	BHChannelsToBeModifiedList              []BHChannelsToBeModifiedItem             `lb:1,ub:maxnoofBHRLCChannels,optional,reject`
	BHChannelsToBeReleasedList              []BHChannelsToBeReleasedItem             `lb:1,ub:maxnoofBHRLCChannels,optional,reject`
	NRV2XServicesAuthorized                 *NRV2XServicesAuthorized                 `optional,ignore`
	LTEV2XServicesAuthorized                *LTEV2XServicesAuthorized                `optional,ignore`
	NRUESidelinkAggregateMaximumBitrate     *NRUESidelinkAggregateMaximumBitrate     `optional,ignore`
	LTEUESidelinkAggregateMaximumBitrate    *LTEUESidelinkAggregateMaximumBitrate    `optional,ignore`
	PC5LinkAMBR                             int64                                    `lb:0,ub:4000000000000,mandatory,ignore,valueExt`
	SLDRBsToBeSetupModList                  []SLDRBsToBeSetupModItem                 `lb:1,ub:maxnoofSLDRBs,optional,reject`
	SLDRBsToBeModifiedList                  []SLDRBsToBeModifiedItem                 `lb:1,ub:maxnoofSLDRBs,optional,reject`
	SLDRBsToBeReleasedList                  []SLDRBsToBeReleasedItem                 `lb:1,ub:maxnoofSLDRBs,optional,reject`
	ConditionalIntraDUMobilityInformation   *ConditionalIntraDUMobilityInformation   `mandatory,reject`
}

func (msg *UEContextModificationRequest) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("UEContextModificationRequest"), err)
		return
	}
	return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_UEContextModification, Criticality_PresentReject, ies)
}
func (msg *UEContextModificationRequest) toIes() (ies []F1apMessageIE, err error) {
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
	if msg.SpCellID != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SpCellID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.SpCellID,
		})
	}
	if msg.ServCellIndex != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ServCellIndex},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 31},
				ext:   false,
				Value: aper.Integer(*msg.ServCellIndex),
			}})
	}
	if msg.SpCellULConfigured != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SpCellULConfigured},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.SpCellULConfigured,
		})
	}
	if msg.DRXCycle != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DRXCycle},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.DRXCycle,
		})
	}
	if msg.CUtoDURRCInformation != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CUtoDURRCInformation},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.CUtoDURRCInformation,
		})
	}
	if msg.TransmissionActionIndicator != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TransmissionActionIndicator},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.TransmissionActionIndicator,
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
	if msg.RRCReconfigurationCompleteIndicator != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RRCReconfigurationCompleteIndicator},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RRCReconfigurationCompleteIndicator,
		})
	}
	if msg.RRCContainer != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RRCContainer},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   false,
				Value: msg.RRCContainer,
			}})
	}
	if len(msg.SCellToBeSetupModList) > 0 {
		tmp_SCellToBeSetupModList := Sequence[*SCellToBeSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSCells},
			ext: false,
		}
		for _, i := range msg.SCellToBeSetupModList {
			tmp_SCellToBeSetupModList.Value = append(tmp_SCellToBeSetupModList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SCellToBeSetupModList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_SCellToBeSetupModList,
		})
	}
	if len(msg.SCellToBeRemovedList) > 0 {
		tmp_SCellToBeRemovedList := Sequence[*SCellToBeRemovedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSCells},
			ext: false,
		}
		for _, i := range msg.SCellToBeRemovedList {
			tmp_SCellToBeRemovedList.Value = append(tmp_SCellToBeRemovedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SCellToBeRemovedList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_SCellToBeRemovedList,
		})
	}
	if len(msg.SRBsToBeSetupModList) > 0 {
		tmp_SRBsToBeSetupModList := Sequence[*SRBsToBeSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSRBs},
			ext: false,
		}
		for _, i := range msg.SRBsToBeSetupModList {
			tmp_SRBsToBeSetupModList.Value = append(tmp_SRBsToBeSetupModList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SRBsToBeSetupModList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_SRBsToBeSetupModList,
		})
	}
	if len(msg.DRBsToBeSetupModList) > 0 {
		tmp_DRBsToBeSetupModList := Sequence[*DRBsToBeSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		for _, i := range msg.DRBsToBeSetupModList {
			tmp_DRBsToBeSetupModList.Value = append(tmp_DRBsToBeSetupModList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DRBsToBeSetupModList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_DRBsToBeSetupModList,
		})
	}
	if len(msg.DRBsToBeModifiedList) > 0 {
		tmp_DRBsToBeModifiedList := Sequence[*DRBsToBeModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		for _, i := range msg.DRBsToBeModifiedList {
			tmp_DRBsToBeModifiedList.Value = append(tmp_DRBsToBeModifiedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DRBsToBeModifiedList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_DRBsToBeModifiedList,
		})
	}
	if len(msg.SRBsToBeReleasedList) > 0 {
		tmp_SRBsToBeReleasedList := Sequence[*SRBsToBeReleasedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSRBs},
			ext: false,
		}
		for _, i := range msg.SRBsToBeReleasedList {
			tmp_SRBsToBeReleasedList.Value = append(tmp_SRBsToBeReleasedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SRBsToBeReleasedList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_SRBsToBeReleasedList,
		})
	}
	if len(msg.DRBsToBeReleasedList) > 0 {
		tmp_DRBsToBeReleasedList := Sequence[*DRBsToBeReleasedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		for _, i := range msg.DRBsToBeReleasedList {
			tmp_DRBsToBeReleasedList.Value = append(tmp_DRBsToBeReleasedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DRBsToBeReleasedList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_DRBsToBeReleasedList,
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
	if msg.DRXConfigurationIndicator != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DRXConfigurationIndicator},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.DRXConfigurationIndicator,
		})
	}
	if msg.RLCFailureIndication != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RLCFailureIndication},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RLCFailureIndication,
		})
	}
	if msg.UplinkTxDirectCurrentListInformation != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UplinkTxDirectCurrentListInformation},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   false,
				Value: msg.UplinkTxDirectCurrentListInformation,
			}})
	}
	if msg.GNBDUConfigurationQuery != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GNBDUConfigurationQuery},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.GNBDUConfigurationQuery,
		})
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
	if msg.NeedforGap != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NeedForGap},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.NeedforGap,
		})
	}
	if msg.FullConfiguration != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_FullConfiguration},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.FullConfiguration,
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
	if msg.LowerLayerPresenceStatusChange != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_LowerLayerPresenceStatusChange},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.LowerLayerPresenceStatusChange,
		})
	}
	if len(msg.BHChannelsToBeSetupModList) > 0 {
		tmp_BHChannelsToBeSetupModList := Sequence[*BHChannelsToBeSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofBHRLCChannels},
			ext: false,
		}
		for _, i := range msg.BHChannelsToBeSetupModList {
			tmp_BHChannelsToBeSetupModList.Value = append(tmp_BHChannelsToBeSetupModList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_BHChannelsToBeSetupModList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_BHChannelsToBeSetupModList,
		})
	}
	if len(msg.BHChannelsToBeModifiedList) > 0 {
		tmp_BHChannelsToBeModifiedList := Sequence[*BHChannelsToBeModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofBHRLCChannels},
			ext: false,
		}
		for _, i := range msg.BHChannelsToBeModifiedList {
			tmp_BHChannelsToBeModifiedList.Value = append(tmp_BHChannelsToBeModifiedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_BHChannelsToBeModifiedList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_BHChannelsToBeModifiedList,
		})
	}
	if len(msg.BHChannelsToBeReleasedList) > 0 {
		tmp_BHChannelsToBeReleasedList := Sequence[*BHChannelsToBeReleasedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofBHRLCChannels},
			ext: false,
		}
		for _, i := range msg.BHChannelsToBeReleasedList {
			tmp_BHChannelsToBeReleasedList.Value = append(tmp_BHChannelsToBeReleasedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_BHChannelsToBeReleasedList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_BHChannelsToBeReleasedList,
		})
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
	if msg.PC5LinkAMBR != 0 {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PC5LinkAMBR},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 4000000000000},
				ext:   true,
				Value: aper.Integer(msg.PC5LinkAMBR),
			},
		})
	}

	if len(msg.SLDRBsToBeSetupModList) > 0 {
		tmp_SLDRBsToBeSetupModList := Sequence[*SLDRBsToBeSetupModItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		for _, i := range msg.SLDRBsToBeSetupModList {
			tmp_SLDRBsToBeSetupModList.Value = append(tmp_SLDRBsToBeSetupModList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SLDRBsToBeSetupModList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_SLDRBsToBeSetupModList,
		})
	}
	if len(msg.SLDRBsToBeModifiedList) > 0 {
		tmp_SLDRBsToBeModifiedList := Sequence[*SLDRBsToBeModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		for _, i := range msg.SLDRBsToBeModifiedList {
			tmp_SLDRBsToBeModifiedList.Value = append(tmp_SLDRBsToBeModifiedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SLDRBsToBeModifiedList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_SLDRBsToBeModifiedList,
		})
	}
	if len(msg.SLDRBsToBeReleasedList) > 0 {
		tmp_SLDRBsToBeReleasedList := Sequence[*SLDRBsToBeReleasedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		for _, i := range msg.SLDRBsToBeReleasedList {
			tmp_SLDRBsToBeReleasedList.Value = append(tmp_SLDRBsToBeReleasedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SLDRBsToBeReleasedList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_SLDRBsToBeReleasedList,
		})
	}
	if msg.ConditionalIntraDUMobilityInformation != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ConditionalIntraDUMobilityInformation},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.ConditionalIntraDUMobilityInformation,
		})
	}
	return
}
func (msg *UEContextModificationRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("UEContextModificationRequest"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := UEContextModificationRequestDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_ExecuteDuplication]; !ok {
		err = fmt.Errorf("Mandatory field ExecuteDuplication is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_ExecuteDuplication},
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
	if _, ok := decoder.list[ProtocolIEID_ConditionalIntraDUMobilityInformation]; !ok {
		err = fmt.Errorf("Mandatory field ConditionalIntraDUMobilityInformation is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_ConditionalIntraDUMobilityInformation},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type UEContextModificationRequestDecoder struct {
	msg      *UEContextModificationRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *UEContextModificationRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_SpCellID:
		var tmp NRCGI
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SpCellID", err)
			return
		}
		msg.SpCellID = &tmp
	case ProtocolIEID_ServCellIndex:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 31},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ServCellIndex", err)
			return
		}
		msg.ServCellIndex = (*int64)(&tmp.Value)
	case ProtocolIEID_SpCellULConfigured:
		var tmp CellULConfigured
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SpCellULConfigured", err)
			return
		}
		msg.SpCellULConfigured = &tmp
	case ProtocolIEID_DRXCycle:
		var tmp DRXCycle
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read DRXCycle", err)
			return
		}
		msg.DRXCycle = &tmp
	case ProtocolIEID_CUtoDURRCInformation:
		var tmp CUtoDURRCInformation
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read CUtoDURRCInformation", err)
			return
		}
		msg.CUtoDURRCInformation = &tmp
	default:
	}
	return
}
