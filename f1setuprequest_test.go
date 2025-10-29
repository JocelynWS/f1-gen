package f1ap

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/lvdund/ngap/aper"
)

func TestF1SetupRequestMandatory(t *testing.T) {
	// ------------------------
	// TransactionID
	// ------------------------
	transactionID := int64(1)
	var bufTr bytes.Buffer
	w := aper.NewWriter(&bufTr)
	if err := w.WriteInteger(transactionID, &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		t.Fatalf("Encode TransactionID err: %v", err)
	}
	w.Close()

	r := aper.NewReader(&bufTr)
	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false)
	if err != nil {
		t.Fatalf("Decode TransactionID err: %v", err)
	}
	if val != transactionID {
		t.Errorf("TransactionID mismatch: got %v, want %v", val, transactionID)
	}
	fmt.Println("TransactionID Encode/Decode pass:", val)

	// ------------------------
	// GNBDUID
	// ------------------------
	gnbduID := int64(1)
	var bufGNB bytes.Buffer
	w = aper.NewWriter(&bufGNB)
	if err := w.WriteInteger(gnbduID, &aper.Constraint{Lb: 0, Ub: 68719476735}, false); err != nil {
		t.Fatalf("Encode GNBDUID err: %v", err)
	}
	w.Close()

	r = aper.NewReader(&bufGNB)
	val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 68719476735}, false)
	if err != nil {
		t.Fatalf("Decode GNBDUID err: %v", err)
	}
	if val != gnbduID {
		t.Errorf("GNBDUID mismatch: got %v, want %v", val, gnbduID)
	}
	fmt.Println("GNBDUID Encode/Decode pass:", val)

	// ------------------------
	// GNBDURRCVersion (BitString)
	// ------------------------
	rrcBytes := []byte{0x02, 0xF8, 0x39}
	rrcBits := uint64(19)
	rrc := aper.BitString{Bytes: rrcBytes, NumBits: rrcBits}

	var bufRRC bytes.Buffer
	w = aper.NewWriter(&bufRRC)
	if err := w.WriteBitString(rrc.Bytes, uint(rrc.NumBits), &aper.Constraint{Lb: 19, Ub: 19}, false); err != nil {
		t.Fatalf("Encode RRCVersion err: %v", err)
	}
	w.Close()

	r = aper.NewReader(&bufRRC)
	bytesOut, nbitsOut, err := r.ReadBitString(&aper.Constraint{Lb: 19, Ub: 19}, false)
	if err != nil {
		t.Fatalf("Decode RRCVersion err: %v", err)
	}
	if uint64(nbitsOut) != rrc.NumBits || !bytes.Equal(bytesOut, rrc.Bytes) {
		t.Errorf("RRCVersion mismatch: got %v %d bits, want %v %d bits",
			bytesOut, nbitsOut, rrc.Bytes, rrc.NumBits)
	}
	fmt.Println("GNBDURRCVersion Encode/Decode pass:", bytesOut, nbitsOut)

	fmt.Println("All mandatory F1SetupRequest IEs Encode/Decode passed!")
}

func Test_F1SetupRequest(t *testing.T) {
	msg := F1SetupRequest{
		TransactionID: 1,
		GNBDUID:       1,
		GNBDURRCVersion: RRCVersion{
			LatestRRCVersion: aper.BitString{
				Bytes:   []byte{0x02, 0xf8, 0x39},
				NumBits: 3,
			},
		},
	}
	var buf bytes.Buffer
	err := msg.Encode(&buf)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(buf.Bytes())
	}
}

// func Test_F1SetupRequest_Decode(t *testing.T) {
// 	// Encode
// 	msg := F1SetupRequest{
// 		TransactionID: 1,
// 		GNBDUID:       1,
// 		GNBDURRCVersion: RRCVersion{
// 			LatestRRCVersion: aper.BitString{
// 				Bytes:   []byte{0x02, 0xf8, 0x39},
// 				NumBits: 3,
// 			},
// 		},
// 	}
// 	var buf bytes.Buffer
// 	err := msg.Encode(&buf)
// 	if err != nil {
// 		fmt.Println("Encode error:", err)
// 		return
// 	}

// 	// Decode
// 	decoded := F1SetupRequest{}
// 	err, diagList := decoded.Decode(buf.Bytes())
// 	if err != nil {
// 		fmt.Println("Decode error:", err)
// 		fmt.Println("Diagnostics:", diagList)
// 	} else {
// 		fmt.Printf("TransactionID: %d\n", decoded.TransactionID)
// 		fmt.Printf("GNBDUID: %d\n", decoded.GNBDUID)
// 	}
// }

func Test_F1SetupRequest_Decode(t *testing.T) {
	msg := F1SetupRequest{
		TransactionID: 1,
		GNBDUID:       1,
		GNBDURRCVersion: RRCVersion{
			LatestRRCVersion: aper.BitString{
				Bytes:   []byte{0x02, 0xf8, 0x39},
				NumBits: 3,
			},
		},
	}

	var buf bytes.Buffer
	err := msg.Encode(&buf)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	fmt.Printf("Encoded: % x\n", buf.Bytes())

	pdu, err, diagnostics := F1apDecode(buf.Bytes())
	if err != nil {
		t.Fatalf("Decode error: %v, diagnostics: %v", err, diagnostics)
	}

	decoded := pdu.Message.Msg.(*F1SetupRequest)
	fmt.Printf("TransactionID: %d\n", decoded.TransactionID)
	fmt.Printf("GNBDUID: %d\n", decoded.GNBDUID)

	if decoded.TransactionID != 1 || decoded.GNBDUID != 1 {
		t.Errorf("Values mismatch!")
	}
}
