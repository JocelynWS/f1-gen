package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type EUTRACoexTDDInfo struct {
	EARFCN                int64                      `lb:0,ub:262143,mandatory`
	TransmissionBandwidth EUTRATransmissionBandwidth `mandatory`
	SubframeAssignment    EUTRASubframeAssignment    `mandatory`
	SpecialSubframeInfo   EUTRASpecialSubframeInfo   `mandatory`
	// IEExtensions * `optional`
}

func (ie *EUTRACoexTDDInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_EARFCN := NewINTEGER(ie.EARFCN, aper.Constraint{Lb: 0, Ub: 262143}, false)
	if err = tmp_EARFCN.Encode(w); err != nil {
		err = utils.WrapError("Encode EARFCN", err)
		return
	}
	if err = ie.TransmissionBandwidth.Encode(w); err != nil {
		err = utils.WrapError("Encode TransmissionBandwidth", err)
		return
	}
	if err = ie.SubframeAssignment.Encode(w); err != nil {
		err = utils.WrapError("Encode SubframeAssignment", err)
		return
	}
	if err = ie.SpecialSubframeInfo.Encode(w); err != nil {
		err = utils.WrapError("Encode SpecialSubframeInfo", err)
		return
	}
	return
}
func (ie *EUTRACoexTDDInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_EARFCN := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 262143},
		ext: false,
	}
	if err = tmp_EARFCN.Decode(r); err != nil {
		err = utils.WrapError("Read EARFCN", err)
		return
	}
	ie.EARFCN = int64(tmp_EARFCN.Value)
	if err = ie.TransmissionBandwidth.Decode(r); err != nil {
		err = utils.WrapError("Read TransmissionBandwidth", err)
		return
	}
	if err = ie.SubframeAssignment.Decode(r); err != nil {
		err = utils.WrapError("Read SubframeAssignment", err)
		return
	}
	if err = ie.SpecialSubframeInfo.Decode(r); err != nil {
		err = utils.WrapError("Read SpecialSubframeInfo", err)
		return
	}
	return
}
