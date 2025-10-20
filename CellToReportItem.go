package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type CellToReportItem struct {
	CellID            NRCGI               `mandatory`
	SSBToReportList   []SSBToReportItem   `lb:1,ub:maxnoofSSBs,optional,valExt`
	SliceToReportList []SliceToReportItem `lb:1,ub:maxnoofSliceItems,optional,valExt`
	// IEExtensions * `optional`
}

func (ie *CellToReportItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.SSBToReportList != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.SliceToReportList != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if err = ie.CellID.Encode(w); err != nil {
		err = utils.WrapError("Encode CellID", err)
		return
	}
	if len(ie.SSBToReportList) > 0 {
		tmp := Sequence[*SSBToReportItem]{
			Value: []*SSBToReportItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofSSBs},
			ext:   true,
		}
		for _, i := range ie.SSBToReportList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode SSBToReportList", err)
			return
		}
	}
	if len(ie.SliceToReportList) > 0 {
		tmp := Sequence[*SliceToReportItem]{
			Value: []*SliceToReportItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofSliceItems},
			ext:   true,
		}
		for _, i := range ie.SliceToReportList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode SliceToReportList", err)
			return
		}
	}
	return
}
func (ie *CellToReportItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if err = ie.CellID.Decode(r); err != nil {
		err = utils.WrapError("Read CellID", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_SSBToReportList := Sequence[*SSBToReportItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSSBs},
			ext: true,
		}
		fn := func() *SSBToReportItem { return new(SSBToReportItem) }
		if err = tmp_SSBToReportList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read SSBToReportList", err)
			return
		}
		ie.SSBToReportList = []SSBToReportItem{}
		for _, i := range tmp_SSBToReportList.Value {
			ie.SSBToReportList = append(ie.SSBToReportList, *i)
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_SliceToReportList := Sequence[*SliceToReportItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofSliceItems},
			ext: true,
		}
		fn := func() *SliceToReportItem { return new(SliceToReportItem) }
		if err = tmp_SliceToReportList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read SliceToReportList", err)
			return
		}
		ie.SliceToReportList = []SliceToReportItem{}
		for _, i := range tmp_SliceToReportList.Value {
			ie.SliceToReportList = append(ie.SliceToReportList, *i)
		}
	}
	return
}
