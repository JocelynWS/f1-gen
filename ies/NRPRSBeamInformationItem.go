package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type NRPRSBeamInformationItem struct {
	PRSResourceSetID PRSResourceSetID `mandatory`
	PRSAngleList     []PRSAngleItem   `lb:1,ub:maxnoofPRSResourcesPerSet,mandatory,valueExt`
	// IEExtensions * `optional`
}

func (ie *NRPRSBeamInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.PRSResourceSetID.Encode(w); err != nil {
		err = utils.WrapError("Encode PRSResourceSetID", err)
		return
	}
	if len(ie.PRSAngleList) > 0 {
		tmp := Sequence[*PRSAngleItem]{
			Value: []*PRSAngleItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofPRSResourcesPerSet},
			ext:   true,
		}
		for _, i := range ie.PRSAngleList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode PRSAngleList", err)
			return
		}
	} else {
		err = utils.WrapError("PRSAngleList is nil", err)
		return
	}
	return
}
func (ie *NRPRSBeamInformationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.PRSResourceSetID.Decode(r); err != nil {
		err = utils.WrapError("Read PRSResourceSetID", err)
		return
	}
	tmp_PRSAngleList := Sequence[*PRSAngleItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofPRSResourcesPerSet},
		ext: true,
	}
	fn := func() *PRSAngleItem { return new(PRSAngleItem) }
	if err = tmp_PRSAngleList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read PRSAngleList", err)
		return
	}
	ie.PRSAngleList = []PRSAngleItem{}
	for _, i := range tmp_PRSAngleList.Value {
		ie.PRSAngleList = append(ie.PRSAngleList, *i)
	}
	return
}
