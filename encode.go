package f1ap

import (
	"bytes"
	"io"
)

func F1apEncode(msg F1apMessageEncoder) (wire []byte, err error) {
	var buf bytes.Buffer
	if err = msg.Encode(&buf); err == nil {
		wire = buf.Bytes()
	}
	return
}

type F1apMessageEncoder interface {
	Encode(io.Writer) error
}
