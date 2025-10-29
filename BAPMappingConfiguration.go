package f1ap

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type BAPMappingConfiguration struct {
	TransactionID                   int64                                 `lb:0,ub:255,mandatory,reject`
	BHRoutingInformationAddedList   []BHRoutingInformationAddedListItem   `lb:1,ub:maxnoofRoutingEntries,mandatory,ignore,valueExt`
	BHRoutingInformationRemovedList []BHRoutingInformationRemovedListItem `lb:1,ub:maxnoofRoutingEntries,mandatory,ignore,valueExt`
	TrafficMappingInformation       *TrafficMappingInfo                   `mandatory,ignore`
}

func (msg *BAPMappingConfiguration) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("BAPMappingConfiguration"), err)
		return
	}
	return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_BAPMappingConfiguration, Criticality_PresentReject, ies)
}
func (msg *BAPMappingConfiguration) toIes() (ies []F1apMessageIE, err error) {
	ies = []F1apMessageIE{}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TransactionID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 255},
			ext:   false,
			Value: aper.Integer(msg.TransactionID),
		}})
	if len(msg.BHRoutingInformationAddedList) > 0 {
		tmp_BHRoutingInformationAddedList := Sequence[*BHRoutingInformationAddedListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofRoutingEntries},
			ext: true,
		}
		for _, i := range msg.BHRoutingInformationAddedList {
			tmp_BHRoutingInformationAddedList.Value = append(tmp_BHRoutingInformationAddedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_BHRoutingInformationAddedList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_BHRoutingInformationAddedList,
		})
	} else {
		err = utils.WrapError("BHRoutingInformationAddedList is nil", err)
		return
	}
	if len(msg.BHRoutingInformationRemovedList) > 0 {
		tmp_BHRoutingInformationRemovedList := Sequence[*BHRoutingInformationRemovedListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofRoutingEntries},
			ext: true,
		}
		for _, i := range msg.BHRoutingInformationRemovedList {
			tmp_BHRoutingInformationRemovedList.Value = append(tmp_BHRoutingInformationRemovedList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_BHRoutingInformationRemovedList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_BHRoutingInformationRemovedList,
		})
	} else {
		err = utils.WrapError("BHRoutingInformationRemovedList is nil", err)
		return
	}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TrafficMappingInformation},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		// SỬA LỖI 1 (Dòng 74): Bỏ & vì TrafficMappingInformation đã là *TrafficMappingInfo
		Value: msg.TrafficMappingInformation,
	})
	return
}
func (msg *BAPMappingConfiguration) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("BAPMappingConfiguration"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := BAPMappingConfigurationDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*F1apMessageIE),
	}
	if _, err = aper.ReadSequenceOf[F1apMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	if _, ok := decoder.list[ProtocolIEID_TransactionID]; !ok {
		err = fmt.Errorf("Mandatory field TransactionID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_TransactionID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_BHRoutingInformationAddedList]; !ok {
		err = fmt.Errorf("Mandatory field BHRoutingInformationAddedList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_BHRoutingInformationAddedList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_BHRoutingInformationRemovedList]; !ok {
		err = fmt.Errorf("Mandatory field BHRoutingInformationRemovedList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_BHRoutingInformationRemovedList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_TrafficMappingInformation]; !ok {
		err = fmt.Errorf("Mandatory field TrafficMappingInformation is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_TrafficMappingInformation},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type BAPMappingConfigurationDecoder struct {
	msg      *BAPMappingConfiguration
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *BAPMappingConfigurationDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_TransactionID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 255},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read TransactionID", err)
			return
		}
		msg.TransactionID = int64(tmp.Value)
	case ProtocolIEID_BHRoutingInformationAddedList:
		tmp := Sequence[*BHRoutingInformationAddedListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofRoutingEntries},
			ext: true,
		}
		fn := func() *BHRoutingInformationAddedListItem { return new(BHRoutingInformationAddedListItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read BHRoutingInformationAddedList", err)
			return
		}
		msg.BHRoutingInformationAddedList = []BHRoutingInformationAddedListItem{}
		for _, i := range tmp.Value {
			msg.BHRoutingInformationAddedList = append(msg.BHRoutingInformationAddedList, *i)
		}
	case ProtocolIEID_BHRoutingInformationRemovedList:
		tmp := Sequence[*BHRoutingInformationRemovedListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofRoutingEntries},
			ext: true,
		}
		fn := func() *BHRoutingInformationRemovedListItem { return new(BHRoutingInformationRemovedListItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read BHRoutingInformationRemovedList", err)
			return
		}
		msg.BHRoutingInformationRemovedList = []BHRoutingInformationRemovedListItem{}
		for _, i := range tmp.Value {
			msg.BHRoutingInformationRemovedList = append(msg.BHRoutingInformationRemovedList, *i)
		}
	case ProtocolIEID_TrafficMappingInformation:
		var tmp TrafficMappingInfo
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read TrafficMappingInformation", err)
			return
		}
		msg.TrafficMappingInformation = &tmp
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
