package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UEContextModificationRequired struct {
	GNBCUUEF1APID                         int64                                `lb:0,ub:4294967295,mandatory,reject`
	GNBDUUEF1APID                         int64                                `lb:0,ub:4294967295,mandatory,reject`
	ResourceCoordinationTransferContainer []byte                               `lb:0,ub:0,optional,ignore`
	DUtoCURRCInformation                  *DUtoCURRCInformation                `mandatory,reject`
	DRBsRequiredToBeModifiedList          []DRBsRequiredToBeModifiedItem       `lb:1,ub:maxnoofDRBs,mandatory,reject`
	SRBsRequiredToBeReleasedList          []SRBsRequiredToBeReleasedItem       `lb:1,ub:maxnoofSRBs,mandatory,reject`
	DRBsRequiredToBeReleasedList          []DRBsRequiredToBeReleasedItem       `lb:1,ub:maxnoofDRBs,mandatory,reject`
	Cause                                 Cause                                `mandatory,ignore`
	BHChannelsRequiredToBeReleasedList    []BHChannelsRequiredToBeReleasedItem `lb:1,ub:maxnoofBHRLCChannels,mandatory,reject`
	SLDRBsRequiredToBeModifiedList        []SLDRBsRequiredToBeModifiedItem     `lb:1,ub:maxnoofSLDRBs,mandatory,reject`
	SLDRBsRequiredToBeReleasedList        []SLDRBsRequiredToBeReleasedItem     `lb:1,ub:maxnoofSLDRBs,mandatory,reject`
	TargetCellsToCancel                   []TargetCellListItem                 `lb:1,ub:maxnoofCHOcells,mandatory,reject`
}

func (msg *UEContextModificationRequired) Encode(w io.Writer) (err error) {
    var ies []F1apMessageIE
    if ies, err = msg.toIes(); err != nil {
        err = msgErrors(fmt.Errorf("UEContextModificationRequired"), err)
        return
    }
    return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_UEContextModificationRequired, Criticality_PresentReject, ies)
}
func (msg *UEContextModificationRequired) toIes() (ies []F1apMessageIE, err error) {
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
	if len(msg.DRBsRequiredToBeModifiedList) > 0 {
		tmp_DRBsRequiredToBeModifiedList := Sequence[*DRBsRequiredToBeModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		for _, i := range msg.DRBsRequiredToBeModifiedList {
			tmp_DRBsRequiredToBeModifiedList.Value = append(tmp_DRBsRequiredToBeModifiedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DRBsRequiredToBeModifiedList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_DRBsRequiredToBeModifiedList,
		})
	} else {
		err = utils.WrapError("DRBsRequiredToBeModifiedList is nil", err)
		return
	}
	if len(msg.SRBsRequiredToBeReleasedList) > 0 {
		tmp_SRBsRequiredToBeReleasedList := Sequence[*SRBsRequiredToBeReleasedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSRBs},
			ext: false,
		}
		for _, i := range msg.SRBsRequiredToBeReleasedList {
			tmp_SRBsRequiredToBeReleasedList.Value = append(tmp_SRBsRequiredToBeReleasedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SRBsRequiredToBeReleasedList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_SRBsRequiredToBeReleasedList,
		})
	} else {
		err = utils.WrapError("SRBsRequiredToBeReleasedList is nil", err)
		return
	}
	if len(msg.DRBsRequiredToBeReleasedList) > 0 {
		tmp_DRBsRequiredToBeReleasedList := Sequence[*DRBsRequiredToBeReleasedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		for _, i := range msg.DRBsRequiredToBeReleasedList {
			tmp_DRBsRequiredToBeReleasedList.Value = append(tmp_DRBsRequiredToBeReleasedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DRBsRequiredToBeReleasedList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_DRBsRequiredToBeReleasedList,
		})
	} else {
		err = utils.WrapError("DRBsRequiredToBeReleasedList is nil", err)
		return
	}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_Cause},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.Cause,
	})
	if len(msg.BHChannelsRequiredToBeReleasedList) > 0 {
		tmp_BHChannelsRequiredToBeReleasedList := Sequence[*BHChannelsRequiredToBeReleasedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofBHRLCChannels},
			ext: false,
		}
		for _, i := range msg.BHChannelsRequiredToBeReleasedList {
			tmp_BHChannelsRequiredToBeReleasedList.Value = append(tmp_BHChannelsRequiredToBeReleasedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_BHChannelsRequiredToBeReleasedList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_BHChannelsRequiredToBeReleasedList,
		})
	} else {
		err = utils.WrapError("BHChannelsRequiredToBeReleasedList is nil", err)
		return
	}
	if len(msg.SLDRBsRequiredToBeModifiedList) > 0 {
		tmp_SLDRBsRequiredToBeModifiedList := Sequence[*SLDRBsRequiredToBeModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		for _, i := range msg.SLDRBsRequiredToBeModifiedList {
			tmp_SLDRBsRequiredToBeModifiedList.Value = append(tmp_SLDRBsRequiredToBeModifiedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SLDRBsRequiredToBeModifiedList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_SLDRBsRequiredToBeModifiedList,
		})
	} else {
		err = utils.WrapError("SLDRBsRequiredToBeModifiedList is nil", err)
		return
	}
	if len(msg.SLDRBsRequiredToBeReleasedList) > 0 {
		tmp_SLDRBsRequiredToBeReleasedList := Sequence[*SLDRBsRequiredToBeReleasedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		for _, i := range msg.SLDRBsRequiredToBeReleasedList {
			tmp_SLDRBsRequiredToBeReleasedList.Value = append(tmp_SLDRBsRequiredToBeReleasedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SLDRBsRequiredToBeReleasedList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_SLDRBsRequiredToBeReleasedList,
		})
	} else {
		err = utils.WrapError("SLDRBsRequiredToBeReleasedList is nil", err)
		return
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
	} else {
		err = utils.WrapError("TargetCellsToCancel is nil", err)
		return
	}
	return
}
func (msg *UEContextModificationRequired) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("UEContextModificationRequired"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := UEContextModificationRequiredDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_DRBsRequiredToBeModifiedList]; !ok {
		err = fmt.Errorf("Mandatory field DRBsRequiredToBeModifiedList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_DRBsRequiredToBeModifiedList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_SRBsRequiredToBeReleasedList]; !ok {
		err = fmt.Errorf("Mandatory field SRBsRequiredToBeReleasedList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_SRBsRequiredToBeReleasedList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_DRBsRequiredToBeReleasedList]; !ok {
		err = fmt.Errorf("Mandatory field DRBsRequiredToBeReleasedList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_DRBsRequiredToBeReleasedList},
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
	if _, ok := decoder.list[ProtocolIEID_BHChannelsRequiredToBeReleasedList]; !ok {
		err = fmt.Errorf("Mandatory field BHChannelsRequiredToBeReleasedList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_BHChannelsRequiredToBeReleasedList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_SLDRBsRequiredToBeModifiedList]; !ok {
		err = fmt.Errorf("Mandatory field SLDRBsRequiredToBeModifiedList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_SLDRBsRequiredToBeModifiedList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_SLDRBsRequiredToBeReleasedList]; !ok {
		err = fmt.Errorf("Mandatory field SLDRBsRequiredToBeReleasedList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_SLDRBsRequiredToBeReleasedList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_TargetCellsToCancel]; !ok {
		err = fmt.Errorf("Mandatory field TargetCellsToCancel is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_TargetCellsToCancel},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type UEContextModificationRequiredDecoder struct {
	msg      *UEContextModificationRequired
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *UEContextModificationRequiredDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_DRBsRequiredToBeModifiedList:
		tmp := Sequence[*DRBsRequiredToBeModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		fn := func() *DRBsRequiredToBeModifiedItem { return new(DRBsRequiredToBeModifiedItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read DRBsRequiredToBeModifiedList", err)
			return
		}
		msg.DRBsRequiredToBeModifiedList = []DRBsRequiredToBeModifiedItem{}
		for _, i := range tmp.Value {
			msg.DRBsRequiredToBeModifiedList = append(msg.DRBsRequiredToBeModifiedList, *i)
		}
	case ProtocolIEID_SRBsRequiredToBeReleasedList:
		tmp := Sequence[*SRBsRequiredToBeReleasedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSRBs},
			ext: false,
		}
		fn := func() *SRBsRequiredToBeReleasedItem { return new(SRBsRequiredToBeReleasedItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SRBsRequiredToBeReleasedList", err)
			return
		}
		msg.SRBsRequiredToBeReleasedList = []SRBsRequiredToBeReleasedItem{}
		for _, i := range tmp.Value {
			msg.SRBsRequiredToBeReleasedList = append(msg.SRBsRequiredToBeReleasedList, *i)
		}
	case ProtocolIEID_DRBsRequiredToBeReleasedList:
		tmp := Sequence[*DRBsRequiredToBeReleasedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		fn := func() *DRBsRequiredToBeReleasedItem { return new(DRBsRequiredToBeReleasedItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read DRBsRequiredToBeReleasedList", err)
			return
		}
		msg.DRBsRequiredToBeReleasedList = []DRBsRequiredToBeReleasedItem{}
		for _, i := range tmp.Value {
			msg.DRBsRequiredToBeReleasedList = append(msg.DRBsRequiredToBeReleasedList, *i)
		}
	case ProtocolIEID_Cause:
		var tmp Cause
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read Cause", err)
			return
		}
		msg.Cause = tmp
	case ProtocolIEID_BHChannelsRequiredToBeReleasedList:
		tmp := Sequence[*BHChannelsRequiredToBeReleasedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofBHRLCChannels},
			ext: false,
		}
		fn := func() *BHChannelsRequiredToBeReleasedItem { return new(BHChannelsRequiredToBeReleasedItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read BHChannelsRequiredToBeReleasedList", err)
			return
		}
		msg.BHChannelsRequiredToBeReleasedList = []BHChannelsRequiredToBeReleasedItem{}
		for _, i := range tmp.Value {
			msg.BHChannelsRequiredToBeReleasedList = append(msg.BHChannelsRequiredToBeReleasedList, *i)
		}
	case ProtocolIEID_SLDRBsRequiredToBeModifiedList:
		tmp := Sequence[*SLDRBsRequiredToBeModifiedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		fn := func() *SLDRBsRequiredToBeModifiedItem { return new(SLDRBsRequiredToBeModifiedItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SLDRBsRequiredToBeModifiedList", err)
			return
		}
		msg.SLDRBsRequiredToBeModifiedList = []SLDRBsRequiredToBeModifiedItem{}
		for _, i := range tmp.Value {
			msg.SLDRBsRequiredToBeModifiedList = append(msg.SLDRBsRequiredToBeModifiedList, *i)
		}
	case ProtocolIEID_SLDRBsRequiredToBeReleasedList:
		tmp := Sequence[*SLDRBsRequiredToBeReleasedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSLDRBs},
			ext: false,
		}
		fn := func() *SLDRBsRequiredToBeReleasedItem { return new(SLDRBsRequiredToBeReleasedItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SLDRBsRequiredToBeReleasedList", err)
			return
		}
		msg.SLDRBsRequiredToBeReleasedList = []SLDRBsRequiredToBeReleasedItem{}
		for _, i := range tmp.Value {
			msg.SLDRBsRequiredToBeReleasedList = append(msg.SLDRBsRequiredToBeReleasedList, *i)
		}
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
