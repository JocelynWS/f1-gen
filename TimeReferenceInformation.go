package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TimeReferenceInformation struct {
	ReferenceTime       []byte              `lb:0,ub:0,mandatory`
	ReferenceSFN        int64               `lb:0,ub:1023,mandatory`
	Uncertainty         int64               `lb:0,ub:32767,mandatory`
	TimeInformationType TimeInformationType `mandatory`
	// IEExtensions * `optional`
}

func (ie *TimeReferenceInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_ReferenceTime := NewOCTETSTRING(ie.ReferenceTime, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_ReferenceTime.Encode(w); err != nil {
		err = utils.WrapError("Encode ReferenceTime", err)
		return
	}
	tmp_ReferenceSFN := NewINTEGER(ie.ReferenceSFN, aper.Constraint{Lb: 0, Ub: 1023}, false)
	if err = tmp_ReferenceSFN.Encode(w); err != nil {
		err = utils.WrapError("Encode ReferenceSFN", err)
		return
	}
	tmp_Uncertainty := NewINTEGER(ie.Uncertainty, aper.Constraint{Lb: 0, Ub: 32767}, false)
	if err = tmp_Uncertainty.Encode(w); err != nil {
		err = utils.WrapError("Encode Uncertainty", err)
		return
	}
	if err = ie.TimeInformationType.Encode(w); err != nil {
		err = utils.WrapError("Encode TimeInformationType", err)
		return
	}
	return
}
func (ie *TimeReferenceInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_ReferenceTime := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_ReferenceTime.Decode(r); err != nil {
		err = utils.WrapError("Read ReferenceTime", err)
		return
	}
	ie.ReferenceTime = tmp_ReferenceTime.Value
	tmp_ReferenceSFN := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 1023},
		ext: false,
	}
	if err = tmp_ReferenceSFN.Decode(r); err != nil {
		err = utils.WrapError("Read ReferenceSFN", err)
		return
	}
	ie.ReferenceSFN = int64(tmp_ReferenceSFN.Value)
	tmp_Uncertainty := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 32767},
		ext: false,
	}
	if err = tmp_Uncertainty.Decode(r); err != nil {
		err = utils.WrapError("Read Uncertainty", err)
		return
	}
	ie.Uncertainty = int64(tmp_Uncertainty.Value)
	if err = ie.TimeInformationType.Decode(r); err != nil {
		err = utils.WrapError("Read TimeInformationType", err)
		return
	}
	return
}
