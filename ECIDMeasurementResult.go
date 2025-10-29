package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ECIDMeasurementResult struct {
	GeographicalCoordinates *GeographicalCoordinates  `optional`
	MeasuredResultsList     []ECIDMeasuredResultsItem `lb:1,ub:maxnoofMeasECID,optional,valExt`
	// IEExtensions * `optional`
}

func (ie *ECIDMeasurementResult) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.GeographicalCoordinates != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.MeasuredResultsList != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if ie.GeographicalCoordinates != nil {
		if err = ie.GeographicalCoordinates.Encode(w); err != nil {
			err = utils.WrapError("Encode GeographicalCoordinates", err)
			return
		}
	}
	if len(ie.MeasuredResultsList) > 0 {
		tmp := Sequence[*ECIDMeasuredResultsItem]{
			Value: []*ECIDMeasuredResultsItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofMeasECID},
			ext:   true,
		}
		for _, i := range ie.MeasuredResultsList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode MeasuredResultsList", err)
			return
		}
	}
	return
}
func (ie *ECIDMeasurementResult) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(GeographicalCoordinates)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read GeographicalCoordinates", err)
			return
		}
		ie.GeographicalCoordinates = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_MeasuredResultsList := Sequence[*ECIDMeasuredResultsItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMeasECID},
			ext: true,
		}
		fn := func() *ECIDMeasuredResultsItem { return new(ECIDMeasuredResultsItem) }
		if err = tmp_MeasuredResultsList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read MeasuredResultsList", err)
			return
		}
		ie.MeasuredResultsList = []ECIDMeasuredResultsItem{}
		for _, i := range tmp_MeasuredResultsList.Value {
			ie.MeasuredResultsList = append(ie.MeasuredResultsList, *i)
		}
	}
	return
}
