package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type NGRANHighAccuracyAccessPointPosition struct {
	Latitude               int64 `lb:2147483648,ub:2147483647,mandatory`
	Longitude              int64 `lb:2147483648,ub:2147483647,mandatory`
	Altitude               int64 `lb:64000,ub:1280000,mandatory`
	UncertaintySemiMajor   int64 `lb:0,ub:255,mandatory`
	UncertaintySemiMinor   int64 `lb:0,ub:255,mandatory`
	OrientationOfMajorAxis int64 `lb:0,ub:179,mandatory`
	HorizontalConfidence   int64 `lb:0,ub:100,mandatory`
	UncertaintyAltitude    int64 `lb:0,ub:255,mandatory`
	VerticalConfidence     int64 `lb:0,ub:100,mandatory`
	// IEExtensions * `optional`
}

func (ie *NGRANHighAccuracyAccessPointPosition) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_Latitude := NewINTEGER(ie.Latitude, aper.Constraint{Lb: 2147483648, Ub: 2147483647}, false)
	if err = tmp_Latitude.Encode(w); err != nil {
		err = utils.WrapError("Encode Latitude", err)
		return
	}
	tmp_Longitude := NewINTEGER(ie.Longitude, aper.Constraint{Lb: 2147483648, Ub: 2147483647}, false)
	if err = tmp_Longitude.Encode(w); err != nil {
		err = utils.WrapError("Encode Longitude", err)
		return
	}
	tmp_Altitude := NewINTEGER(ie.Altitude, aper.Constraint{Lb: 64000, Ub: 1280000}, false)
	if err = tmp_Altitude.Encode(w); err != nil {
		err = utils.WrapError("Encode Altitude", err)
		return
	}
	tmp_UncertaintySemiMajor := NewINTEGER(ie.UncertaintySemiMajor, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_UncertaintySemiMajor.Encode(w); err != nil {
		err = utils.WrapError("Encode UncertaintySemiMajor", err)
		return
	}
	tmp_UncertaintySemiMinor := NewINTEGER(ie.UncertaintySemiMinor, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_UncertaintySemiMinor.Encode(w); err != nil {
		err = utils.WrapError("Encode UncertaintySemiMinor", err)
		return
	}
	tmp_OrientationOfMajorAxis := NewINTEGER(ie.OrientationOfMajorAxis, aper.Constraint{Lb: 0, Ub: 179}, false)
	if err = tmp_OrientationOfMajorAxis.Encode(w); err != nil {
		err = utils.WrapError("Encode OrientationOfMajorAxis", err)
		return
	}
	tmp_HorizontalConfidence := NewINTEGER(ie.HorizontalConfidence, aper.Constraint{Lb: 0, Ub: 100}, false)
	if err = tmp_HorizontalConfidence.Encode(w); err != nil {
		err = utils.WrapError("Encode HorizontalConfidence", err)
		return
	}
	tmp_UncertaintyAltitude := NewINTEGER(ie.UncertaintyAltitude, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_UncertaintyAltitude.Encode(w); err != nil {
		err = utils.WrapError("Encode UncertaintyAltitude", err)
		return
	}
	tmp_VerticalConfidence := NewINTEGER(ie.VerticalConfidence, aper.Constraint{Lb: 0, Ub: 100}, false)
	if err = tmp_VerticalConfidence.Encode(w); err != nil {
		err = utils.WrapError("Encode VerticalConfidence", err)
		return
	}
	return
}
func (ie *NGRANHighAccuracyAccessPointPosition) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_Latitude := INTEGER{
		c:   aper.Constraint{Lb: 2147483648, Ub: 2147483647},
		ext: false,
	}
	if err = tmp_Latitude.Decode(r); err != nil {
		err = utils.WrapError("Read Latitude", err)
		return
	}
	ie.Latitude = int64(tmp_Latitude.Value)
	tmp_Longitude := INTEGER{
		c:   aper.Constraint{Lb: 2147483648, Ub: 2147483647},
		ext: false,
	}
	if err = tmp_Longitude.Decode(r); err != nil {
		err = utils.WrapError("Read Longitude", err)
		return
	}
	ie.Longitude = int64(tmp_Longitude.Value)
	tmp_Altitude := INTEGER{
		c:   aper.Constraint{Lb: 64000, Ub: 1280000},
		ext: false,
	}
	if err = tmp_Altitude.Decode(r); err != nil {
		err = utils.WrapError("Read Altitude", err)
		return
	}
	ie.Altitude = int64(tmp_Altitude.Value)
	tmp_UncertaintySemiMajor := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 255},
		ext: false,
	}
	if err = tmp_UncertaintySemiMajor.Decode(r); err != nil {
		err = utils.WrapError("Read UncertaintySemiMajor", err)
		return
	}
	ie.UncertaintySemiMajor = int64(tmp_UncertaintySemiMajor.Value)
	tmp_UncertaintySemiMinor := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 255},
		ext: false,
	}
	if err = tmp_UncertaintySemiMinor.Decode(r); err != nil {
		err = utils.WrapError("Read UncertaintySemiMinor", err)
		return
	}
	ie.UncertaintySemiMinor = int64(tmp_UncertaintySemiMinor.Value)
	tmp_OrientationOfMajorAxis := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 179},
		ext: false,
	}
	if err = tmp_OrientationOfMajorAxis.Decode(r); err != nil {
		err = utils.WrapError("Read OrientationOfMajorAxis", err)
		return
	}
	ie.OrientationOfMajorAxis = int64(tmp_OrientationOfMajorAxis.Value)
	tmp_HorizontalConfidence := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 100},
		ext: false,
	}
	if err = tmp_HorizontalConfidence.Decode(r); err != nil {
		err = utils.WrapError("Read HorizontalConfidence", err)
		return
	}
	ie.HorizontalConfidence = int64(tmp_HorizontalConfidence.Value)
	tmp_UncertaintyAltitude := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 255},
		ext: false,
	}
	if err = tmp_UncertaintyAltitude.Decode(r); err != nil {
		err = utils.WrapError("Read UncertaintyAltitude", err)
		return
	}
	ie.UncertaintyAltitude = int64(tmp_UncertaintyAltitude.Value)
	tmp_VerticalConfidence := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 100},
		ext: false,
	}
	if err = tmp_VerticalConfidence.Decode(r); err != nil {
		err = utils.WrapError("Read VerticalConfidence", err)
		return
	}
	ie.VerticalConfidence = int64(tmp_VerticalConfidence.Value)
	return
}
