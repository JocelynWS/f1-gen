package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ServedEUTRACellsInformation struct {
	EUTRAModeInfo                    EUTRAModeInfo `mandatory`
	ProtectedEUTRAResourceIndication []byte        `lb:0,ub:0,mandatory`
	// IEExtensions * `optional`
}

func (ie *ServedEUTRACellsInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.EUTRAModeInfo.Encode(w); err != nil {
		err = utils.WrapError("Encode EUTRAModeInfo", err)
		return
	}
	tmp_ProtectedEUTRAResourceIndication := NewOCTETSTRING(ie.ProtectedEUTRAResourceIndication, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_ProtectedEUTRAResourceIndication.Encode(w); err != nil {
		err = utils.WrapError("Encode ProtectedEUTRAResourceIndication", err)
		return
	}
	return
}

func (ie *ServedEUTRACellsInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.EUTRAModeInfo.Decode(r); err != nil {
		err = utils.WrapError("Read EUTRAModeInfo", err)
		return
	}
	tmp_ProtectedEUTRAResourceIndication := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_ProtectedEUTRAResourceIndication.Decode(r); err != nil {
		err = utils.WrapError("Read ProtectedEUTRAResourceIndication", err)
		return
	}
	ie.ProtectedEUTRAResourceIndication = tmp_ProtectedEUTRAResourceIndication.Value
	return
}
