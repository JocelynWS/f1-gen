package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AccessPointPosition struct {
	LatitudeSign           LatitudeSign        `madatory`
	Latitude               int64               `lb:0,ub:8388607,madatory`
	Longitude              int64               `lb:-8388608,ub:8388607,madatory`
	DirectionOfAltitude    DirectionOfAltitude `madatory`
	Altitude               int64               `lb:0,ub:32767,madatory`
	UncertaintySemiMajor   int64               `lb:0,ub:127,madatory`
	UncertaintySemiMinor   int64               `lb:0,ub:127,madatory`
	OrientationOfMajorAxis int64               `lb:0,ub:179,madatory`
	UncertaintyAltitude    int64               `lb:0,ub:127,madatory`
	Confidence             int64               `lb:0,ub:100,madatory`
	// IEExtensions *AccessPointPositionExtIEs `optional`
}

func (ie *AccessPointPosition) Encode(w *aper.AperWriter) (err error) {
	optionals := []byte{0x0}
	if err = w.WriteBits(optionals, 1); err != nil {
		return
	}
	if err = ie.LatitudeSign.Encode(w); err != nil {
		err = utils.WrapError("Encode LatitudeSign", err)
		return
	}
	tmp_Latitude := NewINTEGER(ie.Latitude, aper.Constraint{Lb: 0, Ub: 8388607}, false)
	if err = tmp_Latitude.Encode(w); err != nil {
		err = utils.WrapError("Encode Latitude", err)
		return
	}
	tmp_Longitude := NewINTEGER(ie.Longitude, aper.Constraint{Lb: -8388608, Ub: 8388607}, false)
	if err = tmp_Longitude.Encode(w); err != nil {
		err = utils.WrapError("Encode Longitude", err)
		return
	}
	if err = ie.DirectionOfAltitude.Encode(w); err != nil {
		err = utils.WrapError("Encode DirectionOfAltitude", err)
		return
	}
	tmp_Altitude := NewINTEGER(ie.Altitude, aper.Constraint{Lb: 0, Ub: 32767}, false)
	if err = tmp_Altitude.Encode(w); err != nil {
		err = utils.WrapError("Encode Altitude", err)
		return
	}
	tmp_UncertaintySemiMajor := NewINTEGER(ie.UncertaintySemiMajor, aper.Constraint{Lb: 0, Ub: 127}, false)
	if err = tmp_UncertaintySemiMajor.Encode(w); err != nil {
		err = utils.WrapError("Encode UncertaintySemiMajor", err)
		return
	}
	tmp_UncertaintySemiMinor := NewINTEGER(ie.UncertaintySemiMinor, aper.Constraint{Lb: 0, Ub: 127}, false)
	if err = tmp_UncertaintySemiMinor.Encode(w); err != nil {
		err = utils.WrapError("Encode UncertaintySemiMinor", err)
		return
	}
	tmp_OrientationOfMajorAxis := NewINTEGER(ie.OrientationOfMajorAxis, aper.Constraint{Lb: 0, Ub: 179}, false)
	if err = tmp_OrientationOfMajorAxis.Encode(w); err != nil {
		err = utils.WrapError("Encode OrientationOfMajorAxis", err)
		return
	}
	tmp_UncertaintyAltitude := NewINTEGER(ie.UncertaintyAltitude, aper.Constraint{Lb: 0, Ub: 127}, false)
	if err = tmp_UncertaintyAltitude.Encode(w); err != nil {
		err = utils.WrapError("Encode UncertaintyAltitude", err)
		return
	}
	tmp_Confidence := NewINTEGER(ie.Confidence, aper.Constraint{Lb: 0, Ub: 100}, false)
	if err = tmp_Confidence.Encode(w); err != nil {
		err = utils.WrapError("Encode Confidence", err)
		return
	}
	return
}

func (ie *AccessPointPosition) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.LatitudeSign.Decode(r); err != nil {
		err = utils.WrapError("Read LatitudeSign", err)
		return
	}
	tmp_Latitude := INTEGER{c: aper.Constraint{Lb: 0, Ub: 8388607}}
	if err = tmp_Latitude.Decode(r); err != nil {
		err = utils.WrapError("Read Latitude", err)
		return
	}
	ie.Latitude = int64(tmp_Latitude.Value)
	tmp_Longitude := INTEGER{c: aper.Constraint{Lb: -8388608, Ub: 8388607}}
	if err = tmp_Longitude.Decode(r); err != nil {
		err = utils.WrapError("Read Longitude", err)
		return
	}
	ie.Longitude = int64(tmp_Longitude.Value)
	if err = ie.DirectionOfAltitude.Decode(r); err != nil {
		err = utils.WrapError("Read DirectionOfAltitude", err)
		return
	}
	tmp_Altitude := INTEGER{c: aper.Constraint{Lb: 0, Ub: 32767}}
	if err = tmp_Altitude.Decode(r); err != nil {
		err = utils.WrapError("Read Altitude", err)
		return
	}
	ie.Altitude = int64(tmp_Altitude.Value)
	tmp_UncertaintySemiMajor := INTEGER{c: aper.Constraint{Lb: 0, Ub: 127}}
	if err = tmp_UncertaintySemiMajor.Decode(r); err != nil {
		err = utils.WrapError("Read UncertaintySemiMajor", err)
		return
	}
	ie.UncertaintySemiMajor = int64(tmp_UncertaintySemiMajor.Value)
	tmp_UncertaintySemiMinor := INTEGER{c: aper.Constraint{Lb: 0, Ub: 127}}
	if err = tmp_UncertaintySemiMinor.Decode(r); err != nil {
		err = utils.WrapError("Read UncertaintySemiMinor", err)
		return
	}
	ie.UncertaintySemiMinor = int64(tmp_UncertaintySemiMinor.Value)
	tmp_OrientationOfMajorAxis := INTEGER{c: aper.Constraint{Lb: 0, Ub: 179}}
	if err = tmp_OrientationOfMajorAxis.Decode(r); err != nil {
		err = utils.WrapError("Read OrientationOfMajorAxis", err)
		return
	}
	ie.OrientationOfMajorAxis = int64(tmp_OrientationOfMajorAxis.Value)
	tmp_UncertaintyAltitude := INTEGER{c: aper.Constraint{Lb: 0, Ub: 127}}
	if err = tmp_UncertaintyAltitude.Decode(r); err != nil {
		err = utils.WrapError("Read UncertaintyAltitude", err)
		return
	}
	ie.UncertaintyAltitude = int64(tmp_UncertaintyAltitude.Value)
	tmp_Confidence := INTEGER{c: aper.Constraint{Lb: 0, Ub: 100}}
	if err = tmp_Confidence.Decode(r); err != nil {
		err = utils.WrapError("Read Confidence", err)
		return
	}
	ie.Confidence = int64(tmp_Confidence.Value)
	return
}
