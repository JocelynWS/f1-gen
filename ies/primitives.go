package ies

import (
	"github.com/lvdund/ngap/aper"
)

type AdditionalRRMPriorityIndex struct {
	Value aper.BitString `lb:32,ub:32`
}

type AveragingWindow struct {
	Value int64 `lb:0,ub:4095,ext`
}

type BAPAddress struct {
	Value aper.BitString `lb:10,ub:10`
}

type BAPPathID struct {
	Value aper.BitString `lb:10,ub:10`
}

type BHRLCChannelID struct {
	Value aper.BitString `lb:16,ub:16`
}

type BitRate struct {
	Value int64 `lb:0,ub:4000000000000,ext`
}

type BurstArrivalTime struct {
	Value []byte
}

type CRNTI struct {
	Value int64 `lb:0,ub:65535,ext`
}

type CGConfig struct {
	Value []byte
}

type CGConfigInfo struct {
	Value []byte
}

type CHOProbability struct {
	Value int64 `lb:1,ub:100`
}

type CPTrafficType struct {
	Value int64 `lb:1,ub:3,ext`
}

type CellPortionID struct {
	Value int64 `lb:0,ub:4095,ext`
}

type CellCapacityClassValue struct {
	Value int64 `lb:1,ub:100,ext`
}

type CellGroupConfig struct {
	Value []byte
}

type ConfiguredEPSTAC struct {
	Value []byte `lb:2,ub:2`
}

type CoordinateID struct {
	Value int64 `lb:0,ub:511,ext`
}

type DRBID struct {
	Value int64 `lb:1,ub:32,ext`
}

type DRXConfig struct {
	Value []byte
}

type DRXLongCycleStartOffset struct {
	Value int64 `lb:0,ub:10239`
}

type DUtoCURRCContainer struct {
	Value []byte
}

type EUTRACellID struct {
	Value aper.BitString `lb:28,ub:28`
}

type EUTRANRCellResourceCoordinationReqContainer struct {
	Value []byte
}

type EUTRANRCellResourceCoordinationReqAckContainer struct {
	Value []byte
}

type ExtendedEARFCN struct {
	Value int64 `lb:0,ub:262143`
}

type ExtendedPacketDelayBudget struct {
	Value int64 `lb:1,ub:65535,ext`
}

type FiveGSTAC struct {
	Value []byte `lb:3,ub:3`
}

type GNBCUUEF1APID struct {
	Value int64 `lb:0,ub:4294967295`
}

type GNBDUID struct {
	Value int64 `lb:0,ub:68719476735`
}

type GNBDUUEF1APID struct {
	Value int64 `lb:0,ub:4294967295`
}

type GNBCUMeasurementID struct {
	Value int64 `lb:0,ub:4095,ext`
}

type GNBDUMeasurementID struct {
	Value int64 `lb:0,ub:4095,ext`
}

type GNBSetID struct {
	Value aper.BitString `lb:22,ub:22`
}

type GTPTEID struct {
	Value []byte `lb:4,ub:4`
}

type HandoverPreparationInformation struct {
	Value []byte
}

type InterfacesToTrace struct {
	Value aper.BitString `lb:8,ub:8`
}

type LCID struct {
	Value int64 `lb:1,ub:32,ext`
}

type LMFMeasurementID struct {
	Value int64 `lb:1,ub:65536,ext`
}

type LMFUEMeasurementID struct {
	Value int64 `lb:1,ub:256,ext`
}

type M7period struct {
	Value int64 `lb:1,ub:60,ext`
}

type MIBmessage struct {
	Value []byte
}

type MaskedIMEISV struct {
	Value aper.BitString `lb:64,ub:64`
}

type MaxDataBurstVolume struct {
	Value int64 `lb:0,ub:4095,ext`
}

type MaxPacketLossRate struct {
	Value int64 `lb:0,ub:1000`
}

type MeasConfig struct {
	Value []byte
}

type MeasGapConfig struct {
	Value []byte
}

type MeasGapSharingConfig struct {
	Value []byte
}

type MeasurementTimingConfiguration struct {
	Value []byte
}

type MeasurementsToActivate struct {
	Value aper.BitString `lb:8,ub:8`
}

type MessageIdentifier struct {
	Value aper.BitString `lb:16,ub:16`
}

type NRCellIdentity struct {
	Value aper.BitString `lb:36,ub:36`
}

type NRUERLFReportContainer struct {
	Value []byte
}

type NoofDownlinkSymbols struct {
	Value int64 `lb:0,ub:14`
}

type NoofUplinkSymbols struct {
	Value int64 `lb:0,ub:14`
}

type NumberOfBroadcasts struct {
	Value int64 `lb:0,ub:65535`
}

type NumberofActiveUEs struct {
	Value int64 `lb:0,ub:16777215,ext`
}

type NumberofBroadcastRequest struct {
	Value int64 `lb:0,ub:65535`
}

type OffsetToPointA struct {
	Value int64 `lb:0,ub:2199,ext`
}

type PC5QoSFlowIdentifier struct {
	Value int64 `lb:1,ub:2048`
}

type PDCCHBlindDetectionSCG struct {
	Value []byte
}

type PDCPSN struct {
	Value int64 `lb:0,ub:4095`
}

type PDUSessionID struct {
	Value int64 `lb:0,ub:255`
}

type PERExponent struct {
	Value int64 `lb:0,ub:9,ext`
}

type PERScalar struct {
	Value int64 `lb:0,ub:9,ext`
}

type PRSResourceID struct {
	Value int64 `lb:0,ub:63`
}

type Periodicity struct {
	Value int64 `lb:0,ub:640000,ext`
}

type PhInfoMCG struct {
	Value []byte
}

type PhInfoSCG struct {
	Value []byte
}

type PortNumber struct {
	Value aper.BitString `lb:16,ub:16`
}

type PosAssistanceInformation struct {
	Value []byte
}

type PosAssistanceInformationFailureList struct {
	Value []byte
}

type PriorityLevel struct {
	Value int64 `lb:0,ub:15`
}

type ProtectedEUTRAResourceIndication struct {
	Value []byte
}

type ProtocolExtensionID struct {
	Value int64 `lb:0,ub:65535`
}

type QCI struct {
	Value int64 `lb:0,ub:255`
}

type QoSFlowIdentifier struct {
	Value int64 `lb:0,ub:63`
}

type QoSParaSetIndex struct {
	Value int64 `lb:1,ub:8,ext`
}

type QoSParaSetNotifyIndex struct {
	Value int64 `lb:0,ub:8,ext`
}

type RACHConfigCommon struct {
	Value []byte
}

type RACHConfigCommonIAB struct {
	Value []byte
}

type RACHReportContainer struct {
	Value []byte
}

type RANMeasurementID struct {
	Value int64 `lb:1,ub:65536,ext`
}

type RANUEMeasurementID struct {
	Value int64 `lb:1,ub:256,ext`
}

type RANAC struct {
	Value int64 `lb:0,ub:255`
}

type RANUEID struct {
	Value []byte `lb:8,ub:8`
}

type RATFrequencySelectionPriority struct {
	Value int64 `lb:1,ub:256,ext`
}

type RRCContainer struct {
	Value []byte
}

type RRCContainerRRCSetupComplete struct {
	Value []byte
}

type ReferenceSFN struct {
	Value int64 `lb:0,ub:1023`
}

type ReferenceTime struct {
	Value []byte
}

type RepetitionPeriod struct {
	Value int64 `lb:0,ub:131071,ext`
}

type ReportCharacteristics struct {
	Value aper.BitString `lb:32,ub:32`
}

type ReportingPeriodicityValue struct {
	Value int64 `lb:0,ub:512,ext`
}

type RequestedPDCCHBlindDetectionSCG struct {
	Value []byte
}

type RequestedBandCombinationIndex struct {
	Value []byte
}

type RequestedFeatureSetEntryIndex struct {
	Value []byte
}

type RequestedPMaxFR2 struct {
	Value []byte
}

type ResourceCoordinationTransferContainer struct {
	Value []byte
}

type RoutingID struct {
	Value []byte
}

type SCellIndex struct {
	Value int64 `lb:1,ub:31,ext`
}

type SIB1message struct {
	Value []byte
}

type SIB10message struct {
	Value []byte
}

type SIB12message struct {
	Value []byte
}

type SIB13message struct {
	Value []byte
}

type SIB14message struct {
	Value []byte
}

type SIBTypePWS struct {
	Value int64 `lb:6,ub:8,ext`
}

type SItype struct {
	Value int64 `lb:1,ub:32,ext`
}

type SLConfigDedicatedEUTRAInfo struct {
	Value []byte
}

type SLPHYMACRLCConfig struct {
	Value []byte
}

type SLDRBID struct {
	Value int64 `lb:1,ub:512,ext`
}

type SRBID struct {
	Value int64 `lb:0,ub:3,ext`
}

type SRSResourceID struct {
	Value int64 `lb:0,ub:63`
}

type SRSResourceSetID struct {
	Value int64 `lb:0,ub:15,ext`
}

type SSBtransmissionTimingOffset struct {
	Value int64 `lb:0,ub:127,ext`
}

type SelectedBandCombinationIndex struct {
	Value []byte
}

type SelectedFeatureSetEntryIndex struct {
	Value []byte
}

type SerialNumber struct {
	Value aper.BitString `lb:16,ub:16`
}

type ServCellIndex struct {
	Value int64 `lb:0,ub:31,ext`
}

type ServingCellMO struct {
	Value int64 `lb:1,ub:64,ext`
}

type ShortDRXCycleTimer struct {
	Value int64 `lb:1,ub:16`
}

type SlotNumber struct {
	Value int64 `lb:0,ub:79`
}

type SrsFrequency struct {
	Value int64 `lb:0,ub:3279165`
}

type SubscriberProfileIDforRFP struct {
	Value int64 `lb:1,ub:256,ext`
}

type SystemFrameNumber struct {
	Value int64 `lb:0,ub:1023`
}

type SystemInformationAreaID struct {
	Value aper.BitString `lb:24,ub:24`
}

type TDDULDLConfigCommonNR struct {
	Value []byte
}

type TransactionID struct {
	Value int64 `lb:0,ub:255,ext`
}

type UACReductionIndication struct {
	Value int64 `lb:0,ub:100`
}

type UECapabilityRATContainerList struct {
	Value []byte
}

type UEAssistanceInformation struct {
	Value []byte
}

type UEAssistanceInformationEUTRA struct {
	Value []byte
}

type ULSRSRSRP struct {
	Value int64 `lb:0,ub:126`
}

type Uncertainty struct {
	Value int64 `lb:0,ub:32767,ext`
}

type UplinkTxDirectCurrentListInformation struct {
	Value []byte
}
