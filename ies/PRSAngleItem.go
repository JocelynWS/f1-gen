package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PRSAngleItem struct {
	NRPRSAzimuth       int64 `lb:0,ub:359,mandatory`
	NRPRSAzimuthFine   int64 `lb:0,ub:9,mandatory`
	NRPRSElevation     int64 `lb:0,ub:180,mandatory`
	NRPRSElevationFine int64 `lb:0,ub:9,mandatory`
	// IEExtensions * `optional`
}

func (ie *PRSAngleItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_NRPRSAzimuth := NewINTEGER(ie.NRPRSAzimuth, aper.Constraint{Lb: 0, Ub: 359}, false)
	if err = tmp_NRPRSAzimuth.Encode(w); err != nil {
		err = utils.WrapError("Encode NRPRSAzimuth", err)
		return
	}
	tmp_NRPRSAzimuthFine := NewINTEGER(ie.NRPRSAzimuthFine, aper.Constraint{Lb: 0, Ub: 9}, false)
	if err = tmp_NRPRSAzimuthFine.Encode(w); err != nil {
		err = utils.WrapError("Encode NRPRSAzimuthFine", err)
		return
	}
	tmp_NRPRSElevation := NewINTEGER(ie.NRPRSElevation, aper.Constraint{Lb: 0, Ub: 180}, false)
	if err = tmp_NRPRSElevation.Encode(w); err != nil {
		err = utils.WrapError("Encode NRPRSElevation", err)
		return
	}
	tmp_NRPRSElevationFine := NewINTEGER(ie.NRPRSElevationFine, aper.Constraint{Lb: 0, Ub: 9}, false)
	if err = tmp_NRPRSElevationFine.Encode(w); err != nil {
		err = utils.WrapError("Encode NRPRSElevationFine", err)
		return
	}
	return
}
func (ie *PRSAngleItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_NRPRSAzimuth := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 359},
		ext: false,
	}
	if err = tmp_NRPRSAzimuth.Decode(r); err != nil {
		err = utils.WrapError("Read NRPRSAzimuth", err)
		return
	}
	ie.NRPRSAzimuth = int64(tmp_NRPRSAzimuth.Value)
	tmp_NRPRSAzimuthFine := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 9},
		ext: false,
	}
	if err = tmp_NRPRSAzimuthFine.Decode(r); err != nil {
		err = utils.WrapError("Read NRPRSAzimuthFine", err)
		return
	}
	ie.NRPRSAzimuthFine = int64(tmp_NRPRSAzimuthFine.Value)
	tmp_NRPRSElevation := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 180},
		ext: false,
	}
	if err = tmp_NRPRSElevation.Decode(r); err != nil {
		err = utils.WrapError("Read NRPRSElevation", err)
		return
	}
	ie.NRPRSElevation = int64(tmp_NRPRSElevation.Value)
	tmp_NRPRSElevationFine := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 9},
		ext: false,
	}
	if err = tmp_NRPRSElevationFine.Decode(r); err != nil {
		err = utils.WrapError("Read NRPRSElevationFine", err)
		return
	}
	ie.NRPRSElevationFine = int64(tmp_NRPRSElevationFine.Value)
	return
}
