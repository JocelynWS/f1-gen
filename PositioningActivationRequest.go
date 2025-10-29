package f1ap

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PositioningActivationRequest struct {
	GNBCUUEF1APID  int64           `lb:0,ub:4294967295,mandatory,reject`
	GNBDUUEF1APID  int64           `lb:0,ub:4294967295,mandatory,reject`
	SRSType        SRSType         `mandatory,reject`
	ActivationTime *aper.BitString `lb:64,ub:64,optional,ignore`
}

func (msg *PositioningActivationRequest) Encode(w io.Writer) (err error) {
	var ies []F1apMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("PositioningActivationRequest"), err)
		return
	}
	return encodeMessage(w, F1apPduInitiatingMessage, ProcedureCode_PositioningActivation, Criticality_PresentReject, ies)
}
func (msg *PositioningActivationRequest) toIes() (ies []F1apMessageIE, err error) {
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
		Id:          ProtocolIEID{Value: ProtocolIEID_SRSType},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.SRSType,
	})
	if msg.ActivationTime != nil {
		ies = append(ies, F1apMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ActivationTime},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &BITSTRING{
				c:   aper.Constraint{Lb: 64, Ub: 64},
				ext: false,
				Value: aper.BitString{
					Bytes: msg.ActivationTime.Bytes, NumBits: msg.ActivationTime.NumBits},
			}})
	}
	return
}
func (msg *PositioningActivationRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("PositioningActivationRequest"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := PositioningActivationRequestDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_SRSType]; !ok {
		err = fmt.Errorf("Mandatory field SRSType is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_SRSType},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type PositioningActivationRequestDecoder struct {
	msg      *PositioningActivationRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*F1apMessageIE
}

func (decoder *PositioningActivationRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *F1apMessageIE, err error) {
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
	case ProtocolIEID_SRSType:
		var tmp SRSType
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SRSType", err)
			return
		}
		msg.SRSType = tmp
	case ProtocolIEID_ActivationTime:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 64, Ub: 64},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ActivationTime", err)
			return
		}
		msg.ActivationTime = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
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
