package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SULInformation struct {
	SULNRARFCN               int64                 `lb:0,ub:maxNRARFCN,mandatory`
	SULTransmissionBandwidth TransmissionBandwidth `mandatory`
	// IEExtensions * `optional`
}

func (ie *SULInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_SULNRARFCN := NewINTEGER(ie.SULNRARFCN, aper.Constraint{Lb: 0, Ub: maxNRARFCN}, false)
	if err = tmp_SULNRARFCN.Encode(w); err != nil {
		err = utils.WrapError("Encode SULNRARFCN", err)
		return
	}
	if err = ie.SULTransmissionBandwidth.Encode(w); err != nil {
		err = utils.WrapError("Encode SULTransmissionBandwidth", err)
		return
	}
	return
}
func (ie *SULInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_SULNRARFCN := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: maxNRARFCN},
		ext: false,
	}
	if err = tmp_SULNRARFCN.Decode(r); err != nil {
		err = utils.WrapError("Read SULNRARFCN", err)
		return
	}
	ie.SULNRARFCN = int64(tmp_SULNRARFCN.Value)
	if err = ie.SULTransmissionBandwidth.Decode(r); err != nil {
		err = utils.WrapError("Read SULTransmissionBandwidth", err)
		return
	}
	return
}
