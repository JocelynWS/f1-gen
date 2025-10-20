package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type Paging struct {
	UEIdentityIndexValue UEIdentityIndexValue `mandatory,reject`
	PagingIdentity       PagingIdentity       `mandatory,reject`
	PagingDRX            *PagingDRX           `optional,ignore`
	PagingPriority       *PagingPriority      `optional,ignore`
	PagingCellList       []PagingCellItem     `lb:1,ub:maxnoofPagingCells,mandatory,ignore,valueExt`
	PagingOrigin         *PagingOrigin        `optional,ignore`
}

func (msg *Paging) Encode(w io.Writer) (err error) {
    var ies []F1apMessageIE
    if ies, err = msg.toIes(); err != nil {
        err = msgErrors(fmt.Errorf("Paging"), err)
        return
    }
    return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_Paging, Criticality_PresentIgnore, ies)
}
func (msg *Paging) toIes() (ies []F1apMessageIE, err error) {
	ies = []F1apMessageIE{}
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_UEIdentityIndexValue},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.UEIdentityIndexValue,
	})
	ies = append(ies, F1apMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_PagingIdentity},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.PagingIdentity,
	})
	if msg.PagingDRX != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PagingDRX},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.PagingDRX,
		})
	}
	if msg.PagingPriority != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PagingPriority},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.PagingPriority,
		})
	}
	if len(msg.PagingCellList) > 0 {
		tmp_PagingCellList := Sequence[*PagingCellItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPagingCells},
			ext: true,
		}
		for _, i := range msg.PagingCellList {
			tmp_PagingCellList.Value = append(tmp_PagingCellList.Value, &i)
		}
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PagingCellList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_PagingCellList,
		})
	} else {
		err = utils.WrapError("PagingCellList is nil", err)
		return
	}
	if msg.PagingOrigin != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PagingOrigin},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.PagingOrigin,
		})
	}
	return
}
func (msg *Paging) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("Paging"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := PagingDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*F1apMessageIE),
	}
	if _, err = aper.ReadSequenceOf[F1apMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	if _, ok := decoder.list[ProtocolIEID_UEIdentityIndexValue]; !ok {
		err = fmt.Errorf("Mandatory field UEIdentityIndexValue is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_UEIdentityIndexValue},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_PagingIdentity]; !ok {
		err = fmt.Errorf("Mandatory field PagingIdentity is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_PagingIdentity},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_PagingCellList]; !ok {
		err = fmt.Errorf("Mandatory field PagingCellList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_PagingCellList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type PagingDecoder struct {
	msg      *Paging
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *PagingDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_UEIdentityIndexValue:
		var tmp UEIdentityIndexValue
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read UEIdentityIndexValue", err)
			return
		}
		msg.UEIdentityIndexValue = tmp
	case ProtocolIEID_PagingIdentity:
		var tmp PagingIdentity
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read PagingIdentity", err)
			return
		}
		msg.PagingIdentity = tmp
	case ProtocolIEID_PagingDRX:
		var tmp PagingDRX
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read PagingDRX", err)
			return
		}
		msg.PagingDRX = &tmp
	case ProtocolIEID_PagingPriority:
		var tmp PagingPriority
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read PagingPriority", err)
			return
		}
		msg.PagingPriority = &tmp
	case ProtocolIEID_PagingCellList:
		tmp := Sequence[*PagingCellItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPagingCells},
			ext: true,
		}
		fn := func() *PagingCellItem { return new(PagingCellItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read PagingCellList", err)
			return
		}
		msg.PagingCellList = []PagingCellItem{}
		for _, i := range tmp.Value {
			msg.PagingCellList = append(msg.PagingCellList, *i)
		}
	case ProtocolIEID_PagingOrigin:
		var tmp PagingOrigin
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read PagingOrigin", err)
			return
		}
		msg.PagingOrigin = &tmp
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
