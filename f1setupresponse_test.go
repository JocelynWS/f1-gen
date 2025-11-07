package f1ap

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/JocelynWS/f1-gen/ies"
	"github.com/lvdund/ngap/aper"
)

func TestF1SetupResponseMandatory(t *testing.T) {
	transactionID := int64(2)
	var bufTr bytes.Buffer
	w := aper.NewWriter(&bufTr)
	if err := w.WriteInteger(transactionID, &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		t.Fatalf("Encode TransactionID err: %v", err)
	}
	w.Close()

	fmt.Printf("TransactionID: %d (0x%x)\n", transactionID, transactionID)

	r := aper.NewReader(&bufTr)
	val, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false)
	if err != nil {
		t.Fatalf("Decode TransactionID err: %v", err)
	}
	if val != transactionID {
		t.Errorf("TransactionID mismatch: got %v, want %v", val, transactionID)
	}
	fmt.Printf("TransactionID Decoded: %d (0x%x)\n", val, val)
	fmt.Println("---")

	gnbcuName := []byte("OAI-CU")
	fmt.Printf("GNBCUName: %s (% x)\n", string(gnbcuName), gnbcuName)
	fmt.Println("---")

	rrcBytes := []byte{0x0c, 0x22, 0x38}
	rrcBits := uint64(3)
	rrc := aper.BitString{Bytes: rrcBytes, NumBits: rrcBits}

	var bufRRC bytes.Buffer
	w = aper.NewWriter(&bufRRC)
	if err := w.WriteBitString(rrc.Bytes, uint(rrc.NumBits), &aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		t.Fatalf("Encode RRCVersion err: %v", err)
	}
	w.Close()

	fmt.Printf("RRCVersion: % x (%d bits)\n", rrc.Bytes, rrcBits)

	r = aper.NewReader(&bufRRC)
	bytesOut, nbitsOut, err := r.ReadBitString(&aper.Constraint{Lb: 3, Ub: 3}, false)
	if err != nil {
		t.Fatalf("Decode RRCVersion err: %v", err)
	}
	if uint64(nbitsOut) != rrc.NumBits || !bytes.Equal(bytesOut, rrc.Bytes[:1]) {
		t.Errorf("RRCVersion mismatch: got %v %d bits, want %v %d bits",
			bytesOut, nbitsOut, rrc.Bytes[:1], rrc.NumBits)
	}
	fmt.Printf("RRCVersion Decoded: % x (%d bits)\n", bytesOut, nbitsOut)
	fmt.Println("---")

	nrCellID := int64(123456)
	var bufCell bytes.Buffer
	w = aper.NewWriter(&bufCell)
	if err := w.WriteInteger(nrCellID, &aper.Constraint{Lb: 0, Ub: 68719476735}, false); err != nil {
		t.Fatalf("Encode NRCellID err: %v", err)
	}
	w.Close()

	fmt.Printf("NRCellID: %d (0x%x)\n", nrCellID, nrCellID)

	r = aper.NewReader(&bufCell)
	val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 68719476735}, false)
	if err != nil {
		t.Fatalf("Decode NRCellID err: %v", err)
	}
	if val != nrCellID {
		t.Errorf("NRCellID mismatch: got %v, want %v", val, nrCellID)
	}
	fmt.Printf("NRCellID Decoded: %d (0x%x)\n", val, val)
	fmt.Println("---")

	nrpci := int64(1)
	var bufPCI bytes.Buffer
	w = aper.NewWriter(&bufPCI)
	if err := w.WriteInteger(nrpci, &aper.Constraint{Lb: 0, Ub: 1007}, false); err != nil {
		t.Fatalf("Encode NRPCI err: %v", err)
	}
	w.Close()

	fmt.Printf("NRPCI: %d (0x%x)\n", nrpci, nrpci)

	r = aper.NewReader(&bufPCI)
	val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1007}, false)
	if err != nil {
		t.Fatalf("Decode NRPCI err: %v", err)
	}
	if val != nrpci {
		t.Errorf("NRPCI mismatch: got %v, want %v", val, nrpci)
	}
	fmt.Printf("NRPCI Decoded: %d (0x%x)\n", val, val)
	fmt.Println("---")

	plmn := []byte{0x21, 0x23, 0xf1}
	fmt.Printf("PLMN: % x (MCC=12, MNC=123)\n", plmn)
	fmt.Println("---")

	siType := int64(7)
	fmt.Printf("SI Type: %d (0x%x)\n", siType, siType)
	fmt.Println("---")

	siContainer := []byte("test")
	fmt.Printf("SI Container: %s (% x)\n", string(siContainer), siContainer)
	fmt.Printf("SI Container length: %d bytes\n", len(siContainer))
	fmt.Println("---")
}

func Test_F1SetupResponse(t *testing.T) {
	msg := ies.F1SetupResponse{
		TransactionID: 2,
		GNBCUName:     []byte("OAI-CU"),
		GNBCURRCVersion: ies.RRCVersion{
			LatestRRCVersion: aper.BitString{
				Bytes:   []byte{0x0c, 0x22, 0x38},
				NumBits: 3,
			},
		},
	}

	fmt.Println("=== F1SetupResponse Encode (Mandatory Fields Only) ===")
	fmt.Printf("TransactionID: %d (0x%x)\n", msg.TransactionID, msg.TransactionID)
	fmt.Printf("GNBCUName: %s (% x)\n", string(msg.GNBCUName), msg.GNBCUName)
	fmt.Printf("RRCVersion: % x (%d bits)\n", msg.GNBCURRCVersion.LatestRRCVersion.Bytes, msg.GNBCURRCVersion.LatestRRCVersion.NumBits)

	var buf bytes.Buffer
	err := msg.Encode(&buf)
	if err != nil {
		fmt.Println("Encode error:", err)
	} else {
		fmt.Printf("\nFull Encoded Message: % x\n", buf.Bytes())
		fmt.Printf("Full Encoded Message (length %d bytes)\n", len(buf.Bytes()))
	}
}

func Test_F1SetupResponse_Decode(t *testing.T) {
	cell := ies.CellstobeActivatedListItem{
		NRCGI: ies.NRCGI{
			PLMNIdentity: []byte{0x21, 0x23, 0xf1},
			NRCellIdentity: aper.BitString{
				Bytes:   []byte{0x00, 0x1e, 0x24, 0x00, 0x00},
				NumBits: 36,
			},
		},
		NRPCI: &ies.NRPCI{Value: 1},
	}

	msg := ies.F1SetupResponse{
		TransactionID: 2,
		GNBCUName:     []byte("OAI-CU"),
		CellstobeActivatedList: []ies.CellstobeActivatedListItem{cell},
		GNBCURRCVersion: ies.RRCVersion{
			LatestRRCVersion: aper.BitString{
				Bytes:   []byte{0x0c, 0x22, 0x38},
				NumBits: 3,
			},
		},
	}

	fmt.Println("=== F1SetupResponse Full Message Encode/Decode ===")
	fmt.Println("\n--- Original Message ---")
	fmt.Printf("TransactionID: %d (0x%x)\n", msg.TransactionID, msg.TransactionID)
	fmt.Printf("GNBCUName: %s (% x)\n", string(msg.GNBCUName), msg.GNBCUName)
	fmt.Printf("RRCVersion: % x (%d bits)\n", msg.GNBCURRCVersion.LatestRRCVersion.Bytes, msg.GNBCURRCVersion.LatestRRCVersion.NumBits)
	fmt.Printf("Num Cells to Activate: %d\n", len(msg.CellstobeActivatedList))
	if len(msg.CellstobeActivatedList) > 0 {
		fmt.Printf("Cell[0] NRCGI PLMN: % x\n", msg.CellstobeActivatedList[0].NRCGI.PLMNIdentity)
		fmt.Printf("Cell[0] NRCellIdentity: % x (%d bits)\n",
			msg.CellstobeActivatedList[0].NRCGI.NRCellIdentity.Bytes,
			msg.CellstobeActivatedList[0].NRCGI.NRCellIdentity.NumBits)
		if msg.CellstobeActivatedList[0].NRPCI != nil {
			fmt.Printf("Cell[0] NRPCI: %d (0x%x)\n", msg.CellstobeActivatedList[0].NRPCI.Value, msg.CellstobeActivatedList[0].NRPCI.Value)
		}
	}

	var buf bytes.Buffer
	err := msg.Encode(&buf)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	fmt.Printf("\n--- Encoded Bytes ---\n")
	fmt.Printf("Full Message: % x\n", buf.Bytes())
	fmt.Printf("Length: %d bytes\n", len(buf.Bytes()))

	pdu, err, diagnostics := F1apDecode(buf.Bytes())
	if err != nil {
		t.Fatalf("Decode error: %v, diagnostics: %v", err, diagnostics)
	}

	decoded := pdu.Message.Msg.(*ies.F1SetupResponse)

	fmt.Printf("\n--- Decoded Message ---\n")
	fmt.Printf("TransactionID: %d (0x%x)\n", decoded.TransactionID, decoded.TransactionID)
	if decoded.GNBCUName != nil {
		fmt.Printf("GNBCUName: %s (% x)\n", string(decoded.GNBCUName), decoded.GNBCUName)
	}
	fmt.Printf("RRCVersion: % x (%d bits)\n", decoded.GNBCURRCVersion.LatestRRCVersion.Bytes, decoded.GNBCURRCVersion.LatestRRCVersion.NumBits)
	fmt.Printf("Num Cells to Activate: %d\n", len(decoded.CellstobeActivatedList))
	if len(decoded.CellstobeActivatedList) > 0 {
		fmt.Printf("Cell[0] NRCGI PLMN: % x\n", decoded.CellstobeActivatedList[0].NRCGI.PLMNIdentity)
		fmt.Printf("Cell[0] NRCellIdentity: % x (%d bits)\n",
			decoded.CellstobeActivatedList[0].NRCGI.NRCellIdentity.Bytes,
			decoded.CellstobeActivatedList[0].NRCGI.NRCellIdentity.NumBits)
		if decoded.CellstobeActivatedList[0].NRPCI != nil {
			fmt.Printf("Cell[0] NRPCI: %d (0x%x)\n", decoded.CellstobeActivatedList[0].NRPCI.Value, decoded.CellstobeActivatedList[0].NRPCI.Value)
		}
	}

	if msg.TransactionID != decoded.TransactionID {
		t.Errorf("TransactionID mismatch: got %v, want %v", decoded.TransactionID, msg.TransactionID)
	}
	if !bytes.Equal(msg.GNBCUName, decoded.GNBCUName) {
		t.Errorf("GNBCUName mismatch: got %s, want %s", string(decoded.GNBCUName), string(msg.GNBCUName))
	}
	if !bytes.Equal(msg.GNBCURRCVersion.LatestRRCVersion.Bytes[:1], decoded.GNBCURRCVersion.LatestRRCVersion.Bytes) {
		t.Errorf("RRCVersion bytes mismatch: got %x, want %x",
			decoded.GNBCURRCVersion.LatestRRCVersion.Bytes,
			msg.GNBCURRCVersion.LatestRRCVersion.Bytes[:1])
	}
	if msg.GNBCURRCVersion.LatestRRCVersion.NumBits != decoded.GNBCURRCVersion.LatestRRCVersion.NumBits {
		t.Errorf("RRCVersion NumBits mismatch: got %d, want %d",
			decoded.GNBCURRCVersion.LatestRRCVersion.NumBits,
			msg.GNBCURRCVersion.LatestRRCVersion.NumBits)
	}
	if len(msg.CellstobeActivatedList) != len(decoded.CellstobeActivatedList) {
		t.Errorf("CellstobeActivatedList length mismatch")
	}
}