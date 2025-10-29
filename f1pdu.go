package f1ap

import "github.com/JocelynWS/f1-gen/ies"

type F1apPdu struct {
	Present uint8
	Message F1apMessage
}

type F1apMessage struct {
	ProcedureCode ies.ProcedureCode
	Criticality   ies.Criticality
	Msg           MessageUnmarshaller
}

func createMessage(present uint8, procedureCode ies.ProcedureCode) MessageUnmarshaller {
	switch present {
	case ies.F1apPduInitiatingMessage:
		switch int64(procedureCode.Value) {
		case ies.ProcedureCode_Reset:
			return new(ies.Reset)
		case ies.ProcedureCode_F1Setup:
			return new(ies.F1SetupRequest)
		case ies.ProcedureCode_ErrorIndication:
			return new(ies.ErrorIndication)
		case ies.ProcedureCode_GNBDUConfigurationUpdate:
			return new(ies.GNBDUConfigurationUpdate)
		case ies.ProcedureCode_GNBCUConfigurationUpdate:
			return new(ies.GNBCUConfigurationUpdate)
		case ies.ProcedureCode_UEContextSetup:
			return new(ies.UEContextSetupRequest)
		case ies.ProcedureCode_UEContextRelease:
			return new(ies.UEContextReleaseCommand)
		case ies.ProcedureCode_UEContextModification:
			return new(ies.UEContextModificationRequest)
		case ies.ProcedureCode_UEContextModificationRequired:
			return new(ies.UEContextModificationRequired)
		case ies.ProcedureCode_UEContextReleaseRequest:
			return new(ies.UEContextReleaseRequest)
		case ies.ProcedureCode_InitialULRRCMessageTransfer:
			return new(ies.InitialULRRCMessageTransfer)
		case ies.ProcedureCode_DLRRCMessageTransfer:
			return new(ies.DLRRCMessageTransfer)
		case ies.ProcedureCode_ULRRCMessageTransfer:
			return new(ies.ULRRCMessageTransfer)
		case ies.ProcedureCode_UEInactivityNotification:
			return new(ies.UEInactivityNotification)
		case ies.ProcedureCode_GNBDUResourceCoordination:
			return new(ies.GNBDUResourceCoordinationRequest)
		case ies.ProcedureCode_SystemInformationDeliveryCommand:
			return new(ies.SystemInformationDeliveryCommand)
		case ies.ProcedureCode_Paging:
			return new(ies.Paging)
		case ies.ProcedureCode_Notify:
			return new(ies.Notify)
		case ies.ProcedureCode_WriteReplaceWarning:
			return new(ies.WriteReplaceWarningRequest)
		case ies.ProcedureCode_PWSCancel:
			return new(ies.PWSCancelRequest)
		case ies.ProcedureCode_PWSRestartIndication:
			return new(ies.PWSRestartIndication)
		case ies.ProcedureCode_PWSFailureIndication:
			return new(ies.PWSFailureIndication)
		case ies.ProcedureCode_GNBDUStatusIndication:
			return new(ies.GNBDUStatusIndication)
		case ies.ProcedureCode_RRCDeliveryReport:
			return new(ies.RRCDeliveryReport)
		case ies.ProcedureCode_F1Removal:
			return new(ies.F1RemovalRequest)
		case ies.ProcedureCode_NetworkAccessRateReduction:
			return new(ies.NetworkAccessRateReduction)
		case ies.ProcedureCode_TraceStart:
			return new(ies.TraceStart)
		case ies.ProcedureCode_DeactivateTrace:
			return new(ies.DeactivateTrace)
		case ies.ProcedureCode_DUCURadioInformationTransfer:
			return new(ies.DUCURadioInformationTransfer)
		case ies.ProcedureCode_CUDURadioInformationTransfer:
			return new(ies.CUDURadioInformationTransfer)
		case ies.ProcedureCode_BAPMappingConfiguration:
			return new(ies.BAPMappingConfiguration)
		case ies.ProcedureCode_GNBDUResourceConfiguration:
			return new(ies.GNBDUResourceConfiguration)
		case ies.ProcedureCode_IABTNLAddressAllocation:
			return new(ies.IABTNLAddressRequest)
		case ies.ProcedureCode_IABUPConfigurationUpdate:
			return new(ies.IABUPConfigurationUpdateRequest)
		case ies.ProcedureCode_ResourceStatusReportingInitiation:
			return new(ies.ResourceStatusRequest)
		case ies.ProcedureCode_ResourceStatusReporting:
			return new(ies.ResourceStatusUpdate)
		case ies.ProcedureCode_AccessAndMobilityIndication:
			return new(ies.AccessAndMobilityIndication)
		case ies.ProcedureCode_AccessSuccess:
			return new(ies.AccessSuccess)
		case ies.ProcedureCode_CellTrafficTrace:
			return new(ies.CellTrafficTrace)
		case ies.ProcedureCode_PositioningMeasurementExchange:
			return new(ies.PositioningMeasurementRequest)
		case ies.ProcedureCode_PositioningAssistanceInformationControl:
			return new(ies.PositioningAssistanceInformationControl)
		case ies.ProcedureCode_PositioningAssistanceInformationFeedback:
			return new(ies.PositioningAssistanceInformationFeedback)
		case ies.ProcedureCode_PositioningMeasurementReport:
			return new(ies.PositioningMeasurementReport)
		case ies.ProcedureCode_PositioningMeasurementAbort:
			return new(ies.PositioningMeasurementAbort)
		case ies.ProcedureCode_PositioningMeasurementFailureIndication:
			return new(ies.PositioningMeasurementFailureIndication)
		case ies.ProcedureCode_PositioningMeasurementUpdate:
			return new(ies.PositioningMeasurementUpdate)
		case ies.ProcedureCode_TRPInformationExchange:
			return new(ies.TRPInformationRequest)
		case ies.ProcedureCode_PositioningInformationExchange:
			return new(ies.PositioningInformationRequest)
		case ies.ProcedureCode_PositioningActivation:
			return new(ies.PositioningActivationRequest)
		case ies.ProcedureCode_PositioningDeactivation:
			return new(ies.PositioningDeactivation)
		case ies.ProcedureCode_ECIDMeasurementInitiation:
			return new(ies.ECIDMeasurementInitiationRequest)
		case ies.ProcedureCode_ECIDMeasurementFailureIndication:
			return new(ies.ECIDMeasurementFailureIndication)
		case ies.ProcedureCode_ECIDMeasurementReport:
			return new(ies.ECIDMeasurementReport)
		case ies.ProcedureCode_ECIDMeasurementTermination:
			return new(ies.ECIDMeasurementTerminationCommand)
		case ies.ProcedureCode_PositioningInformationUpdate:
			return new(ies.PositioningInformationUpdate)
		case ies.ProcedureCode_ReferenceTimeInformationReport:
			return new(ies.ReferenceTimeInformationReport)
		case ies.ProcedureCode_ReferenceTimeInformationReportingControl:
			return new(ies.ReferenceTimeInformationReportingControl)
		}

	case ies.F1apPduSuccessfulOutcome:
		switch int64(procedureCode.Value) {
		case ies.ProcedureCode_Reset:
			return new(ies.ResetAcknowledge)
		case ies.ProcedureCode_F1Setup:
			return new(ies.F1SetupResponse)
		case ies.ProcedureCode_GNBDUConfigurationUpdate:
			return new(ies.GNBDUConfigurationUpdateAcknowledge)
		case ies.ProcedureCode_GNBCUConfigurationUpdate:
			return new(ies.GNBCUConfigurationUpdateAcknowledge)
		case ies.ProcedureCode_UEContextSetup:
			return new(ies.UEContextSetupResponse)
		case ies.ProcedureCode_UEContextRelease:
			return new(ies.UEContextReleaseComplete)
		case ies.ProcedureCode_UEContextModification:
			return new(ies.UEContextModificationResponse)
		case ies.ProcedureCode_UEContextModificationRequired:
			return new(ies.UEContextModificationConfirm)
		case ies.ProcedureCode_WriteReplaceWarning:
			return new(ies.WriteReplaceWarningResponse)
		case ies.ProcedureCode_PWSCancel:
			return new(ies.PWSCancelResponse)
		case ies.ProcedureCode_GNBDUResourceCoordination:
			return new(ies.GNBDUResourceCoordinationResponse)
		case ies.ProcedureCode_F1Removal:
			return new(ies.F1RemovalResponse)
		case ies.ProcedureCode_BAPMappingConfiguration:
			return new(ies.BAPMappingConfigurationAcknowledge)
		case ies.ProcedureCode_GNBDUResourceConfiguration:
			return new(ies.GNBDUResourceConfigurationAcknowledge)
		case ies.ProcedureCode_IABTNLAddressAllocation:
			return new(ies.IABTNLAddressResponse)
		case ies.ProcedureCode_IABUPConfigurationUpdate:
			return new(ies.IABUPConfigurationUpdateResponse)
		case ies.ProcedureCode_ResourceStatusReportingInitiation:
			return new(ies.ResourceStatusResponse)
		case ies.ProcedureCode_PositioningMeasurementExchange:
			return new(ies.PositioningMeasurementResponse)
		case ies.ProcedureCode_TRPInformationExchange:
			return new(ies.TRPInformationResponse)
		case ies.ProcedureCode_PositioningInformationExchange:
			return new(ies.PositioningInformationResponse)
		case ies.ProcedureCode_PositioningActivation:
			return new(ies.PositioningActivationResponse)
		case ies.ProcedureCode_ECIDMeasurementInitiation:
			return new(ies.ECIDMeasurementInitiationResponse)
		}

	case ies.F1apPduUnsuccessfulOutcome:
		switch int64(procedureCode.Value) {
		case ies.ProcedureCode_F1Setup:
			return new(ies.F1SetupFailure)
		case ies.ProcedureCode_GNBDUConfigurationUpdate:
			return new(ies.GNBDUConfigurationUpdateFailure)
		case ies.ProcedureCode_GNBCUConfigurationUpdate:
			return new(ies.GNBCUConfigurationUpdateFailure)
		case ies.ProcedureCode_UEContextSetup:
			return new(ies.UEContextSetupFailure)
		case ies.ProcedureCode_UEContextModification:
			return new(ies.UEContextModificationFailure)
		case ies.ProcedureCode_UEContextModificationRequired:
			return new(ies.UEContextModificationRefuse)
		case ies.ProcedureCode_IABUPConfigurationUpdate:
			return new(ies.IABUPConfigurationUpdateFailure)
		case ies.ProcedureCode_ResourceStatusReportingInitiation:
			return new(ies.ResourceStatusFailure)
		case ies.ProcedureCode_PositioningMeasurementExchange:
			return new(ies.PositioningMeasurementFailure)
		case ies.ProcedureCode_TRPInformationExchange:
			return new(ies.TRPInformationFailure)
		case ies.ProcedureCode_PositioningInformationExchange:
			return new(ies.PositioningInformationFailure)
		case ies.ProcedureCode_PositioningActivation:
			return new(ies.PositioningActivationFailure)
		case ies.ProcedureCode_ECIDMeasurementInitiation:
			return new(ies.ECIDMeasurementInitiationFailure)
		}

	default:
		return nil
	}
	return nil
}
