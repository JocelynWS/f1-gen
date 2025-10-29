package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type RelativeGeodeticLocation struct {
	MilliArcSecondUnits MilliArcSecondUnits `madatory,valExt`
	HeightUnits         HeightUnits         `madatory,valExt`
	DeltaLatitude       int64               `lb:-1024,ub:1023,madatory`
	DeltaLongitude      int64               `lb:-1024,ub:1023,madatory`
	DeltaHeight         int64               `lb:-1024,ub:1023,madatory`
	LocationUncertainty LocationUncertainty `madatory`
	// IEExtensions *RelativeGeodeticLocationExtIEs `optional`
}

func (ie *RelativeGeodeticLocation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if err = w.WriteBits(optionals, 1); err != nil {
		return
	}
	if err = ie.MilliArcSecondUnits.Encode(w); err != nil {
		err = utils.WrapError("Encode MilliArcSecondUnits", err)
		return
	}
	if err = ie.HeightUnits.Encode(w); err != nil {
		err = utils.WrapError("Encode HeightUnits", err)
		return
	}
	tmp_DeltaLatitude := NewINTEGER(ie.DeltaLatitude, aper.Constraint{Lb: -1024, Ub: 1023}, false)
	if err = tmp_DeltaLatitude.Encode(w); err != nil {
		err = utils.WrapError("Encode DeltaLatitude", err)
		return
	}
	tmp_DeltaLongitude := NewINTEGER(ie.DeltaLongitude, aper.Constraint{Lb: -1024, Ub: 1023}, false)
	if err = tmp_DeltaLongitude.Encode(w); err != nil {
		err = utils.WrapError("Encode DeltaLongitude", err)
		return
	}
	tmp_DeltaHeight := NewINTEGER(ie.DeltaHeight, aper.Constraint{Lb: -1024, Ub: 1023}, false)
	if err = tmp_DeltaHeight.Encode(w); err != nil {
		err = utils.WrapError("Encode DeltaHeight", err)
		return
	}
	if err = ie.LocationUncertainty.Encode(w); err != nil {
		err = utils.WrapError("Encode LocationUncertainty", err)
		return
	}
	return
}

func (ie *RelativeGeodeticLocation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.MilliArcSecondUnits.Decode(r); err != nil {
		err = utils.WrapError("Read MilliArcSecondUnits", err)
		return
	}
	if err = ie.HeightUnits.Decode(r); err != nil {
		err = utils.WrapError("Read HeightUnits", err)
		return
	}
	tmp_DeltaLatitude := INTEGER{c: aper.Constraint{Lb: -1024, Ub: 1023}}
	if err = tmp_DeltaLatitude.Decode(r); err != nil {
		err = utils.WrapError("Read DeltaLatitude", err)
		return
	}
	ie.DeltaLatitude = int64(tmp_DeltaLatitude.Value)
	tmp_DeltaLongitude := INTEGER{c: aper.Constraint{Lb: -1024, Ub: 1023}}
	if err = tmp_DeltaLongitude.Decode(r); err != nil {
		err = utils.WrapError("Read DeltaLongitude", err)
		return
	}
	ie.DeltaLongitude = int64(tmp_DeltaLongitude.Value)
	tmp_DeltaHeight := INTEGER{c: aper.Constraint{Lb: -1024, Ub: 1023}}
	if err = tmp_DeltaHeight.Decode(r); err != nil {
		err = utils.WrapError("Read DeltaHeight", err)
		return
	}
	ie.DeltaHeight = int64(tmp_DeltaHeight.Value)
	if err = ie.LocationUncertainty.Decode(r); err != nil {
		err = utils.WrapError("Read LocationUncertainty", err)
		return
	}
	return
}
