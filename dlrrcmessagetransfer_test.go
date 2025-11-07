package f1ap

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/JocelynWS/f1-gen/ies"
	//"github.com/lvdund/ngap/aper"
)

func Test_DLRRCMessageTransfer_Decode(t *testing.T) {
	msg := ies.DLRRCMessageTransfer{
		GNBCUUEF1APID:        12,
		GNBDUUEF1APID:        12,
		SRBID:                1,
		RRCContainer:         []byte("RRC Container"),
		RedirectedRRCmessage: []byte("Redirected Message"),
		OldgNBDUUEF1APID:     new(int64),
	}
	*msg.OldgNBDUUEF1APID = 0

	msg.ExecuteDuplication = &ies.ExecuteDuplication{
		Value: ies.ExecuteDuplicationTrue,
	}

	fmt.Println("=== DLRRCMessageTransfer Full Message Encode/Decode ===")
	fmt.Println("\n--- Original Message ---")
	fmt.Printf("GNBCUUEF1APID: %d (0x%x)\n", msg.GNBCUUEF1APID, msg.GNBCUUEF1APID)
	fmt.Printf("GNBDUUEF1APID: %d (0x%x)\n", msg.GNBDUUEF1APID, msg.GNBDUUEF1APID)
	if msg.OldgNBDUUEF1APID != nil {
		fmt.Printf("OldgNBDUUEF1APID: %d (0x%x)\n", *msg.OldgNBDUUEF1APID, *msg.OldgNBDUUEF1APID)
	}
	fmt.Printf("SRBID: %d (0x%x)\n", msg.SRBID, msg.SRBID)
	if msg.ExecuteDuplication != nil {
		fmt.Printf("ExecuteDuplication: %d\n", msg.ExecuteDuplication.Value)
	}
	fmt.Printf("RRCContainer (hex): % x\n", msg.RRCContainer)
	fmt.Printf("RRCContainer (string): %s\n", string(msg.RRCContainer))
	fmt.Printf("RRCContainer length: %d bytes\n", len(msg.RRCContainer))
	fmt.Printf("RedirectedRRCmessage (hex): % x\n", msg.RedirectedRRCmessage)
	fmt.Printf("RedirectedRRCmessage (string): %s\n", string(msg.RedirectedRRCmessage))
	fmt.Printf("RedirectedRRCmessage length: %d bytes\n", len(msg.RedirectedRRCmessage))

	var buf bytes.Buffer
	err := msg.Encode(&buf)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	encoded := buf.Bytes()
	fmt.Printf("\n--- Encoded Bytes ---\n")
	fmt.Printf("Full Message: % x\n", encoded)
	fmt.Printf("Length: %d bytes\n", len(encoded))

	pdu, err, diagnostics := F1apDecode(encoded)
	if err != nil {
		t.Fatalf("F1apDecode error: %v, diagnostics: %v", err, diagnostics)
	}

	fmt.Printf("\n--- Decoded PDU ---\n")
	fmt.Printf("Present: %d\n", pdu.Present)
	fmt.Printf("ProcedureCode: %d\n", pdu.Message.ProcedureCode.Value)
	fmt.Printf("Criticality: %d\n", pdu.Message.Criticality.Value)

	decoded, ok := pdu.Message.Msg.(*ies.DLRRCMessageTransfer)
	if !ok {
		t.Fatalf("Failed to cast to DLRRCMessageTransfer")
	}

	fmt.Printf("\n--- Decoded Message ---\n")
	fmt.Printf("GNBCUUEF1APID: %d (0x%x)\n", decoded.GNBCUUEF1APID, decoded.GNBCUUEF1APID)
	fmt.Printf("GNBDUUEF1APID: %d (0x%x)\n", decoded.GNBDUUEF1APID, decoded.GNBDUUEF1APID)
	if decoded.OldgNBDUUEF1APID != nil {
		fmt.Printf("OldgNBDUUEF1APID: %d (0x%x)\n", *decoded.OldgNBDUUEF1APID, *decoded.OldgNBDUUEF1APID)
	}
	fmt.Printf("SRBID: %d (0x%x)\n", decoded.SRBID, decoded.SRBID)
	if decoded.ExecuteDuplication != nil {
		fmt.Printf("ExecuteDuplication: %d\n", decoded.ExecuteDuplication.Value)
	}
	fmt.Printf("RRCContainer (hex): % x\n", decoded.RRCContainer)
	fmt.Printf("RRCContainer (string): %s\n", string(decoded.RRCContainer))
	fmt.Printf("RRCContainer length: %d bytes\n", len(decoded.RRCContainer))
	fmt.Printf("RedirectedRRCmessage (hex): % x\n", decoded.RedirectedRRCmessage)
	fmt.Printf("RedirectedRRCmessage (string): %s\n", string(decoded.RedirectedRRCmessage))
	fmt.Printf("RedirectedRRCmessage length: %d bytes\n", len(decoded.RedirectedRRCmessage))

	if msg.GNBCUUEF1APID != decoded.GNBCUUEF1APID {
		t.Errorf("GNBCUUEF1APID mismatch: got %v, want %v", decoded.GNBCUUEF1APID, msg.GNBCUUEF1APID)
	}
	if msg.GNBDUUEF1APID != decoded.GNBDUUEF1APID {
		t.Errorf("GNBDUUEF1APID mismatch: got %v, want %v", decoded.GNBDUUEF1APID, msg.GNBDUUEF1APID)
	}
	if msg.OldgNBDUUEF1APID != nil && decoded.OldgNBDUUEF1APID != nil {
		if *msg.OldgNBDUUEF1APID != *decoded.OldgNBDUUEF1APID {
			t.Errorf("OldgNBDUUEF1APID mismatch: got %v, want %v", *decoded.OldgNBDUUEF1APID, *msg.OldgNBDUUEF1APID)
		}
	} else if msg.OldgNBDUUEF1APID != nil || decoded.OldgNBDUUEF1APID != nil {
		t.Errorf("OldgNBDUUEF1APID presence mismatch")
	}
	if msg.SRBID != decoded.SRBID {
		t.Errorf("SRBID mismatch: got %v, want %v", decoded.SRBID, msg.SRBID)
	}
	if msg.ExecuteDuplication != nil && decoded.ExecuteDuplication != nil {
		if msg.ExecuteDuplication.Value != decoded.ExecuteDuplication.Value {
			t.Errorf("ExecuteDuplication mismatch: got %v, want %v", decoded.ExecuteDuplication.Value, msg.ExecuteDuplication.Value)
		}
	} else if msg.ExecuteDuplication != nil || decoded.ExecuteDuplication != nil {
		t.Errorf("ExecuteDuplication presence mismatch")
	}
	if !bytes.Equal(msg.RRCContainer, decoded.RRCContainer) {
		t.Errorf("RRCContainer content mismatch:\ngot  (hex): % x\nwant (hex): % x\ngot  (string): %s\nwant (string): %s",
			decoded.RRCContainer, msg.RRCContainer,
			string(decoded.RRCContainer), string(msg.RRCContainer))
	}
	if !bytes.Equal(msg.RedirectedRRCmessage, decoded.RedirectedRRCmessage) {
		t.Errorf("RedirectedRRCmessage content mismatch:\ngot  (hex): % x\nwant (hex): % x\ngot  (string): %s\nwant (string): %s",
			decoded.RedirectedRRCmessage, msg.RedirectedRRCmessage,
			string(decoded.RedirectedRRCmessage), string(msg.RedirectedRRCmessage))
	}
}