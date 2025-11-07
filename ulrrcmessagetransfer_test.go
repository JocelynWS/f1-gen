package f1ap

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/JocelynWS/f1-gen/ies"
	"github.com/lvdund/ngap/aper"
)

func TestDLRRCMessageTransferMandatory(t *testing.T) {
	gnbcuueF1APID := int64(12)
	var bufCU bytes.Buffer
	w := aper.NewWriter(&bufCU)
	if err := w.WriteInteger(gnbcuueF1APID, &aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		t.Fatalf("Encode GNBCUUEF1APID err: %v", err)
	}
	w.Close()

	fmt.Printf("GNBCUUEF1APID: %d (0x%x)\n", gnbcuueF1APID, gnbcuueF1APID)

	r := aper.NewReader(&bufCU)
	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false)
	if err != nil {
		t.Fatalf("Decode GNBCUUEF1APID err: %v", err)
	}
	if val != gnbcuueF1APID {
		t.Errorf("GNBCUUEF1APID mismatch: got %v, want %v", val, gnbcuueF1APID)
	}
	fmt.Printf("GNBCUUEF1APID Decoded: %d (0x%x)\n", val, val)
	fmt.Println("---")

	gnbduueF1APID := int64(12)
	var bufDU bytes.Buffer
	w = aper.NewWriter(&bufDU)
	if err := w.WriteInteger(gnbduueF1APID, &aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		t.Fatalf("Encode GNBDUUEF1APID err: %v", err)
	}
	w.Close()

	fmt.Printf("GNBDUUEF1APID: %d (0x%x)\n", gnbduueF1APID, gnbduueF1APID)

	r = aper.NewReader(&bufDU)
	val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false)
	if err != nil {
		t.Fatalf("Decode GNBDUUEF1APID err: %v", err)
	}
	if val != gnbduueF1APID {
		t.Errorf("GNBDUUEF1APID mismatch: got %v, want %v", val, gnbduueF1APID)
	}
	fmt.Printf("GNBDUUEF1APID Decoded: %d (0x%x)\n", val, val)
	fmt.Println("---")

	srbID := int64(1)
	var bufSRB bytes.Buffer
	w = aper.NewWriter(&bufSRB)
	if err := w.WriteInteger(srbID, &aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		t.Fatalf("Encode SRBID err: %v", err)
	}
	w.Close()

	fmt.Printf("SRBID: %d (0x%x)\n", srbID, srbID)

	r = aper.NewReader(&bufSRB)
	val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 3}, false)
	if err != nil {
		t.Fatalf("Decode SRBID err: %v", err)
	}
	if val != srbID {
		t.Errorf("SRBID mismatch: got %v, want %v", val, srbID)
	}
	fmt.Printf("SRBID Decoded: %d (0x%x)\n", val, val)
	fmt.Println("---")

	execDup := ies.ExecuteDuplication{Value: ies.ExecuteDuplicationTrue}
	var bufExec bytes.Buffer
	w = aper.NewWriter(&bufExec)
	if err := execDup.Encode(w); err != nil {
		t.Fatalf("Encode ExecuteDuplication err: %v", err)
	}
	w.Close()

	fmt.Printf("ExecuteDuplication: %d\n", execDup.Value)

	r = aper.NewReader(&bufExec)
	var decodedExec ies.ExecuteDuplication
	if err := decodedExec.Decode(r); err != nil {
		t.Fatalf("Decode ExecuteDuplication err: %v", err)
	}
	if decodedExec.Value != execDup.Value {
		t.Errorf("ExecuteDuplication mismatch: got %v, want %v", decodedExec.Value, execDup.Value)
	}
	fmt.Printf("ExecuteDuplication Decoded: %d\n", decodedExec.Value)
	fmt.Println("---")

	rrcContainer := []byte("RRC Container")
	fmt.Printf("RRCContainer: % x\n", rrcContainer)
	fmt.Printf("RRCContainer length: %d bytes\n", len(rrcContainer))
	fmt.Println("---")

	redirectedRRC := []byte("Redirected Message")
	fmt.Printf("RedirectedRRCmessage: % x\n", redirectedRRC)
	fmt.Printf("RedirectedRRCmessage length: %d bytes\n", len(redirectedRRC))
	fmt.Println("---")
}

func TestULRRCMessageTransferMandatory(t *testing.T) {
	gnbcuueF1APID := int64(12)
	var bufCU bytes.Buffer
	w := aper.NewWriter(&bufCU)
	if err := w.WriteInteger(gnbcuueF1APID, &aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		t.Fatalf("Encode GNBCUUEF1APID err: %v", err)
	}
	w.Close()

	fmt.Printf("GNBCUUEF1APID: %d (0x%x)\n", gnbcuueF1APID, gnbcuueF1APID)

	r := aper.NewReader(&bufCU)
	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false)
	if err != nil {
		t.Fatalf("Decode GNBCUUEF1APID err: %v", err)
	}
	if val != gnbcuueF1APID {
		t.Errorf("GNBCUUEF1APID mismatch: got %v, want %v", val, gnbcuueF1APID)
	}
	fmt.Printf("GNBCUUEF1APID Decoded: %d (0x%x)\n", val, val)
	fmt.Println("---")

	gnbduueF1APID := int64(12)
	var bufDU bytes.Buffer
	w = aper.NewWriter(&bufDU)
	if err := w.WriteInteger(gnbduueF1APID, &aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		t.Fatalf("Encode GNBDUUEF1APID err: %v", err)
	}
	w.Close()

	fmt.Printf("GNBDUUEF1APID: %d (0x%x)\n", gnbduueF1APID, gnbduueF1APID)

	r = aper.NewReader(&bufDU)
	val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false)
	if err != nil {
		t.Fatalf("Decode GNBDUUEF1APID err: %v", err)
	}
	if val != gnbduueF1APID {
		t.Errorf("GNBDUUEF1APID mismatch: got %v, want %v", val, gnbduueF1APID)
	}
	fmt.Printf("GNBDUUEF1APID Decoded: %d (0x%x)\n", val, val)
	fmt.Println("---")

	srbID := int64(2)
	var bufSRB bytes.Buffer
	w = aper.NewWriter(&bufSRB)
	if err := w.WriteInteger(srbID, &aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		t.Fatalf("Encode SRBID err: %v", err)
	}
	w.Close()

	fmt.Printf("SRBID: %d (0x%x)\n", srbID, srbID)

	r = aper.NewReader(&bufSRB)
	val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 3}, false)
	if err != nil {
		t.Fatalf("Decode SRBID err: %v", err)
	}
	if val != srbID {
		t.Errorf("SRBID mismatch: got %v, want %v", val, srbID)
	}
	fmt.Printf("SRBID Decoded: %d (0x%x)\n", val, val)
	fmt.Println("---")

	rrcContainer := []byte("RRC Container")
	fmt.Printf("RRCContainer: % x\n", rrcContainer)
	fmt.Printf("RRCContainer length: %d bytes\n", len(rrcContainer))
	fmt.Println("---")
}

func Test_ULRRCMessageTransfer_Decode(t *testing.T) {
	msg := ies.ULRRCMessageTransfer{
		GNBCUUEF1APID:    12,
		GNBDUUEF1APID:    12,
		SRBID:            1,
		RRCContainer:     []byte("RRC Container"),
		SelectedPLMNID:   []byte{0x22, 0xf2, 0x10},
		NewgNBDUUEF1APID: new(int64),
	}
	*msg.NewgNBDUUEF1APID = 98765

	fmt.Println("=== ULRRCMessageTransfer Full Message Encode/Decode ===")
	fmt.Println("\n--- Original Message ---")
	fmt.Printf("GNBCUUEF1APID: %d (0x%x)\n", msg.GNBCUUEF1APID, msg.GNBCUUEF1APID)
	fmt.Printf("GNBDUUEF1APID: %d (0x%x)\n", msg.GNBDUUEF1APID, msg.GNBDUUEF1APID)
	fmt.Printf("SRBID: %d (0x%x)\n", msg.SRBID, msg.SRBID)
	fmt.Printf("RRCContainer (hex): % x\n", msg.RRCContainer)
	fmt.Printf("RRCContainer (string): %s\n", string(msg.RRCContainer))
	fmt.Printf("RRCContainer length: %d bytes\n", len(msg.RRCContainer))
	if msg.SelectedPLMNID != nil {
		fmt.Printf("SelectedPLMNID: % x\n", msg.SelectedPLMNID)
	}
	if msg.NewgNBDUUEF1APID != nil {
		fmt.Printf("NewgNBDUUEF1APID: %d (0x%x)\n", *msg.NewgNBDUUEF1APID, *msg.NewgNBDUUEF1APID)
	}

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

	decoded, ok := pdu.Message.Msg.(*ies.ULRRCMessageTransfer)
	if !ok {
		t.Fatalf("Failed to cast to ULRRCMessageTransfer")
	}

	fmt.Printf("\n--- Decoded Message ---\n")
	fmt.Printf("GNBCUUEF1APID: %d (0x%x)\n", decoded.GNBCUUEF1APID, decoded.GNBCUUEF1APID)
	fmt.Printf("GNBDUUEF1APID: %d (0x%x)\n", decoded.GNBDUUEF1APID, decoded.GNBDUUEF1APID)
	fmt.Printf("SRBID: %d (0x%x)\n", decoded.SRBID, decoded.SRBID)
	fmt.Printf("RRCContainer (hex): % x\n", decoded.RRCContainer)
	fmt.Printf("RRCContainer (string): %s\n", string(decoded.RRCContainer))
	fmt.Printf("RRCContainer length: %d bytes\n", len(decoded.RRCContainer))
	if decoded.SelectedPLMNID != nil {
		fmt.Printf("SelectedPLMNID: % x\n", decoded.SelectedPLMNID)
	}
	if decoded.NewgNBDUUEF1APID != nil {
		fmt.Printf("NewgNBDUUEF1APID: %d (0x%x)\n", *decoded.NewgNBDUUEF1APID, *decoded.NewgNBDUUEF1APID)
	}

	if msg.GNBCUUEF1APID != decoded.GNBCUUEF1APID {
		t.Errorf("GNBCUUEF1APID mismatch: got %v, want %v", decoded.GNBCUUEF1APID, msg.GNBCUUEF1APID)
	}
	if msg.GNBDUUEF1APID != decoded.GNBDUUEF1APID {
		t.Errorf("GNBDUUEF1APID mismatch: got %v, want %v", decoded.GNBDUUEF1APID, msg.GNBDUUEF1APID)
	}
	if msg.SRBID != decoded.SRBID {
		t.Errorf("SRBID mismatch: got %v, want %v", decoded.SRBID, msg.SRBID)
	}
	if !bytes.Equal(msg.RRCContainer, decoded.RRCContainer) {
		t.Errorf("RRCContainer content mismatch:\ngot  (hex): % x\nwant (hex): % x\ngot  (string): %s\nwant (string): %s",
			decoded.RRCContainer, msg.RRCContainer,
			string(decoded.RRCContainer), string(msg.RRCContainer))
	}
	if !bytes.Equal(msg.SelectedPLMNID, decoded.SelectedPLMNID) {
		t.Errorf("SelectedPLMNID mismatch: got %x, want %x", decoded.SelectedPLMNID, msg.SelectedPLMNID)
	}
	if msg.NewgNBDUUEF1APID != nil && decoded.NewgNBDUUEF1APID != nil {
		if *msg.NewgNBDUUEF1APID != *decoded.NewgNBDUUEF1APID {
			t.Errorf("NewgNBDUUEF1APID mismatch: got %v, want %v", *decoded.NewgNBDUUEF1APID, *msg.NewgNBDUUEF1APID)
		}
	} else if msg.NewgNBDUUEF1APID != nil || decoded.NewgNBDUUEF1APID != nil {
		t.Errorf("NewgNBDUUEF1APID presence mismatch")
	}
}