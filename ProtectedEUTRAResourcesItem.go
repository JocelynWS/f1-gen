package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ProtectedEUTRAResourcesItem struct {
	SpectrumSharingGroupID int64                `lb:0,ub:maxCellineNB,mandatory`
	EUTRACellsList         []EUTRACellsListItem `lb:1,ub:maxCellineNB,mandatory,valExt`
	// IEExtensions * `optional`
}

func (ie *ProtectedEUTRAResourcesItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_SpectrumSharingGroupID := NewINTEGER(ie.SpectrumSharingGroupID, aper.Constraint{Lb: 0, Ub: maxCellineNB}, false)
	if err = tmp_SpectrumSharingGroupID.Encode(w); err != nil {
		err = utils.WrapError("Encode SpectrumSharingGroupID", err)
		return
	}
	if len(ie.EUTRACellsList) > 0 {
		tmp := Sequence[*EUTRACellsListItem]{
			Value: []*EUTRACellsListItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxCellineNB},
			ext:   true,
		}
		for _, i := range ie.EUTRACellsList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode EUTRACellsList", err)
			return
		}
	} else {
		err = utils.WrapError("EUTRACellsList is nil", err)
		return
	}
	return
}
func (ie *ProtectedEUTRAResourcesItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_SpectrumSharingGroupID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: maxCellineNB},
		ext: false,
	}
	if err = tmp_SpectrumSharingGroupID.Decode(r); err != nil {
		err = utils.WrapError("Read SpectrumSharingGroupID", err)
		return
	}
	ie.SpectrumSharingGroupID = int64(tmp_SpectrumSharingGroupID.Value)
	tmp_EUTRACellsList := Sequence[*EUTRACellsListItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxCellineNB},
		ext: true,
	}
	fn := func() *EUTRACellsListItem { return new(EUTRACellsListItem) }
	if err = tmp_EUTRACellsList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read EUTRACellsList", err)
		return
	}
	ie.EUTRACellsList = []EUTRACellsListItem{}
	for _, i := range tmp_EUTRACellsList.Value {
		ie.EUTRACellsList = append(ie.EUTRACellsList, *i)
	}
	return
}
