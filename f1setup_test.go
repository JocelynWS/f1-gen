package f1ap

import (
	"fmt"
	"testing"

	"github.com/JocelynWS/f1-gen/ies"
	"github.com/lvdund/ngap/aper"
)

func TestF1SetupRequest(t *testing.T) {
	msg := ies.F1SetupRequest{
		TransactionID: 12,
		GNBDUID:       33,
		GNBDUName:     []byte("DU-UE"),
		GNBDURRCVersion: ies.RRCVersion{
			LatestRRCVersion: aper.BitString{
				Bytes:   []byte{2, 248, 57},
				NumBits: 3,
			},
		},
		GNBDUServedCellsList: []ies.GNBDUServedCellsItem{
			ies.GNBDUServedCellsItem{
				ServedCellInformation: ies.ServedCellInformation{
					NRCGI: ies.NRCGI{
						PLMNIdentity:   []byte{2, 248, 57},
						NRCellIdentity: aper.BitString{Bytes: []byte{0x0F, 0xFF, 0xFF, 0xFF, 0xFF}, NumBits: 36},
					},
					NRPCI:     ies.NRPCI{Value: 123},
					FiveGSTAC: []byte{0, 0, 1},
					ServedPLMNs: []ies.ServedPLMNsItem{
						ies.ServedPLMNsItem{
							PLMNIdentity: []byte{2, 248, 57},
						},
					},
					NRModeInfo: ies.NRModeInfo{
						Choice: ies.NRModeInfoPresentFDD,
						FDD: &ies.FDDInfo{
							ULNRFreqInfo: ies.NRFreqInfo{
								NRARFCN: 1,
								FreqBandListNr: []ies.FreqBandNrItem{
									ies.FreqBandNrItem{
										FreqBandIndicatorNr: 1,
										SupportedSULBandList: []ies.SupportedSULFreqBandItem{
											ies.SupportedSULFreqBandItem{
												FreqBandIndicatorNr: 1,
											},
										},
									},
								},
							},
							DLNRFreqInfo: ies.NRFreqInfo{
								NRARFCN: 44,
								FreqBandListNr: []ies.FreqBandNrItem{
									ies.FreqBandNrItem{
										FreqBandIndicatorNr: 1,
										SupportedSULBandList: []ies.SupportedSULFreqBandItem{
											ies.SupportedSULFreqBandItem{
												FreqBandIndicatorNr: 1,
											},
										},
									},
								},
							},
							ULTransmissionBandwidth: ies.TransmissionBandwidth{
								NRSCS: ies.NRSCS{Value: 1},
								NRNRB: ies.NRNRB{Value: 6},
							},
							DLTransmissionBandwidth: ies.TransmissionBandwidth{
								NRSCS: ies.NRSCS{Value: 2},
								NRNRB: ies.NRNRB{Value: 5},
							},
						},
					},
					MeasurementTimingConfiguration: []byte{1, 2, 3, 13, 43, 245, 12},
				},
			},
		},
	}

	buf, err := F1apEncode(&msg)
	if err != nil {
		fmt.Printf("Encode F1AP PDU: %s", err.Error())
		return
	}

	pdu, err, _ := F1apDecode(buf)
	if err != nil {
		fmt.Printf("Decode F1AP PDU: %s", err.Error())
		return
	}
	decoded := pdu.Message.Msg.(*ies.F1SetupRequest)

	// print all field here
	fmt.Printf("\n================== Decoded Fields =================\n")
	fmt.Printf("TransactionID: %d\n", decoded.TransactionID)
	fmt.Printf("GNBDUID: %d\n", decoded.GNBDUID)
	if decoded.GNBDUName != nil {
		fmt.Printf("GNBDUName: %s\n", string(decoded.GNBDUName))
	} else {
		fmt.Printf("GNBDUName: <nil>\n")
	}
	fmt.Printf("GNBDURRCVersion: Bytes=%v, NumBits=%d\n",
		decoded.GNBDURRCVersion.LatestRRCVersion.Bytes,
		decoded.GNBDURRCVersion.LatestRRCVersion.NumBits)

	if len(decoded.GNBDUServedCellsList) > 0 {
		decodedCell := decoded.GNBDUServedCellsList[0]
		fmt.Printf("\n--- ServedCellInformation ---\n")
		fmt.Printf("NRCGI.PLMNIdentity: %v\n", decodedCell.ServedCellInformation.NRCGI.PLMNIdentity)
		fmt.Printf("NRCGI.NRCellIdentity: Bytes=%v, NumBits=%d\n",
			decodedCell.ServedCellInformation.NRCGI.NRCellIdentity.Bytes,
			decodedCell.ServedCellInformation.NRCGI.NRCellIdentity.NumBits)
		fmt.Printf("NRPCI: %d\n", decodedCell.ServedCellInformation.NRPCI.Value)
		fmt.Printf("FiveGSTAC: %v\n", decodedCell.ServedCellInformation.FiveGSTAC)
		if len(decodedCell.ServedCellInformation.ServedPLMNs) > 0 {
			fmt.Printf("ServedPLMNs[0].PLMNIdentity: %v\n",
				decodedCell.ServedCellInformation.ServedPLMNs[0].PLMNIdentity)
		}
		fmt.Printf("NRModeInfo.Choice: %d\n", decodedCell.ServedCellInformation.NRModeInfo.Choice)

		if decodedCell.ServedCellInformation.NRModeInfo.FDD != nil {
			fmt.Printf("\n--- FDD Info ---\n")
			fmt.Printf("ULNRFreqInfo.NRARFCN: %d\n",
				decodedCell.ServedCellInformation.NRModeInfo.FDD.ULNRFreqInfo.NRARFCN)
			fmt.Printf("ULNRFreqInfo.FreqBandListNr length: %d\n",
				len(decodedCell.ServedCellInformation.NRModeInfo.FDD.ULNRFreqInfo.FreqBandListNr))
			if len(decodedCell.ServedCellInformation.NRModeInfo.FDD.ULNRFreqInfo.FreqBandListNr) > 0 {
				fmt.Printf("ULNRFreqInfo.FreqBandListNr[0].FreqBandIndicatorNr: %d\n",
					decodedCell.ServedCellInformation.NRModeInfo.FDD.ULNRFreqInfo.FreqBandListNr[0].FreqBandIndicatorNr)
				fmt.Printf("ULNRFreqInfo.FreqBandListNr[0].SupportedSULBandList length: %d\n",
					len(decodedCell.ServedCellInformation.NRModeInfo.FDD.ULNRFreqInfo.FreqBandListNr[0].SupportedSULBandList))
			}

			fmt.Printf("DLNRFreqInfo.NRARFCN: %d\n",
				decodedCell.ServedCellInformation.NRModeInfo.FDD.DLNRFreqInfo.NRARFCN)
			fmt.Printf("DLNRFreqInfo.FreqBandListNr length: %d\n",
				len(decodedCell.ServedCellInformation.NRModeInfo.FDD.DLNRFreqInfo.FreqBandListNr))
			if len(decodedCell.ServedCellInformation.NRModeInfo.FDD.DLNRFreqInfo.FreqBandListNr) > 0 {
				fmt.Printf("DLNRFreqInfo.FreqBandListNr[0].FreqBandIndicatorNr: %d\n",
					decodedCell.ServedCellInformation.NRModeInfo.FDD.DLNRFreqInfo.FreqBandListNr[0].FreqBandIndicatorNr)
			}

			fmt.Printf("ULTransmissionBandwidth.NRSCS: %d\n",
				decodedCell.ServedCellInformation.NRModeInfo.FDD.ULTransmissionBandwidth.NRSCS.Value)
			fmt.Printf("ULTransmissionBandwidth.NRNRB: %d\n",
				decodedCell.ServedCellInformation.NRModeInfo.FDD.ULTransmissionBandwidth.NRNRB.Value)
			fmt.Printf("DLTransmissionBandwidth.NRSCS: %d\n",
				decodedCell.ServedCellInformation.NRModeInfo.FDD.DLTransmissionBandwidth.NRSCS.Value)
			fmt.Printf("DLTransmissionBandwidth.NRNRB: %d\n",
				decodedCell.ServedCellInformation.NRModeInfo.FDD.DLTransmissionBandwidth.NRNRB.Value)
		}

		fmt.Printf("\nMeasurementTimingConfiguration: %v\n",
			decodedCell.ServedCellInformation.MeasurementTimingConfiguration)
	}
	fmt.Printf("==============================================\n")
}

func TestF1SetupResponse(t *testing.T) {
	msg := ies.F1SetupResponse{
		TransactionID: 12,
		GNBCUName:     []byte("CU-UE"),
		CellstobeActivatedList: []ies.CellstobeActivatedListItem{
			ies.CellstobeActivatedListItem{
				NRCGI: ies.NRCGI{
					PLMNIdentity:   []byte{2, 248, 57},
					NRCellIdentity: aper.BitString{Bytes: []byte{0x0F, 0xFF, 0xFF, 0xFF, 0xFF}, NumBits: 36},
				},
				NRPCI: &ies.NRPCI{Value: 123},
			},
		},
		GNBCURRCVersion: ies.RRCVersion{
			LatestRRCVersion: aper.BitString{
				Bytes:   []byte{2, 248, 57},
				NumBits: 3,
			},
		},
		TransportLayerAddressInfo: &ies.TransportLayerAddressInfo{
			TransportUPLayerAddressInfoToAddList: []ies.TransportUPLayerAddressInfoToAddItem{
				ies.TransportUPLayerAddressInfoToAddItem{
					IPSecTransportLayerAddress: aper.BitString{Bytes: []byte{0x0F, 0xFF, 0xFF, 0xFF, 0xFF}, NumBits: 36},
					GTPTransportLayerAddressToAdd: []ies.GTPTLAItem{ies.GTPTLAItem{
						GTPTransportLayerAddress: aper.BitString{Bytes: []byte{0x0F, 0xFF, 0xFF, 0xFF, 0xFF}, NumBits: 36},
					}},
				},
			},
			TransportUPLayerAddressInfoToRemoveList: []ies.TransportUPLayerAddressInfoToRemoveItem{
				ies.TransportUPLayerAddressInfoToRemoveItem{
					IPSecTransportLayerAddress: aper.BitString{Bytes: []byte{0x0F, 0xFF, 0xFF, 0xFF, 0xFF}, NumBits: 36},
					GTPTransportLayerAddressToRemove: []ies.GTPTLAItem{ies.GTPTLAItem{
						GTPTransportLayerAddress: aper.BitString{Bytes: []byte{0x0F, 0xFF, 0xFF, 0xFF, 0xFF}, NumBits: 36},
					}},
				},
			},
		},
		ULBHNonUPTrafficMapping: &ies.ULBHNonUPTrafficMapping{
			ULBHNonUPTrafficMappingList: []ies.ULBHNonUPTrafficMappingItem{
				ies.ULBHNonUPTrafficMappingItem{
					NonUPTrafficType: ies.NonUPTrafficType{Value: 2},
					BHInfo: ies.BHInfo{
						BAProutingID: &ies.BAPRoutingID{
							BAPAddress: aper.BitString{Bytes: []byte{0x03, 0xFF}, NumBits: 10},
							BAPPathID:  aper.BitString{Bytes: []byte{0x03, 0xFF}, NumBits: 10},
						},
					},
				},
			},
		},
		BAPAddress: &aper.BitString{Bytes: []byte{0x0F, 0xFF, 0xFF, 0xFF, 0xFF}, NumBits: 36},
		ExtendedGNBDUName: &ies.ExtendedGNBDUName{
			GNBDUNameVisibleString: &ies.GNBDUNameVisibleString{Value: "test1"},
			GNBDUNameUTF8String:    &ies.GNBDUNameUTF8String{Value: "test2"},
		},
	}
	buf, err := F1apEncode(&msg)
	if err != nil {
		fmt.Printf("Encode F1AP PDU: %s", err.Error())
		return
	}

	pdu, err, _ := F1apDecode(buf)
	if err != nil {
		fmt.Printf("Decode F1AP PDU: %s", err.Error())
		return
	}

	fmt.Println("Decode F1AP PDU successfully")
	decoded := pdu.Message.Msg.(*ies.F1SetupResponse)
	fmt.Printf("\n================== Decoded Fields =================\n")
	fmt.Printf("TransactionID: %d\n", decoded.TransactionID)
	if decoded.GNBCUName != nil {
		fmt.Printf("GNBCUName: %s\n", string(decoded.GNBCUName))
	} else {
		fmt.Printf("GNBCUName: <nil>\n")
	}
	fmt.Printf("GNBCURRCVersion: Bytes=%v, NumBits=%d\n",
		decoded.GNBCURRCVersion.LatestRRCVersion.Bytes,
		decoded.GNBCURRCVersion.LatestRRCVersion.NumBits)

	fmt.Printf("\n--- CellstobeActivatedList ---\n")
	fmt.Printf("CellstobeActivatedList length: %d\n", len(decoded.CellstobeActivatedList))
	for i, cell := range decoded.CellstobeActivatedList {
		fmt.Printf("CellstobeActivatedList[%d].NRCGI.PLMNIdentity: %v\n", i, cell.NRCGI.PLMNIdentity)
		fmt.Printf("CellstobeActivatedList[%d].NRCGI.NRCellIdentity: Bytes=%v, NumBits=%d\n",
			i, cell.NRCGI.NRCellIdentity.Bytes, cell.NRCGI.NRCellIdentity.NumBits)
		if cell.NRPCI != nil {
			fmt.Printf("CellstobeActivatedList[%d].NRPCI: %d\n", i, cell.NRPCI.Value)
		} else {
			fmt.Printf("CellstobeActivatedList[%d].NRPCI: <nil>\n", i)
		}
	}

	if decoded.TransportLayerAddressInfo != nil {
		fmt.Printf("\n--- TransportLayerAddressInfo ---\n")
		fmt.Printf("TransportUPLayerAddressInfoToAddList length: %d\n",
			len(decoded.TransportLayerAddressInfo.TransportUPLayerAddressInfoToAddList))
		for i, item := range decoded.TransportLayerAddressInfo.TransportUPLayerAddressInfoToAddList {
			fmt.Printf("TransportUPLayerAddressInfoToAddList[%d].IPSecTransportLayerAddress: Bytes=%v, NumBits=%d\n",
				i, item.IPSecTransportLayerAddress.Bytes, item.IPSecTransportLayerAddress.NumBits)
			fmt.Printf("TransportUPLayerAddressInfoToAddList[%d].GTPTransportLayerAddressToAdd length: %d\n",
				i, len(item.GTPTransportLayerAddressToAdd))
			for j, gtp := range item.GTPTransportLayerAddressToAdd {
				fmt.Printf("TransportUPLayerAddressInfoToAddList[%d].GTPTransportLayerAddressToAdd[%d].GTPTransportLayerAddress: Bytes=%v, NumBits=%d\n",
					i, j, gtp.GTPTransportLayerAddress.Bytes, gtp.GTPTransportLayerAddress.NumBits)
			}
		}
		fmt.Printf("TransportUPLayerAddressInfoToRemoveList length: %d\n",
			len(decoded.TransportLayerAddressInfo.TransportUPLayerAddressInfoToRemoveList))
		for i, item := range decoded.TransportLayerAddressInfo.TransportUPLayerAddressInfoToRemoveList {
			fmt.Printf("TransportUPLayerAddressInfoToRemoveList[%d].IPSecTransportLayerAddress: Bytes=%v, NumBits=%d\n",
				i, item.IPSecTransportLayerAddress.Bytes, item.IPSecTransportLayerAddress.NumBits)
			fmt.Printf("TransportUPLayerAddressInfoToRemoveList[%d].GTPTransportLayerAddressToRemove length: %d\n",
				i, len(item.GTPTransportLayerAddressToRemove))
			for j, gtp := range item.GTPTransportLayerAddressToRemove {
				fmt.Printf("TransportUPLayerAddressInfoToRemoveList[%d].GTPTransportLayerAddressToRemove[%d].GTPTransportLayerAddress: Bytes=%v, NumBits=%d\n",
					i, j, gtp.GTPTransportLayerAddress.Bytes, gtp.GTPTransportLayerAddress.NumBits)
			}
		}
	} else {
		fmt.Printf("\nTransportLayerAddressInfo: <nil>\n")
	}

	if decoded.ULBHNonUPTrafficMapping != nil {
		fmt.Printf("\n--- ULBHNonUPTrafficMapping ---\n")
		fmt.Printf("ULBHNonUPTrafficMappingList length: %d\n",
			len(decoded.ULBHNonUPTrafficMapping.ULBHNonUPTrafficMappingList))
		for i, item := range decoded.ULBHNonUPTrafficMapping.ULBHNonUPTrafficMappingList {
			fmt.Printf("ULBHNonUPTrafficMappingList[%d].NonUPTrafficType: %d\n",
				i, item.NonUPTrafficType.Value)
			if item.BHInfo.BAProutingID != nil {
				fmt.Printf("ULBHNonUPTrafficMappingList[%d].BHInfo.BAProutingID.BAPAddress: Bytes=%v, NumBits=%d\n",
					i, item.BHInfo.BAProutingID.BAPAddress.Bytes, item.BHInfo.BAProutingID.BAPAddress.NumBits)
				fmt.Printf("ULBHNonUPTrafficMappingList[%d].BHInfo.BAProutingID.BAPPathID: Bytes=%v, NumBits=%d\n",
					i, item.BHInfo.BAProutingID.BAPPathID.Bytes, item.BHInfo.BAProutingID.BAPPathID.NumBits)
			} else {
				fmt.Printf("ULBHNonUPTrafficMappingList[%d].BHInfo.BAProutingID: <nil>\n", i)
			}
		}
	} else {
		fmt.Printf("\nULBHNonUPTrafficMapping: <nil>\n")
	}

	if decoded.BAPAddress != nil {
		fmt.Printf("\nBAPAddress: Bytes=%v, NumBits=%d\n",
			decoded.BAPAddress.Bytes, decoded.BAPAddress.NumBits)
	} else {
		fmt.Printf("\nBAPAddress: <nil>\n")
	}

	if decoded.ExtendedGNBDUName != nil {
		fmt.Printf("\n--- ExtendedGNBDUName ---\n")
		if decoded.ExtendedGNBDUName.GNBDUNameVisibleString != nil {
			fmt.Printf("GNBDUNameVisibleString: %s\n", decoded.ExtendedGNBDUName.GNBDUNameVisibleString.Value)
		} else {
			fmt.Printf("GNBDUNameVisibleString: <nil>\n")
		}
		if decoded.ExtendedGNBDUName.GNBDUNameUTF8String != nil {
			fmt.Printf("GNBDUNameUTF8String: %s\n", decoded.ExtendedGNBDUName.GNBDUNameUTF8String.Value)
		} else {
			fmt.Printf("GNBDUNameUTF8String: <nil>\n")
		}
	} else {
		fmt.Printf("\nExtendedGNBDUName: <nil>\n")
	}

	fmt.Printf("==============================================\n")
}
