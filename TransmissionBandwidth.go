package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TransmissionBandwidth struct {
	NRSCS NRSCS `mandatory`
	NRNRB NRNRB `mandatory`
	// IEExtensions * `optional`
}

func (ie *TransmissionBandwidth) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.NRSCS.Encode(w); err != nil {
		err = utils.WrapError("Encode NRSCS", err)
		return
	}
	if err = ie.NRNRB.Encode(w); err != nil {
		err = utils.WrapError("Encode NRNRB", err)
		return
	}
	return
}
func (ie *TransmissionBandwidth) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.NRSCS.Decode(r); err != nil {
		err = utils.WrapError("Read NRSCS", err)
		return
	}
	if err = ie.NRNRB.Decode(r); err != nil {
		err = utils.WrapError("Read NRNRB", err)
		return
	}
	return
}
