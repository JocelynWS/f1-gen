package f1ap

import (
	"bytes"
	"fmt"

	"github.com/JocelynWS/f1-gen/ies"
	"github.com/lvdund/ngap/aper"
)

func F1apDecode(buf []byte) (pdu F1apPdu, err error, diagnostics *ies.CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewBuffer(buf))
	var b bool
	if b, err = r.ReadBool(); err != nil {
		return
	}
	_ = b
	c, err := r.ReadChoice(2, false)
	if err != nil {
		return
	}
	present := uint8(c)
	v, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false)
	if err != nil {
		return
	}
	var procedureCode = ies.ProcedureCode{Value: aper.Integer(v)}
	e, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return
	}
	var criticality = ies.Criticality{Value: aper.Enumerated(e)}
	var containerBytes []byte
	if containerBytes, err = r.ReadOpenType(); err != nil {
		return
	}
	message := createMessage(present, procedureCode)
	if message == nil {
		err = fmt.Errorf("Unknown F1AP message (present=%d, procedureCode=%d)", present, int64(procedureCode.Value))
		return
	}

	var diagnosticsItems []ies.CriticalityDiagnosticsIEItem
	if err, diagnosticsItems = message.Decode(containerBytes); err != nil {
		return
	}

	pdu = F1apPdu{
		Present: present,
		Message: F1apMessage{
			ProcedureCode: ies.ProcedureCode{Value: procedureCode.Value},
			Criticality:   ies.Criticality{Value: criticality.Value},
			Msg:           message,
		},
	}
	if len(diagnosticsItems) > 0 {
		diagnostics = ies.BuildDiagnostics(present, procedureCode, criticality, int64(procedureCode.Value), diagnosticsItems)
	}
	return
}

type MessageUnmarshaller interface {
	Decode(buf []byte) (error, []ies.CriticalityDiagnosticsIEItem)
}
