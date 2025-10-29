package f1ap

import (
	"io"

	"github.com/JocelynWS/f1-gen/ies"
)

type F1apPdu struct {
	Present uint8
	Message F1apMessage
}

type F1apMessage struct {
	ProcedureCode ies.ProcedureCode
	Criticality   ies.Criticality
	Msg           MessageUnmarshaller
}

// type MessageUnmarshaller interface {
// 	Decode([]byte) (error, []ies.CriticalityDiagnosticsIEItem)
// }

type F1apMessageEncoder interface {
	Encode(io.Writer) error
}
